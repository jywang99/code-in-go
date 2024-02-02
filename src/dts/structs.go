package dts

import (
	"fmt"
	"strings"
)

type Creature struct {
	Species string
	Greeting string
}

func (c *Creature) Reset() {
    c.Species = ""
}

func TestMethods() {
    var creature Creature = Creature{Species: "shark"}
    fmt.Printf("1) %+v\n", creature)
    creature.Reset()
    fmt.Printf("2) %+v\n", creature)
}

/* Implementing an interface */

type Ocean struct {
    Creatures []string
}

// implements Stringer
func (o Ocean) String() string {
    return strings.Join(o.Creatures, ", ")
}

func log2(header string, s fmt.Stringer) { // log already imported elsewhere, avoid conflict
    fmt.Println(header, ":", s)
}

func TestStringer() {
    o := Ocean{
        Creatures: []string{
            "sea urchin",
            "lobster",
            "shark",
        },
    }
    log2("ocean contains", o)
}

/* Polymorphism */

type Submersible interface {
    Dive()
}

type Shark struct {
    Name string
    isUnderwater bool
}

func (s Shark) String() string {
    if s.isUnderwater {
        return fmt.Sprintf("%s is underwater", s.Name)
    }
    return fmt.Sprintf("%s is on the surface", s.Name)
}

func (s *Shark) Dive() {
    s.isUnderwater = true
}

func submerge(s Submersible) {
    s.Dive()
}

func TestPolymorphism() {
    s := &Shark{
        Name: "Sammy",
    }
    fmt.Println(s)
    submerge(s)
    fmt.Println(s)
}
