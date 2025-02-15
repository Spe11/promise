### Example

```
package main

import (
	"errors"
	"fmt"
	"time"
	"github.com/Spe11/promise"
)

func main() {
	promise1 := success()
	promise1.OnSuccess(func(data int) {
		fmt.Println(data)
	}).OnError(func (err error) {
		fmt.Println(err.Error())
	})

	promise2 := fail()
	promise2.OnSuccess(func(data string) {
		fmt.Println(data)
	}).OnError(func (err error) {
		fmt.Println(err.Error())
	})
}

func success() promise.Promise[int] {
	promise := promise.CreatePromise[int]()
	go func () {
		time.Sleep(time.Second)
		promise.Resolve(123)
	}()

	return promise
}

func fail() promise.Promise[string] {
	promise := promise.CreatePromise[string]()
	go func ()  {
		time.Sleep(time.Second)
		promise.Reject(errors.New("error"))
	}()

	return promise
}
```
