package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/crowi/go-crowi"
	"github.com/k0kubun/pp"
)

func main() {
	client, err := crowi.NewClient("http://localhost:3000", os.Getenv("CROWI_ACCESS_TOKEN"))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var (
		path = fmt.Sprintf("/user/%s/go-crowi-test-%d", os.Getenv("USER"), time.Now().UnixNano())
		body = "# this is a sample\n\ntest"
	)

	res1, err := client.PagesCreate(ctx, path, body)
	if err != nil {
		panic(err)
	}

	res2, err := client.AttachmentsAdd(ctx, res1.Page.ID, "_example/attachments/sample.png")
	if err != nil {
		panic(err)
	}
	pp.Println(res2)

	res3, err := client.AttachmentsList(ctx, res1.Page.ID)
	if err != nil {
		panic(err)
	}

	// body = fmt.Sprintf("![](%s)", res3.URL)
	_, err = client.PagesUpdate(ctx, res1.Page.ID, body)
	if err != nil {
		panic(err)
	}

	pp.Println(res3)
}