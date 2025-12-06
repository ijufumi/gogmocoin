# GMOコインAPIのラッパーコード
## APIドキュメント
https://api.coin.z.com/docs

## 実装済みAPI一覧
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

## 使い方
### Public API
```golang
client := rest.New()

orderbooks, err := client.OrderBooks(consts.SymbolBCHJPY)

if err != nil {
    log.Fatalln(err)
}

log.Printf("[result]%+v", orderbooks)
```

[サンプルコード](httpss://github.com/ijufumi/gogmocoin-examples/tree/main/app/public/rest)

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

[サンプルコード](httpss://github.com/ijufumi/gogmocoin-examples/tree/main/app/public/ws)

### Private API
#### 1. .envファイルを作成する
`.env.example`を`.env`にコピーし、あなたの`API_KEY`と`API_SECRET`を`.env`ファイルに記述してください。

```.env
API_KEY=YOUR_API_KEY
API_SECRET=YOUR_API_SECRET
```

#### 2. APIを実行する

##### REST API
```golang
client := rest.New()
ordersRes, err := client.Orders(12345676879)
if err != nil {
    log.Fatalln(err)
}
log.Printf("ordersRes:%+v", ordersRes)
```

[サンプルコード](httpss://github.com/ijufumi/gogmocoin-examples/tree/main/app/private/rest)

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

[サンプルコード](httpss://github.com/ijufumi/gogmocoin-examples/tree/main/app/private/ws)


## コントリビュートをお待ちしています
もしあなたが何かしらの理由（タイポ、イケてないコード、機能追加など）でコードを修正した場合は、`Issue`と`Pull Request`を作成してください。
