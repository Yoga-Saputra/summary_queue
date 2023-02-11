package router

import (
	"log"
	"summary/app/helpers"
	"summary/app/models"
	"summary/app/tasks"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gookit/validate"
	"github.com/hibiken/asynq"
)

type Response struct {
	Success bool        `json:"success"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Error   interface{} `json:"error"`
}

const (
	DDMMYYYY = "2006-01-02"
)

func SummaryRefetch(c *fiber.Ctx) error {

	client := asynq.NewClient(asynq.RedisClientOpt{Addr: "localhost:6379"})
	var input *models.CreateSummaryInput
	c.BodyParser(&input)

	v := validate.Struct(input)
	if !v.Validate() {
		return c.JSON(Response{
			Success: false,
			Code:    fiber.StatusBadRequest,
			Data:    v.Errors.One(),
		})
	}

	inputStartDate, _ := time.Parse(DDMMYYYY, input.RangeDate)
	m := inputStartDate.Month()

	t1 := tasks.NewRefetchSummaryDeliveryTask(inputStartDate, input.ProviderCode)

	// Process the task immediately.
	_, err := client.Enqueue(t1)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf(" [*] Successfully enqueued task")

	tableName := helpers.GetTablename(input.ProviderCode, m)

	return c.JSON(Response{
		Success: true,
		Code:    1,
		Data:    "success re summary table => " + tableName,
	})

}

func Route(app fiber.Router) {
	app.Post("/summary-fetcher", SummaryRefetch) //using queues
}
