package formatter

import "fmt"

// büyük harfli olmalıdır diğer package'ların buna ulaşması için
func Format(num int) string {
	return fmt.Sprintf("The number is %d", num)
}
