package faker

import (
	"math/rand"
	"time"
)

var (
	LangThai             = []int64{3585, 3654}
	LangArmenian         = []int64{1328, 1423}
	LangChinese          = []int64{19968, 40869}
	LangJapaneseKatakana = []int64{12449, 12531}
	LangJapaneseHiragana = []int64{12353, 12435}
	LangKoreanHangul     = []int64{12593, 12686}
	LangCyrillianRussian = []int64{1025, 1169}
	LangGreek            = []int64{884, 974}
)

func RandInt(start, end int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return start + rand.Int63n(end-start)
}

func RandCJK(size int, start, end int64) string {
	randRune := make([]rune, size)
	for i := range randRune {
		randRune[i] = rune(RandInt(start, end))
	}
	return string(randRune)
}
