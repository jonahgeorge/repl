package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/siddontang/go-mysql/mysql"
	"github.com/siddontang/go-mysql/replication"
	"github.com/sirupsen/logrus"
)

var (
	log        = logrus.New()
	DBFlavor   = "mariadb"
	DBServerID uint
	DBHost     string
	DBPort     uint
	DBUser     string
	DBPassword string
	GTID       string
)

func init() {
	flag.UintVar(&DBPort, "port", 3306, "")
	flag.UintVar(&DBServerID, "server_id", 666, "")
	flag.StringVar(&DBHost, "host", "localhost", "")
	flag.StringVar(&DBUser, "user", "root", "")
	flag.StringVar(&DBPassword, "password", "root", "")
	flag.StringVar(&GTID, "gtid", "", "")
	flag.Parse()
}

func GetStringVariable(db *sql.DB, name string) (string, error) {
	log.Info("SELECT " + name)

	var gtid string
	err := db.QueryRow("SELECT " + name).Scan(&gtid)
	return gtid, err
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/", DBUser, DBPassword, DBHost, DBPort))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if GTID == "" {
		GTID, err = GetStringVariable(db, "@@gtid_current_pos")
		if err != nil {
			log.Errorf("failed to fetch variable: %s", err)
		}
	}

	gtidSet, err := mysql.ParseGTIDSet(mysql.MariaDBFlavor, GTID)
	if err != nil {
		log.Errorf("failed to parse gtid set: %s", err)
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
		log.WithField("gtid", GTID).Fatalf("failed to begin replicating: %v", err)
	}

	ctx := context.Background()
	for {
		event, err := streamer.GetEvent(ctx)
		if err != nil {
			log.Fatal(err)
		}

		log.Infof("%+v: %+v", event.Header.EventType, event.Event)
	}
}
