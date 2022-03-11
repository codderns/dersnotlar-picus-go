package main

/*GO'da birçok keywords vardır, break, continue, const, default, defer, struct, switch, var, for vs...
Bu anahtar kelimelerden biri de interface'dir. 
 */

// import (
// 	"fmt"
// )

// func main() {
// 	var i interface{} // i'nin türü interface olduktan sonra i'yi birçok farklı türdeki değere atayabildik
// 	/*Bazı koşullarda bu durum olabilir. Değişken interface olarak tanımlanırsa istenildiği gibi 
// 	kullanılabilir. 
// 	*/
// 	i = 20
// 	i = "hello"
// 	i = struct {
// 		FirstName string
// 		LastName  string
// 	}{"Mustafa", "Kemal"}

// 	fmt.Println(i) // satırda son alınan değer neyse ona eşittir i.
// }


////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

// import (
// 		"fmt"
// 		"io/ioutil"
// 		"encoding/json"
// 	)

// /* Bir JSON dosyasında key veya value değerleri int, string, dizi veya bool olabilir. Bunları aynı 
// değişken için farklı değerler alması olanaksız gelir GO'da eğer interface olmasaydı. Interface tanımı
// ile bu sorun çözülür. */

// func main() {
// 	data := map[string]interface{}{} //map karşısına interface de alabilir.
// 	contents, err := ioutil.ReadFile("data.json") //ReadFile bir function'dur, dosya okuma adı üstünde
// 	if err != nil {
// 		panic(err) //dosya yoksa panic ver. Alta devam etmesin.
// 	}

// //json.Unmarshal json dosyası içindeki datayı alır ve okur ve üsteki map'in içinde gömer.
// 								//burada & ile ref değere pass edilmiş. Eğer referans etmezsek 
// // çalıştırıncaboş bir data göreceğiz. Biz datanın adresini veriyoruz, git buranın adresinin olduğu 
// // value'yü al değiştir. //error varsa panic, yoksa devam.
// 	if err := json.Unmarshal(contents, &data); err != nil { //bu tarz yukarıdaki gibinin daha şekillisi
// 		panic(err)  
// 	}

// 	fmt.Println(data)
// 	/*Çıktı: 
// map[age:21 details:map[firstKey:firstValue secondKey:secondValue] isFemale:false name:Mustafa test:123]*/

// 	fmt.Println(data["age"].(float64) + 1)//age'de 21 değerini 22 yapmak için o key'in value'sünün türünü
// 	//yazmak gerekir. // eğer int verse idik panic verir. GO işini garantiye alan dil olduğu için 
// 	// anlayabileceği ölçeği geniş tutar. O yüzden sayı varsa float64 olarak anlar.

// 	fmt.Println(len(data["name"].(string)))// aynı şekilde value'de değişiklik için türünü belirtmek lazım
// // Data tipinin ne olduğunu belirtmeye "type casting" denir.

// 	fmt.Println(!data["isFemale"].(bool))

// 					//json'da details içindeki map olduğu için türüne map yazılmışş map içindeki ise string
// 					//ve karşılığı interface olacağı için böyle belirtilmiş
// 	details := data["details"].(map[string]interface{})
// 	for _, d := range details {
// 		fmt.Println(d)
// 	}

// 	if err := json.Unmarshal(contents, &data); err != nil {
// 		panic(err)
// 	}

// 	byteData, err := json.Marshal(data)
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println(byteData)
// 	fmt.Println(string(byteData))
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////


//  import (
// 	"fmt" 
// 	"io/ioutil"
// 	"encoding/json"
// 	)

// type Person struct { /*Hep aynı yapıda ise json, data.json'u struct'a çevirdik 
// https://json2struct.mervine.net/ bu siteden. */ 
// 	Name     string `json:"name"`
// 	Age      int    `json:"age"`
// 	IsFemale bool   `json:"isFemale"`
// 	Details  struct {
// 		FirstKey  string `json:"firstKey"`
// 		SecondKey string `json:"secondKey"`
// 	} `json:"details"`
// }
// func main(){
// 	data := &Person{}  // data adında Person'dan ref aldı, atandı. Initializing
// 	contents,err := ioutil.ReadFile("data.json") //json dosyasını oku
// 	if err != nil{
// 		panic(err)
// 	}			
// 	if err:= json.Unmarshal(contents,data); err != nil{ // byte et json'u. &data yazsa idik yukarıdaki
// 	// &Person, Person şeklinde olacaktı
// 		panic(err)
// 	}
// 	fmt.Println(data.Name,data.Age,data.IsFemale,data.Details) //Mustafa 21 false {firstValue secondValue}

// 	byteData, err := json.Marshal(data) // Struct'u veya metni json formatına dönüştürür Marshal
// 	// bir err de döner
// 	if err != nil{
// 		panic(err)
// 	}
// 	fmt.Println(byteData) /*Çıktı:
// 	[123 34 110 97 109 101 34 58 34 77 117 115 116 97 102 97 34 44 34 97 103 101 34 58 50 49 44 34 105 
// 	115 70 101 109 97 108 101 34 58 102 97 108 115 101 44 34 100 101 116 97 105 108 115 34 58 123 34 
// 	102 105 114 115 116 75 101 121 34 58 34 102 105 114 115 116 86 97 108 117 101 34 44 34 115 101 99 
// 	111 110 100 75 101 121 34 58 34 115 101 99 111 110 100 86 97 108 117 101 34 125 125]*/
// 	fmt.Println(string(byteData)) 
// 	// Çıktı: 
// //{"name":"Mustafa","age":21,"isFemale":false,"details":{"firstKey":"firstValue","secondKey":"secondValue"}}
// }

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

//Genel Bilgi Not: Interface struct'a bağlanır, struct'un metodu oluşturulur. Arada interface bir layer
// olarak düşün.


// Soyut tip interface

import "fmt"

 //  	Direk struct ile ilgilidir. Üstteki örneklerdeki interface datanın ne türünde olduğuyla alakalı

 // buradaki interface farklı bir amacı vardır. Struct'ın davranışı ile alakalıdır. 
type Animal interface {
	Sound() interface{}  // Sound() fonks ve Eat() davranışları var Animal'ın. // alttaki return türleri
// farklı olduğu için Sound'un signature'ı interface ve alttakilerde de belirtimeli Sound kullanılanlarda
	Eat(bool) string
	// Aşağıdaki diğer struct'ları buna implement ederiz
}

type Dog struct {
}

func (d Dog) Sound() interface{} {
	return 123
}

func (d Dog) Eat(b bool) string {
	return "dog-eat!"
}

type Cat struct {
}

func (c Cat) Sound() interface{} {
	return "Miyav!"
}
func (c Cat) Eat(b bool) string {
	return "cat-eat!"
}

type Cow struct {
}

func (l Cow) Sound() interface{} {
	return "Moo!"
}
func (l Cow) Eat(b bool) string {
	return "cow-eat!"
}

//if you got method you are my type
func main() {
	animals := []Animal{Dog{}, Cat{}, Cow{}}

	//dog:=Dog{}
	//dog.Sound()
	//
	//cat := Cat{}
	//cat.Sound()

	for _, animal := range animals {
		fmt.Println(animal.Sound())
		fmt.Println(animal.Eat(true))
	}
}
