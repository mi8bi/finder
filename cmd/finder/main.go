package main

import (
	"finder/internal/fileio"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// file path
	path := flag.String("path", "", "ファイルパスを指定")
	// ext
	ext := flag.String("ext", "", "拡張子でフィルタ")
	// size
	size := flag.Int64("size", 1000, "サイズ以上のファイルのみ")
	// パースを実行
	flag.Parse()

	// path check
	if _, err := fileio.IsDir(*path); err != nil {
		flag.Usage()
		os.Exit(1)
	}
	rootPath, err := filepath.Abs(*path)
	if err != nil {
		flag.Usage()
		os.Exit(1)
	}

	filepath.Walk(rootPath, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return fmt.Errorf("failed filepath.Walk: %v", err)
		}
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == *ext && info.Size() >= *size {
			fmt.Fprintf(
				os.Stdout,
				"%s | %d | %s\n",
				info.Name(),
				info.Size(),
				info.ModTime().Format("2006/1/2 15:04:05"),
			)
		}
		return nil
	})
}
