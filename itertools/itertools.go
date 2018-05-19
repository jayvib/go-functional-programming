package itertools



type Iter chan interface{}

func New(els ...interface{}) Iter {
	outstream := make(chan interface{})

	go func() {
		defer close(outstream)

		for _, v := range els {
			outstream <- v
		}
	}()

	return outstream
}

