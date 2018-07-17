For SQL Database :
  connectivity will use (https://golang.org/pkg/database/sql/)
  1. database/sql :   Package sql provides a generic interface around SQL (or SQL-like) databases. 
  2. github.com/mattn/go-sqlite3  : The driver that supports the database/sql interface standard in its SQLite driver, 
                go get github.com/mattn/go-sqlite3

  For GRPC

  1. GRPC : go get -u google.golang.org/grpc
  2. protoc Go plugin : go get -u github.com/golang/protobuf/protoc-gen-go
  3. export PATH=$PATH:$GOPATH/bin
  4. Run protoc command (with the grpc plugin) : protoc GRPC/protobuff/*.proto --go_out=plugins=grpc:./GRPC/protobuff











