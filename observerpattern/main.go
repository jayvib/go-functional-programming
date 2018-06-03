package main

import "github.com/go-functional-programming/observerpattern/observer"

func main() {
	subject := observer.Subject{}

	oa := observer.Observable{Name: "A"}
	ob := observer.Observable{Name: "B"}

	subject.AddObserver(&observer.Observer{})
	subject.NotifyObservers(oa, ob)

	oc := observer.Observable{Name: "C"}
	subject.NotifyObservers(oa, ob, oc)

	subject.DeleteObserver(&observer.Observer{})
	subject.NotifyObservers(oa, ob, oc)

	od := observer.Observable{Name: "D"}
	subject.NotifyObservers(oa, ob, oc, od)
}
