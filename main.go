package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/A11Might/lark_docx_md"
	lark "github.com/larksuite/oapi-sdk-go/v3"
)

func main() {
	processor := lark_docx_md.NewDocxMarkdownProcessor(
		lark.NewClient(os.Getenv("APP_ID"), os.Getenv("APP_SECRET")),
		os.Getenv("DOCUMENT_TYPE"), os.Getenv("DOCUMENT_ID"),
		lark_docx_md.DownloadStatic("static", "static"),
		lark_docx_md.UseGhCalloutStyle(),
	)
	md, err := processor.DocxMarkdown(context.Background())
	if err != nil {
		log.Fatal(err.Error())
	}
	f, err := os.Create("README.md")
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()
	if _, err := f.WriteString(md); err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("main run done")
}
