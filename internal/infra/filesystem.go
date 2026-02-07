package infra

import (
	"finder/internal/domain"
	"io/fs"
	"path/filepath"
)

type LocalFileSystem struct{}

func (l *LocalFileSystem) FetchFiles(root string) ([]domain.FileInfo, error) {
	var files []domain.FileInfo
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		info, err := d.Info()
		if err != nil {
			return nil
		}

		files = append(files, domain.FileInfo{
			Name:    info.Name(),
			Size:    info.Size(),
			ModTime: info.ModTime(),
			Path:    path,
		})
		return nil
	})
	return files, err
}
