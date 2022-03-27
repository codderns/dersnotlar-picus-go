package main

import "fmt"

func main() {
	//TEMEL VERİ TİPLERİ İÇİN BAKINIZ : BASIC TYPES (BUILT-IN TYPES FOR GOLANG)

	// go, değişken atama ve değer verme hususunda statik yöne sahiptir, net olarak türünü belirtmek gerekir değişkenlerin

	// Boolean
	var flag bool
	fmt.Println(flag)

	var patika bool = true
	fmt.Println(patika)
	//patika = "patika" //bir bool int'e bir int string'e vs dönüşümler olamaz
	// Ancak int64, int32'ye dönüşüm gibi değişiklikler olabilir.

	//var str = "10"
	//var Integer = int(str) //cannot convert str (type string) to type int

	var integer int64 = 45
	var yeni = int32(integer)
	fmt.Printf("%T \n", yeni) //int32 olduğu görülür.

	// Integer Types
	var num int64 = 10

	// Float
	var num1 float32 = 10.2
	var num2 float64 = 10.2

	//var vs := // hatalı
	var xa int = 10
	//	var x = 10 // bu şekilde de atama olabilir
	//	var x int // bir kez deklare edilir bir değişken
	fmt.Println(xa)

	//	var ya float64 = 20
	ya := 20

	// Çoklu Değişken Tanımlama

	var x, y int = 10, 20
	var a, b = 10, "patika"

	var (
		c    int
		d        = 20
		e    int = 30
		f, g     = 40, "hello"
		h, i string
	)
	fmt.Println(num, num1, num2, ya, x, y, a, b, c, d, e, f, g, h, i)
	//const
	const pi = 3
	fmt.Println(pi) // kesinlikle değiştirilemez, hata verir denenirse

	//Zero Values
	//numerik değerlerin go'daki değeri sıfırdır :
	var zeroint int
	fmt.Println(zeroint) //varsayılan olarak sıfır alırlar

	var zerostring string
	fmt.Println(zerostring) // varsayılan olarak boş string

	var zerobool bool
	fmt.Println(zerobool) // varsayılan olarak false

	// Const asla değiştirilemez değişkenlerdir.
	//const x float64 = 10
	//
	//const max_value = 1
	//const MAX_VALUE = 1
	//
	//const MAXVALUE = 1
	//const maxValue = 1
}
