# repl

A small utility for inspecting binary log events.

## Usage

```sh
$ ./repl -h
Usage of ./repl:
  -gtid string
    	
  -host string
    	 (default "localhost")
  -password string
    	 (default "root")
  -port uint
    	 (default 3306)
  -server_id uint
    	 (default 666)
  -user string
    	 (default "root")
```

## Development

```sh
$ docker run --rm -it -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root \
	mariadb:10.1.28 --server-id=1 --log-bin --binlog_format=ROW

$ go run repl.go
```
