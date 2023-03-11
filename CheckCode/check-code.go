package checkCode

import (
	"encoding/json"
	"fmt"
	"net/http"

	appCore "github.com/Taigore/ticket-go--core"
)

type ticketBody struct {
	TicketNumber string
}

type validationResponse struct {
	IsValid bool
}

type errorResponse struct {
	Error string
}

func Handle(resp http.ResponseWriter, req *http.Request) {
	status, body := handleInner(req)
	enc := json.NewEncoder(resp)

	resp.WriteHeader(status)
	resp.Header().Set("Content-Type", "application/json")
	enc.Encode(&body)
}

func handleInner(req *http.Request) (status int, result any) {
	dec := json.NewDecoder(req.Body)
	dec.DisallowUnknownFields()

	var body ticketBody
	err := dec.Decode(&body)

	if err != nil {
		return http.StatusBadRequest, newErrorJson(err)
	}

	validationResult := appCore.CheckTicketCode(body.TicketNumber)

	validationResp := validationResponse{
		IsValid: validationResult.IsValid(),
	}

	return http.StatusOK, validationResp
}

func newErrorJson(err error) errorResponse {
	return errorResponse{
		Error: fmt.Sprint(err),
	}
}
