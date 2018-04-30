package etcd

import (
	"github.com/vibakar/golang-movieapp/models/dbConnection"
	"context"
	"github.com/coreos/etcd/client"
)

func Set(key, value string) (*client.Response, error)  {
	con, err := dbConnection.ConnectToETCD()
	resp, err := con.Set(context.Background(), key, value, nil)
	return resp, err
}

func Get(key string) (*client.Response, error)  {
	con, err := dbConnection.ConnectToETCD()
	resp, err := con.Get(context.Background(), key, nil)
	return resp, err
}

func Delete(key string) (*client.Response, error)  {
	con, err := dbConnection.ConnectToETCD()
	resp, err := con.Delete(context.Background(), key, nil)
	return resp, err
}