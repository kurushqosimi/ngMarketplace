package translit

import (
	"strings"
	"unicode"
)

// translitTable - таблица соответствия кириллических символов латинским.
var translitTable = map[rune]string{
	'А': "A", 'а': "a",
	'Б': "B", 'б': "b",
	'В': "V", 'в': "v",
	'Г': "G", 'г': "g",
	'Д': "D", 'д': "d",
	'Е': "E", 'е': "e",
	'Ё': "Yo", 'ё': "yo",
	'Ж': "Zh", 'ж': "zh",
	'З': "Z", 'з': "z",
	'И': "I", 'и': "i",
	'Й': "J", 'й': "j",
	'К': "K", 'к': "k",
	'Л': "L", 'л': "l",
	'М': "M", 'м': "m",
	'Н': "N", 'н': "n",
	'О': "O", 'о': "o",
	'П': "P", 'п': "p",
	'Р': "R", 'р': "r",
	'С': "S", 'с': "s",
	'Т': "T", 'т': "t",
	'У': "U", 'у': "u",
	'Ф': "F", 'ф': "f",
	'Х': "Kh", 'х': "kh",
	'Ц': "Ts", 'ц': "ts",
	'Ч': "Ch", 'ч': "ch",
	'Ш': "Sh", 'ш': "sh",
	'Щ': "Shch", 'щ': "shch",
	'Ъ': "", 'ъ': "",
	'Ы': "Y", 'ы': "y",
	'Ь': "", 'ь': "",
	'Э': "E", 'э': "e",
	'Ю': "Yu", 'ю': "yu",
	'Я': "Ya", 'я': "ya",
	// Таджикские специфические символы (если используются)
	'Ғ': "G", 'ғ': "g",
	'Қ': "Q", 'қ': "q",
	'Ҳ': "H", 'ҳ': "h",
	'Ҷ': "J", 'ҷ': "j",
}

func TranslitFieldName(name string) string {
	var result strings.Builder
	hasCyrillic := false

	for _, r := range name {
		if unicode.Is(unicode.Cyrillic, r) {
			hasCyrillic = true
			if translit, ok := translitTable[r]; ok {
				result.WriteString(translit)
			} else {
				result.WriteRune(r)
			}
		} else {
			result.WriteRune(r)
		}
	}

	if hasCyrillic {
		return result.String()
	}
	return name
}
