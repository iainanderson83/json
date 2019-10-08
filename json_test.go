package json

import (
	"testing"

	"github.com/pendo-io/jsonparser"
	"github.com/tidwall/gjson"
)

func BenchmarkJsonParserSmall(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(smallFixture, "uuid")
		jsonparser.GetInt(smallFixture, "tz")
		jsonparser.Get(smallFixture, "ua")
		jsonparser.GetInt(smallFixture, "st")
	}
}

func BenchmarkJsonParserMedium(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		jsonparser.Get(mediumFixture, "person", "name", "fullName")
		jsonparser.GetInt(mediumFixture, "person", "github", "followers")
		jsonparser.Get(mediumFixture, "company")

		jsonparser.ArrayEach(mediumFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "url")
		}, "person", "gravatar", "avatars")
	}
}

func BenchmarkJsonParserLarge(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.Get(value, "username")
		}, "users")

		jsonparser.ArrayEach(largeFixture, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
			jsonparser.GetInt(value, "id")
			jsonparser.Get(value, "slug")
		}, "topics", "topics")
	}
}

func BenchmarkGJsonSmall(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.GetBytes(smallFixture, "uuid")
		gjson.GetBytes(smallFixture, "tz").Int()
		gjson.GetBytes(smallFixture, "ua")
		gjson.GetBytes(smallFixture, "st").Int()
	}
}

func BenchmarkGJsonSmallString(b *testing.B) {
	s := string(smallFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.Get(s, "uuid")
		gjson.Get(s, "tz").Int()
		gjson.Get(s, "ua")
		gjson.Get(s, "st").Int()
	}
}

func BenchmarkGJsonSmallParse(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := gjson.ParseBytes(smallFixture)

		r.Get("uuid")
		r.Get("tz").Int()
		r.Get("ua")
		r.Get("st").Int()
	}
}

func BenchmarkGJsonSmallParseString(b *testing.B) {
	s := string(smallFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := gjson.Parse(s)

		r.Get("uuid")
		r.Get("tz").Int()
		r.Get("ua")
		r.Get("st").Int()
	}
}

func BenchmarkGJsonMedium(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.GetBytes(mediumFixture, "person.name.fullName")
		gjson.GetBytes(mediumFixture, "person.github.followers").Int()
		gjson.GetBytes(mediumFixture, "company")

		gjson.GetBytes(mediumFixture, "person.gravatar.avatars").ForEach(func(key, value gjson.Result) bool {
			value.Value()
			return true
		})
	}
}

func BenchmarkGJsonMediumString(b *testing.B) {
	s := string(mediumFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.Get(s, "person.name.fullName")
		gjson.Get(s, "person.github.followers").Int()
		gjson.Get(s, "company")

		gjson.Get(s, "person.gravatar.avatars").ForEach(func(key, value gjson.Result) bool {
			value.Value()
			return true
		})
	}
}

func BenchmarkGJsonMediumParseString(b *testing.B) {
	s := string(mediumFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := gjson.Parse(s)

		r.Get("person.name.fullName")
		r.Get("person.github.followers").Int()
		r.Get("company")

		r.Get("person.gravatar.avatars").ForEach(func(key, value gjson.Result) bool {
			_ = value
			return true
		})
	}
}

func BenchmarkGJsonLarge(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.GetBytes(largeFixture, "users.#.username").ForEach(func(key, value gjson.Result) bool {
			value.Value()
			return true
		})

		gjson.GetBytes(largeFixture, "topics.topics").ForEach(func(key, value gjson.Result) bool {
			value.Get("id")
			value.Get("slug").Int()
			return true
		})
	}
}

func BenchmarkGJsonLargeString(b *testing.B) {
	s := string(largeFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.Get(s, "users.#.username").ForEach(func(key, value gjson.Result) bool {
			value.Value()
			return true
		})

		gjson.Get(s, "topics.topics").ForEach(func(key, value gjson.Result) bool {
			value.Get("id")
			value.Get("slug").Int()
			return true
		})
	}
}

func BenchmarkGJsonLargeStringNoPath(b *testing.B) {
	s := string(largeFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		gjson.Get(s, "users").ForEach(func(key, value gjson.Result) bool {
			value.Get("username")
			return true
		})

		gjson.Get(s, "topics.topics").ForEach(func(key, value gjson.Result) bool {
			value.Get("id")
			value.Get("slug").Int()
			return true
		})
	}
}

func BenchmarkGJsonLargeParseString(b *testing.B) {
	s := string(largeFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := gjson.Parse(s)

		r.Get("users.#.username").ForEach(func(key, value gjson.Result) bool {
			_ = value
			return true
		})

		r.Get("topics.topics").ForEach(func(key, value gjson.Result) bool {
			value.Get("id")
			value.Get("slug").Int()
			return true
		})
	}
}

func BenchmarkGJsonLargeParseStringNoPath(b *testing.B) {
	s := string(largeFixture)
	b.ResetTimer()

	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		r := gjson.Parse(s)

		r.Get("users").ForEach(func(key, value gjson.Result) bool {
			value.Get("username")
			return true
		})

		r.Get("topics.topics").ForEach(func(key, value gjson.Result) bool {
			value.Get("id")
			value.Get("slug").Int()
			return true
		})
	}
}