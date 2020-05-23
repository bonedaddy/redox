package main

import (
	"fmt"
	//"github.com/sdcoffey/big"
	//"github.com/sdcoffey/techan"
	//"github.com/stretchr/testify/assert"
	"os"
	"github.com/tobgu/qframe"
	//"time"
	//"github.com/tobgu/qframe/types"

	//"strings"
)

func main()  {
	f, _ := os.Open("btc.csv")
	fmt.Print(f)

//	input := `COL1,COL2
//a,1.5
//b,2.25
//c,3.0`

	//f1 := qframe.ReadCSV(strings.NewReader(input))
	qf := qframe.ReadCSV(f)
	 	fmt.Println(qf.Select("Date").StringView("Date"))



	//ts := techan.NewTimeSeries()

	//named := qf.ColumnTypeMap()

	//fmt.Println(named)
	//for _, col := range qf.() {
	//	if named[col] == types.Int {
	//		view := qf.MustIntView(col)
	//		for i := 0; i < view.Len(); i++ {
	//			fmt.Println(view.ItemAt(i))
	//		}
	//	}
	//}
	//for _, val := range values {
	//	candle := NewCandle(NewTimePeriod(time.Unix(int64(candleIndex), 0), time.Second))
	//	candle.OpenPrice = big.NewFromString(val)
	//	candle.ClosePrice = big.NewFromString(val)
	//	candle.MaxPrice = big.NewFromString(val)
	//	candle.MinPrice = big.NewFromString(val)
	//	candle.Volume = big.NewFromString(val)
	//
	//	ts.AddCandle(candle)
	//
	//	candleIndex++
	//}
	//
	//cci := techan.NewCCIIndicator(series, 20)
	//
	//results := []string{"101.9185", "31.1946", "6.5578", "33.6078", "34.9686", "13.6027",
	//	"-10.6789", "-11.4710", "-29.2567", "-128.6000", "-72.7273"}
	//
	//for i, result := range results {
	//	assert.EqualValues(t, result, cci.Calculate(i+19).FormattedString(4))
	//}
}