package staticPersistence

import "github.com/buger/jsonparser"

func FindJsonVersion(data []byte) int {
	if data == nil {
		return v1
	}
	j := new(Json)
	v := j.ReadInt(data, "version")
	return v
}

// Json
type Json struct{}

func (j *Json) ReadString(value []byte, keys ...string) string {
	v, err := jsonparser.GetString(value, keys...)
	if err != nil {
		return ""
	}
	return v
}

func (j *Json) ReadInt(value []byte, keys ...string) int {
	v, err := jsonparser.GetInt(value, keys...)
	if err != nil {
		return 0
	}
	return int(v)
}
