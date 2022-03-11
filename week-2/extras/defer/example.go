/// defer'ı daha iyi anlamak için örnek:

package main

import (
	"fmt"
)

func main() {
	B()  //B fonksiyonuna git ↓
}

func A() {
	defer D()  // defer var. Önce A içindeki işlemler çalışır sonra D'ye gider. ↓
	fmt.Print("P")  // Bu sebepten P harfini yazdırır.
}

func B() { 
	C()    // C fonksiyonunu çalıştır  ↓
	fmt.Print("E") // üsteki çalışır ve bunu yazmaz tabi ki (şimdilik), ilk görev gerçekleşince geri gelir 
	//bu fonksiyondaki tüm işlemi tamamlamak için
}

func C() {
	defer E()  // E ye değil A 'ya gider. ↑ Defer'ın özelliği budur. Bu yüzden satır sondaki işlem görür.
	defer A()
}

func D() {
	fmt.Print("A") // A harfini yazdırır.  Şimdi ise takip edince, D'yi çağıran A idi. A'yı çağıran
	// C idi. C'de defer A çalışıtırılmıştı. Şimdi ise E yi çalıştıracaktır. ↓
}

func E() {
	fmt.Print("C") // C yazdırır. Buradaki işlem de bitince B fonksiyonunda çalışmayan print çalışır.
	// Çünkü C fonksiyonu çalıştıktan sonra bir sürü işlem oldu, bunlar gerçekleşti ve oradaki print 
	// işlemi olur.
}

func F() {
	fmt.Print("B")
}

// Nihayetinde ekrana P A C E yazılır
