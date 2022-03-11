package main

import "fmt"

func main(){

	//////////////Composite Types////////////////

	// Bunlar => Array, Slice, Struct, Pointer, Function, Interface, Map, Channel 

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

	//Arrays

	// var x [3]int // içine maks üç değer gireceği belirtilir.  eğer girilmez ise 0 dolar içi
	// fmt.Println(x) // Çıktı : [0 0 0]

	// var y = [3]int{10, 20, 30} // içine değerler verilebilir
	// y = append(y, 40) 		 // hata verir. slice'a append edebiliriz.
	// fmt.Println(y)
	//var x = []int{10, 20, 30}
	//var x [][]int // iki boyutlu bir dizi. matrix.

	
////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

	// Slices
	// [...] -> array
	// [] -> slice

	// Slice'lar sabit değil esnektir değer alma konusunda.

	// x := []int{10, 20, 30} // []int bir değişkenin slice olduğunu tanımlatır go'ya.
	// fmt.Println(len(x))

	// var x []int
	// y := []int{40, 50, 60}
	// x = append(x, y...) // Bir slice'taki veriyi başka bir slice'a eklemek bu şekilde olur. Sona eklenir.
	// fmt.Println(x)

/// Capacity ////
// => Bir slice'a veri ekledikçe go o değişkene bir mantık ile ram'de ekstra yer ayırır.
// => bu ayrılan alan capasity'dir.

	// var x []int
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 10)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 20)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 30)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 40)
	// fmt.Println(x, len(x), cap(x))
	// x = append(x, 50)
	// fmt.Println(x, len(x), cap(x))


/// Make ///
// => Bir slice yaratmak için kullanılır. Başlangıç tanımlaması için tür(type) ve boyut(size) 
// => girişi tanımlanır.

	// x := make([]int, 5)
	// fmt.Println(x)		// Çıktı : [0 0 0 0 0]
	// x = append(x, 99)	
	// fmt.Println(x)		// Çıktı : [0 0 0 0 0 99]

	// x := make([]int, 5, 10) // Tür, boyut tanımlaması ile birlikte kapasite (10) değeri de girilebilir.
	// x = append(x, 1)
	// x = append(x, 2)
	// fmt.Println(x, len(x), cap(x)) //[0 0 0 0 0 1 2] 7 10 
	// x = append(x, 3)
	// x = append(x, 4)
	// x = append(x, 5)
	// x = append(x, 6)
	// fmt.Println(x, len(x), cap(x)) //Çıktı [0 0 0 0 0 1 2 3 4 5 6 7] 11 20 //cap dolunca +10 arttırıldı
	// x = append(x, 7)
	// fmt.Println(x, len(x), cap(x)) //Çıktı [0 0 0 0 0 1 2 3 4 5 6 7] 12 20 
	

/// Slicing Slices ///
// => Bir slice'ın bir bölümü bir değişkene atılabilir. Veya tamamı da olabilir

	//x := []int{1, 2, 3, 4}
	// y := x[:2]  // baştan 2. index'e kadar (dahil değil)
	// fmt.Println(y)
	
	// z := x[1:] // 1. indexten (dahil) sona kadar.
	// fmt.Println(z)

	//e := x[:]   // tamamını al
	//fmt.Println(e)

	//x := [4]int{1, 2, 3, 4}
	//y := x[:2]


/// copy ///
// => Bir veriyi kopyalamak ve kopyalanan veri önceki veriden bağımsız şekilde saklamak istersek kullanılır

	//x := []int{1, 2, 3, 4}
	//y := x
	//y := make([]int, len(x)) // bu önemlidir yani len(). Çünkü kopyalanacak verinin boyutu bilinmeyebilir.
	//num := copy(y, x)
	//fmt.Println(y, num)

	//y:=x => bu şekilde de y'deki veriler x'e geçerdi ancak, RAM'in kendisinde aynı referans bazlı bir
	//gösterim olur bu. y'de yapılan değişiklik bu durumda x'i etkileyebilir.


////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

	//Maps

	// Key Value diye de bilinir. Bir datayı bir data ile onun değerini bir yerde birleştirmek isteriz.
	// Bu gibi durumlarda map kullanılır

//Tanımlama 1 
	//var x map[string]int  // Tanımlama bu şekilde olur. // string key, int ise value'nün türüdür
//Tanımlama 2
	//totalWins := map[string]int{} // {} belirtmek önemli burada, RAM'e içi boş şekilde koy anlamını taşır.
//Tanımlama 3
	// teams := map[string][]string{
	// 	"Ahmet": []string{"a", ""},
	// 	"Ayşe":  []string{"e", "f", "g"},
	// } //içi değerler alarak yapıldı. Key değer, karşısındaki string slice alır. Bir function bile karşısında alabilir.
	//Karşısına channel bile alabilir. Map, karşılıklı dataları tutan bir yapıdır.

	// teams["Sevde"] = []string{"a", "y"}
	// teams["Sevde"] = []string{"x", "o"}  // value değerleri "a y" den "x o" ile değişti.
	// fmt.Println(teams)
	//

	//y := teams["Ayşe"]
	//y = append(y, "k")
	//teams["Ayşe"] = y
	//fmt.Println(teams)

	//ages := make(map[string]string, 10) // bu şekilde de tanımlama olabilir. [key'in türü] value'nün türü
	// 10 ise initial olarak 10 değer alabilir anlamında. Bir yerden data geldi, performans için yer ayırtmış olduk
	//fmt.Println(ages)

	 
/// Reading and Writing a Map ///
// => Key'in Karşısındaki Value Değeri ile Oynanabilir.

	//totalWins := map[string]int{}
	//totalWins["Ayşe"] = 1
	//totalWins["Osman"] = 2
	//totalWins["Ayşe"]++
	//fmt.Println(totalWins)

	// The comma ok idiom
	m := map[string]string{
		"hello": "5",
		"world": "0",
	}
	deger, var_mi := m["goodbye"]  // m map'inin içinde "goodbye" adında key var mı? varsa var_mi, true 
	//olsun. Bu key'in eğer value'sü varsa deger, o value'ye eşitlensin, yokda boş değer alsın.

	fmt.Println(deger, var_mi)
	delete(m,"hello") //hello key'ini sil.
	fmt.Println(m)

}