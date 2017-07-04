package uuider

import "github.com/google/uuid"

// V4Producer Version4-UUIDを生成するProducer
type V4Producer struct{}

// New V4Producerを利用してVersion４−UUIDを生成する
func (v *V4Producer) New() string {
	return uuid.New().String()
}
