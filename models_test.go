package validator

import (
	"encoding/json"
	"testing"
)

const jsonStreamInvalidDate = `
{
	"basics": {
		"name": "John Doe",
		"label": "Programmer",
		"email": "john@gmail.com",
		"summary": "A summary of John Doe...",
		"profiles": [{
		  "network": "Twitter",
		  "username": "john"
		}]
  	},
	"work": [{
		"startDate": "2013-01-01",
		"endDate": "2014/01/01",
		"highlights": [
		  "Started the company"
		]
  	}]
}
`

func TestJsonTime_UnmarshalJSON_Successful(t *testing.T) {
	var resume Resume
	err := json.Unmarshal([]byte(jsonStream), &resume)

	if err != nil {
		t.Fatalf("unexpected err: %+v\n", err)
	}
}

func TestJsonTime_UnmarshalJSON_Unsuccessful(t *testing.T) {
	var resume Resume
	err := json.Unmarshal([]byte(jsonStreamInvalidDate), &resume)

	if err == nil {
		t.Fatalf("expected error but got nil")
	}
}
