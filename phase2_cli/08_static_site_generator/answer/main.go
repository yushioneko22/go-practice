package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Post struct {
	Title   string
	Content string
}

const htmlTemplate = `
<!DOCTYPE html>
<html>
<head><title>{{.Title}}</title></head>
<body>
	<h1>{{.Title}}</h1>
	<p>{{.Content}}</p>
</body>
</html>
`

func main() {
	// 出力先フォルダの作成
	os.MkdirAll("dist", 0755)

	// 仮の入力データ (本来はファイルを読み込む)
	files := map[string]string{
		"hello.md": "これは最初の投稿です。",
		"go.md":    "Go言語はパワフルなツールです。",
	}

	tmpl, _ := template.New("post").Parse(htmlTemplate)

	for filename, body := range files {
		title := strings.TrimSuffix(filename, ".md")
		outName := filepath.Join("dist", title+".html")

		f, _ := os.Create(outName)
		defer f.Close()

		p := Post{Title: title, Content: body}
		tmpl.Execute(f, p)
		fmt.Println("Generated:", outName)
	}
}
