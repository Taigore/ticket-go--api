package checkCode

import (
	"encoding/json"
	"net/http"

	appCore "github.com/Taigore/ticket-go/app-core"
)

type validationJson struct {
	IsValid bool
}

func Handle(resp http.ResponseWriter, req *http.Request) {
	validationResult := appCore.CheckTicketCode(" ")
	validationJson := new(validationJson)
	validationJson.IsValid = validationResult.IsValid()

	enc := json.NewEncoder(resp)

	resp.WriteHeader(http.StatusOK)
	enc.Encode(&validationJson)
}
