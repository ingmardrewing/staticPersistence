package staticPersistence

import (
	"path"
	"testing"
)

func TestReadConfig(t *testing.T) {
	p := path.Join(currentDir(), "testResources")
	name := "configNew.json"

	c := ReadConfig(p, name)

	actual := c[0].Domain
	expected := "drewing.de"

	if actual != expected {
		t.Error("Expected", expected, "but got", actual)
	}
}
