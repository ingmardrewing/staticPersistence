package staticPersistence

import "testing"

func TestFindVersion_returns_zero_if_no_version_defined(t *testing.T) {
	json := []byte(`{"noversiongiven":""}`)
	expected := 0

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_version_one_is_defined(t *testing.T) {
	json := []byte(`{"version":1}`)
	expected := 1

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}

func TestFindVersion_returns_one_if_new_empty_dao_is_needed(t *testing.T) {
	json := []byte{}
	expected := 1

	actual := FindJsonVersion(json)

	if expected != actual {
		t.Error("Expected", expected, ", but got", actual)
	}
}
