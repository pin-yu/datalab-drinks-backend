# DataLab Drinks Backend

## Run the server

```bash
GIN_MODE=release go run src/main.go
```

The server will start listening on port 5000

## Run the unit and integration test

```bash
GIN_MODE=test go test ./...
```

## API

- GET: `/v1/menus`
    - return Cama's menu in json format

```json
{
    "status_message": "ok",
    "payload": {
        "menu_version": "2020W",
        "menu": [
            {
                "series": "現烘義式",
                "items": [
                    {
                        "item_id": 1,
                        "item": "黑咖啡",
                        "medium_price": 45,
                        "large_price": 60,
                        "cold": true,
                        "hot": true
                    }
                ]
            }
        ],
        "sugar": [
            {
                "sugar_id": 1,
                "sugar_tag": "無糖"
            },
            {
                "sugar_id": 2,
                "sugar_tag": "微糖"
            },
            {
                "sugar_id": 3,
                "sugar_tag": "半糖"
            },
            {
                "sugar_id": 4,
                "sugar_tag": "正常糖"
            }
        ],
        "ice": [
            {
                "ice_id": 1,
                "ice_tag": "熱"
            },
            {
                "ice_id": 2,
                "ice_tag": "去冰"
            },
            {
                "ice_id": 3,
                "ice_tag": "少冰"
            },
            {
                "ice_id": 4,
                "ice_tag": "正常冰"
            }
        ]
    }
}
```

- GET: `/v1/orders`
    - return the orders within `this week`!
    - the new week is defined if time pasts at 16:00 Friday
        - 15:59 Friday (this week)
        - 16:00 Friday (next week)
    - in json format
    - order_time follows RFC3339 format

```json
{
    "status_message": "ok",
    "payload": {
        "meeting_time": "2021-02-12T13:00:00+08:00",
        "total_price": 330,
        "aggregate_orders": [
            {
                "item": "特調咖啡",
                "size": "large",
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "sub_total_price": 140,
                "number": 2
            },
            {
                "item": "特調咖啡",
                "size": "medium",
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "sub_total_price": 55,
                "number": 1
            },
            {
                "item": "特調咖啡",
                "size": "medium",
                "sugar_tag": "半糖",
                "ice_tag": "去冰",
                "sub_total_price": 55,
                "number": 1
            },
            {
                "item": "紅茶拿鐵",
                "size": "large",
                "sugar_tag": "微糖",
                "ice_tag": "熱",
                "sub_total_price": 80,
                "number": 1
            }
        ],
        "week_orders": [
            {
                "order_by": "王祥郁",
                "item": "特調咖啡",
                "size": "medium",
                "price": 55,
                "sugar_tag": "半糖",
                "ice_tag": "去冰",
                "order_time": "2021-02-10T23:36:27+08:00"
            },
            {
                "order_by": "王平郁",
                "item": "紅茶拿鐵",
                "size": "large",
                "price": 80,
                "sugar_tag": "微糖",
                "ice_tag": "熱",
                "order_time": "2021-02-10T23:36:49+08:00"
            },
            {
                "order_by": "何星緯",
                "item": "特調咖啡",
                "size": "large",
                "price": 70,
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "order_time": "2021-02-11T11:42:00+08:00"
            },
            {
                "order_by": "吳義路",
                "item": "特調咖啡",
                "size": "large",
                "price": 70,
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "order_time": "2021-02-11T11:42:13+08:00"
            },
            {
                "order_by": "羅昱喬",
                "item": "特調咖啡",
                "size": "medium",
                "price": 55,
                "sugar_tag": "無糖",
                "ice_tag": "熱",
                "order_time": "2021-02-11T11:42:35+08:00"
            }
        ]
    }
}
```

- POST: `/v1/oders`
    - requires json body
    - if you want to update the order, just re-post
        - The server considers the value of `order_by` as key

```json
{
    "order_by": "平郁",
    "size": "medium",
    "item_id": 10,
    "sugar_id": 3,
    "ice_id": 2
}
```

## TODO:
- order validation (done)
- complete integrateion test (done)
- return an order report (done)
- dockerize (in development)
- get history orders, not only just get this week's order (in development)

