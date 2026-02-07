package usecase

import (
	"finder/internal/domain"
	"path/filepath"
	"strings"
)

// Repository 外部(ファイルシステム)へのインタフェース
type FileRepository interface {
	FetchFiles(root string) ([]domain.FileInfo, error)
}

type FileFinder struct {
	repo FileRepository
}

func NewFileFinder(r FileRepository) *FileFinder {
	return &FileFinder{repo: r}
}

// Execute フィルタリングのロジック
func (f *FileFinder) Execute(root, ext string, minSizeKB int64) ([]domain.FileInfo, error) {
	files, err := f.repo.FetchFiles(root)
	if err != nil {
		return nil, err
	}

	var result []domain.FileInfo
	minSizeBytes := minSizeKB * 1024

	for _, file := range files {
		// 拡張子チェック
		if ext != "" && !strings.EqualFold(filepath.Ext(file.Path), ext) {
			continue
		}
		// サイズチェック
		if file.Size < minSizeBytes {
			continue
		}
		result = append(result, file)
	}
	return result, nil
}
