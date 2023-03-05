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

type validationJson struct {
	IsValid bool
}

type errorJson struct {
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
		errorResp := new(errorJson)
		errorResp.Error = fmt.Sprint(err)

		return http.StatusBadRequest, errorResp
	}

	validationResult := appCore.CheckTicketCode(body.TicketNumber)

	validationResp := new(validationJson)
	validationResp.IsValid = validationResult.IsValid()

	return http.StatusOK, validationResp
}
