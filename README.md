# DataLab Drinks Backend

## Run the server

```bash
go run src/main.go
```

then the server starts listening on port 5000

## API

- GET: `/v1/menus` (done)
    - return Cama's menu in json format

<details><summary>click me to see detail json</summary>
<p>
```json
{
    "menu_version": "2020W",
    "menu": [
        {
            "series": "現烘義式",
            "items": [
                {
                    "id": 1,
                    "item": "黑咖啡",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 2,
                    "item": "特調咖啡",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 3,
                    "item": "卡布奇諾",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 4,
                    "item": "焦糖瑪琪朵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 5,
                    "item": "香草拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 6,
                    "item": "榛果拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 7,
                    "item": "海鹽焦糖拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 8,
                    "item": "黑糖拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 9,
                    "item": "蜂蜜拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 10,
                    "item": "拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                }
            ]
        },
        {
            "series": "精選茶飲",
            "items": [
                {
                    "id": 11,
                    "item": "蜂蜜鮮奶茶",
                    "flavor": null,
                    "price": {
                        "large": 85,
                        "medium": 65
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 12,
                    "item": "紅茶拿鐵",
                    "flavor": null,
                    "price": {
                        "large": 80,
                        "medium": 60
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 13,
                    "item": "皇家奶茶",
                    "flavor": null,
                    "price": {
                        "large": 50,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 14,
                    "item": "抹茶鮮奶",
                    "flavor": null,
                    "price": {
                        "large": 80,
                        "medium": 65
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 15,
                    "item": "芒果烤茶",
                    "flavor": null,
                    "price": {
                        "large": 65,
                        "medium": 0
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 16,
                    "item": "英倫早餐茶",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 17,
                    "item": "花草茶",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 18,
                    "item": "日式烘烤煎茶",
                    "flavor": null,
                    "price": {
                        "large": 0,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": true
                }
            ]
        },
        {
            "series": "其他飲品",
            "items": [
                {
                    "id": 19,
                    "item": "經典巧克力",
                    "flavor": null,
                    "price": {
                        "large": 95,
                        "medium": 75
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 20,
                    "item": "純釀烏梅汁",
                    "flavor": null,
                    "price": {
                        "large": 50,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": true
                },
                {
                    "id": 21,
                    "item": "風味果醋",
                    "flavor": [
                        "青蘋",
                        "野櫻梅"
                    ],
                    "price": {
                        "large": 50,
                        "medium": 40
                    },
                    "cold": true,
                    "hot": false
                }
            ]
        }
    ],
    "sugar": [
        {
            "id": "1",
            "tag": "無糖"
        },
        {
            "id": "2",
            "tag": "微糖"
        },
        {
            "id": "3",
            "tag": "半糖"
        },
        {
            "id": "4",
            "tag": "正常糖"
        }
    ],
    "ice": [
        {
            "id": "1",
            "tag": "熱"
        },
        {
            "id": "2",
            "tag": "少冰"
        },
        {
            "id": "3",
            "tag": "正常冰"
        }
    ]
}
```
</p>
</details>

- GET: `/v1/orders` (in development)
    - return the orders within `this week`!
    - the new week is defined if time pasts at 12:00 Friday
        - 11:59 Friday (this week)
        - 12:01 Friday (next week)

- POST: `/v1/oders`
    - require parameters
        - `/v1/orders?order_by=平郁&item_id=1&sugar=3&ice=4`

    - if you want to update the order, just re-post
        - The server considers the value of `order_by` as key

## Architecture
Basically, this is a Domain Drive Design project. There are three layers in the architecture.

- api
    - routes
        - every route should have one or more than one service.
        - routes `should not` talk to the domain objects
    - services
        - services are the place where to talk to the domain objects

- domain
    - menus
        - handle the menus entity
    - orders
        - handle the orders entity
    - users (I not sure whether we need this)
        - handle the users entity
- infra
    - domain logic `should not` appear here.
    - orm
        - migration or simple database communication

- suggestion
    - please follow the design rules
    - each route has one or more than one service
    - service can use functions in domain package
    - domain can use functions in infra package
    - `all the logic should place in domain package`

## TODO:
- get history orders, not only just get this week's order
- complete test cases