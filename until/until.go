package until

import "fmt"
import "time"

// var (
//     a = b + c
//     b = f()
//     c = f()
//     d = 3
// )

// func f() int {
//     d++
//     return d
// }

// func main() {
//     fmt.Println(a, b, c, d)
// }

func Get_ct_time() (string){
	now := time.Now()
	year,mon,day := now.UTC().Date()
	hour,min,sec := now.UTC().Clock()
	zone,_ := now.UTC().Zone()
	now_time := fmt.Sprintf("%d-%d-%d %02d:%02d:%02d %s", year, mon, day, hour, min, sec, zone)
	fmt.Println("UTC time is ",now_time) 
	return now_time
}

    
    
    
    



