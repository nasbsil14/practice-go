package main

import "fmt"

type MySuccess struct {
	msg string
}
func (s *MySuccess) String() string {
	return s.msg
}

type MyError struct {
	code string
	msg string
}
//errorIF override method
func (e *MyError) Error() string {
	return fmt.Sprintf("[error]code:%v msg:%v", e.code, e.msg)
}

func run(isErr bool) (*MySuccess, error) {
	if isErr {
		//MyError implements errorIF
		return nil, &MyError{"0001", "my error!!"}
	} else {
		return &MySuccess{"OK"}, nil
	}
}

func main() {
	if result, err := run(false); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
