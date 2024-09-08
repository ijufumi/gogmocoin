# Wrapper code of GMO Coin's API
## API document
https://api.coin.z.com/docs

## Implemented API list
### Public API
* `/public/v1/status`
* `/public/v1/ticker`
* `/public/v1/orderbooks`
* `/public/v1/trades`
* `/public/v1/klines`
* `/public/v1/symbols`

### Public Websocket API
* `ticker`
* `orderbooks`
* `trades`

### Private API
* `/private/v1/account/margin`
* `/private/v1/account/assets`
* `/private/v1/orders`
* `/private/v1/activeOrders`
* `/private/v1/executions`
* `/private/v1/latestExecutions`
* `/private/v1/openPositions`
* `/private/v1/positionSummary`
* `/private/v1/order`
* `/private/v1/changeOrder`
* `/private/v1/cancelOrder`
* `/private/v1/closeOrder`
* `/private/v1/closeBulkOrder`
* `/private/v1/changeLosscutPrice`
* `/private/v1/ws-auth`

### Private Websocket API
* `executionEvents`
* `orderEvents`
* `positionEvents`
* `positionSummaryEvents`

## How to use
### Public API
```golang
client := rest.New()

orderbooks, err := client.OrderBooks(consts.SymbolBCHJPY)

if err != nil {
    log.Fatalln(err)
}

log.Printf("[result]%+v", orderbooks)
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/main/app/public/rest)

### Public Websocket API
```golang
client := ws.NewTicker(consts.SymbolBTCJPY)
timeoutCnt := 0
err := client.Subscribe()
if err != nil {
    log.Fatalln(err)
}
for {
    select {
    case v := <-client.Receive():
        log.Printf("msg:%+v", v)
    case <-time.After(180 * time.Second):
        log.Println("timeout...")
        timeoutCnt++
    }
    if timeoutCnt > 20 {
        break
    }
}
e := client.Unsubscribe()
if e != nil {
    log.Fatalln(e)
}
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/main/app/public/ws)

### Private API
#### 1. Create .env file
Copy `.env.example` to `.env` ant then input your `API_KEY` and `API_SECRET` to `.env` file.

```.env
API_KEY=YOUR_API_KEY
API_SECRET=YOUR_API_SECRET
```

#### 2. Execute API

##### REST API
```golang
client := rest.New()
ordersRes, err := client.Orders(12345676879)
if err != nil {
    log.Fatalln(err)
}
log.Printf("ordersRes:%+v", ordersRes)
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/main/app/private/rest)

##### Websocket API
```golang
client := ws.NewExecutionEvents(true)
if err := client.Subscribe(); err != nil {
    log.Fatal(err)
}
timeoutCnt := 0
for {
    select {
    case v := <-client.Receive():
        log.Printf("msg:%+v\n", v)
    case <-time.After(180 * time.Second):
        log.Println("timeout...")
        timeoutCnt++
    }
    if timeoutCnt > 20 {
        break
    }
}
if err := client.Unsubscribe(); err != nil {
    log.Fatal(err)
}
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/main/app/private/ws)


## Welcome your contribution.
If you modified code by anything reasons (typo, bad coding, implements of features, etc...), please make `Issue` and `Pull Request`.
