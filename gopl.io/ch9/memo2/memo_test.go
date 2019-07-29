package memo_test

import (
	"testing"

	memo "go-starter/gopl.io/ch9/memo2"
	"go-starter/gopl.io/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func TestSequential(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	memotest.Concurrent(t, m)
}
