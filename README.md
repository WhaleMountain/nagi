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

パッケージのインストールを```go get```でもいいが、```"github.com/docker/go-connections/nat"```が入らないです。```"github.com/docker/docker/vendor```の中に同じ名前のやつがいるから。なのでパッケージのインストールは```dep```を使っています。