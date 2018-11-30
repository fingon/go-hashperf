/*
 * Author: Markus Stenberg <fingon@iki.fi>
 *
 * Copyright (c) 2018 Markus Stenberg
 *
 * Created:       Fri Nov 30 10:02:25 2018 mstenber
 * Last modified: Fri Nov 30 10:09:48 2018 mstenber
 * Edit time:     5 min
 *
 */

package hashperf

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"testing"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/blake2s"
)

func BenchmarkSha256(b *testing.B) {
	for _, s := range []int{40, 200, 500, 1200} {
		data := bytes.Repeat([]byte{0}, s)
		b.Run(fmt.Sprintf("sha256-%d", s), func(b *testing.B) {
			b.SetBytes(int64(s))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				sha256.Sum256(data)
			}
		})

	}
}

func BenchmarkBlake2B(b *testing.B) {
	for _, s := range []int{40, 200, 500, 1200} {
		data := bytes.Repeat([]byte{0}, s)
		b.Run(fmt.Sprintf("b2b-%d", s), func(b *testing.B) {
			b.SetBytes(int64(s))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				blake2b.Sum256(data)
			}
		})

	}
}

func BenchmarkBlake2S(b *testing.B) {
	for _, s := range []int{40, 200, 500, 1200} {
		data := bytes.Repeat([]byte{0}, s)
		b.Run(fmt.Sprintf("b2s-%d", s), func(b *testing.B) {
			b.SetBytes(int64(s))
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				blake2s.Sum256(data)
			}
		})

	}
}
