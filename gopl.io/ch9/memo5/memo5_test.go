package memo5_test

import (
	"gopl.io/ch9/memo5"
	memo_test "gopl.io/ch9/memotest"
	"testing"
)

func Test(t *testing.T) {
	m := memo5.New(memo_test.HTTPGetBody)
	memo_test.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo5.New(memo_test.HTTPGetBody)
	memo_test.Concurrent(t, m)
}
