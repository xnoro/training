package main
 
import "fmt"
 
type student struct {
   name string
   rollNumber int
}
 
func main() {
    var student1 = student{name:"Anu", rollNumber:14}
    // get the member value
    fmt.Printf("student1 details\n--------------\n")
    fmt.Printf("Name : %s\n", student1.name)
    fmt.Printf("Rollnumber : %d\n", student1.rollNumber)
     
    var student2 = student{name:"Arjit", rollNumber:13}
    // set the member value
    student2.rollNumber=24
    student2.name = "Bungi"
     
    fmt.Printf("\nstudent2 details\n--------------\n")
    fmt.Printf("Name : %s\n", student2.name)
    fmt.Printf("Rollnumber : %d\n", student2.rollNumber)
}
