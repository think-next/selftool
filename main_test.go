package main

import (
	"fmt"
	"testing"
)

func TestBar(t *testing.T) {
	fmt.Println("testing")
	t.Error("err")
}
