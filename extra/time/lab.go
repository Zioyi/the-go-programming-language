package main

import (
	"fmt"
	"time"
)

func getTimeOfDay() string {
	now := time.Now()
	hour := now.Hour()
	if hour >= 0 && hour < 6 {
		return "Night"
	} else if hour >= 6 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 18 {
		return "Afternoon"
	}
	return "Evening"
}

func getTimeOfDay2(curTime time.Time) string {
	hour := curTime.Hour()
	if hour >= 0 && hour < 6 {
		return "Night"
	} else if hour >= 6 && hour < 12 {
		return "Morning"
	} else if hour >= 12 && hour < 18 {
		return "Afternoon"
	}
	return "Evening"
}

func main() {

	base := time.Now()
	fmt.Println(base.Hour())
	fmt.Println(getTimeOfDay())
	//t3 := bas:wqe.Add(time.Hour * 1).Unix()
	//fmt.Println(t3)
	//
	//fmt.Println(base.Add(time.Hour * 24).Unix())
	//t := base
	//for i := 0; i < 100; i++ {
	//	t = t.Add(time.Hour * 1)
	//	fmt.Println(t.Unix(), t)
	//
	//}

}
