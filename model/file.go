package model

import "time"

type File struct {
	Name         string    `json:"name"`
	Path         string    `json:"path"`
	Size         int64     `json:"size"`
	LastModified time.Time `json:"lastModified" type:"timestamp" timestampFormat:"iso8601"`
}
