package staticPersistence

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) PageDao {
	d := new(abstractPageDao)

	dto := NewFilledDto(0,
		"", "", "", "", "",
		"", "", "", "", "",
		path, filename, "", "")

	d.Dto(dto)
	d.Data(data)

	return d
}
