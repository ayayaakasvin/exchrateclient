package example

import (
	"testing"

	"github.com/ayayaakasvin/exchrateclient"
)

/*
	You can run this tests, using go test -v ./example/tests
	Make sure that you have downloaded this package into your local PC and set your api key inside example_test.go file
	Otherwise this will not work
	Feel free to change or add some test cases that will help to check client for issues and options
	Documentions of responce structs should be inside README.md file
*/

// Index checks if the index response is correct.
func Codes (cl exchrateclient.Client, tt *testing.T)  {
	codesMap, err := cl.FetchCodes()
	if err != nil {
		tt.Errorf("failed to retrieve codes: %s", err.Error())
		return
	}

	if codesMap == nil {
		tt.Errorf("failed to retrieve codes: empty map")
		return
	}
}

// Index checks if the index response is correct.
func Pair (cl exchrateclient.Client, tt *testing.T)	 {
	base, target := "USD",  "KZT"
	pairResp, err := cl.FetchPair(base, target)
	if err != nil {
		tt.Errorf("failed to retrieve pair: %s", err.Error())
		return
	}

	if pairResp == nil {
		tt.Errorf("failed to retrieve pair: nil struct responce")
		return
	}

	if pairResp.Rate <= 0 {
		tt.Errorf("unexpected rate: %.4f, expected > 0", pairResp.Rate)
	}

	if pairResp.BaseCode != base {
		tt.Errorf("unexpected code: %s, expected %s", pairResp.BaseCode, base)
	}

	if pairResp.TargetCode != target {
		tt.Errorf("unexpected code: %s, expected %s", pairResp.TargetCode, target)
	}
}

// Index checks if the index response is correct.
func Index (cl exchrateclient.Client, tt *testing.T)  {
	base := "USD"
	indexResponce, err := cl.FetchIndex(base)
	if err != nil {
		tt.Errorf("failed to retrieve index: %s", err.Error())
		return
	}

	if indexResponce == nil {
		tt.Errorf("failed to retrieve index: nil struct responce")
		return
	}

	if indexResponce.BaseCode != base {
		tt.Errorf("unexpected code: %s, expected %s", indexResponce.BaseCode, base)
	}

	if indexResponce.RateMap == nil {
		tt.Errorf("failed to retrieve index: nil map responce")
	}
}

func TestFetches (t *testing.T) {
	cl := exchrateclient.New("your_api_key")

	tests := []struct {
		name string
		test func(exchrateclient.Client, *testing.T)
	} {
		{"Test for code fetching", Codes},
		{"Test for index fetching", Index},
		{"Test for pair fetching", Pair},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.test(cl, t)
		})
	}
}
