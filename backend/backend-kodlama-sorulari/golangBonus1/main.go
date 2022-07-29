package golangBonus1

// Bonus #1-1: Aşağıdaki Kod nerede kullanılabilir açıklayınız.
// Bonus #1-2: Bu fonksiyonun maliyetini hesaplayınız ve daha düşük maliyet için fikir yürütünüz.
var a = "lorem ipsum dolor sit amet consectetur adipiscing elit"
var c int32 = 10

func Ceaser(s *string, i int32) string {
	tmp := ""
	for _, v := range *s {
		character := v
		if character < 0 && character > 127 {
			return ""
		}
		for j := 0; j < int(i); j++ {
			character++
		}
		tmp += string(character)
	}
	return tmp
}
