package faker

import (
	"fmt"
	"testing"
)

func TestCJK(t *testing.T) {
	fmt.Println("Random Chinese :\n ", RandCJK(10, LangChinese[0], LangChinese[1]))
	fmt.Println("Random Thai :\n ", RandCJK(10, LangThai[0], LangThai[1]))
	fmt.Println("Random Japanese Katakana :\n ", RandCJK(10, LangJapaneseKatakana[0], LangJapaneseKatakana[1]))
	fmt.Println("Random Japanese Hiragana :\n ", RandCJK(10, LangJapaneseHiragana[0], LangJapaneseHiragana[1]))
	fmt.Println("Random Korean :\n ", RandCJK(10, LangKoreanHangul[0], LangKoreanHangul[1]))
	fmt.Println("Random Russian :\n ", RandCJK(10, LangCyrillianRussian[0], LangCyrillianRussian[1]))
	fmt.Println("Random Armenian :\n ", RandCJK(10, LangArmenian[0], LangArmenian[1]))
}
