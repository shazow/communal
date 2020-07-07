BINARY = build/news

run:
	go run *.go

deploy:
	GOARCH=amd64 GOOS=linux go build -o build/news-linux_amd64
	rsync -bavz --progress build/news-linux_amd64 ip.shazow.net:projects/news/news-linux_amd64
	ssh ip.shazow.net fuser -s -k -HUP projects/news/news-linux_amd64~ || echo "binary did not change"
