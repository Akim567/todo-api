package storage

import (
	"encoding/json"
	"io"
)

// Serializer описывает поведение сериализации и десериализации
type Serializer interface {
	Serialize(data any, w io.Writer) error
	Deserialize(r io.Reader, v any) error
}

// JSONSerializer — реализация Serializer для формата JSON
type JSONSerializer struct{}

func (s *JSONSerializer) Serialize(data any, w io.Writer) error {
	return json.NewEncoder(w).Encode(data)
}

func (s *JSONSerializer) Deserialize(r io.Reader, v any) error {
	return json.NewDecoder(r).Decode(v)
}

// NewJSONSerializer возвращает новый экземпляр JSON-сериализатора
func NewJSONSerializer() *JSONSerializer {
	return &JSONSerializer{}
}
