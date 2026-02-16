package dlms_tests

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/Gurux/gxdlms-go/types"
	"golang.org/x/text/language"
)

func CurrentLanguage() language.Tag {
	langEnv := os.Getenv("LANG")
	if langEnv == "" {
		return language.AmericanEnglish
	}
	langEnv = strings.Split(langEnv, ".")[0]
	tag, err := language.Parse(langEnv)
	if err != nil {
		return language.AmericanEnglish
	}
	return tag
}

func isSupportedLanguage(tag language.Tag) bool {
	supportedLanguages := []language.Tag{
		language.AmericanEnglish,
		language.Finnish,
		language.German,
		language.French,
		language.Spanish,
		language.Italian,
		language.Swedish,
		language.Hindi, //India
		language.Estonian,
	}
	for _, supported := range supportedLanguages {
		if tag == supported {
			return true
		}
	}
	return false
}

func Test_DateTimeFromStringAll(t *testing.T) {
	expected := time.Date(2024, 12, 31, 15, 4, 5, 0, time.Local)
	for _, base := range language.Supported.BaseLanguages() {
		tag := language.Make(base.String())
		if !isSupportedLanguage(tag) {
			continue
		}
		ret := types.NewGXDateTimeFromTime(expected)
		str := ret.ToString(&tag, true)
		dt, err := types.NewGXDateTimeFromString(str, &tag)
		if err != nil {
			t.Errorf("Test_DateTimeFromStringAll failed for language %s: %v", tag, err)
			continue
		}
		if dt.Value != expected {
			t.Errorf("Test_DateTimeFromStringGlobal failed for language %s. Expected: %s, Actual: %s", tag, expected, dt.Value)
		}
		t.Logf("Test_DateTimeFromStringGlobal passed for language %s.", tag)
	}
}

func Test_DateTimeFromString(t *testing.T) {
	expected := "31.12.2024 15.04.05"
	dt, err := types.NewGXDateTimeFromString(expected, &language.Finnish)
	if err != nil {
		t.Errorf("Error parsing date: %v", err)
	}
	actual := dt.ToString(&language.Finnish, true)
	if actual != expected {
		t.Errorf("DateTimeFromString failed. Expected: %s, Actual: %s", expected, actual)
	}
}

func Test_DateFromString(t *testing.T) {
	expected := "31.12.2024"
	dt, err := types.NewGXDateFromString(expected, &language.Finnish)
	if err != nil {
		t.Errorf("Error parsing date: %v", err)
	}
	actual := dt.ToString(&language.Finnish, false)
	if actual != expected {
		t.Errorf("DateTimeFromString failed. Expected: %s, Actual: %s", expected, actual)
	}
}
