package routers

import (
	"encoding/json"
	"testing"

	"github.com/pin-yu/datalab-drinks-backend/src/interface/responses"
	"github.com/stretchr/testify/assert"
)

func TestReadMenu(t *testing.T) {
	router := setupRouter()

	w := performRequest(router, "GET", "/v2/menus/", nil)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, "ok", body.StatusMessage)

	assert.Equal(t, "2021-spring", body.Payload.(map[string]interface{})["menu_version"])

	testWuLong(t, &body)
	testVanillaLatte(t, &body)
}

func testWuLong(t *testing.T, body *responses.Body) {
	// { "menu": [{...}, {"series": "精選茶飲", "items": [{}, {"item":"鮮奶清焙烏龍"...}]}]}
	menu := body.Payload.(map[string]interface{})["menu"].([]interface{})[1].(map[string]interface{})
	assert.Equal(t, "精選茶飲", menu["series"])

	// check item
	item := menu["items"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(14), uint(item["item_id"].(float64)))
	assert.Equal(t, "鮮奶清焙烏龍", item["item"])
	assert.Equal(t, uint(70), uint(item["large_price"].(float64)))
	assert.Equal(t, uint(55), uint(item["medium_price"].(float64)))

	sugar := item["sugars"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(sugar["sugar_id"].(float64)))
	assert.Equal(t, "正常糖", sugar["sugar_tag"])
	assert.True(t, sugar["enable"].(bool))

	ice := item["ices"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(ice["ice_id"].(float64)))
	assert.Equal(t, "正常冰", ice["ice_tag"])
	assert.True(t, ice["enable"].(bool))
}

func testVanillaLatte(t *testing.T, body *responses.Body) {
	// { "menu": [{...}, {...}, {"series": "其他飲品", "items": [{}, {"item":"純釀烏梅汁"...}]}]}
	menu := body.Payload.(map[string]interface{})["menu"].([]interface{})[0].(map[string]interface{})
	assert.Equal(t, "咖啡系列", menu["series"])

	// check item
	item := menu["items"].([]interface{})[4].(map[string]interface{})
	assert.Equal(t, uint(5), uint(item["item_id"].(float64)))
	assert.Equal(t, "香草拿鐵", item["item"])
	assert.Equal(t, uint(95), uint(item["large_price"].(float64)))
	assert.Equal(t, uint(75), uint(item["medium_price"].(float64)))

	sugar := item["sugars"].([]interface{})[2].(map[string]interface{})
	assert.Equal(t, uint(3), uint(sugar["sugar_id"].(float64)))
	assert.Equal(t, "半糖", sugar["sugar_tag"])
	assert.False(t, sugar["enable"].(bool))

	sugar = item["sugars"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(sugar["sugar_id"].(float64)))
	assert.Equal(t, "正常糖", sugar["sugar_tag"])
	assert.True(t, sugar["enable"].(bool))

	ice := item["ices"].([]interface{})[3].(map[string]interface{})
	assert.Equal(t, uint(4), uint(ice["ice_id"].(float64)))
	assert.Equal(t, "正常冰", ice["ice_tag"])
	assert.True(t, ice["enable"].(bool))
}
