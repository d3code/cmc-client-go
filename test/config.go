package test

import (
	"github.com/d3code/cmc-client-go/cmc"
	"testing"
)

var testClient = cmc.Client("").Test(true).PrintResponse(true)

func checkResponse(t *testing.T, status cmc.Status) {
	if status.ErrorCode != 0 {
		t.Error("Non zero error code in response: ", status.ErrorMessage)
	}
}
