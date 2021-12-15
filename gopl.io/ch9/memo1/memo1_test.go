package memo1_test

import (
	"gopl.io/ch9/memo1"
	memo_test "gopl.io/ch9/memotest"
	"testing"
)

func Test(t *testing.T) {
	m := memo1.New(memo_test.HTTPGetBody)
	memo_test.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo1.New(memo_test.HTTPGetBody)
	memo_test.Concurrent(t, m)
}
