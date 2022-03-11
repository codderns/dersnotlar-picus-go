package main

import (
	"fmt"
	"os"  //os argümanları alır. terminale go run 1-args.go twitter yahoo google yazınca
	// [C:\Users\ENES\AppData\Local\Temp\go-build3134769506\b001\exe\1-args.exe twitter yahoo google]
	//bastırır. 1-args.go 'dan sonra ne yazarsak yazalım oradaki argümanları bizim için toplayacaktır.
)

func main() {
	fmt.Println(os.Args) 
}
