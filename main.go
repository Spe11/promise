package promise

type result[T any] struct {
	Data T
	Err error
}
type Promise[T any] struct {
	channel chan result[T]
	result result[T]
	recevied bool
}

func CreatePromise[T any]() Promise[T] {
	return Promise[T]{channel: make(chan result[T])}
}

func (p *Promise[T]) Resolve(data T) {
	p.channel <- result[T]{Data: data}
}
func (p *Promise[T]) Reject(err error) {
	p.channel <- result[T]{Err: err}
}
func (p *Promise[T]) OnSuccess(callback func(data T)) *Promise[T] {
	p.getResult()
	if (p.result.Err == nil) {
		callback(p.result.Data)
	}

	return p
}
func (p *Promise[T]) OnError(callback func(err error)) *Promise[T] {
	p.getResult()
	if (p.result.Err != nil) {
		callback(p.result.Err)
	}

	return p
}
func (p *Promise[T]) getResult() {
	if (!p.recevied) {
		p.result = <-p.channel
		p.recevied = true
	}
}
