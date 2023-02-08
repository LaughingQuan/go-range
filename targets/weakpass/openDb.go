package weakpass

import (
	"database/sql"

	"github.com/go-redis/redis"
	_ "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var redisCf = struct {
	Addr     string
	Password string
	DB       int
}{
	Addr:     "127.0.0.1:6379", // 指定
	Password: "password",
	DB:       0, // redis一共16个库，指定其中一个库即可
}

var sqlCf = []struct {
	name string

	driverName string
	uri        string
}{
	{
		name:       "MySQL",
		driverName: "mysql",
		uri:        "root:root@tcp(127.0.0.1:3306)/irisapp?charset=utf8&parseTime=True&loc=Local",
	},
	{
		name:       "godror", //"github.com/godror/godror"
		driverName: "godror",
		uri:        `user="testq" password="123456" connectString="192.169.123.72:1521/orcl"`,
	}, {
		name:       "oci8",
		driverName: "oci8", //"github.com/mattn/go-oci8"
		uri:        "liang/liang@192.168.0.4:1521/orcl",
	}, {
		name:       "PgSql",
		driverName: "postgres", //"github.com/lib/pq"
		uri:        "postgres://pqgotest:password@localhost:9090/pqgotest?sslmode=verify-full",
	}, {
		name:       "PgSql",
		driverName: "postgres", // "github.com/lib/pq"
		uri:        "user=bob password=secret host=1.2.3.4 port=5432 dbname=mydb sslmode=verify-full",
	},
}
var mongoCf = []struct {
	name string
	uri  string
}{
	{
		name: "mongodb",
		uri:  "mongodb://admin:password@host:port",
	},
}

func WeakPass() {
	redis.NewClient(&redis.Options{
		Addr:     redisCf.Addr,
		Password: redisCf.Password,
		DB:       redisCf.DB,
	})
	for _, cf := range sqlCf {
		sql.Open(cf.driverName, cf.uri)
	}
	for _, cf := range mongoCf {
		options.Client().ApplyURI(cf.uri)
	}
}
