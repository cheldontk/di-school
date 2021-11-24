package model

import (
	"time"

	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(false)
}

type Base struct {
	ID        string    `json: "id" valid:"required, uuid"`
	CreatedAt time.Time `json: "created_at" valid:"-"`
	UpdatedAt time.Time `json: "updated_at" valid:"-"`
	DeletedAt time.Time `json: "deleted_at" valid:"-"`
}
