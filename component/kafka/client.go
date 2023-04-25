package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"strings"
	"sync"
	"time"
)

type Client struct {
	conf      Config
	writerLock      sync.Mutex
	readerLock sync.Mutex
	writers   map[string]*kafka.Writer
	readers   map[string]*kafka.Reader
}

type Config struct {
	Host string `json:"host"`
	GroupId string `json:"group_id"`
}

func NewClient(conf Config) *Client {
	return &Client{
		conf: conf,
		writers: map[string]*kafka.Writer{},
		readers: map[string]*kafka.Reader{},
	}
}


func (client *Client) Write(topic string, message []byte) error {
	client.writerLock.Lock()
	defer client.writerLock.Unlock()
	writer, ok := client.writers[topic]
	if !ok {
		writer = &kafka.Writer{
			Addr:         kafka.TCP(strings.Split(client.conf.Host, ",")...),
			Topic:        topic,
		}
	}


	err := writer.WriteMessages(context.Background(), kafka.Message{Value: message})
	return err
}

func (client *Client) Read(topic string) ([]byte, error) {
	client.readerLock.Lock()
	defer client.readerLock.Unlock()
	reader, ok := client.readers[topic]
	if !ok {
		reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:        strings.Split(client.conf.Host, ","),
			GroupID:        client.conf.GroupId,
			Topic:          topic,
			MinBytes:       10e3, // 10KB
			MaxBytes:       10e6, // 10MB
			CommitInterval: 500,
		})
		client.readers[topic] = reader
	}

	// 读取超过1秒则自动取消
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	m, err := reader.ReadMessage(ctx)
	if err != nil {
		return nil, err
	}

	return m.Value, nil
}
