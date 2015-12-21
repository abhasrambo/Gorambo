package main
import "fmt"
type info struct{
id string
name string
sal float64
age int
}
func main(){
c1:=new(info)
c1.name="abhas kumar"
c1.id="m1234"
c1.sal=15000.00
c1.age=30
fmt.Println(c1.name)
fmt.Println(*c1)
fmt.Println("before")
c1.name="Martian"
fmt.Println("after")
fmt.Println(c1.name)
fmt.Println(*c1)
}
