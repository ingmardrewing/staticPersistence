package staticPersistence

// marginalDAO
func NewMarginalDAO(data []byte, path, filename string) PageDao {
	return NewPageDaoReader(data, path, filename)
}
