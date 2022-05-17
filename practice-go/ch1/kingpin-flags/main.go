package main

import (
	"context"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

// 標準のFlagは、 POSIXスタイルの短縮記号と長い引数名への対応、短縮記号をつなげた記法、必須フラグ
// のチェック、gitコマンドのようなサブコマンドを持つ複雑なケースには向いていません。
var (
	defaultLanguage = kingpin.Flag("default-language", "Default Language").String()
	generateCmd     = kingpin.Command("create-index", "Generate Index")
	inputFolder     = generateCmd.Arg("INPUT", "Input Folder").Required().ExistingDir()

	searchCmd   = kingpin.Command("search", "Search")
	inputFile   = searchCmd.Flag("input", "Input index file").Short('i').File()
	searchWords = searchCmd.Arg("WORDS", "Search words").Strings()
)

func main() {
	ctx := context.Background()

	switch kingpin.Parse() {
	case generateCmd.FullCommand():
		err := generate(ctx)
		if err != nil {
			os.Exit(1)
		}
	case searchCmd.FullCommand():
		err := search(ctx)
		if err != nil {
			os.Exit(1)
		}
	}

}

func generate(ctx context.Context) error {
	return nil
}

func search(ctx context.Context) error {
	return nil
}
