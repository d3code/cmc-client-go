package test

import (
	"github.com/d3code/cmc-client-go/cmc"
	"testing"
)

var testClient = cmc.Client("").Test(true)

func checkResponse(t *testing.T, err error, status cmc.Status) {
	if err != nil {
		t.Error(err)
	}

	if status.ErrorCode != 0 {
		t.Error("Non zero error code in response: ", status.ErrorMessage)
	}
}
