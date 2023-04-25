package es

import (
	"github.com/olivere/elastic/v7"
)
import "context"

type Client struct {
	conf     *Config
	instance *elastic.Client
}

const (
	TypeDoc = "_doc"
)

func (esClient *Client) Instance() *elastic.Client {
	return esClient.instance
}

// 使用索引前缀
func (esClient *Client) WithPrefix(index string) string {
	return esClient.conf.IndexPrefix + index
}

func (esClient *Client) InsertDoc(index string, data interface{}) (*elastic.IndexResponse, error) {
	ctx := context.Background()
	put1, err := esClient.instance.Index().
		Index(esClient.WithPrefix(index)).
		Type(TypeDoc).
		BodyJson(data).
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return put1, nil
}

func (esClient *Client) SearchCountWithTime(index string, beginTime int64, endTime int64) (int64, error) {
	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewRangeQuery("created_at").Gt(beginTime)) //>
	boolQuery.Filter(elastic.NewRangeQuery("created_at").Lt(endTime))

	count, err := esClient.instance.Count(esClient.WithPrefix(index)).Query(boolQuery).Pretty(true).Do(context.Background())

	return count, err
}

// SearchByTime
func (esClient *Client) SearchByTime(index string, beginTime int64, endTime int64) (*elastic.SearchResult, error) {

	boolQuery := elastic.NewBoolQuery()
	boolQuery.Filter(elastic.NewRangeQuery("created_at").Gt(beginTime))
	boolQuery.Filter(elastic.NewRangeQuery("created_at").Lt(endTime))

	searchResult, err := esClient.instance.Search(esClient.WithPrefix(index)).Query(boolQuery).Pretty(true).Do(context.Background())

	return searchResult, err
}

//MultiInsertDoc 批量插入
func (esClient *Client) MultiInsertDoc(index string, items []string) (bool, error) {
	bulkRequest := esClient.instance.Bulk()

	for _, data := range items {
		doc := elastic.NewBulkIndexRequest().Index(esClient.WithPrefix(index)).Type(TypeDoc).Doc(data)
		bulkRequest = bulkRequest.Add(doc)
	}

	response, err := bulkRequest.Do(context.TODO())
	if err != nil {
		return false, err
	}
	return !response.Errors, nil
}
