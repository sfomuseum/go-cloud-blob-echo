# gocloud-blob-echo

## Tools

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