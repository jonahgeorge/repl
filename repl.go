package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
)

var (
	DBFlavor   string
	DBServerID uint
	DBHost     string
	DBPort     uint
	DBUser     string
	DBPassword string
	GTID       string
	output     string
)

func init() {
	flag.UintVar(&DBPort, "port", 3306, "")
	flag.UintVar(&DBServerID, "server_id", 666, "")
	flag.StringVar(&DBHost, "host", "localhost", "")
	flag.StringVar(&DBUser, "user", "root", "")
	flag.StringVar(&DBPassword, "password", "root", "")
	flag.StringVar(&GTID, "gtid", "", "")
	flag.StringVar(&DBFlavor, "flavor", "mariadb", "mariadb or mysql")
	flag.StringVar(&output, "output", "pretty", "Output format. json or pretty")
	flag.Parse()
}

var logger = log.New(os.Stderr, os.Args[0], log.Ldate|log.Ltime)

func GetStringVariable(db *sql.DB, name string) (string, error) {
	logger.Println("SELECT " + name)

	var gtid string
	err := db.QueryRow("SELECT " + name).Scan(&gtid)
	return gtid, err
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", DBUser, DBPassword, DBHost, DBPort))
	if err != nil {
		logger.Fatalf("failed to open db connection: %s", err)
	}
	defer db.Close()

	if GTID == "" {
		GTID, err = GetStringVariable(db, "@@gtid_current_pos")
		if err != nil {
			logger.Fatalf("failed to fetch variable: %s", err)
		}
	}

	gtidSet, err := mysql.ParseGTIDSet(DBFlavor, GTID)
	if err != nil {
		logger.Fatalf("failed to parse gtid set: %s", err)
	}

	syncer := replication.NewBinlogSyncer(replication.BinlogSyncerConfig{
		ServerID: uint32(DBServerID),
		Flavor:   DBFlavor,
		Host:     DBHost,
		Port:     uint16(DBPort),
		User:     DBUser,
		Password: DBPassword,
	})
	defer syncer.Close()

	streamer, err := syncer.StartSyncGTID(gtidSet)
	if err != nil {
		logger.Fatalf("failed to begin replicating: %v", err)
	}

	for {
		event, err := streamer.GetEvent(context.Background())
		if err != nil {
			logger.Fatalf("failed to fetch binlog event: %s", err)
		}

		switch output {
		case "json":
			fmt.Printf("%+v: %+v", event.Header.EventType, event.Event)
		default:
			os.Stdout.WriteString(event.Header.EventType.String() + "\n")
			event.Event.Dump(os.Stdout)
		}
	}
}
