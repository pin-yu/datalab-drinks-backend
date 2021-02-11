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
	"github.com/pinyu/datalab-drinks-backend/src/utils"
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
	performOrderRequests(router)

	// get the orders
	w := performRequest(router, "GET", "/v1/orders/", nil)
	assert.Equal(t, 200, w.Code)

	// parse status message
	body := responses.Body{}
	json.Unmarshal([]byte(w.Body.String()), &body)

	// check status message
	assert.Equal(t, "ok", body.StatusMessage)

	// parse payload
	b, _ := json.Marshal(body.Payload)
	orders := responses.OrdersResponse{}
	json.Unmarshal(b, &orders)

	// -------------------------
	//       check payload
	// -------------------------

	// check meeting time
	assert.Equal(t, utils.MeetingStartTime().Format(time.RFC3339), orders.MeetingTime)

	// check totalPrice
	assert.Equal(t, uint(240), orders.TotalPrice)

	// check aggregateOrders
	assert.Len(t, orders.AggregateOrders, 3)

	assertAggregateOrder(t, &orders.AggregateOrders[0], "黑咖啡", "large", "無糖", "熱", uint(60), 1)
	assertAggregateOrder(t, &orders.AggregateOrders[1], "特調咖啡", "large", "微糖", "去冰", uint(70), 1)
	assertAggregateOrder(t, &orders.AggregateOrders[2], "特調咖啡", "medium", "微糖", "去冰", uint(110), 2)

	// check detailOrders
	assert.Len(t, orders.DetailOrders, 4)
	assertDetailOrder(t, &orders.DetailOrders[0], "pinyu", "large", "黑咖啡", uint(60), "無糖", "熱")
	assertDetailOrder(t, &orders.DetailOrders[1], "hsinwei", "medium", "特調咖啡", uint(55), "微糖", "去冰")
	assertDetailOrder(t, &orders.DetailOrders[2], "yilu", "medium", "特調咖啡", uint(55), "微糖", "去冰")
	assertDetailOrder(t, &orders.DetailOrders[3], "yuchiao", "large", "特調咖啡", uint(70), "微糖", "去冰")
}

func performOrderRequests(router *gin.Engine) {
	// create the second order
	reqBody := newOrdersRequestBody("hsinwei", 2, "medium", 2, 2)
	performRequest(router, "POST", "/v1/orders/", reqBody)

	// create the third order
	reqBody = newOrdersRequestBody("yilu", 2, "medium", 2, 2)
	performRequest(router, "POST", "/v1/orders/", reqBody)

	// create the fourth order
	reqBody = newOrdersRequestBody("yuchiao", 2, "large", 2, 2)
	performRequest(router, "POST", "/v1/orders/", reqBody)
}

func assertAggregateOrder(t *testing.T, order *responses.AggregateOrderResponse, item string, size string, sugarTag string, iceTag string, subTotalPrice uint, number uint) {
	assert.Equal(t, order.Item, item)
	assert.Equal(t, size, order.Size)
	assert.Equal(t, sugarTag, order.SugarTag)
	assert.Equal(t, iceTag, order.IceTag)
	assert.Equal(t, subTotalPrice, order.SubTotalPrice)
	assert.Equal(t, number, order.Number)
}

func assertDetailOrder(t *testing.T, order *responses.OrderResponse, orderBy string, size string, item string, price uint, sugarTag string, iceTag string) {
	assert.Equal(t, orderBy, order.OrderBy)
	assert.Equal(t, size, order.Size)
	assert.Equal(t, item, order.Item)
	assert.Equal(t, price, order.Price)
	assert.Equal(t, sugarTag, order.SugarTag)
	assert.Equal(t, iceTag, order.IceTag)
	_, err := time.Parse(time.RFC3339, order.OrderTime)
	assert.NoError(t, err, "time is not valid")
}
