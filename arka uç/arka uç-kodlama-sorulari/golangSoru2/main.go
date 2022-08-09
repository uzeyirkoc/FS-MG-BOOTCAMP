package golangSoru2

// Soru 2: SwapInt isimli fonksiyonu, parametre olarak girilen i ve i2 değişkenlerinin değerlerini birbiri ile değiştirecek şekilde oluşturunuz.

func SwapInt(i *int, i2 *int) {
	// Kodu geliştiriniz
	i = 33
	i2 = 11

	i = i + i2
	i2 = i - i2
	i = i - i2
	fmt.Printf("Değişim Sonucu : i = %v, i2 = %v", i, i2)

}
