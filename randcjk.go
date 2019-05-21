package faker

import (
	"math/rand"
	"time"
)

var (
	// LangThai repreents the Thai range
	LangThai = []int64{3585, 3654}
	// LangArmenian repreents the Armenian range
	LangArmenian = []int64{1328, 1423}
	// LangChinese repreents the Chinese range
	LangChinese = []int64{19968, 40869}
	// LangJapaneseKatakana repreents the Japanese Katakana range
	LangJapaneseKatakana = []int64{12449, 12531}
	// LangJapaneseHiragana repreents the Japanese Hiragana range
	LangJapaneseHiragana = []int64{12353, 12435}
	// LangKoreanHangul repreents the Korean Hangul range
	LangKoreanHangul = []int64{12593, 12686}
	// LangCyrillianRussian repreents the Cyrillian Russian range
	LangCyrillianRussian = []int64{1025, 1169}
	// LangGreek repreents the Greek range
	LangGreek = []int64{884, 974}
)

// RandInt64 returns a random int64 between start and end.
func RandInt64(start, end int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return start + rand.Int63n(end-start)
}

// RandCJK returns a string of CJK in lenght of size between start and end range in unicode.
func RandCJK(size int, start, end int64) string {
	randRune := make([]rune, size)
	for i := range randRune {
		randRune[i] = rune(RandInt64(start, end))
	}
	return string(randRune)
}
