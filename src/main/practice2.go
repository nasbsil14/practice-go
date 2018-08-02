package main
import "fmt"

func main() {
	fmt.Println("START")

	str1 := "test"
	str2 := str1
	//byRef
	str3 := &str1

	fmt.Printf("str1=%v\n", str1)
	fmt.Printf("str2=%v\n", str2)
	//fmt.Printf("str3=%v\n", str3)
	fmt.Printf("*str3=%v\n", *str3)
	str1 = "---"
	fmt.Printf("str1=%v\n", str1)
	fmt.Printf("str2=%v\n", str2)
	fmt.Printf("str3=%v\n", str3)
	fmt.Printf("*str3=%v\n", *str3)

	//ary
	strs := []string{"1", "2", "3"}
	for s := range strs {
		fmt.Println(s)
	}
	//sprit
	fmt.Println(fmt.Sprintf("No.%v No.%v No.%v", strs[0], strs[1], strs[2]))

	var sliceBlank []int = make([]int, 0)
	sliceBlank[0] = 0
	sliceBlank[1] = 1
	sliceBlank[2] = 2

	mStr := map[string]string {
		"key1":"val1",
		"key2":"val2",
		"key3":"val3",
	}
	fmt.Println(mStr)
	var mBlank map[string]int = make(map[string]int)
	mBlank["A"] = 0

	fmt.Println("END")
}
