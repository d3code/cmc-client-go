package test

import (
	"fmt"
	"testing"

	"github.com/d3code/cmc-client-go/pkg/cmc"
)

func TestClient(t *testing.T) {
	testExchangeClient := cmc.Client("").Test(true).ExchangeClient()
	exchanges, err := testExchangeClient.GetExchangeMap()

	if err != nil {
		t.Error(err)
	}

	fmt.Println(exchanges)
}
