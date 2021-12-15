// Package memo1 提供了一个对类型 Func 并发不安全的函数记忆功能
package memo1

// Memo 缓存了调用 Func 的结果
type Memo struct {
	f     Func
	cache map[string]result
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

// Get 注意： 非并发安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}
