package main

import "fmt"

type Eraser interface {
	Erase(string)
}

type Writer interface {
	Write(string)
}

type EraseWriter interface {
	Eraser
	Writer
}

// Defining an object that can satisfy the Eraser interface
type AmazingEraser struct{}

func (AmazingEraser) Erase(str string) {
	fmt.Println("[Amazing Eraser] Erasing: ", str)
}

type NiceWriter struct{}

func (NiceWriter) Write(str string) {
	fmt.Println("[Pencil] writing: ", str)
}

type student struct {
	EraseWriter
}

func NewStudent(tool EraseWriter) student {
	return student{
		EraseWriter: tool,
	}
}

func (student) PerformWrite(writer Writer, str string) {
	writer.Write(str)
}

func (student) PerformErase(eraser Eraser, str string) {
	eraser.Erase(str)
}

func (s student) do(do EraseWriter, str string) {
	s.PerformWrite(do, str) // facade pattern
	fmt.Println("then...")
	s.PerformErase(do, str) // facade pattern
}

func (s student) Do(str string) {
	s.do(s.EraseWriter, str)
}

type Case struct {
	Writer
	Eraser
}

type WeirdEraser struct{}

func (WeirdEraser) Erase(str string) {
	fmt.Printf("[WierdEraser] Erasing \"%s\"...whoahhhhh the paper vanish!!", str)
}

type FlyingPen struct{}

func (FlyingPen) Write(str string) {
	fmt.Printf("[FlyingPen] Writing \"%s\"....whoahhhhhh I'm flying towards the universe\n", str)
}

func main() {

	newCase := Case{
		Writer: NiceWriter{},
		Eraser: AmazingEraser{},
	}

	pupil := NewStudent(newCase)

	pupil.Do("Interface is amazing")

	newWierdCase := Case{
		Writer: FlyingPen{},
		Eraser: WeirdEraser{},
	}

	pupil = NewStudent(newWierdCase)
	pupil.Do("Interface is weird but fun to use.")
}
