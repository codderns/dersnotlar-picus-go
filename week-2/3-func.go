package main

//import "errors"


import (
	"errors"
	"fmt"
)

// func main() {
// 	fmt.Println(sum(99, 1)) // alttaki sum fonksiyonu burada çalıştırıldı.
// }

// // Go'da function tanımlarken türüne de belirtmek gerekir. func yanındaki bilgilere signature denir.
// // signature yanına hangi tür dönecekse o yazılır. 

// func sum(a, b int) int {  
// 	return a + b
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// Fonksiyon içinde ... yazmanın []int 'dan bir farkı yoktur. Parametrelerden birinin slice olduğu
// belirtilir. Bu fonksiyonda base adında int parametresi var, girilen diğer tüm int parametreler 
// ikinci parametre olan diziye ait olacak. Eğer varsa dizi içindeki değerleri base'deki değerler ile
// toplayıp değerleri toplanmış "dizi"yi return etmek isteniyor.
// func addTo(base int, vals ...int) []int {
// 	out := make([]int, 0, len(vals))
// 	for _, v := range vals {
// 		out = append(out, base+v)
// 	}
// 	return out
// }
// //
// func main() {
// 	fmt.Println(addTo(3)) //dizi için input yok ve dizi boş döner. []
// 	fmt.Println(addTo(3, 2)) // dizinin tek değeri 3 olan base değeri ile toplanır ve [5] return eder 
// 	fmt.Println(addTo(3, 2, 4, 6, 8)) // => [5 7 9 11]
// 	a := []int{4, 3} // bir dizi oluştur
// 	fmt.Println(addTo(3, a...)) // o diziyi parametre olarak gir fonksiyona [7 6]
// 	fmt.Println(addTo(3, []int{1, 2, 3, 4, 5}...)) // yine dizi parametre olarak girilmiş [4 5 6 7 8]
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// 0'a bölme hatasını denetleyen fonksiyon
// func divide(numerator, denominator float64) (float64, error) {
// 	if denominator == 0 {
// 		return 0, errors.New("cannot divide by zero") // boş ve errors değeri oluşturup string'ini dön
// 		// error varsa ilk değer kullanılmayacağı için bir değer girmen yeterlidir.
// 	}

// 	return numerator / denominator, nil  //ilk return float değer ikincisi ise nil dönsün, hiç hata yok.
// }
// // GO'da hatanızı try catch gibi bir ifade, blok ile yakalamak yerine programcı bunu kendi oluşturur. //
// func main() {
// 	result, err := divide(5, 3)  //fonksiyon iki değişkene eşitlenmeli zira float64 ve error olarak 
// 	// return ediliyor
// 	if err != nil {
// 		fmt.Println(err)
// 	}
	
// 	resultDiv, _ := divide(5, 3) // error'u check etmek istemiyorsak _ koyarak bunla işimizin olmadığını belirtiriz okuyan için. error'u görmemiş olur. ama error'larda pek kullanılmaz bu durum.
// 	fmt.Println(resultDiv)

// 	fmt.Println(result)
// 	result, err = divide(10, 2)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result)

// 	result, err = divide(8, 0)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	fmt.Println(result) 
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// Function'lar type olarak tanımlanabilir. Böyle olunca parametreleri gerek yoktur yazmaya. 

// type divider func(int, int) int  // tanımlanır ve parametreleri yazmaya gerek yoktur.

// var opMap = map[string]divider{}  // map'te key'ler karşısına herşeyi alabilir. Burada da yukarıdaki 
// //fonksiyonu almış.
// // Şöyle de olabilirdi:
// var opMap2 = map[string]func(int,int) int{}  //ama bunun yerine yukarıdakini uygulamak okunaklık için
// // çok daha iyidir.


////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////


// type Person struct {
// 	FirstName string
// 	LastName  string
// 	Age       int
// }

// func main() {
// 	people := []Person{ // Person yapısından dizi.
// 		{"Zeytin", "Home", 21},
// 		{"Ahmet", "Home", 21},
// 		{"Fatma", "Home", 21},
// 	}
// 	// Bazı noktalarda bazı funcion'lar parametre olarak function alabilir. Mesela:
// // GO'da sort adında bir packace vardır. Interface olarak people'ı aldı burada. 
// 	sort.Slice(people, func(i int, j int) bool {  //2. parametre bir func ve func i,j parametresi ve bool dönüyor.
// 		return people[i].LastName < people[j].LastName
// 	})
// }
// Yani bir fonksiyon fonksiyon olarak parametre alabilir. Son kısım ileriki kısımlarda ayrıntılandırılacak.
