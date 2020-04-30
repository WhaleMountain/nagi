# なんちゃってCompose

docker client を使って docker の操作を楽しむ。
でも docker compose ができないから実装してみよう的な。 ```libcompose```は保守されてないので使わない方向で。

## 環境

- Go
    - version go1.14.2 darwin/amd64
- Package 管理
    - dep
- Package
    - "github.com/docker/docker/client"
    - "github.com/docker/docker/api"
    - "github.com/docker/go-connections/nat"
    - "github.com/gin-gonic/gin"

パッケージのインストールを```go get```でもいいが、```"github.com/docker/go-connections/nat"```が入らないです。```"github.com/docker/docker/vendor```の中に同じ名前のやつがいるから。なのでパッケージのインストールは```dep```を使っています。

## 追記

curl でdocker container と compose ができるようになりました。

### 例

- Container
    - ```curl http://localhost:8080/container -X POST -H "Content-Type: application/json" -d '{"conname": "nagi","images": "mysql:5.7","environment": ["-e","MYSQL_ROOT_PASSWORD=mysql"],"guestport": "3306","hostport": "3306","driver": "bridge"}'```

- Compose
    - ```curl http://localhost:8080/compose -X POST -H "Content-Type: application/json" -d '{"conname": ["nagi-db","nagi-word"],"images": ["mysql:5.7","wordpress:latest"],"environment": [["-e","MYSQL_ROOT_PASSWORD=mysql"],["-e", "WORDPRESS_DB_HOST=nagi-db", "WORDPRESS_DB_USER=root", "WORDPRESS_DB_PASSWORD=mysql"]],"guestports": ["3306","80"],"hostports": ["3306","1270"],"composename":"nagi","driver": "bridge"}'```

### 未実装
- エラー処理
    - 何か1つでもエラーの場合 500 code のみを返している。そのためどこのエラーなのかわかりづらい。
- イメージのpull機能
    - docker image は pull している前提で動きます。