// Package memo5 提供了一个对类型 Func 并发不安全的函数记忆功能
package memo5

type entry struct {
	res   result
	ready chan struct{} // res 准备好会被关闭
}

// Memo 缓存了调用 Func 的结果
type Memo struct {
	requests chan request
}

type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// request是一条请求消息，key需要用Func来调用
type request struct {
	key      string
	response chan<- result // 客户端需要单个result
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}

// Get 注意： 非并发安全
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}
func (memo *Memo) Close() {
	close(memo.requests)
}

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			// 对这个 key 的第一次请求
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key) // 调用f(key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	// 执行函数
	e.res.value, e.res.err = f(key)
	// 通知数据已准备完毕
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	// 等待数据准备完毕
	<-e.ready
	// 向客户端发送结果
	response <- e.res
}
