package staticPersistence

func NewImageDto(
	title,
	w80Square,
	w185Square,
	w390Square,
	w800Square,
	w800,
	w1600,
	maxResolution string) *imageDTO {
	return &imageDTO{
		title,
		w80Square,
		w185Square,
		w390Square,
		w800Square,
		w800,
		w1600,
		maxResolution}
}

// imageDTO
type imageDTO struct {
	title         string
	w80Square     string
	w185Square    string
	w390Square    string
	w800Square    string
	w800          string
	w1600         string
	maxResolution string
}

func (i imageDTO) W80Square() string  { return i.w80Square }
func (i imageDTO) W185Square() string { return i.w185Square }
func (i imageDTO) W390Square() string { return i.w390Square }
func (i imageDTO) W800Square() string { return i.w800Square }

func (i imageDTO) W800() string  { return i.w800 }
func (i imageDTO) W1600() string { return i.w1600 }

func (i imageDTO) MaxResolution() string { return i.maxResolution }
func (i imageDTO) Title() string         { return i.title }
