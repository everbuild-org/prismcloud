package model

import "gorm.io/gorm"

type Namespace struct {
	gorm.Model

	Name     string
	RamLimit int64
}
