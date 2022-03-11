package main

import "fmt"

func main(){

	// if

	//n := 10
	//if n == 0{
	//	//
	//}else if n == 1{
	//
	//}else{
	//
	//}

////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////

	// for


	// c-style //

	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	//


	// condition-only //

	//i := 1
	//for i < 100 {
	//	fmt.Println(i)
	//	i = i * 2
	//}
	//


	// infinite //

	//for {
	//	fmt.Println("HELLO")
	//}
	//


	// break //
	
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		break
	//	}
	//}


	// continue //
	
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//	if i == 5{
	//		fmt.Println("finished")
	//		continue
	//	}
	//}


	// for-range //

	// ciftsayilar := []int{2, 4, 6, 8, 10}
	// for index, value := range ciftsayilar {
	// 	fmt.Println(index,value)
	// }

	// evenVals := []int{2, 4, 6, 8, 10}
	// for _, v := range evenVals {  //eğer index print etmesin dilersek
	// 	fmt.Println(v)
	// }


	// iterate //

	// m := map[string]int{
	// 	"a": 1,
	// 	"c": 3,
	// 	"b": 2,
	// }
	
	// for i := 0; i < 3; i++ {
	// 	fmt.Println(m)  //çıktı bu kısmın : map[a:1 b:2 c:3] . Bir order var.
	
	// 	fmt.Println("Loop", i) 
	// 	for k, v := range m {
	// 		fmt.Println(k, v) // index'ler yani map'teki değerlerin sırası değişir her loop'ta

	// 		/*Sebebi security ile ilgili. Biz map'e 4. veri eklesek. Sıralama a:1 b:2 c:3 d:4 olur.
	// 		Güvenlik için hash table'lar farklı çalışır. Ne kadar eklersek ekleyelim, biz buraya sıralı
	// 		eklesek de c'den sonra b'nin geldiğini bilmemek gerekir. Go'nun hash algoritması rastgele
	// 		verebilir diğer hash'in ne olduğunu bilmemek için. */
	// 	}
	// }

	evenVals := []int{2, 4, 6, 8, 10}
	for _, v := range evenVals {
		v *= 2
	}
	fmt.Println(evenVals)
	// for bir scope'tur ve bu scope'un dışına çıkınca değerlerin değişmez. Scope içinde değişiklikler
	// dışarıya aksetmez.

	
////////////////////////////////////////////****************////////////////////////////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////****************////////////////////////////////****************////////////////////
////////////////////////////////////////////****************////////////////////////////////////////////


	// switch case
	
	//words := []string{"a","cow","smile","gopher", "octopus"}
	//for _, word := range words{
	//	switch word{
	//	case "a":
	//		fmt.Println("its a")
	//	case "cow":
	//		fmt.Println("cow")
	//	default:
	//		fmt.Println("Default value")
	//	}
	//}
}