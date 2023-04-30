package models

import (
    "time"
)

type CreateProduct struct {
    Name          string    `json:"name" binding:"required"`
    Slug          string    `json:"slug" binding:"required"`
    Image         string    `json:"image" binding:"required"`
    Images        []string  `json:"images"`
    Brand         string    `json:"brand" binding:"required"`
    Category      string    `json:"category" binding:"required"`
    Description   string    `json:"description" binding:"required"`
    Price         float64   `json:"price" binding:"required"`
    CountInStock  int       `json:"countInStock" binding:"required"`
    Reviews       []Review  `json:"reviews"`
    CreatedAt     time.Time `json:"createdAt"`
    UpdatedAt     time.Time `json:"updatedAt"`
}
