package utils

func GetColor(index int) string {
	colors := []string{
		"#00FF00", // 1. renk: Yeşil
		"#0000FF", // 2. renk: Mavi
		"#FFA500", // 3. renk: Turuncu
		"#FF0000", // 4. renk: Kırmızı
		"#800080", // 5. renk: Mor
		"#FFFF00", // 6. renk: Sarı
		"#00FFFF", // 7. renk: Camgöbeği
		"#FFC0CB", // 8. renk: Pembe
	}
	return colors[index%len(colors)]
}
