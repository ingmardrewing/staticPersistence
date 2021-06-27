package staticPersistence

func NewImageDto(
	title,
	w80Square,
	w100Square,
	w185Square,
	w200Square,
	w390Square,
	w400Square,
	w800Square,
	w800,
	w1600,
	maxResolution string) *imageDto {
	return &imageDto{
		title,
		w80Square,
		w100Square,
		w185Square,
		w200Square,
		w390Square,
		w400Square,
		w800Square,
		w800,
		w1600,
		maxResolution}
}

// imageDto
type imageDto struct {
	title         string
	w80Square     string
	w100Square    string
	w185Square    string
	w200Square    string
	w390Square    string
	w400Square    string
	w800Square    string
	w800          string
	w1600         string
	maxResolution string
}

func (i imageDto) W80Square() string  { return i.w80Square }
func (i imageDto) W100Square() string { return i.w100Square }
func (i imageDto) W185Square() string { return i.w185Square }
func (i imageDto) W200Square() string { return i.w200Square }
func (i imageDto) W390Square() string { return i.w390Square }
func (i imageDto) W400Square() string { return i.w400Square }
func (i imageDto) W800Square() string { return i.w800Square }

func (i imageDto) W800() string  { return i.w800 }
func (i imageDto) W1600() string { return i.w1600 }

func (i imageDto) MaxResolution() string { return i.maxResolution }
func (i imageDto) Title() string         { return i.title }
