package domain

import "time"

type FileInfo struct {
	Name    string
	Size    int64
	ModTime time.Time
	Path    string
}
