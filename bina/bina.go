package main

import (
	"fmt"
	"github.com/koinotice/redox/packages/goex"
	"github.com/koinotice/redox/packages/goex/binance"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
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
var bnWs = binance.NewBinanceWs()

var (
	CAN_OPEN_BUY bool =true
	CAN_CLOSE_BUY bool =false
	CAN_OPEN_SELL bool =true
	CAN_CLOSE_SELL bool =false
)

func init() {
	bnWs.SetBaseUrl("wss://stream.binancefuture.com/ws")
	bnWs.SetCombinedBaseURL("wss://stream.binancefuture.com/stream?streams=")

}

var bs = binance.NewBinanceSwap(&goex.APIConfig{
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
func stage1() {

	//fmt.Println(series.LastIndex())
	cci := techan.NewCCIIndicator(series, 55)
	//fmt.Println(series.LastCandle())
	//for i:=range series.Candles{
	//	fmt.Println(i)
	//}
	for i := series.LastIndex(); i > series.LastIndex()-5; i-- {
		fmt.Println(i, cci.Calculate(i).FormattedString(4))
	}

	//indicator := techan.NewCCIIndicator(series, 20)
	// record trades on this object
	//record := techan.NewTradingRecord()

	buyConstant := techan.NewConstantIndicator(100)
	sellConstant := techan.NewConstantIndicator(-100)

	openBuyRule := techan.NewCrossUpIndicatorRule(buyConstant, cci)
	closeBuyRule := techan.NewCrossDownIndicatorRule(cci, buyConstant)
	openSellRule := techan.NewCrossDownIndicatorRule(cci, sellConstant)
	closeSellRule := techan.NewCrossUpIndicatorRule(sellConstant, cci)



	openBuy:=openBuyRule.IsSatisfied( series.LastIndex() , nil)
	closeBuy:=closeBuyRule.IsSatisfied( series.LastIndex() , nil)
	openSell:=openSellRule.IsSatisfied( series.LastIndex() , nil)
	closeSell:=closeSellRule.IsSatisfied( series.LastIndex() , nil)

	if openBuy&&CAN_OPEN_BUY{
		order,_:=bs.PlaceFutureOrder(goex.BTC_USDT, "", "", "1", goex.OPEN_BUY, 1, 0)
		CAN_OPEN_BUY=false //已开多单
		CAN_CLOSE_BUY=true
		fmt.Println("-----------")
		fmt.Println("开多",order)
	}
	if closeBuy&&CAN_CLOSE_BUY{
		order,_:=bs.PlaceFutureOrder(goex.BTC_USDT, "", "", "1", goex.CLOSE_BUY, 1, 0)
		CAN_OPEN_BUY=true
		CAN_CLOSE_BUY=false

		fmt.Println("-----------")

		fmt.Println("平多",order)
	}
	if openSell&&!CAN_OPEN_SELL{
		order,_:=bs.PlaceFutureOrder(goex.BTC_USDT, "", "", "1", goex.OPEN_SELL, 1, 0)
		CAN_OPEN_SELL=false
		CAN_CLOSE_SELL=true
		fmt.Print("开空",order)
	}
	if closeSell&&CAN_CLOSE_SELL{
		order,_:=bs.PlaceFutureOrder(goex.BTC_USDT, "", "", "1", goex.CLOSE_SELL, 1, 0)
		CAN_OPEN_SELL=true
		CAN_CLOSE_SELL=false
		fmt.Print("平空",order)
	}

	// Is satisfied when the price ema moves above 30 and the current position is new

	//exitRule := exitConstant..NewCrossDownIndicatorRule(indicator, exitConstant)
	////techan.And(
	////techan.NewCrossDownIndicatorRule(indicator, exitConstant),
	////techan.PositionOpenRule{}) // Is satisfied when the price ema moves below 10 and the current position is open
	//
	//strategy := techan.RuleStrategy{
	//	UnstablePeriod: 10,
	//	EntryRule:      entryRule,
	//	ExitRule:       exitRule,
	//}
	//isBuy := strategy.ShouldEnter(series.LastIndex(), record)
	//isSell := strategy.ShouldExit(series.LastIndex(), record)
	//fmt.Print(isBuy,isSell)
	//fmt.Println(cci.Calculate(series.LastIndex()-1))
	// fmt.Sprintf("%v",cci.Calculate(55).FormattedString(4))

}



var isExsit = techan.TimePeriod{}

func getHistory() {
	kline, _ := bs.GetKlineRecords("", goex.BTC_USDT, goex.KLINE_PERIOD_1MIN, 1500, 0)
	for _, dt := range kline {
		//fmt.Println(dt.Kline)
		period := techan.NewTimePeriod(time.Unix(dt.Kline.Timestamp, 0), time.Minute*1)

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
	isExsit = series.LastCandle().Period

	watch()
	//fmt.Println(series.LastIndex())

}

func watch() {

	bnWs.KlineCallback = func(kline *goex.Kline, p int) {

		period := techan.NewTimePeriod(time.Unix(kline.Timestamp, 0), time.Minute*1)
		//fmt.Print(period.String())
		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewDecimal(kline.Open)
		candle.ClosePrice = big.NewDecimal(kline.Close)
		candle.MaxPrice = big.NewDecimal(kline.High)
		candle.MinPrice = big.NewDecimal(kline.Low)
		candle.Volume = big.NewDecimal(kline.Vol)
		if isExsit == period {
			fmt.Printf("update:::%s,%s \n", isExsit, period)

			series.Candles[len(series.Candles)-1] = candle

		} else {
			fmt.Printf("new:::%s,%s \n", isExsit, period)
			isExsit = period
			series.AddCandle(candle)
		}
		stage1()

	}

	err := bnWs.SubscribeKline(goex.BTC_USDT, goex.KLINE_PERIOD_1MIN)
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
