package main

import (
	"errors"
	"fmt"
)

func test() error {
	return nil
}

func foo() (int, error) {
	return 0, nil
}

func main() {

	//errors.Is()

	var x, err = foo()
	if err != nil {
		// to do sth...
	}

	_ = x

	//var err = test()
	//switch err := err.(type) {
	//case nil:
	//	// call success, noting to do
	//case *MyError:
	//	fmt.Println("error occurred on line:", err.Line)
	//default:
	//	// unknown error
	//}

	//var reader = bufio.NewReader(os.Stdin)
	//_, _, err := reader.ReadLine()
	//if err != io.EOF {
	//	// to do sth...
	//}

	//if ErrNamedType == New("EOF") {
	//	fmt.Println("named type error")
	//}
	//
	//if ErrMyStructType == NewMyStructError("EOF") {
	//	fmt.Println("struct type error")
	//}
	//
	//if ErrStructType == errors.New("EOF") {
	//	fmt.Println("struct type error")
	//}
}

type MyError struct {
	Msg  string
	File string
	Line int
}

func (m *MyError) Error() string {
	return fmt.Sprintf("%s:%d %s", m.File, m.Line, m.Msg)
}

var ErrNamedType = New("EOF")
var ErrMyStructType = NewMyStructError("EOF")
var ErrStructType = errors.New("EOF")

func New(e string) error {
	return errorString(e)
}

type errorString string

func (e errorString) Error() string {
	return string(e)
}

func NewMyStructError(e string) error {
	return myStructError{
		s: e,
	}
}

type myStructError struct {
	s string
}

func (e myStructError) Error() string {
	return e.s
}
