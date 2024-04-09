package models

import "time"

type Banner struct {
    ID         int64                  `json:"banner_id"`
    TagIDs     []int64                `json:"tag_ids"`
    FeatureID  int64                  `json:"feature_id"`
    Content    map[string]interface{} `json:"content"`
    IsActive   bool                   `json:"is_active"`
    CreatedAt  time.Time              `json:"created_at"`
    UpdatedAt  time.Time              `json:"updated_at"`
}
