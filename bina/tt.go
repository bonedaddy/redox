package main

import (
	"fmt"
	"github.com/sdcoffey/techan"
)

func main()  {
	upInd := techan.NewConstantIndicator(5)
	dnInd := techan.NewFixedIndicator(1,2,3,4,5,6,7,9)
	//fmt.Println(dnInd.Calculate(0).Cmp(upInd.Calculate(0)))


	rule := techan.NewCrossUpIndicatorRule(upInd, dnInd)
	for i:=0;i<8;i++{
		 a:=rule.IsSatisfied(i, nil)
		//fmt.Println(dnInd.Calculate(i).Cmp(upInd.Calculate(i)))

		 fmt.Println(a)
	}

}