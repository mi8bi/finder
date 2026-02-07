package main

import (
	"finder/internal/infra"
	"finder/internal/usecase"
	"flag"
	"fmt"
)

func main() {
	// フラグ定義
	pathPtr := flag.String("path", ".", "探索を開始するファイルパス")
	extPtr := flag.String("ext", "", "抽出する拡張子 (例: .jpg)")
	sizeKBPtr := flag.Int64("size", 0, "指定したKB以上のファイルのみ表示")
	flag.Parse()

	repo := &infra.LocalFileSystem{}
	finder := usecase.NewFileFinder(repo)

	results, _ := finder.Execute(*pathPtr, *extPtr, *sizeKBPtr)

	fmt.Printf("Searching in: %s (Ext: %s, MinSize: %d KB)\n", *pathPtr, *extPtr, *sizeKBPtr)
	fmt.Println("Name | Size(Bytes) | ModTime")
	fmt.Println("---------------------------------")
	for _, f := range results {
		fmt.Printf("%s | %.1fKB | %s\n", f.Name, float64(f.Size)/1024.0, f.ModTime)
	}
}
