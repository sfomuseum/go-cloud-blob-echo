package echo

import (
	"context"
	"gocloud.dev/blob"
	"io"
	_ "log"
	"net/url"
	"os"
)

func Echo(ctx context.Context, from string, to string) (int64, error) {

	var reader io.ReadCloser
	var writer io.WriteCloser

	switch from {
	case "":
		reader = os.Stdin
	default:

		u, err := url.Parse(from)

		if err != nil {
			return -1, err
		}

		path := u.Path
		u.Path = ""

		reader_bucket_uri := u.String()
		reader_bucket, err := blob.OpenBucket(ctx, reader_bucket_uri)

		if err != nil {
			return -1, err
		}

		defer reader_bucket.Close()

		reader, err = reader_bucket.NewReader(ctx, path, nil)

		if err != nil {
			return -1, err
		}

		defer reader.Close()
	}

	switch to {
	case "":

		writer = os.Stdout

	default:

		u, err := url.Parse(to)

		if err != nil {
			return -1, err
		}

		path := u.Path
		u.Path = "/"

		writer_bucket_uri := u.String()
		writer_bucket, err := blob.OpenBucket(ctx, writer_bucket_uri)

		if err != nil {
			return -1, err
		}

		defer writer_bucket.Close()

		writer, err = writer_bucket.NewWriter(ctx, path, nil)

		if err != nil {
			return -1, err
		}
	}

	n, err := io.Copy(writer, reader)

	if err != nil {
		return n, err
	}

	err = writer.Close()

	if err != nil {
		return n, err
	}

	return n, nil
}
