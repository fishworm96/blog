CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o ./bin/blog
ssh root@ "supervisorctl stop blog"
scp ./bin/blog root@:/data/app/blog/bin/
scp ./conf/config-docker.yaml root@:/data/app/blog/conf/config.yaml
ssh root@ "supervisorctl start blog"
rm ./bin/blog