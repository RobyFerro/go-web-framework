package gwf

import (
	"github.com/elastic/go-elasticsearch/v8"
)

func ConnectElastic(conf Conf) *elasticsearch.Client {
	cfg := elasticsearch.Config{
		Addresses: conf.Elastic.Hosts,
	}
	es, err := elasticsearch.NewClient(cfg)

	if err != nil {
		ProcessError(err)
	}

	return es
}
