package main_test

import ("testing"
		"./kadai")
func TestHex_String(t *testing.T) {
	expect := "a"
	actual := kadai.MyOct(10).String()
	if actual != expect {
		t.Errorf(`expect="%s" actual="%s"`, expect, actual)
	}
}