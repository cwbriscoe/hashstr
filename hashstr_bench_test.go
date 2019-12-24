// Copyright 2019 Christopher Briscoe.  All rights reserved. 

package hashstr

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"
	"testing"
)

func BenchmarkSprintfSha1(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%x", sum)
		}
	})
}

func BenchmarkSprintfMd5(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = fmt.Sprintf("%x", sum)
		}
	})
}

func BenchmarkSha1(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := sha1.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Sha1ToLowerString(sum)
		}
	})
}
func BenchmarkMd5(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = Md5ToLowerString(sum)
		}
	})
}

func BenchmarkHexEncoder(b *testing.B) {
	v := []byte(strings.Repeat("X", 2000))
	sum := md5.Sum(v)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_ = hex.EncodeToString(sum[:])
		}
	})
}