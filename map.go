package main
import "fmt"
func main(){
map1 :=map[int]string{
0:"PPL",
1:"gogo",
2:"java",
3:"c",
4:"Python",
}
fmt.Println(len(map1))
for _,val:=range map1{
fmt.Println(val)
}
delete(map1,1)
fmt.Println("after deleting")
fmt.Println("deleted")
fmt.Println(map1)
fmt.Println("\nChecking for 1st element")
if _,exists:=map1[1];exists{
fmt.Println("value exists")
fmt.Println(map1)
}else{
fmt.Println("value NULL")
fmt.Println(map1)
}
}
