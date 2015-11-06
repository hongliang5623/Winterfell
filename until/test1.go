package until

import (
    "fmt"
)
 
type Doaction interface {
    sayHi_i()
    eatFood_i()
    sing_i(lyrics string)
}

type People struct {
    Name   string
    Age    int
    Weight int
}

func (p People) SayHi_p() {
    fmt.Println("People",p.Name)
}
 
func (p People) EatFood_p() {
    fmt.Println("People",2)
}
 
func (p People) Sing_p(lyrics string) {
    fmt.Printf("People,sing", lyrics)
}
 
 
type Student struct {
    // until.People
    // specialty string
}

func (p Student) sayHi_i() {
    fmt.Println("Student")
}
 
func (p Student) eatFood_i() {
    fmt.Println("Student",2)
}
 
func (p Student) sing_i(lyrics string) {
    fmt.Printf("Student,sing", lyrics)
}
