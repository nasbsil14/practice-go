package main

import (
	"fmt"
	"time"
)

var pStr string = "public"

func main() {
	fmt.Println("START")

	var sStr string = "private"
	sStr2 := "implicit type. private only!!"

	fmt.Printf("pStr:[%v] sStr:[%v] sStr2:[%v]", pStr, sStr, sStr2)

	if len("test") > 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	flg := true
	for flg {
		flg = false
	}
	//無限ループ
	//for {
	//}

	switch str := "test"; str {
	case "ttt":
		fmt.Println(str)
	case "sss":
		fmt.Println(str)
	default:
		fmt.Println(str)
	}
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	default:
		fmt.Println("Too far away.")
	}
	//type switch
	var iStr interface{} = sStr
	switch iStr.(type) {
	case string:
		fmt.Println("string type")
	case *string:
		fmt.Println("string pointer type")
	case int:
		fmt.Println("int type")
	case bool:
		fmt.Println("bool type")
	default:
		fmt.Println("any type")
	}

	fmt.Println("END")
}
