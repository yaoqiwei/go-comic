package cache_test

import (
	. "fehu/util/cache"
	"fmt"
	"testing"
	"time"
)

var i int = 1

func tr(e int) {

	i2, _ := GetCache("key1").(*int)
	if i2 == nil {
		SetCache("key1", &i, 100*time.Millisecond)
		fmt.Println(e, ":", i, 'b')
	} else {
		fmt.Println(e, ":", *i2)
	}

}

type testStr struct {
	A int
}

type testStrList []*testStr

func Test(t *testing.T) {

	for i := 0; i < 100; i++ {
		go tr(i)
	}

}

func Test2(t *testing.T) {

	// var v testStrList
	// SetCache("gift", v, 5*time.Second)

	inter := GetCache("gift")
	gift, ok := inter.(testStrList)
	if ok && gift != nil {
		t.Errorf("err")
	}

}
