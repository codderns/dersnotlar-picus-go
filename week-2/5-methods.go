package main

import "fmt"

/*Method, struct'a bağlı olan bir fonksiyondur. Sadece struct için kullanılır.*/

//import "fmt"

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// 	TestF     func() bool
// }

// //aşağıdaki kuralda yazılır
// 						 //döndürdüğü tür
// //  (structname) functionName() string
// func (p Person) String() string {
// 	return fmt.Sprintf("%s %s, age %d", p.FirstName, p.LastName, p.Age)
// 				//Sprintf formatlanmış string
// }

// func (p Person) ChangeAge() int {
// 	return p.Age
// }

// func main() {
// 	person := Person{
// 		FirstName: "Mustafa",
// 		LastName:  "Kemal",
// 		Age:       30,
// 	}
// 	fmt.Println(person.String()) // structadi.fonksiyonu
// 	fmt.Println(person.ChangeAge())

// 	// Bu başka bir olay. Struct içinde bir func var. Bu struct'a ait bir mehod gibi.
// 	person.TestF = func() bool {
// 		return false
// 	}
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// type Adder struct {
// 	start int
// }

// //↓ bu struct'a bağlı birşey eklemek değiştirmek istiyoruz fonksiyonlar ↓

// //a-> Adder'a çalıştırmak için bir değişken adı vermek gerekli
// func (a Adder) AddTo(val int) int {
// 	return a.start + val
// }

// func main() {
// 	myAdder := Adder{start: 10} // bir struct oluşturduk Adder'dan
// 	fmt.Println(myAdder.AddTo(5))
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

type Employee struct {
	Name string
	ID   string
}

func (e Employee) Description() string {
	return fmt.Sprintf("%s (%s)", e.Name, e.ID)
}

type Manager struct {
	Employee            // Employee struct'ını buraya al (içine gömdük)
	Reports  []Employee // manager altında çalışanların rapor dizisi
}

// Manager'ın kendi methodu var burada
func (m Manager) FindNewEmployees() []Employee {
	// do business logic
	return []Employee{}
}

func main() {
	m := Manager{
		Employee: Employee{ // Burada bir ayrı struct var. 83. satırdaki struct yani.
			//Employee bilgisi yani ayrıca.
			Name: "Bob Bobson", // manager hakkında bilgi
			ID:   "12345",
		},
		Reports: []Employee{
			{Name: "Ahmet", ID: "123"}, // manager altında çalışanların bilgisi
			{Name: "Fatma", ID: "9659"},
		},
	}
	fmt.Println(m.ID)            //12345
	fmt.Println(m.Description()) //manager'ın description diye bir methodu aslında yok. Manager'ın
	//description metodu aslında Employee içindir. Manager aynı zamanda Employee'ı embed edip içine
	// aldığı için Employee'ın func metoduna da erişim sağlayabiliyor. // Bob Bobson (12345)
	fmt.Println(m.Employee) // {Bob Bobson 12345}
	fmt.Println(m.Reports)  //[{Ahmet 123} {Fatma 9659}]
}

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// type Inner struct {
// 	ID int
// }

// type Outer struct {
// 	Inner
// 	ID int
// }

// func main() {
// 	inner := Inner{
// 		ID: 10,
// 	}

// 	o := Outer{
// 		Inner: inner,
// 		ID:    20,
// 	}

// 	fmt.Println(o.ID)
// 	fmt.Println(o.Inner.ID) // GO hangi ID'den bahsedildiğini anlaması için böyle belirtilmesi zorunlu
// }
