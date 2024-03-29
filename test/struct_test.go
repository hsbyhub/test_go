package test

import (
	"fmt"
	"testing"
)

type Man struct {
	Name string
	Age  int
}

func testStructChange() {
	m := Man{
		Name: "cobe",
		Age:  42,
	}
	fmt.Println(m)
	changeStruct(&m)
	fmt.Println(m)
}

func changeStruct(m *Man) {
	*m = Man{
		Name: "james",
		Age:  36,
	}
}

func TestStruct(t *testing.T) {
	testStructChange()
}
