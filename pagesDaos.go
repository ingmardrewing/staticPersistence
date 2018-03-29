package staticPersistence

// Page Dao
func NewPageDAO(data []byte, path, filename string) PageDao {
	return NewPageDaoReader(data, path, filename)
}
