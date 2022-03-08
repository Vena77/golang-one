package main

import (
	"testing"
)

func SetDataTest(b *testing.B) {
	resp := In{}
	SetData(&resp)
}
