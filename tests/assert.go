package tests

import (
	"reflect"
	"regexp"
	"strings"
	"testing"
	"time"
)

// AssertMatches 文字列が正規表現にマッチするか判定する
func AssertMatches(t *testing.T, title string, actual, pattern string) {
	matched, err := regexp.MatchString(pattern, actual)
	if err != nil {
		t.Errorf("%s: illegal pattern:%v", title, pattern)
	}
	if !matched {
		t.Errorf("%s: not matched, actual: `%v`, pattern: `%v`", title, actual, pattern)
	}
}

// AssertEquals 実値と期待値が同値か判定を実施する
func AssertEquals(t *testing.T, title string, actual, expected interface{}) {
	if actual != expected {
		t.Errorf("%s: unexpected, actual: `%v`, expected: `%v`", title, actual, expected)
	}
}

// AssertStringExist 実値が存在するか判定を実施する
func AssertStringExist(t *testing.T, title string, actual string) {
	if actual == "" {
		t.Errorf("%s: unexpected, actual: `%v`", title, actual)
	}
}

// AssertTimeEquals time.Time型の実値と期待値が同値か判定を実施する
func AssertTimeEquals(t *testing.T, title string, actual, expected time.Time) {
	if !actual.Equal(expected) {
		t.Errorf("%s: unexpected, actual: `%s`, expected: `%s`", title, actual, expected)
	}
}

// AssertDeepEquals 実値と期待値が同値か判定を実施する
func AssertDeepEquals(t *testing.T, title string, actual, expected interface{}) {
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("%s: unexpected, actual: `%v`, expected: `%v`", title, actual, expected)
	}
}

// AssertStringContains 実値文字列が期待値文字列を含むかどうか判定する
func AssertStringContains(t *testing.T, title string, actual, expected string) {
	if !strings.Contains(actual, expected) {
		t.Errorf("%s: not contained, actual: `%s`, expected: `%s`", title, actual, expected)
	}
}
