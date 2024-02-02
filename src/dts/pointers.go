package dts

import "fmt"

func PrintPointers() {
	var creature string = "shark"
	var pointer *string = &creature
	fmt.Println("creature =", creature)
	fmt.Println("pointer =", pointer)   // print pointer address
	fmt.Println("*pointer =", *pointer) // print stored value at the address

	// modify through pointer
	*pointer = "jellyfish" // changes value of 'creature' as well
	fmt.Println("*pointer =", *pointer)
	fmt.Println("creature =", creature)
}

func TestChangeStructFn() {
	// var creature Creature = Creature{Species: "shark"} // also valid
    var creature *Creature
    creature = &Creature{Species: "shark"}
	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature) // dont forget & if creature is the struct, not pointer
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
	if creature == nil {
		fmt.Println("creature is nil")
		return
	}
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
