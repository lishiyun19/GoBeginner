package main
import "fmt"

//结构体
type person struct {
    name string
    age int
}

func main() {
    //结构体初始化
    fmt.Println(person{name:"bob", age:20})
    fmt.Println(person{"bob", 20})
    fmt.Println(person{name:"fred"})
    s := person{name:"Ann", age:40}

	//sp 是结构体指针
    sp := &s
    fmt.Println(sp.age)
}
