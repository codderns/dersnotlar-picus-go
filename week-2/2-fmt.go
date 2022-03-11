package main

import (
	"fmt"
	"os"
)

// eğer terminale şunu yazarsak: \go run 2-fmt.go yahoo google microsoft

func main() {
	fmt.Println(os.Args)
// ilk çıktı: 
// [C:\Users\ENES\AppData\Local\Temp\go-build4022411997\b001\exe\2-fmt.exe yahoo google microsoft]

// bu bir string dizisidir : 
	args := os.Args[1:] //1'den(yahoo) başlayarak sırala

	for i, arg := range args {
		fmt.Printf("index: %d, value: %s\n", i, arg) //printf format halinde bilgileri gösterebilir
						// %d integer %s string.
		
		/*
		index: 0, value: yahoo
		index: 1, value: google
		index: 2, value: microsoft
		*/
		// Bir çıktı alırız.
	}
	
}
