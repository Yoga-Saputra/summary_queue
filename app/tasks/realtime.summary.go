package tasks

import (
	"context"
	"encoding/json"
	"log"
	"summary/app/controller"
	"time"

	"github.com/hibiken/asynq"
)

// A list of task types.
const (
	TypeSummaryDelivery = "summary:deliver"
	DDMMYYYY            = "RFC1123Z"
)

type (
	SummaryDeliveryPayload struct {
		DateStart    time.Time
		ProviderCode string
	}
)

//----------------------------------------------
// Write a function NewXXXTask to create a task.
// A task consists of a type and a payload.
//----------------------------------------------

func NewRealTimeSummaryDeliveryTask(DateStart time.Time, ProviderCode string) *asynq.Task {
	payload, err := json.Marshal(SummaryDeliveryPayload{DateStart: DateStart, ProviderCode: ProviderCode})
	if err != nil {
		return nil
	}
	return asynq.NewTask(TypeSummaryDelivery, payload)
}

// ---------------------------------------------------------------
// Write a function HandleXXXTask to handle the input task.
// Note that it satisfies the asynq.HandlerFunc interface.
//
// Handler doesn't need to be a function. You can define a type
// that satisfies asynq.Handler interface. See examples below.
// ---------------------------------------------------------------

func HandleRealTimeSummaryDeliveryTask(ctx context.Context, t *asynq.Task) error {
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

	_, err = controller.CreateOrUpdate(inputStartDate, providerCode)

	if err != nil {
		//("info", err.Error())
		return err
	}

	return nil
}
