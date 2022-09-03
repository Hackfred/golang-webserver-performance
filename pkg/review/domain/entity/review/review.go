package review

import "time"

type ReviewedEntityID string

type Review struct {
	ID               string
	ReviewedEntityID ReviewedEntityID
	AuthorID         string
	Text             string
	Stars            uint8
	Created          time.Time
	Updated          *time.Time
	Payload          [10]uint64
}
