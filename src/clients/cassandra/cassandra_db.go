package cassandra

import (
	"github.com/acargorkem/ecommerce_oauth-api/src/utils/config"
	"github.com/gocql/gocql"
)

var (
	session *gocql.Session
)

func init() {
	cluster := gocql.NewCluster(config.CASSANDRA_URL)
	cluster.Keyspace = ""
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() *gocql.Session {
	return session
}
