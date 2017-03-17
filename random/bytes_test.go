package random_test

import (
	. "github.com/firelyu/go_sample/random"
	"strconv"
	"testing"
)

var lengthList = []int{128, 256, 512, 1024, 2048}

func benchmarkRand(fn func(int) string, length int, b *testing.B) {
	for i := 0; i < b.N; i++ {
		fn(length)
	}
}

func benchmarkRandGroup(length int, b *testing.B) {
	b.Run("RandStringRunes"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringRunes, length, b)
	})
	b.Run("RandStringBytes"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringBytes, length, b)
	})
	b.Run("RandStringBytesReminder"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringBytesReminder, length, b)
	})
	b.Run("RandStringBytesMask"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringBytesMask, length, b)
	})
	b.Run("RandStringBytesMaskImpr"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringBytesMaskImpr, length, b)
	})
	b.Run("RandStringBytesMaskImprSrc"+strconv.Itoa(length), func(b *testing.B) {
		benchmarkRand(RandStringBytesMaskImprSrc, length, b)
	})
}

func BenchmarkMain(b *testing.B) {
	for _, len := range lengthList {
		benchmarkRandGroup(len, b)
	}
}

//func BenchmarkRandStringRunes(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringRunes, length, b)
//		})
//	}
//
//}
//
//func BenchmarkRandStringBytes(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringBytes, length, b)
//		})
//	}
//}
//
//func BenchmarkRandStringBytesReminder(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringBytesReminder, length, b)
//		})
//	}
//}
//
//func BenchmarkRandStringBytesMask(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringBytesMask, length, b)
//		})
//	}
//}
//
//func BenchmarkRandStringBytesMaskImpr(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringBytesMaskImpr, length, b)
//		})
//	}
//}
//
//func BenchmarkRandStringBytesMaskImprSrc(b *testing.B)  {
//	for _, length := range lengthList {
//		b.Run(strconv.Itoa(length), func(b *testing.B) {
//			benchmarkRand(RandStringBytesMaskImprSrc, length, b)
//		})
//	}
//}
