package main

import (
	"io"
	"log"
	"os"
)
// defer, func kapanmadan önce çalışan kod satırıdır.
func main() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
 
	defer f.Close() // bazı noktalarda bir bağlantı açılır, dosya açılır... bunlar iş bitildikten sonra 
	// kapatılmalıdır. yoksa optimizasyonda sorun olabilir, ramde yer kaplar vs. İşi garantiye almak için
	// burada defer kullanıp func bitince dosyayı kapa emri oluştururuz.
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
			break
		}
	}
}
