package model

import (
	"time"

	"gorm.io/gorm"
)

// Scope interface for db query
type Scope interface {
	Scope(*gorm.DB) *gorm.DB
}

// DateRangeOpts options to paginate results
type DateRangeOpts struct {
	Begin *time.Time `url:"from" binding:"omitempty"`
	End   *time.Time `url:"to" binding:"omitempty"`
}

// Scope implment of model.Scope
func (o DateRangeOpts) Scope(x *gorm.DB) *gorm.DB {
	if o.Begin != nil {
		x = x.Where("created_at >= ?", o.Begin)
	}
	if o.End != nil {
		x = x.Where("created_at < ?", o.End)
	}
	return x
}
