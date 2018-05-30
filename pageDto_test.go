package staticPersistence

import "testing"

func TestNewFilledDto(t *testing.T) {
	dto := NewFilledDto(42,
		"titleValue",
		"titlePlainValue",
		"thumbUrlValue",
		"imageUrlValue",
		"descriptionValue",
		"disqusIdValue",
		"createDateValue",
		"contentValue",
		"urlValue",
		"domainValue",
		"pathValue",
		"fspathValue",
		"htmlfilenameValue",
		"thumbBase64Value",
		"categoryValue",
		"microThumbValue")

	if dto.Id() != 42 {
		t.Error("Expected 42, but got ", dto.Id())
	}
	if dto.FsPath() != "fspathValue" {
		t.Error("Expected fspathValue, but got ", dto.FsPath())
	}
	if dto.HtmlFilename() != "htmlfilenameValue" {
		t.Error("Expected htmlfilenameValue, but got ", dto.HtmlFilename())
	}
	if dto.Title() != "titleValue" {
		t.Error("Expected titleValue, but got ", dto.Title())
	}
	if dto.Domain() != "domainValue" {
		t.Error("Expected domainValue, but got ", dto.Domain())
	}
	if dto.TitlePlain() != "titlePlainValue" {
		t.Error("Expected titlePlainValue, but got ", dto.TitlePlain())
	}
	if dto.ThumbUrl() != "thumbUrlValue" {
		t.Error("Expected thumbUrlValue, but got ", dto.ThumbUrl())
	}
	if dto.ImageUrl() != "imageUrlValue" {
		t.Error("Expected imageUrlValue, but got ", dto.ImageUrl())
	}
	if dto.Description() != "descriptionValue" {
		t.Error("Expected descriptionValue, but got ", dto.Description())
	}
	if dto.DisqusId() != "disqusIdValue" {
		t.Error("Expected disqusIdValue, but got ", dto.DisqusId())
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
	if dto.ThumbBase64() != "thumbBase64Value" {
		t.Error("Expected thumbBase64Value, but got ", dto.ThumbBase64())
	}
	if dto.Url() != "urlValue" {
		t.Error("Expected thumbBase64Value, but got ", dto.ThumbBase64())
	}
	if dto.PathFromDocRoot() != "pathValue" {
		t.Error("Expected pathValue, but got ", dto.PathFromDocRoot())
	}
	if dto.Category() != "categoryValue" {
		t.Error("Expected categoryValue, but got ", dto.Category())
	}
	if dto.MicroThumbUrl() != "microThumbValue" {
		t.Error("Expected microThumbValue, but got ", dto.MicroThumbUrl())
	}
}
