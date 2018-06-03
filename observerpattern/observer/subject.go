package observer

type Subject struct {
	callback []Callback
}

func (o *Subject) AddObserver(c Callback) {
	o.callback = append(o.callback, c)
}

func (o *Subject) DeleteObserver(c Callback) {
	o.callback = append(o.callback, c)
	newCallbacks := make([]Callback, 0)
	for _, cp := range o.callback {
		if cp != c {
			newCallbacks = append(newCallbacks, cp)
		}
	}
	o.callback = newCallbacks
}

func (o *Subject) NotifyObservers(oes ...Observable) {
	for _, oe := range oes {
		for _, c := range o.callback {
			c.Notify(&oe)
		}
	}
}
