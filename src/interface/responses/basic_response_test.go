package responses

import (
	"encoding/json"
	"net/http"
	"testing"
)

func TestNewResponse(t *testing.T) {
	var testPayload struct {
		Test1 string `json:"test1"`
		Test2 int    `json:"test2"`
	}

	testPayload.Test1 = "test1"
	testPayload.Test2 = 99

	res := NewResponse(http.StatusAccepted, "system accepted", testPayload)
	status, body := res.Resolve()

	if status != http.StatusAccepted {
		t.Error("bad status in NewResponse")
	}

	marshalBody, _ := json.Marshal(body)
	expectedJSONString := `{"status_message":"system accepted","payload":{"test1":"test1","test2":99}}`

	if string(marshalBody) != expectedJSONString {
		t.Error("bad payload in responses.Body")
	}
}
