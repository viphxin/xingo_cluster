# xingo_cluster
xingo cluster 分布式集群 示例代码
示例配置:
```json
{
"master":{"host": "192.168.2.225","rootport":9999},
"servers":{
"gate":{"host": "192.168.2.225", "rootport":10000,"name":"gate", "module": "gate", "log": "gate.log"},
"gate1":{"host": "192.168.2.225", "rootport":10001,"name":"gate1", "module": "gate", "log": "gate1.log"},
"net1":{"host": "192.168.2.225", "netport":11009,"name":"net1","remotes":["gate", "gate1", "admin"], "module": "net", "log": "net.log"},
"net2":{"host": "192.168.2.225", "netport":11009,"name":"net2","remotes":["gate", "gate1", "admin"], "module": "net", "log": "net.log"},
"net3":{"host": "192.168.2.225", "netport":11009,"name":"net3","remotes":["gate", "gate1", "admin"], "module": "net", "log": "net.log"},
"net4":{"host": "192.168.2.225", "netport":11009,"name":"net4","remotes":["gate", "gate1", "admin"], "module": "net", "log": "net.log"},
"admin":{"host": "192.168.2.225", "remotes":["gate", "gate1"], "name":"admin", "module": "admin", "http": [8888, "/static"]},
"game1":{"host": "192.168.2.225", "remotes":["gate", "gate1"], "name":"game1", "module": "game"}
}
}
```
架构图：
![alt text](https://github.com/viphxin/xingo_cluster/blob/master/conf/xingo_cluster_架构.png)
