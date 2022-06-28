package convert_test

import (
	. "fehu/util/convert"
	"log"
	"testing"
)

func Test(t *testing.T) {

	var s string
	TypeChange(&s, 1)
	if s != "1" {
		t.Error("err")
	}
	log.Println("success", s)

	TypeChange(&s, []byte("12d"))
	if s != "12d" {
		t.Error("err")
	}
	log.Println("success", s)

	TypeChange(&s, []string{"1", "2"}, "&")
	if s != "1&2" {
		t.Error("err")
	}
	log.Println("success", s)

	var i int
	TypeChange(&i, []byte("123154"))
	if i != 123154 {
		t.Error("err")
	}
	log.Println("success", i)

}
