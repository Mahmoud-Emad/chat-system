Elasticsearch::Model.client = Elasticsearch::Client.new(url: ENV["ELASTICSEARCH_URL"] || "http://0.0.0.0:9200")
