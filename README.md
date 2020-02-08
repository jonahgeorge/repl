# repl

A small utility for inspecting MySQL/MariaDB binary log streams.

## Usage

```sh
$ ./repl -h
Usage of ./repl:
  -flavor string
    	mariadb or mysql (default "mariadb")
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
```
$ repl -gtid 0-1-7102
[2020/02/08 15:08:59] [info] binlogsyncer.go:141 create BinlogSyncer with config {666 mariadb localhost 3306 root    false false <nil> false UTC false 0 0s 0s 0 false 0}
[2020/02/08 15:08:59] [info] binlogsyncer.go:380 begin to sync binlog from GTID set 0-1-7102
[2020/02/08 15:08:59] [info] binlogsyncer.go:211 register slave for master server localhost:3306
[2020/02/08 15:08:59] [info] binlogsyncer.go:731 rotate to (mysqld-bin.000002, 4)
INFO[0000] RotateEvent: &{Position:4 NextLogName:[109 121 115 113 108 100 45 98 105 110 46 48 48 48 48 48 50]} 
INFO[0000] FormatDescriptionEvent: &{Version:4 ServerVersion:[49 48 46 49 46 50 56 45 77 97 114 105 97 68 66 45 49 126 106 101 115 115 105 101 0 108 111 103 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0] CreateTimestamp:1581024620 EventHeaderLength:19 EventTypeHeaderLengths:[56 13 0 8 0 18 0 4 4 4 4 18 0 0 221 0 4 26 8 0 0 0 8 8 8 2 0 0 0 10 10 10 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 4 19 4 0] ChecksumAlgorithm:0} 
INFO[0000] MariadbGTIDListEvent: &{GTIDs:[{DomainID:0 ServerID:1 SequenceNumber:7102}]} 
INFO[0000] MariadbBinLogCheckPointEvent: &{Info:[17 0 0 0 109 121 115 113 108 100 45 98 105 110 46 48 48 48 48 48 50]} 
INFO[0000] MariadbGTIDEvent: &{GTID:{DomainID:0 ServerID:1 SequenceNumber:7103}} 
INFO[0000] QueryEvent: &{SlaveProxyID:6 ExecutionTime:0 ErrorCode:0 StatusVars:[0 0 0 0 0 1 0 0 0 80 0 0 0 0 6 3 115 116 100 4 33 0 33 0 8 0] Schema:[116 101 115 116] Query:[99 114 101 97 116 101 32 100 97 116 97 98 97 115 101 32 116 101 115 116] GSet:0-1-7103} 
INFO[0000] MariadbGTIDEvent: &{GTID:{DomainID:0 ServerID:1 SequenceNumber:7104}} 
INFO[0000] QueryEvent: &{SlaveProxyID:6 ExecutionTime:0 ErrorCode:0 StatusVars:[0 0 0 0 0 1 0 0 0 80 0 0 0 0 6 3 115 116 100 4 33 0 33 0 8 0] Schema:[116 101 115 116] Query:[99 114 101 97 116 101 32 116 97 98 108 101 32 117 115 101 114 115 32 40 105 100 32 105 110 116 32 97 117 116 111 95 105 110 99 114 101 109 101 110 116 32 112 114 105 109 97 114 121 32 107 101 121 44 32 110 97 109 101 32 116 101 120 116 41] GSet:0-1-7104} 
INFO[0000] MariadbGTIDEvent: &{GTID:{DomainID:0 ServerID:1 SequenceNumber:7105}} 
INFO[0000] TableMapEvent: &{tableIDSize:6 TableID:18 Flags:1 Schema:[116 101 115 116] Table:[117 115 101 114 115] ColumnCount:2 ColumnType:[3 252] ColumnMeta:[0 2] NullBitmap:[2]} 
INFO[0000] WriteRowsEventV1: &{Version:1 tableIDSize:6 tables:map[18:0xc0001160a0] needBitmap2:false Table:0xc0001160a0 TableID:18 Flags:1 ExtraData:[] ColumnCount:2 ColumnBitmap1:[255] ColumnBitmap2:[] Rows:[[1 [74 111 110 97 104]]] parseTime:false timestampStringLocation:<nil> useDecimal:false ignoreJSONDecodeErr:false} 
INFO[0000] XIDEvent: &{XID:22 GSet:0-1-7105}
```

## Development

```sh
$ docker run --rm -it -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root \
	mariadb:10.1.28 --server-id=1 --log-bin --binlog_format=ROW

$ go run repl.go
```
