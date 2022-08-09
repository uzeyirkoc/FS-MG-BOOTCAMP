//package golangSoru3

// Soru 4: IkiUstuN'i Geliştiriniz.

//func IkiUstuN(n int) int {
	// 0, 1, 2, 4, ... , 2^n-1 toplamlarını gösterecek bir yazılımı yazınız.
	// Ipucu: Olabildiğince hızlı çalışması önemlidir.
//	return -1
//}
package mgbootcamp

func IkiUstun(n ...int) int {
	toplam := 0
	for i := 0; i < len(n); i++ {
		toplam = toplam + 2 ^ i - 1
	}
	return toplam
}
