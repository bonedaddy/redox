package main

import (
	"fmt"
	"github.com/koinotice/redox/packages/goex"
	"github.com/koinotice/redox/packages/goex/binance"
	//"fmt"

	"net/http"
	"time"
)

var bina = binance.NewBinanceSwap(&goex.APIConfig{
	Endpoint: "https://testnet.binancefuture.com",
	HttpClient: &http.Client{
		//Transport: &http.Transport{
		//	Proxy: func(req *http.Request) (*url.URL, error) {
		//		return url.Parse("socks5://127.0.0.1:15235")
		//		return nil, nil
		//	},
		//},
		Timeout: 10 * time.Second,
	},
	ApiKey:       "f3891cec51cbdf6c34b00bd5a50fa0ab4558e990a7d3c85fa5a193b10767f788",
	ApiSecretKey: "c9e79f5f1812029c09e6f33ec00edf16ace3294bc23a90d29d4380f36b976fcb",
})

func main() {

	 order,_:=bina.PlaceFutureOrder(goex.BTC_USDT, "", "9150", "0.01", goex.CLOSE_SELL, 1, 0)
	//fmt.Print(order)

	//order,_:=bina.GetFutureOrder("2315175995", goex.BTC_USDT, "")
	fmt.Print(order)
}
