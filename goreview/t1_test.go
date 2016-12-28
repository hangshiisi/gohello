package main

import (
	//"fmt"
	"testing"
)

func TestNo1(t *testing.T) {
	t.Log("Get into testing function 1\n")
}

func TestNo2(t *testing.T) {
	t.Log("Get into testing function 2\n")
	t.Fail()
}
