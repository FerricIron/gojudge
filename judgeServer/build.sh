docker rm  gojudge
rm -rf judgeServer
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' . && docker build . -t gojudge
docker run -it  --net=maymomo --name=gojudge -p 127.0.0.1:8080:8080 gojudge
