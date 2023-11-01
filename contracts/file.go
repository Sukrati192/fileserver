package contracts

import "time"

type File struct {
	FileID    string    `json:"fileID"`
	Size      string    `json:"size"`
	Timestamp time.Time `json:"timestamp"`
}
