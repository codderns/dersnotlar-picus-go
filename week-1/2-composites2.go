package main

import "fmt"

func main(){
	
	// Struct
// Birbirleri ile alakalı field'ları tek bir grup altında toplamaktır.

// Bir struct tanımlamak istersek bu type ile başlar, sonra struct diye belirtilir.


	// type person struct {
	// 	name string  // name, age ve pet , person struct'ının field'larıdır.
	// 	age  int
	// 	pet  string
	// }
	
	//var ahmet person // => ahmet değişkeni initial olarak bir type'a sahip ve field'ları name
	//age ve pet olan bir struct ile başlatıldı

	// // person struct'ındaki tüm field'lar doldurulmak zorunluluğu yoktur.
	// osman := person{
	// 	"Osman",
	// 	21,
	// }
	//ayse := person{} // ayse artık bir person ve bunla uğraşabiliriz. yukarıdaki(ahmet) ile aynı amaç

	//
	//
	//osman := person{ 
	//	name: "Osman",
	//	age:  21,
	//	pet:  "Zeytin",
	//}
	//
	//osman.age = 22
	//osman.pet = "turuncu"
	//


// anonymous structs //
//  => Sadece bulunduğu struct içinde kullanılabilen yapılardır.

	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//
	//var person struct {
	//	name string
	//	age  int
	//	pet  string
	//}

	//person.name = "Osman" //person field'larına value verildi
	//person.age = 21
	//person.name = "Zeytin"
	//
	//pet := struct { 		// pet field'ına anonymous struct açıldı
	//	name string
	//	kind string
	//}{
	//	name: "Zeytin", 	// pet struct'ına değer verildi.
	//	kind: "Cat",
	//}
	//

	//
	//type firstPerson struct{
	//	name string
	//	age int
	//}
	//f = firstPerson{
	//	name: "Osman",
	//	age:  21,
	//}
	//

	//type person struct {
	//	name string
	//	age  int
	//	pet  string
	//}
	//var g struct {
	//	name string
	//	age  int
	//}
	//g.age = 21
	//g.name = "test"
	//g = f
	//
	//
	//fmt.Println(f == g)
	//

/// Shadowing Variables ///
// => Var olan değişkeni gölgelemek bir alan dışındaki değişken ve bu alanın içindeki aynı isimde 
// => bir değişken girmeye denir. Bu durumda aşağıdaki örnekteki gibi, x'e 10 deklare edilmiş. Sonra
// => bir koşul sonucunda print edildikten sonra aynı isimde bir başka değişken o alan içinde deklare 
// => edilmiş := komutu ile. Koşul alanı dışına çıkınca print edince x'i ismi ilk x belirlenen yani 10
// => görünecektir. Ancak x:=5 yerine x=5 yazsaydık var olan değişkenin değerini değiştirmiş olurduk.

// => Shadowing variables sorunlara yol açabilen bir durumdur. O yüzden değişken isimlerinin farklı
// => olması mühimdir.

	// x := 10
	// if x > 5 {
	// 	fmt.Println(x)
	// 	x := 5
	// 	fmt.Println(x)
	// }
	// fmt.Println(x)
}