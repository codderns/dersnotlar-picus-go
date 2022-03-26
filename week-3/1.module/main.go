package main

import (
	"fmt"
	"log"

	"github.com/google/uuid" // kod çalışıyor
)

// go mod init github.com/nsgnd/dersnotlar-picus-go
// go mod tidy
// go get github.com/google/uuid
func main() {
	u, err := uuid.NewUUID()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(u.String()) //  fa582db5-ad06-11ec-ba48-0071c214b6ee
	// her çalıştığında değişir
}
