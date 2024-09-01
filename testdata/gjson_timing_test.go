package testdata

import (
	"strconv"
	"testing"

	. "github.com/tidwall/gjson"
	"github.com/tidwall/gjson/internal/fast"
)

func BenchmarkGetComplexPath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(basicJSON, `loggy.programmers.#[tag="good"]#.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
		}
	})
	b.Run("large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(twitterLarge, `statuses.#[friends_count>100]#.id`)
		}
	})
}

func BenchmarkGetSimplePath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(basicJSON, `loggy.programmers.0.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			_ = Get(TwitterJsonMedium, `statuses.3.id`)
		}
	})
	b.Run("Large", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			x := Get(twitterLarge, `statuses.50.id`)
			if !x.Exists() {
				b.Fatal()
			}
		}
	})
}

func BenchmarkFastPath(b *testing.B) {
	opt := fast.FastPathEnable
	b.Run("normal", func(b *testing.B) {
		fast.FastPathEnable = false
		b.Run("small", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(basicJSON, `loggy.programmers.0.firstName`)
			}
		})
		b.Run("medium", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(TwitterJsonMedium, `statuses.3.id`)
			}
		})
		b.Run("Large", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Get(twitterLarge, `statuses.50.id`)
				if !x.Exists() {
					b.Fatal()
				}
			}
		})
	})
	b.Run("fast-path", func(b *testing.B) {
		fast.FastPathEnable = true
		b.Run("small", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(basicJSON, `loggy.programmers.0.firstName`)
			}
		})
		b.Run("medium", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = Get(TwitterJsonMedium, `statuses.3.id`)
			}
		})
		b.Run("Large", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := Get(twitterLarge, `statuses.50.id`)
				if !x.Exists() {
					b.Fatal()
				}
			}
		})
	})
	fast.FastPathEnable = opt
}

func TestEscapeT(t *testing.T) {
	_ = Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
	// _ = gjson.Get(TwitterJsonMedium, `statuses.#[friends_count>100]#.id`)
}

func BenchmarkParseString(b *testing.B) {
	opt := fast.FastStringEnable
	opt2 := fast.ValidStringEnable
	b.Run("normal", func(b *testing.B) {
		fast.FastStringEnable = false
		fast.ValidStringEnable = false
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	b.Run("fast-string", func(b *testing.B) {
		fast.FastStringEnable = true
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	b.Run("validate-string", func(b *testing.B) {
		fast.FastStringEnable = true
		fast.ValidStringEnable = true
		b.Run("small", func(b *testing.B) {
			var str = `"<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>"`
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
		b.Run("medium", func(b *testing.B) {
			var str = strconv.Quote(complicatedJSON)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)

			}
		})
		b.Run("large", func(b *testing.B) {
			var str = strconv.Quote(TwitterJsonMedium)
			for i := 0; i < b.N; i++ {
				_ = Parse(str)
			}
		})
	})
	fast.FastStringEnable = opt
	fast.ValidStringEnable = opt2
}
