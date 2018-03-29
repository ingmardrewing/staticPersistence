package staticPersistence

// narrativeDAO
func NewNarrativeDAO(data []byte, path, filename string) PageDao {
	return NewPageDaoReader(data, path, filename)
}
