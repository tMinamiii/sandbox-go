package main

import (
	"log"
	"strings"
)

func badConcat() {
	src := []string{"Back", "To", "The", "Future", "Part", "III"}

	var title string

	for i, word := range src {
		if i != 0 {
			title += " "
		}
		title += word
	}
	log.Println(title)
}

func goodConcat() {
	src := []string{"Back", "To", "The", "Future", "Part", "III"}
	var builder strings.Builder
	builder.Grow(100) // 最大100文字以下と仮定できる場合
	for i, word := range src {
		if i != 0 {
			builder.WriteByte(' ')
		}
		builder.WriteString(word)
	}
	log.Println(builder.String())
}

func good2() {
	// また数が少なく、結合する文字列の数が少なければ、1つの式でまとめて書いてしまうのも効率は悪くありません。
	title := "abc"
	displayTitle := "1990年7月6日公開 - " + title + " - ロバート・ゼメキス"
	log.Println(displayTitle)
}
