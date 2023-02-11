package tasks

import (
	"context"
	"encoding/json"
	"log"
	"summary/app/controller"
	"summary/config"
	"time"

	"github.com/hibiken/asynq"
)

const (
	TypeReSummaryDelivery = "resummary:deliver"
)

type (
	ReSummaryDeliveryPayload struct {
		DateStart    time.Time
		ProviderCode string
	}
)

// ----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
// ----------------------------------------------
func NewRefetchSummaryDeliveryTask(DateStart time.Time, ProviderCode string) *asynq.Task {
	payload, err := json.Marshal(ReSummaryDeliveryPayload{DateStart: DateStart, ProviderCode: ProviderCode})
	if err != nil {
		return nil
	}
	return asynq.NewTask(TypeReSummaryDelivery, payload)
}

// ---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
// ---------------------------------------------------------------

func HandleRefetchSummaryDeliveryTask(ctx context.Context, t *asynq.Task) error {

	arr := make(map[string]string)
	payl := t.Payload()

	// byte to array
	json.Unmarshal([]byte(payl), &arr)
	fromDate := arr["DateStart"]
	providerCode := arr["ProviderCode"]
	inputStartDate, err := time.Parse(time.RFC3339, string(fromDate))

	if err != nil {
		log.Println(err)
	}

	payloadReport, err := controller.CreateOrUpdate(inputStartDate, providerCode)

	if err != nil {
		//("info", err.Error())
		return err
	}

	countData := len(payloadReport)
	data := map[string]interface{}{
		"countData": countData,
		"data":      payloadReport,
	}

	dataJson, _ := json.Marshal(data)

	config.Loggers("info", string(dataJson))

	return nil
}
