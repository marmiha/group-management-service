package model

import "time"

// swagger:model EntityID
//
// example: 1
type EntityID int64

// This is a common struct field amongst models that need to be persisted.
// swagger:model Entity
type Entity struct {

	// when it was created
	//
	// example: 2021-02-05T16:12:21.385747Z
	CreatedAt time.Time `json:"created_at"`

	// last time it was updated
	//
	// example: 2021-03-05T16:12:21.385747Z
	UpdatedAt time.Time `json:"updated_at"`
}