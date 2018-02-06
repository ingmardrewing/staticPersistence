package staticPersistence

const (
	v0 = iota
	v1 = iota
)

func FindJsonVersion(data []byte) int {
	if data == nil {
		return v1
	}
	j := new(Json)
	v := j.ReadInt(data, "version")
	return v
}

func NewDto() *docDTO {
	return new(docDTO)
}

// Post Dao
func NewPostDAO(data []byte, path, filename string) DAO {
	var d DAO
	switch FindJsonVersion(data) {
	case v1:
		d = new(postDAOv1)
	default:
		d = new(postDAOv0)
	}
	dto := NewDto()
	dto.FsPath(path)
	dto.HtmlFilename(filename)

	d.Dto(dto)
	d.Data(data)

	return d
}

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) DAO {
	var d DAO
	switch FindJsonVersion(data) {
	case v1:
		d = new(marginalDAOv1)
	default:
		d = new(marginalDAOv0)
	}
	dto := NewDto()
	dto.FsPath(path)
	dto.HtmlFilename(filename)

	d.Dto(dto)
	d.Data(data)

	return d
}

type DAO interface {
	Dto(...DTO) DTO
	ExtractFromJson()
	FillJson() []byte
	Data([]byte)
}
