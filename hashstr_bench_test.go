// Copyright 2019 Christopher Briscoe.  All rights reserved. 
package hashstr

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSprintfSha1Upper(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%X", sum)
		}
	})
}

func BenchmarkSprintfSha1Lower(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%x", sum)
		}
	})
}

func BenchmarkSprintfMd5Upper(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%X", sum)
		}
	})
}

func BenchmarkSprintfMd5Lower(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%x", sum)
		}
	})
}

func BenchmarkSha1ToUpperString(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Sha1ToUpperString(sum)
		}
	})
}

func BenchmarkSha1ToLowerString(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Sha1ToLowerString(sum)
		}
	})
}

func BenchmarkMd5ToUpperString(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Md5ToUpperString(sum)
		}
	})
}

func BenchmarkMd5ToLowerString(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Md5ToLowerString(sum)
		}
	})
}