# go-cloud-blob-echo

Go package for echoing data to and from Go Cloud `blob` sources.

## Example

```
import (
       "context"
       "github.com/sfomuseum/go-cloud-blob-echo"
)

ctx := context.Background()

from := ""	// STDIN
to := "file:///tmp/test"

echo.Echo(ctx, from, to)
```

## Tools

To build binary versions of these tools run the `cli` Makefile target. For example:

```
$> make cli
go build -mod vendor -o bin/echo cmd/echo/main.go
```

### echo

```
$> bin/echo -h
Usage of ./bin/echo:
  -from string
    	A valid Go Cloud blob URI. If empty data will be read from STDIN.
  -to string
    	A valid Go Cloud blob URI. If empty data will be written to STDOUT.
```

```
$> echo 'hello world' | bin/echo 
hello world
```

```
$> echo 'hello world' | bin/echo -to 'file:///usr/local/data/test'

$> cat /usr/local/data/test
hello world
```

```
$> bin/echo -from 'file:///usr/local/data/test'
hello world
```

This packages uses the [go-cloud-s3blob](https://github.com/aaronland/go-cloud-s3blob) library to handle reads and writes to S3. `go-cloud-s3blob` is thing wrapper around the default go-cloud S3 blob opener to check for a credentials parameter (in blob URIs) and use it to assign AWS S3 session credentials.

```
$> echo 'this is a test' \

   | ./bin/echo \
   	-to 's3blob://s3-bucket/misc/test.txt?region=us-east-1&credentials=default'

$> bin/echo \
	-from 's3blob://s3-bucket/misc/test.txt?region=us-east-1&credentials=default'
this is a test
```

#### Known knowns

* By default the `cmd/echo/main.go` tool has support for the Go Cloud `fileblob://`, `s3://` and `s3blob://` URI schemes. This reflects the nature of the work at SFO Museum.
* By default the Go Cloud `blob` writers create files with restrictive permissions. The Go Cloud `blob` interfaces don't have an interface for setting permissions (by design) to doing so is left to developers to implement themselves.

## See also

* https://godoc.org/gocloud.dev/blob
* https://github.com/aaronland/go-cloud-s3blob