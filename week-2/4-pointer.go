package main


////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// Tanım

// import "fmt"

// type person struct {
// 	age  int
// 	name string
// }

// func main() {
// 	p := person{}
// 	i := 2
// 	s := "Hello"
// 	tryToModify(i, s, p)

// 	fmt.Println(i, s, p)/*Bu örnekte çıktımız: 2 Hello {0 } olur. Oysaki alttaki fonksiyonda s, "Goodbye"
// 	olarak değişmişti. Ayrıca p.name ="Zeytin" idi. i ise 2 ile çarpılmıştı. Yani main içindeki fonksiyon
// 	pek bir işe yaramadı.
// 	GO'da callbyvalue adında bir özellik vardır. Biz tryToModify(i,s,p) yazınca GO bunun her zaman bir
// 	kopyasını oluşturur. Bizim değiştirmek istediğimiz değişken gider ve kopyamız var onu değiştirdi ve
// 	geriye birşey de dönmedi. Sonuç olarak herhangi bir değişkeni değiştiremedik. Alttaki fonksiyon
// 	içinde değişti ve o scope içindeki olay bittiyse değiştirilen şey yok olud.
// 	Pointer'ın işe yarar noktalarından biri budur. Biz birşey pass edip bunun değişmesini istersek lazım.
// 	https://www.tutorialspoint.com/go/go_function_call_by_value.htm
// 	*/ 

// }

// func tryToModify(i int, s string, p person) {
// 	i = i * 2
// 	s = "Goodbye"
// 	p.name = "Zeytin"
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// MAP ve SLICE için POINTER DURUMU //

//import "fmt"

// func main() {
// 	m := map[int]string{
// 		1: "first",
// 		2: "second",
// 	}
// 	modMap(m)
// 	fmt.Println(m)

// 	s := []int{1, 2, 3}
// 	modSlice(s)
// 	fmt.Println(s)
// }

// func modMap(m map[int]string) {
// 	m[3] = "hello"  // eğer 2:"hello" gibi birşey yazsa idik değişirdi aynı key'in value'sü o zaman çıktı:
// 	// map[2:hello 4:goodbye]  // yazmayınca böyle çıktı : map[2:second 3:hello 4:goodbye]
// 	m[4] = "goodbye" 
// 	delete(m, 1)
// }
// /*Bu olayın sebebi map'ler, slice'lar için olan bir durumdur. Bunlar GO'nun kendisinde olan çekirdeğine
// pointer ile implement edildiği için GO bunu her zaman biz sanki pointer ile pass ediyormuşuz gibi çalışır.
// Map ve slice'de değiştirme yapar, RAM'deki adresi referans alır. Yani kopyasını oluşturmaz, direk alır. */

// func modSlice(s []int) {
// 	for k, v := range s {
// 		s[k] = v * 2
// 	}
// 	s = append(s, 10) //s 'e s'i yeniden ekle // yazmasak da çalışıyor ? Galiba yerini netleştiriyoruz
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

//-------POINTERS--------//

// Kabaca elimizdeki değerin RAM'deki yeridir. Bir değişken yazsak bu bilgi bilgisayarda RAM'de bir 
// adres olarak kayıtlı olur. Bu adres değişkeni referans gösterir. Bazı noktalarda bizim o değişkenin
// adresine ihtiyacım olur. O durumda biz bu adresi kullanarak var olan değişkenin değerini değiştirmek
// isteriz. 
// İşte o bazı durumlarda direk olarak değişkenin değeri değişemez ve biz bu durumda değişkenin adresini 
// alırız, bu adres ile o değişkene ulaşıp değerini değiştirebiliriz. 


// İlk örnekte tryToModify fonksiyonu içinde olan değişiklik main'de olmamıştı. 
// Şimdi olacak

// import ("fmt")

// type person struct {
// 	age  int
// 	name string
// }

// func main() {
// 	p := person{}
// 	i := 2
// 	s := "Hello"
// 	tryToModify(&i, &s, &p)
// 	fmt.Println(i, s, p) 	//değişti. Çıktı: 4 Goodbye {0 Zeytin}
// }

// func tryToModify(i *int, s *string, p *person) { /* yukarıda bahsedilen... adresi kabul ederken * işareti
// 	ile kabul edilir. * indirection operation olarak geçer. Adresi aldık, * ile o adresin karşılığını
// 	okuruz. */
// 	text := "Goodbye"
// 	*i = *i * 2
// 	*s = text
// 	p.name = "Zeytin"
// }


////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

/*Aşağıda bir struct'ımız var. Middlename adındaki field pointer alıyor.*/

// import "fmt"

// type person struct {
// 	FirstName  string
// 	MiddleName *string //pointer kasıtı
// 	LastName   string
// }

// func main() {
// p := person{
// 	FirstName:  "Pat",
// 	MiddleName: "Perry", // Bu satır compile(derlenme) olmaz. Çünkü biz burada biz string'in kendisini
// 	// veririz. GO bunu istemez, pointer verecektik ancak string vermişiz. Eğer func, struct bizden 
// 	// referans isterse direk olarak string veya integer vs. pass edilemez. O string'i, int'i vs.'yi RAM'e
// 	// yazmak gereklidir. O string'i memory'e yazıp adresi pass edebiliriz.	
// 	LastName:   "Peterson",
// }

// p := person{
// 	FirstName:  "Pat",
// 	MiddleName: stringp("Perry"), /*Biz direk "Perry" yazamıyoruz.Bir tane helper func var aşağıda(stringp).
// 	Bu helper function s diye string değer alıyor ve dönüyor. Sonra buraya onu çalıştırız ve bu sayede
// 	referansı döndürmüş oluruz.*/
// 	LastName:   "Peterson",
// }

// üsttekini yorum alak ve bunu yazıp altta print edek.
// p := person{
// 	FirstName: "Pat",
// 	LastName:  "Peterson",
// }

// fmt.Println(p.FirstName, p.MiddleName, p.LastName) /*Çıktı: Pat <nil> Peterson
// p.MiddleName bir pointer alır ve o yüzden <nil> olarak gözükür. Yoksa ekranda herhangi birşey gözükmezdi
// */
// }

// func stringp(s string) *string {
// 	return &s
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

//------ DO NOT FEAR THE POINTERS------//

/*Değişkenin adresini vererek değiştirmek istiyorsak aşağıdaki gibi.
f ile içine pas ettik, * ile yakaladık ve value'sunu değiştirdik.*/
// import "fmt"

// func main() { 
// 	var f *int
// 	// burada ilk değerimiz nil olduğu için f'yi alıp g'ye yani 8'e eşitleyemiyoruz.  
// 	// RAM'de bir f var ama değeri belirsiz. Tanımlamak lazım, initial etmek gerek.  
// 	willItUpdate(f) 
// 	fmt.Println(f) 	// Bu yüzden de <nil> döner. 
// }

// func willItUpdate(g *int) {
// 	// bazı func'lar referans adresi ister, referans değerli şekilde atama ister.
// 	x := 8 // önce x değişkeni oluşturup 8 e eşitleriz
// 	g = &x // sonra g'yi x'e atarız. Bu şekilde olur. 
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////


// import "fmt"
// type Person struct {
// 	FirstName  string
// 	MiddleName string
// 	LastName   string
// }

//BUNU KESİNLİKLE YAPMAYINIZ. Bir struct değiştirebilirsiniz ama bunu return etmeniz lazım performan için
// func MakeFoo(p *Person) error {
// 	p.FirstName = "mustafa"
// 	p.MiddleName = "kemal"
// 	p.LastName = "atatürk"
// 	return nil
// }

// // BUNU YAPINIZ ÜSTTEKİ YERİNE
// func MakePerson(p Person) (Person, error) {
// 	p.FirstName = "mustafa"
// 	p.MiddleName = "kemal"
// 	p.LastName = "atatürk"
// 	return p, nil
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

/*TAG'LER
Struct'ların tag'leri vardır. 
*/
/*Elimizde bir json ve field'ı var. Aşağıdakini düşünelim, name ve age'i var json'un. Bizim uygulama
dışında başka bir yerde json varsa.... Bir yere istek atıyoruz ve bize name ve age dönüyor. 
Bizim bu data'yı struct'ımıza bind etmemiz gerekiyor. Bind etmek ...
Mesela f.name yazınca Mustafa ismini almak isteriz. f.age yazınca 30'u almak isteiz. Unmarshal yapmak
gerekir. Unmarshal, JSON datayı alır, struct'ımıza gömer eğer veriler tutuyorsa. JSON valid ise tabiki.
JSON içindeki name ve age'i struct içindeki
name ve age'e eşitlememiz gerekiyor. Bunun için TAG'ler kullanılır. */

// import (
// 	"fmt"
// 	"encoding/json"	
// 	"os"
// )

// func main() {
// 	f := struct {
// 		//json datasını bu struct'a eşitleyecek json data mevcut ise sorunlu değilse.
// 		Isim string `json:"name"` // json:"name" bizim tag'imizinden biri alttaki ile birlikte. 
// 		Yas  int    `json:"age"`  
// 	}{}
// 			// önemli olan data'dan gelen json'ın isimlerinin (name ve age) buradaki struct'daki ile 
// 			// eşleşmesi
// 								/// içine bu data alınmış, encode eder.
// 	err := json.Unmarshal([]byte(`{"name": "Mustafa", "age": 30}`), &f) //f referans verilir.
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	fmt.Println(f)
// }
