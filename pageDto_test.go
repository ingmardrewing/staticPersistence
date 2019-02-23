package staticPersistence

import (
	"testing"

	"github.com/ingmardrewing/staticIntf"
)

func TestNewFilledDto(t *testing.T) {
	dto := NewFilledDto(
		"titleValue",
		"descriptionValue",
		"contentValue",
		"categoryValue",
		"createDateValue",
		"pathValue",
		"htmlfilenameValue",
		[]string{},
		[]staticIntf.Image{})

	if dto.PathFromDocRoot() != "pathValue" {
		t.Error("Expected pathValue, but got ", dto.PathFromDocRoot())
	}
	if dto.Filename() != "htmlfilenameValue" {
		t.Error("Expected htmlfilenameValue, but got ", dto.Filename())
	}
	if dto.Title() != "titleValue" {
		t.Error("Expected titleValue, but got ", dto.Title())
	}
	if dto.Description() != "descriptionValue" {
		t.Error("Expected descriptionValue, but got ", dto.Description())
	}
	if dto.CreateDate() != "createDateValue" {
		t.Error("Expected createDateValue, but got ", dto.CreateDate())
	}
	if dto.Content() != "contentValue" {
		t.Error("Expected contentValue, but got ", dto.Content())
	}
	if dto.Category() != "categoryValue" {
		t.Error("Expected categoryValue, but got ", dto.Category())
	}
}
