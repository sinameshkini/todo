package io

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryID  uint   `json:"category_id"`
	Star        uint8  `json:"star"`
	Complete    bool   `json:"complete"`
	ParentID    uint   `json:"parent_id"`
	gorm.Model
}

type TodoCategory struct {
	Name     string `json:"name"`
	ParentID uint   `json:"parent_id"`
	gorm.Model
}

func (t Todo) String() string {
	b, err := json.Marshal(t)
	if err != nil {
		return "unsupported value type"
	}
	return string(b)
}
