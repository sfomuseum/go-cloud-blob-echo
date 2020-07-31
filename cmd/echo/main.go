package main

import (
	_ "gocloud.dev/blob/fileblob"
	_ "github.com/aaronland/go-cloud-s3blob"
)

import (
	"context"
	"flag"
	"github.com/sfomuseum/gocloud-blob-echo"
	"log"
)

func main() {

	from := flag.String("from", "", "A valid Go Cloud blob URI. If empty data will be read from STDIN.")
	to := flag.String("to", "", "A valid Go Cloud blob URI. If empty data will be written to STDOUT.")

	flag.Parse()

	ctx := context.Background()

	_, err := echo.Echo(ctx, *from, *to)

	if err != nil {
		log.Fatal(err)
	}

}
