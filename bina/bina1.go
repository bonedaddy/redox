package main

import (
	"fmt"
	"github.com/koinotice/redox/packages/goex"
	"github.com/koinotice/redox/packages/goex/binance"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"net/url"

	//"fmt"

	"net/http"
	"time"
)

//func stage(series *techan.TimeSeries) {
//
//	cci := techan.NewCCIIndicator(series, 55)
//	fmt.Println(series.LastCandle())
//	fmt.Println(cci.Calculate(55))
//	// fmt.Sprintf("%v",cci.Calculate(55).FormattedString(4))
//
//}

var series = techan.NewTimeSeries()

func stage1(  ){

	fmt.Println(series.LastIndex())
	cci := techan.NewCCIIndicator(series, 55)
    fmt.Println(series.LastCandle())
	//for i:=range series.Candles{
	//	fmt.Println(i)
	//}
	for i := series.LastIndex(); i > series.LastIndex()-5; i-- {
		fmt.Println(i,cci.Calculate(i).FormattedString(4))
	}
	//fmt.Println(cci.Calculate(series.LastIndex()-1))
	// fmt.Sprintf("%v",cci.Calculate(55).FormattedString(4))

}

var bnWs = binance.NewBinanceWs()

func init(){
	//bnWs.SetBaseUrl("wss://stream.binancefuture.com/ws")
	bnWs.ProxyUrl ( "socks5://127.0.0.1:15235")

	bnWs.SetBaseUrl("wss://fstream.binance.com/ws")

	bnWs.SetCombinedBaseURL("wss://fstream.binance.com/stream?streams=")

}
var bs = binance.NewBinanceSwap(&goex.APIConfig{
	//Endpoint: goex.US_API_BASE_URL,
	HttpClient: &http.Client{
		Transport: &http.Transport{
			Proxy: func(req *http.Request) (*url.URL, error) {
				return url.Parse("socks5://127.0.0.1:15235")
				return nil, nil
			},
		},
		Timeout: 10 * time.Second,
	},
	ApiKey:       "f3891cec51cbdf6c34b00bd5a50fa0ab4558e990a7d3c85fa5a193b10767f788",
	ApiSecretKey: "c9e79f5f1812029c09e6f33ec00edf16ace3294bc23a90d29d4380f36b976fcb",
})


var isExsit= techan.TimePeriod{}
func getHistory(){
	kline, _ := bs.GetKlineRecords("", goex.BTC_USDT, goex.KLINE_PERIOD_5MIN, 1500, 0)

	//fmt.Print(kline)
	for _, dt := range kline {
		//fmt.Println(dt.Kline)
		period := techan.NewTimePeriod(time.Unix(dt.Kline.Timestamp, 0), time.Minute*5)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(dt.Kline.Open)
		candle.ClosePrice = big.NewDecimal(dt.Kline.Close)
		candle.MaxPrice = big.NewDecimal(dt.Kline.High)
		candle.MinPrice = big.NewDecimal(dt.Kline.Low)
		candle.Volume = big.NewDecimal(dt.Kline.Vol)

		//series.Candles = append(series.Candles, candle)
		//fmt.Println(candle.String())
		series.AddCandle(candle)
	}
	isExsit=series.LastCandle().Period

	watch()
	//fmt.Println(series.LastIndex())



}

func watch(){

	bnWs.KlineCallback = func(kline *goex.Kline, p int) {

		period := techan.NewTimePeriod(time.Unix(kline.Timestamp, 0), time.Minute*5)
		//fmt.Print(period.String())
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(kline.Open)
		candle.ClosePrice = big.NewDecimal(kline.Close)
		candle.MaxPrice = big.NewDecimal(kline.High)
		candle.MinPrice = big.NewDecimal(kline.Low)
		candle.Volume = big.NewDecimal(kline.Vol)
		if isExsit==period{
			fmt.Printf("update:::%s,%s \n", isExsit,period)

			series.Candles[len(series.Candles)-1]=candle

		}else{
			fmt.Printf("new:::%s,%s \n", isExsit,period)
			isExsit=period
			series.AddCandle(candle)
		}
		  stage1()

	}

	err := bnWs.SubscribeKline(goex.BTC_USDT, goex.KLINE_PERIOD_5MIN)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {



	  getHistory()



	for {
		select {
		case <-time.After(time.Second * time.Duration(1)):



		}
	}
	//
	//select {}
}
