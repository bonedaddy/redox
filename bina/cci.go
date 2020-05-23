package main

import (
	"fmt"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
	"time"

	"github.com/gocarina/gocsv"
	//"github.com/sdcoffey/big"
	//"github.com/sdcoffey/techan"
	//"github.com/stretchr/testify/assert"
	"os"

	//"time"
	//"github.com/tobgu/qframe/types"

	//"strings"
)
type Ohlc struct { // Our example struct, you can use "-" to ignore a field
	Date      string `csv:"Date"`
	Open      string `csv:"Open"`
	Close    string `csv:"Close"`
	Height     string `csv:"Height"`
	Low     string `csv:"Low"`
	NotUsed string `csv:"-"`
}
func main()  {


	clientsFile, err := os.OpenFile("in/ccc.csv", os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer clientsFile.Close()

	clients := []*Ohlc{}

	if err := gocsv.UnmarshalFile(clientsFile, &clients); err != nil { // Load clients from file
		panic(err)
	}

	 ts := techan.NewTimeSeries()

	var candleIndex int



	for _, val := range clients {
		candle := techan.NewCandle(techan.NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))

	//	Date,Open,High,Low,Close

		candle.OpenPrice = big.NewFromString(val.Open)
		candle.ClosePrice = big.NewFromString(val.Close)
		candle.MaxPrice = big.NewFromString(val.Height)
		candle.MinPrice = big.NewFromString(val.Low)
		//candle.Volume = big.NewFromString(val)
		//fmt.Println(candle.String())
		ts.AddCandle(candle)

		candleIndex++
	}


	//
	 cci := techan.NewCCIIndicator(ts, 14)
	 fmt.Println(cci.Calculate(0))

	//results := []string{"101.9185", "31.1946", "6.5578", "33.6078", "34.9686", "13.6027",
	//	"-10.6789", "-11.4710", "-29.2567", "-128.6000", "-72.7273"}
	//

	for i := 0; i < len(clients); i++ {


		//fmt.Println(v)
		fmt.Println(cci.Calculate(i).FormattedString(4))



	}




	values := []string{
		"23.98", "23.92", "23.79", "23.67", "23.54",
		"23.36", "23.65", "23.72", "24.16", "23.91",
		"23.81", "23.92", "23.74", "24.68", "24.94",
		"24.93", "25.10", "25.12", "25.20", "25.06",
		"24.50", "24.31", "24.57", "24.62", "24.49",
		"24.37", "24.41", "24.35", "23.75", "24.09",
	}

	ts1 := techan.NewTimeSeries()
	for _, val := range values {
		candle := techan.NewCandle(techan.NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))
		candle.OpenPrice = big.NewFromString(val)
		candle.ClosePrice = big.NewFromString(val)
		candle.MaxPrice = big.NewFromString(val)
		candle.MinPrice = big.NewFromString(val)
		candle.Volume = big.NewFromString(val)

		ts.AddCandle(candle)

		candleIndex++
	}



	cci2 := techan.NewCCIIndicator(ts1, 14)

	fmt.Print(cci2.Calculate(14))




}