package tiny

import (
    "encoding/json"
    "ioutils"
)

type Config struct {
    ListenPort int    // 监听端口
    MaxConn int       // 最大连接数
    MongoUrl string   // mongodb url
}

var cfg Config = Cinfig{
    ListenPort: 6666,
    MaxConn: 1024,
    M
}

func loadconfig(file string){
    data,err := ioutil.ReadFile(file)
    if err != nil {
        panic(err)
    }
    err = json.Unmarshal(data,&cfg)
    if err != nil{
        panic(err)
    }
}
