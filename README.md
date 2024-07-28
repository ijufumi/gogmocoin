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

## How to use
### Public API
```golang
client := rest.New()

orderbooks, err := client.OrderBooks(consts.SymbolBCHJPY)

if err != nil {
    log.Println(err)
    return
}

log.Printf("[result]%+v", orderbooks)
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/master/app/public/rest)

### Public Websocket API
```golang
ticker := client.NewTicker(consts.SymbolBTCJPY)
timeoutCnt := 0
for {
    select {
    case v := <-ticker.Receive():
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
    log.Println(e)
    return
}
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/master/app/public/ws)

### Private API
#### 1. Create .env file
Copy `.env.example` to `.env` ant then input your `API_KEY` and `API_SECRET` to `.env` file.

```.env
API_KEY=YOUR_API_KEY
API_SECRET=YOUR_API_SECRET
```

#### 2. Execute API

```golang
client := private.New()
ordersRes, err := client.Orders(12345676879)
if err != nil {
    log.Println(err)
    return
}
log.Printf("ordersRes:%+v", ordersRes)
```

[Examples](https://github.com/ijufumi/gogmocoin-examples/tree/master/app/private)

## Welcome your contribution.
If you modified code by anything reasons (typo, bad coding, implements of features, etc...), please make `Issue` and `Pull Request`.
