package compress 

import (
	"time"
)

type categoriesRequest struct {
	BaseModel
}

type Category struct {
	ID          int       `json:"id" `
	Name        string    `json:"name" `
	Created     time.Time `json:"created" gorm:"type:timestamp;" `
	VideoAmount int       `json:"video_amount" `
	ClientID    string    `json:"client_id" `
}

type createCategoryRequest struct {
	BaseModel
	CategoryName string `json:"category_name" validate:"nonnil,min=1" required:"true"`
}