package employee

import "fmt"

// Struct Bình thường
type Employee struct {
	FirstName string
	LastName  string
	Age       int
	Salary    int
	Address Address
}

type Address struct {
	City string
	State string
}

// Struct ẩn danh
var employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func SamplePrint() {
	fmt.Println("print from person package")
}

//viết method trực tiếp trên các model để có các hành vi tương tự như class
func (t Employee) DisplayEmployeeName() {
	fmt.Println(t.FirstName, t.LastName)
}

//thay đổi giá trị trong model bằng con trỏ
func (e *Employee) ChangeFirstName(newName string) {  
    e.FirstName = newName
}