package mongo

import (
	"errors"
	"github.com/globalsign/mgo"
	"time"
)

var (
	ErrUnknownHosts = errors.New("unknown hosts")
)

// mongo客户端
type Client struct {
	*mgo.Session

	db   string
	conf *Config
}

// 连接
func (client *Client) connect() error {

	if len(client.conf.Hosts) == 0 {
		return ErrUnknownHosts
	}

	client.db = client.conf.Database
	dialInfo := &mgo.DialInfo{
		Addrs:    client.conf.Hosts,
		Source:   client.conf.Source,
		Username: client.conf.Username,
		Password: client.conf.Password,
		Timeout:  time.Second * time.Duration(client.conf.Timeout),
	}

	s, err := mgo.DialWithInfo(dialInfo)

	if err != nil {
		return err
	}
	client.Session = s
	return nil
}

// UseCollection 选择集合
func (client *Client) UseCollection(collection string) (*mgo.Session, *mgo.Collection) {
	s := client.Session.Copy()
	c := s.DB(client.db).C(collection)
	return s, c
}

// Insert 插入
func (client *Client) Insert(collection string, docs ...interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Insert(docs...)
}

// FindOne 查找
func (client *Client) FindOne(collection string, query, selector, result interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).One(result)
}

// FindOneWithSort 查找并排序
func (client *Client) FindOneWithSort(collection string, query interface{}, selector interface{}, sort string, result interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).Sort(sort).One(result)
}

// FindAll 查找全部
func (client *Client) FindAll(collection string, query, selector, result interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Find(query).Select(selector).All(result)
}

// Update 更新
func (client *Client) Update(collection string, query, update interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Update(query, update)
}

// UpdateById 根据ID更新
func (client *Client) UpdateById(collection string, id, update interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.UpdateId(id, update)
}

// Remove 删除
func (client *Client) Remove(collection string, query interface{}) error {
	ms, c := client.UseCollection(collection)
	defer ms.Close()
	return c.Remove(query)
}
