package domain

import "fmt"
import "../until"

type Chinese struct {
    num     int
    Name    string
	Gender  string
	Age     uint8
    Address string
}

type Beijings struct {
    num      int
    brithday string
    Chinese
}


func (p *Chinese) Move(add string) (string){
    var old_address string
    old_address = p.Address
    p.Address = add
    return old_address
    }

func Cheinese_people(){
	p := Chinese{"Robert", "Male", 33, "Beijing"}
    fmt.Printf("i am Chinese people, my name is %s\n", p.Name)
	oldAddress := p.Move("ShangHai")
	fmt.Printf("%s moved from %s to %s.\n", p.Name, oldAddress, p.Address)
    now_time := until.Get_ct_time()
    fmt.Printf("现在时刻:%s",now_time)
}

func Get_peoples(num int){
    for i := 0; num; i++ {
        now_time := until.Get_ct_time()
        p := Beijings{1000+1, now_time, Chinese{"Robert", "Male", 33, "Beijing"}}
}