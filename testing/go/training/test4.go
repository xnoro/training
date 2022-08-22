package main
 
import "fmt"
 
func main() {
    arr := [5]string{"apple", "banana", "cherry"}
 
    for i := 0; i < len(arr); i++ {
        fmt.Println(arr[i])
    }
}
