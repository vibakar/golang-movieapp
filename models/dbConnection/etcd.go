package dbConnection

import (
	"github.com/coreos/etcd/client"
	"time"
	"github.com/astaxie/beego"
)

func ConnectToETCD() (client.KeysAPI, error) {
	ETCDEndpoints := beego.AppConfig.String("ETCDEndpoints")
	cfg := client.Config{
		Endpoints:               []string{ETCDEndpoints},
		Transport:               client.DefaultTransport,
		HeaderTimeoutPerRequest: time.Second,
	}
	c, err := client.New(cfg)
	if err != nil {
		return nil, err
	}
	conn := client.NewKeysAPI(c)
	return conn, nil
}