package main

import (
	"fmt"
	"os"
	"testing"

	"github.com/jhq0113/fasthttpunit"
)

func mockA() {
	fmt.Printf("mock:a\n")
}

func mockB() {
	fmt.Printf("mock:b\n")
}

func TestUnit(t *testing.T) {
	fmt.Println(os.Args)

	r := loadRouter()

	casePath := fasthttpunit.BinPath() + "/case"

	conf, err := fasthttpunit.LoadConfByPath(casePath)
	if err != nil {
		t.Fatal(fasthttpunit.Red("load conf err: %s", err.Error()))
	}

	conf.Delay = 3

	u := fasthttpunit.NewUnitWithRouter(conf, t, r)
	u.Test(mockA, mockB)
}
