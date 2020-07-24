package tests

import (
	"strconv"
	"testing"

	"github.com/gcp-kit/line-bot-boilerplate-go/util"
)

func TestRandomString(t *testing.T) {
	length := 10
	randomString := util.RandomString(length)

	AssertEquals(t, "generate a random string", len(randomString), 10)
}

func TestComprehension(t *testing.T) {
	slice := make([]string, 0)

	for i := 0; i < 100; i++ {
		slice = append(slice, strconv.Itoa(i))
		if (i % 10) == 0 {
			slice = append(slice, "")
		}
	}

	AssertEquals(t, "slice size", len(slice), 110)

	slice = util.Comprehension(slice)

	AssertEquals(t, "remove blank strings in slices", len(slice), 100)
}

func TestBase64(t *testing.T) {
	base := "Test Base64"
	data := "VGVzdCBCYXNlNjQ="

	t.Run("B64Encode", func(t *testing.T) {
		encode := util.B64Encode(base)
		AssertEquals(t, "util.B64Encode", encode, data)
	})

	t.Run("B64Decode", func(t *testing.T) {
		decode := util.B64Decode(data)
		AssertEquals(t, "util.B64Decode", decode, base)
	})
}
