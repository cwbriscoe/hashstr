// Copyright 2019 Christopher Briscoe.  All rights reserved.

package hashstr

import (
	"crypto/md5"
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestSha1ToUpperString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(1000)
		ch := byte(rand.Intn(255))
		val := []byte(strings.Repeat(string(ch), sz))
		sum := sha1.Sum(val)
		hash1 := Sha1ToUpperString(sum)
		hash2 := fmt.Sprintf("%X", sum)
		if hash1 != hash2 {
			t.Errorf("Expected hash string of %s instead of %s.", hash2, hash1)
		}
	}
}

func TestSha1ToLowerString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(1000)
		ch := byte(rand.Intn(255))
		val := []byte(strings.Repeat(string(ch), sz))
		sum := sha1.Sum(val)
		hash1 := Sha1ToLowerString(sum)
		hash2 := fmt.Sprintf("%x", sum)
		if hash1 != hash2 {
			t.Errorf("Expected hash string of %s instead of %s.", hash2, hash1)
		}
	}
}

func TestMd5ToUpperString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(1000)
		ch := byte(rand.Intn(255))
		val := []byte(strings.Repeat(string(ch), sz))
		sum := md5.Sum(val)
		hash1 := Md5ToUpperString(sum)
		hash2 := fmt.Sprintf("%X", sum)
		if hash1 != hash2 {
			t.Errorf("Expected hash string of %s instead of %s.", hash2, hash1)
		}
	}
}

func TestMd5ToLowerString(t *testing.T) {
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(1000)
		ch := byte(rand.Intn(255))
		val := []byte(strings.Repeat(string(ch), sz))
		sum := md5.Sum(val)
		hash1 := Md5ToLowerString(sum)
		hash2 := fmt.Sprintf("%x", sum)
		if hash1 != hash2 {
			t.Errorf("Expected hash string of %s instead of %s.", hash2, hash1)
		}
	}
}

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