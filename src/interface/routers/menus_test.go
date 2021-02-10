package routers

import (
	"encoding/json"
	"testing"

	"github.com/pinyu/datalab-drinks-backend/src/interface/responses"
	"github.com/stretchr/testify/assert"
)

func TestReadMenu(t *testing.T) {
	router := setupRouter()

	w := performRequest(router, "GET", "/v1/menus/", nil)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)
	assert.Equal(t, "ok", body.StatusMessage)
	assert.NotNil(t, body.Payload)
}
