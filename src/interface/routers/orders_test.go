package routers

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pinyu/datalab-drinks-backend/src/application/services"
	"github.com/pinyu/datalab-drinks-backend/src/interface/requests"
	"github.com/pinyu/datalab-drinks-backend/src/interface/responses"
	"github.com/stretchr/testify/assert"
)

func setupDB() {
	if os.Getenv("GIN_MODE") != "test" {
		log.Fatal("please set GIN_MODE=test")
	}

	services.DropTable()
	services.MigrateTable()
}

func cleanDB() {
	services.DeleteDB()
}

func newOrdersRequestBody(orderBy string, itemID uint, size string, sugarID uint, iceID uint) io.Reader {
	orderRequest := requests.OrderRequestBody{
		OrderBy: orderBy,
		ItemID:  itemID,
		Size:    size,
		SugarID: sugarID,
		IceID:   iceID,
	}

	marshalData, _ := json.Marshal(orderRequest)
	b := bytes.NewBuffer(marshalData)

	return b
}

func TestOrderDrinks(t *testing.T) {
	setupDB()

	router := setupRouter()

	// test POST orders
	testFirstOrder(t, router)
	testUpdateOrder(t, router)
	testSchema(t, router)
	testBadValue(t, router)

	// test GET orders
	testListOrders(t, router)

	cleanDB()
}

func testFirstOrder(t *testing.T, router *gin.Engine) {
	reqBody := newOrdersRequestBody("pinyu", 1, "medium", 1, 1)

	w := performRequest(router, "POST", "/v1/orders/", reqBody)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	assert.Equal(t, "drinks has been ordered", body.StatusMessage)
	assert.Nil(t, body.Payload)
}

func testUpdateOrder(t *testing.T, router *gin.Engine) {
	reqBody := newOrdersRequestBody("pinyu", 1, "large", 1, 1)

	w := performRequest(router, "POST", "/v1/orders/", reqBody)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	assert.Equal(t, "drinks has been updated", body.StatusMessage)
	assert.Nil(t, body.Payload)
}

func testSchema(t *testing.T, router *gin.Engine) {
	// empty order_by will be consider as bad schema
	reqBody := newOrdersRequestBody("", 1, "large", 1, 1)
	badSchema(t, router, reqBody)

	// item_id 0 will be consider as bad schema
	reqBody = newOrdersRequestBody("pinyu", 0, "large", 1, 1)
	badSchema(t, router, reqBody)

	// size which is not medium and not large will be consider as bad schema
	reqBody = newOrdersRequestBody("pinyu", 1, "L", 1, 1)
	badSchema(t, router, reqBody)

	// sugar_id 0 will be consider as bad schema
	reqBody = newOrdersRequestBody("pinyu", 1, "large", 0, 1)
	badSchema(t, router, reqBody)

	// ice_id 0 will be consider as bad schema
	reqBody = newOrdersRequestBody("pinyu", 1, "large", 1, 0)
	badSchema(t, router, reqBody)
}

func badSchema(t *testing.T, router *gin.Engine, reqBody io.Reader) {
	w := performRequest(router, "POST", "/v1/orders/", reqBody)
	assert.Equal(t, 400, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	assert.Equal(t, "bad schema of order request body", body.StatusMessage)
	assert.Nil(t, body.Payload)
}

func testBadValue(t *testing.T, router *gin.Engine) {
	// the range of item_id is [1, 22]
	reqBody := newOrdersRequestBody("pinyu", 23, "large", 1, 1)
	statusMessage := "invalid item_id"
	badValue(t, router, reqBody, statusMessage)

	// the range of sugar_id [1, 4]
	reqBody = newOrdersRequestBody("pinyu", 22, "large", 5, 2)
	statusMessage = "invalid sugar_id"
	badValue(t, router, reqBody, statusMessage)

	// drinks (item_id 22) cannot be made as hot
	reqBody = newOrdersRequestBody("pinyu", 22, "large", 2, 1)
	statusMessage = "the drinks should be ice"
	badValue(t, router, reqBody, statusMessage)

	// the range of ice_id [1, 4]
	reqBody = newOrdersRequestBody("pinyu", 1, "large", 1, 5)
	statusMessage = "invalid ice_id"
	badValue(t, router, reqBody, statusMessage)
}

func badValue(t *testing.T, router *gin.Engine, reqBody io.Reader, statusMessage string) {
	w := performRequest(router, "POST", "/v1/orders/", reqBody)
	assert.Equal(t, 400, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	assert.Equal(t, statusMessage, body.StatusMessage)
	assert.Nil(t, body.Payload)
}

func testListOrders(t *testing.T, router *gin.Engine) {
	// create the second order
	reqBody := newOrdersRequestBody("hsinwei", 2, "medium", 2, 2)
	performRequest(router, "POST", "/v1/orders/", reqBody)

	w := performRequest(router, "GET", "/v1/orders/", nil)
	assert.Equal(t, 200, w.Code)

	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	assert.Equal(t, "ok", body.StatusMessage)

	b, _ := json.Marshal(body.Payload)
	weekOrders := responses.WeekOrdersResponse{}
	json.Unmarshal(b, &weekOrders)

	order1 := weekOrders.DetailOrders[0]
	order2 := weekOrders.DetailOrders[1]

	assert.Equal(t, order1.OrderBy, "pinyu")
	assert.Equal(t, order1.Size, "large")
	assert.Equal(t, order1.Item, "黑咖啡")
	assert.Equal(t, order1.Price, uint(60))
	assert.Equal(t, order1.SugarTag, "無糖")
	assert.Equal(t, order1.IceTag, "熱")
	_, err := time.Parse(time.RFC3339, order1.OrderTime)
	assert.NoError(t, err, "time is not valid")

	assert.Equal(t, order2.OrderBy, "hsinwei")
	assert.Equal(t, order2.Size, "medium")
	assert.Equal(t, order2.Item, "特調咖啡")
	assert.Equal(t, order2.Price, uint(55))
	assert.Equal(t, order2.SugarTag, "微糖")
	assert.Equal(t, order2.IceTag, "去冰")
	_, err = time.Parse(time.RFC3339, order2.OrderTime)
	assert.NoError(t, err, "time is not valid")

}
