package main

import "fmt"

// Board is the concrete type where all the interface will interact to.
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// ======================================
// these are the set of behavior that is in the
// tool that I will use later.

type NailDriver interface {
	DriveNail(nailSupply *int, b *Board)
}

type NailPuller interface {
	PullNail(nailSupply *int, b *Board)
}

// NailDriverPuller is the behavior of the new tool
// that I will be using and it's composed by the
// NailDriver and NailPuller into it.
type NailDrivePuller interface {
	NailDriver
	NailPuller
}

// =======================================
// These are the actual tools objects that
// implements the behavior that I need for
// the list of interfaces defined above.

type Mallet struct{}

func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.NailsDriven++
	fmt.Println("Mallet: pounded nail into the board.")
}

type Crowbar struct{}

func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

// ====================================
// These is the object that will be using the tool object
// defined above.

type Contractor struct{}

func (Contractor) Fasten(d NailDriver, nailSupply *int, b *Board) { // form of Facade Pattern
	for b.NailsDriven < *nailSupply {
		d.DriveNail(nailSupply, b)
	}
}

func (Contractor) Unfasten(p NailPuller, nailSupply *int, b *Board) { // form of Facade Pattern
	for b.NailsDriven > b.NailsNeeded {
		p.PullNail(nailSupply, b)
	}
}

func (c Contractor) ProcessBoards(dp NailDrivePuller, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(dp, nailSupply, b)
		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(dp, nailSupply, b)
		}
	}
}

type Toolbox struct {
	NailDriver
	NailPuller

	nails int
}

func displayState(tb *Toolbox, boards []Board) {
	fmt.Printf("Box: %#v\n", tb)
	fmt.Println("Boards:")

	for _, b := range boards {
		fmt.Printf("\t%+v\n", b)
	}

	fmt.Println()
}

func main() {
	boards := []Board{
		{NailsDriven: 3},
		{NailsDriven: 1},
		{NailsDriven: 6},

		{NailsNeeded: 6},
		{NailsNeeded: 9},
		{NailsNeeded: 4},
	}

	tb := Toolbox{
		NailDriver: Mallet{},
		NailPuller: Crowbar{},
		nails:      10,
	}

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)

	displayState(&tb, boards)
}
