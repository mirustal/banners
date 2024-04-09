package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	"time"

	"github.com/lib/pq"
)

type JSONB map[string]string

func (j JSONB) Value() (driver.Value, error) {
    return json.Marshal(j)
}

func (j *JSONB) Scan(src interface{}) error {
    if src == nil {
        *j = nil
        return nil
    }
    switch src := src.(type) {
    case []byte:
        return json.Unmarshal(src, &j)
    default:
        return errors.New("incompatible type for JSONB")
    }
}

type Banner struct {
	ID        int             `gorm:"type:int;primary_key" json:"banner_id"`
	TagIDs    pq.Int64Array           `gorm:"type:integer[]" json:"tag_ids"`
	FeatureID int             `json:"feature_id"`
	Content   JSONB `gorm:"type:jsonb" json:"content"`
	IsActive  bool              `json:"is_active"`
	CreatedAt time.Time         `json:"created_at"`
	UpdatedAt time.Time         `json:"updated_at"`
}
