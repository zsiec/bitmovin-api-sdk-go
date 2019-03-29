package model
import (
	"time"
)

type LiveEncodingStatsEvent struct {
	// Timestamp of the event expressed in UTC: YYYY-MM-DDThh:mm:ssZ
	Time *time.Time `json:"time,omitempty"`
	Details *LiveEncodingStatsEventDetails `json:"details,omitempty"`
}

