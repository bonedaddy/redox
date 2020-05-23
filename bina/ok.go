package main

import (
	"fmt"
	"github.com/koinotice/redox/packages/goex"
	"github.com/koinotice/redox/packages/goex/okex"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	//"fmt"

	"net/http"
	"net/url"
	"time"
)

func stage( series *techan.TimeSeries){


	 cci := techan.NewCCIIndicator(series, 55)
	 fmt.Println(series.LastCandle())
	 fmt.Println(cci.Calculate(55))
	// fmt.Sprintf("%v",cci.Calculate(55).FormattedString(4))

}

func main()  {
	apiConfig:=goex.APIConfig{
		Endpoint: "https://www.okex.com",
		HttpClient: &http.Client{
			Transport: &http.Transport{
				Proxy: func(req *http.Request) (*url.URL, error) {
					return &url.URL{
						Scheme: "socks5",
						Host:   "127.0.0.1:15235"}, nil
				},
			},
		},
		ApiKey:        "",
		ApiSecretKey:  "",
		ApiPassphrase: "",
	}
	var okex = okex.NewOKEx(&apiConfig)
	var (
		//okexSpot = okex.OKExSpot
		okexSwap = okex.OKExSwap   //永续合约实现
		//okexFuture=okex.OKExFuture //交割合约实现
		//okexWallet =okex.OKExWallet //资金账户（钱包）操作
	)


	i := 0

	//ok.OKExV3FutureWs.KlineCallback(func(ticker *goex.FutureKline, period goex.KlinePeriod) {
	//	fmt.Println()
	//})

	okex.OKExV3FutureWs.KlineCallback(func(kline *goex.FutureKline,s int) {

		fmt.Print(  kline,s)

		fmt.Printf("%v,%s \n", kline.Kline,s)
	})
	//ok.OKExV3FutureWs.DepthCallback(func(depth *goex.Depth) {
	//	Log(depth)
	//})
	//ok.OKExV3FutureWs.TradeCallback(func(trade *goex.Trade, s string) {
	//	Log(s, trade)
	//})
	//ok.OKExV3FutureWs.OrderCallback(func(order *goex.FutureOrder, s string) {
	//	Log(s, order)
	//})
	okex.OKExV3FutureWs.SubscribeKline(goex.BTC_USDT, goex.QUARTER_CONTRACT,goex.KLINE_PERIOD_5MIN)


	for {
		select {
		case <-time.After(time.Second * time.Duration(1)):
			i++
			//if i == 5{
			//	fmt.Println("break now")
			//	goto ForEnd
			//}

			since := time.Now().Add(-200 * time.Minute * 1).Unix()
			//fmt.Println(since)
			kline, err := okexSwap.GetKlineRecords(goex.SWAP_CONTRACT, goex.BTC_USDT, goex.KLINE_PERIOD_1MIN, 0, int(since))
			if err!=nil{
				fmt.Println(err)
				goto ForEnd

			}else{
				series := techan.NewTimeSeries()
				for _, dt := range kline {

					period := techan.NewTimePeriod(time.Unix(dt.Kline.Timestamp, 0), time.Minute*1)

					candle := techan.NewCandle(period)
					candle.OpenPrice =big.NewDecimal(dt.Kline.Open )
					candle.ClosePrice = big.NewDecimal(dt.Kline.Close)
					candle.MaxPrice = big.NewDecimal(dt.Kline.High)
					candle.MinPrice = big.NewDecimal(dt.Kline.Low)
					candle.Volume=big.NewDecimal(dt.Kline.Vol)

					series.Candles = append(series.Candles, candle)
					//fmt.Println(candle.String())
					//series.AddCandle(candle)
				}



				go stage(series)


			}
			//fmt.Println("inside the select: ")
		}
		//fmt.Println("inside the for: ")
	}
	ForEnd:
	//okexSwap.GetKlineRecords("BTC-USD-SWAP","BTC-USD-SWAP",)
	////接口调用,更多接口调用请看代码
	//log.Println(okexSpot.GetAccount()) //获取账户资产信息
	////okexSpot.BatchPlaceOrders([]goex.Order{...}) //批量下单,单个交易对同时最大只能下10笔
	//log.Println(okexSwap.GetFutureUserinfo()) //获取账户权益信息
	//log.Println(okexFuture.GetFutureUserinfo())//获取账户权益信息
}
