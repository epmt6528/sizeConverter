package main

import (
	"fmt"
)

type Student struct {
	name string
}

func (s Student) calAvg(data []float64) (avgResult float64){
	sum := 0.0
	for i := 0; i < len(data); i++{
		sum += data[i]
	}
	avgResult = sum/float64(len(data))
	return
}

func (s Student) judge(avg float64) (judgeResult string){
	if avg>=60{
		judgeResult = "Passed"
	}else{
		judgeResult = "Failed"
	}
	return
}

func main(){
	a001 := Student{"Murai"}
	data := []float64{70, 8, 0, 69, 90}

	var avg float64 = a001.calAvg(data)
	result := a001.judge(avg)

	fmt.Println(avg)
	fmt.Println(a001.name + " " + result)
}