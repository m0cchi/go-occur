package main

import (
	"fmt"
	"github.com/m0cchi/go-occur/model"
)

var S model.I

func init() {
	S = SuccessPlugin{}
}

type SuccessPlugin struct {
}

func (p SuccessPlugin) M() {
	fmt.Println("success plugin")
}
