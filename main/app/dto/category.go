package dto

import (
	"selfx/app/model"
)

type CategoryTree struct {
	model.Category
	Children []CategoryTree `json:"children"`
}
