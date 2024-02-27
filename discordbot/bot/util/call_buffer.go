package util

func FunctionCallBuffer() (ch chan func()) {
	ch = make(chan func())

	go func() {
		// Remember the functions as they are submitted
		funcs := []func(){}
		for fn := range ch {
			funcs = append(funcs, fn)
		}
		// Once the channel is closed, call all the funcs
		for _, fn := range funcs {
			fn()
		}

	}()
	return
}
