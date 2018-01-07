package staticPersistence

import "testing"

func TestNewPostAdder(t *testing.T) {
	dir := "./testResources/dir/"
	p := NewPostAdder(dir)
	if p.(*postAdder).dirpath != dir {
		t.Error("PostAdder not properly initiated")
	}
}

func TestPostAdder_for_selecting_the_right_image(t *testing.T) {
	p := NewPostAdder("./testResources/imagedir/")
	p.Read()

	expected := "test.png"
	actual := p.(*postAdder).imgfilename

	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPostAdder_for_selecting_the_right_md_file(t *testing.T) {
	p := NewPostAdder("./testResources/textdir/")
	p.Read()

	expected := "test.md"
	actual := p.(*postAdder).mdfilename

	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPostAdder_for_selecting_the_right_json_file(t *testing.T) {
	p := NewPostAdder("./testResources/dir/")
	p.Read()

	expected := "test.json"
	actual := p.(*postAdder).imgjsonfilename

	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPostAdder_read_md_file_contents(t *testing.T) {
	p := NewPostAdder("./testResources/textdir/")
	p.Read()

	expected := "Hello World!"
	actual := p.(*postAdder).mdinitcontent

	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}

func TestPostAdder_read_json_file_contents(t *testing.T) {
	p := NewPostAdder("./testResources/dir/")
	p.Read()

	expected := `{"hello":"world"}`
	actual := p.(*postAdder).imgjsoncontent

	if expected != actual {
		t.Error("Expected", expected, "but got", actual)
	}
}
