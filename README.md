# なんちゃってCompose

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