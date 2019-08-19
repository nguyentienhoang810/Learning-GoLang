package main

import "fmt"
import "HelloWorld/utility"
import "HelloWorld/packages/employee"

// type Employee struct {
// 	firstName string
// 	lastName string
// 	age int
// 	salary int
// 	Address
// }

// type Address struct {
// 	city, state string
// }

func calculate(price int, amount int) (int, int) {
	return price * amount, amount
}

func getBytesOfString(str string) {
	for i := 0; i < len(str); i++ {
		fmt.Printf("%x ", str[i])
	}
}

func getCharsOfString(str string) {
	//Một rune là một kiểu dữ liệu có sẵn trong Go, một cách gọi khác của kiểu dữ liệu int32. Rune đại diện cho một điểm mã trong Go. Không quan trọng điểm mã chứa bao nhiêu byte,
	runes := []rune(str)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func getCharsByRangeOfString(str string) {
	for index, rune := range str {
		fmt.Printf("index %d is %c ", index, rune)
		fmt.Printf("\n")
	}
}

func getStringFromByteSlice(b []byte) {
	str := string(b)
	fmt.Println(str)
}

func getStringFromRune(runes []rune) {
	str := string(runes)
	fmt.Println(str)
}


func main() {
	fmt.Println("Hi there, im go newbie !!")
	
	// variable
	var age int
	fmt.Println("first variable: ", age)
	age = 20
	fmt.Println("shit age: ", age)
	age = 30
	fmt.Println("now age: ", age)

	//function
	var calculationResult, amount = calculate(2000, 4)
	fmt.Println("calculation result: ", calculationResult, "by amount: ", amount)

	//function from package
	fmt.Println(utility.ShowJob())
	employee.SamplePrint()

	//conditions
	num := 10
	if num < 10 {
		fmt.Println("number is less than 10")
	} else {
		fmt.Println("number is greater than 10")
	}

	//loop
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue //pass this condition
		}
		fmt.Printf("value: %d. ", i)
	}

	fmt.Println(" ")
	j := 1
	for j <= 20 {
		fmt.Printf("%d. ", j)
		j += 2
	}

	//Array
	fmt.Println("array started. Array is value type")
	var arr [3]int
	fmt.Println(arr)
	arr[0] = 11
	arr[1] = 22
	arr[2] = 33
	fmt.Println(arr)
	arr2 := []int{12, 13, 14}
	fmt.Println(arr2)

	//Slice

	//map = Dictionary
	//map is reference type
	var sampleMap map[string]int //cách khai báo này chưa được cấp vùng nhớ => không thể thêm phần tử vào map -> phải cấp phát vùng nhớ bằng make như dưới.
	if sampleMap == nil {
		fmt.Println("map is nil")
		sampleMap = make(map[string]int)
	}
	sampleMap["value1"] = 200
	fmt.Println(sampleMap)

	personSalary := make(map[string]int)
	personSalary["one"] = 12000
	personSalary["two"] = 3000
	if personSalary == nil {
		fmt.Println("map is nil")
	} else {
		fmt.Println(personSalary)
	}

	sampleMap2 := map[string]int{ //khởi tạo kèm giá trị
		"steve": 12000,
		"jamie": 15000,
	}
	sampleMap2["mike"] = 9000
	fmt.Println("sampleMap2 map contents:", sampleMap2)

	person1 := "steve"
	fmt.Println("value of", person1, "is: ", sampleMap2["steve"])

	person2 := "job" //không tồn tại cặp key job -> lấy giá trị mặc định là 0
	fmt.Println("value of", person2, "is: ", sampleMap2["job"])

	value, exist := sampleMap2[person2] //kiểm tra sự tồn tại của key-value trong map
	fmt.Println(value, exist)

	for key, value := range sampleMap2 { //duyệt toàn bộ phần tử trong map bằng range
		fmt.Println(key, value)
	}

	//remove phần tử trong map
	delete(sampleMap2, "mike")
	fmt.Println("after delete mike")
	for key, value := range sampleMap2 {
		println(key, value)
	}

	//sample of reference type of map
	//khi một map được gán cho một biến mới, cả hai đều trỏ đến cùng một cấu trúc dữ liệu nội bộ. Do đó những thay đổi được thực hiện ở một biến sẽ ánh xạ đến biến kia.
	fmt.Println("original sampleMap2: ", sampleMap2)
	newReferenceMap := sampleMap2
	newReferenceMap["steve"] = 9999
	fmt.Println("original sampleMap2 has been changed: ", sampleMap2)

	//STRING
	name := "hello world"
	fmt.Println(name)
	getBytesOfString(name) //print out the hexadecimal of string
	fmt.Printf("\n")
	getCharsOfString(name)
	fmt.Printf("\n")
	fmt.Printf("\n")

	name = "Señor"
	getBytesOfString(name) //print out the hexadecimal of string
	fmt.Printf("\n")
	getCharsOfString(name) //result = S e Ã ± o r :vì Unicode của ký tự ñ là U+00F1 => sửa bằng rune
	fmt.Printf("\n")
	getCharsByRangeOfString(name)
	fmt.Printf("\n")

	utf8byteSlice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	fmt.Println("string from utf8 bytes slice")
	getStringFromByteSlice(utf8byteSlice)
	fmt.Printf("\n")

	decimaBytesSlice := []byte{67, 97, 102, 195, 169}
	fmt.Println("string from decima bytes slice")
	getStringFromByteSlice(decimaBytesSlice)
	fmt.Printf("\n")

	runes := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	fmt.Println("string from runes")
	getStringFromRune(runes)
	fmt.Printf("\n")

	//Kiểu dữ liệu của string là bất biến. Không thể thay đổi trực tiếp
	strDemo := "Gemo string"
	// strDemo[0] = "H" -> immutable
	//-> muốn thay đổi chuyển string thành slice chứa các rune, chúng ta thay đổi slice và sau đó chuyển lại thành string mới.
	runedString := []rune(strDemo)
	runedString[0] = 'D'
	convertedString := string(runedString)
	fmt.Println(convertedString)
	fmt.Printf("\n")
	fmt.Printf("\n")

	//POINTER

	//& được sử dụng để lấy địa chỉ của một biến.
	fmt.Println("POINTER")
	fmt.Printf("\n")
	b := 1989
	var a = &b
	fmt.Println("address of b is: ", &b)
    fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of pointer a is", a)
	fmt.Println("reference value of pointer is", *a)
	//nil pointer
	m := 25
    var n *int
    if n == nil {
        fmt.Println("n is", n)
        n = &m
		fmt.Println("n after initialization is", n)
		fmt.Println("referenced value", *n)
		*n++
		fmt.Println("changed value at pointer:", m)
	}
	
	//STRUCT
	fmt.Println("STRUCT")
	fmt.Println("\n")
	//create struct
	
	emp1 := employee.Employee {
		FirstName : "first name",
		LastName: "last name",
		Age: 20,
		Salary: 2000,
	}

	//viết method trực tiếp trên các model để có các hành vi tương tự như class
	fmt.Println("display emp1 name")
	emp1.DisplayEmployeeName()
	emp1.ChangeFirstName("new emp1 first name")
	emp1.DisplayEmployeeName()
	

	emp1.Address = employee.Address {"sample city", "sample state"}
	fmt.Println(emp1)
	fmt.Println(emp1.Address.City, emp1.Address.State)
	fmt.Println("emp1 name:", emp1.FirstName, emp1.LastName)

	emp2 := employee.Employee {"emp2 first" , "emp2 last", 22, 3000, employee.Address {"next city", "next State"}}
	fmt.Println(emp2)

	var emp3 employee.Employee
	emp3.FirstName = "emp3 first"
	emp3.LastName = "emp3 last"
	fmt.Println(emp3)

	//pointer to struct
	empPointer := &emp1
	fmt.Println("pointer name value: ", (*empPointer).FirstName)
	fmt.Println("OR")
	fmt.Println("pointer name value: ", (empPointer).FirstName)
}
