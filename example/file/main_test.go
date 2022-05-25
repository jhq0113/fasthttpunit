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

	confPath := fasthttpunit.BinPath() + "/case/case.yml"

	conf, err := fasthttpunit.LoadConf(confPath)
	if err != nil {
		t.Fatal(fasthttpunit.Red("load conf err: %s", err.Error()))
	}

	u := fasthttpunit.NewUnitWithRouter(conf, t, r)
	u.Test(mockA, mockB)
}
