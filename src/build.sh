# install dep
curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.3.2/dep-linux-amd64
chmod +x /usr/local/bin/dep

# build the project
dep ensure
go build -o main main.go
cp main /usr/local/bin