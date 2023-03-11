package newTicket

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/data/aztables"
)

type TicketEntity struct {
	aztables.Entity
	Code string
}

type successResponse struct {
}

type errorResponse struct {
	Error string
}

func Handle(resp http.ResponseWriter, req *http.Request) {
	status, body := handleInternal(req)
	enc := json.NewEncoder(resp)

	resp.WriteHeader(status)
	resp.Header().Set("Content-Type", "application/json")
	enc.Encode(&body)
}

func handleInternal(req *http.Request) (status int, result any) {
	defer errorHandler(&status, &result)

	identity, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return http.StatusInternalServerError, newErrorResponse(err)
	}

	serviceURL := fmt.Sprintf("https://%s.table.core.windows.net/%s", "febeticketstorage", "tickets")
	tables, err := aztables.NewClient(serviceURL, identity, nil)
	if err != nil {
		return http.StatusInternalServerError, newErrorResponse(err)
	}

	err = tryAddEntity(tables)
	if err != nil {
		return http.StatusInternalServerError, newErrorResponse(err)
	}

	return http.StatusOK, successResponse{}
}

func errorHandler(statusOut *int, resultOut *any) {
	panicArg := recover()
	if panicArg != nil {
		*statusOut = http.StatusInternalServerError
		*resultOut = errorResponse{
			Error: fmt.Sprint(panicArg),
		}
	}
}

func newErrorResponse(err error) errorResponse {
	errorMsg := fmt.Sprint(err)

	return errorResponse{
		Error: errorMsg,
	}
}

func tryAddEntity(tables *aztables.Client) error {
	ticket := TicketEntity{
		Entity: aztables.Entity{
			PartitionKey: "partkey",
			RowKey:       "EFGH",
		},
		Code: "EFGH",
	}

	ticketJs, err := json.Marshal(ticket)
	if err != nil {
		return err
	}

	_, err = tables.AddEntity(context.TODO(), ticketJs, nil)

	return err
}
