package routers

import (
	"encoding/json"
	"testing"

	"github.com/pin-yu/datalab-drinks-backend/src/interface/responses"
	"github.com/stretchr/testify/assert"
)

func TestReadMenu(t *testing.T) {
	router := setupRouter()

	w := performRequest(router, "GET", "/v1/menus/", nil)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)
	assert.Equal(t, "ok", body.StatusMessage)

	assert.Equal(t, "2020W", body.Payload.(map[string]interface{})["menu_version"])

	menu := body.Payload.(map[string]interface{})["menu"].([]interface{})[2].(map[string]interface{})
	assert.Equal(t, "其他飲品", menu["series"])

	// check item
	item := menu["items"].([]interface{})[1].(map[string]interface{})
	assert.Equal(t, uint(20), uint(item["item_id"].(float64)))
	assert.Equal(t, "純釀烏梅汁", item["item"])
	assert.Equal(t, uint(50), uint(item["large_price"].(float64)))
	assert.Equal(t, uint(40), uint(item["medium_price"].(float64)))
	assert.Equal(t, true, item["sugar"])
	assert.Equal(t, true, item["cold"])
	assert.Equal(t, true, item["hot"])

	// check sugar_id and sugar_tag
	sugar := body.Payload.(map[string]interface{})["sugar"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(sugar["sugar_id"].(float64)))
	assert.Equal(t, "正常糖", sugar["sugar_tag"])

	// check ice_id and ice_tag
	ice := body.Payload.(map[string]interface{})["ice"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(ice["ice_id"].(float64)))
	assert.Equal(t, "正常冰", ice["ice_tag"])
}
