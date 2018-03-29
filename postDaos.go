package staticPersistence

// Post DAOs
func NewPostDAO(data []byte, path, filename string) PageDao {
	return NewPageDaoReader(data, path, filename)
}
