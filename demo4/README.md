## demo4
### 读取consul配置、api代理、pub/sub等
1. 环境配置：
	consul地址 127.0.0.1：8500(服务发现，配置管理)

 	mysql地址 127.0.0.1：3306a(数据库)

	nats地址：127.0.0.1：4222(broker)

2. user服务发布消息，其他服务订阅

3. consul k/v 设置：

  key:/micro/config/database

  value:
  ```
{
    "user": {
        "db": {
            "master": {
                "host": "127.0.0.1",
                "port": 3306,
                "user": "testuser",
                "password": "test123",
                "dbname": "test"
            },
            "slave": {
                "host": "127.0.0.1",
                "port": 3306,
                "user": "testuser",
                "password": "test123",
                "dbname": "test"
            }
        }
    }
}
```
4. 用户服务：

```
cd user-service
go build -o user-service
./user-service
```

5. api服务：
```
cd api
go build -o api-service
./api-service
```

6. api-gateway
```
cd apigw
go build -o micro
./micro --broker=nats --registry=consul api --handler=http  --namespace=chope.co.api
```

7. dinner服务负责接收broker的消息
```
cd diner-service
go run main.go handler.go
```

8. 访问
```
curl localhost:8080/user/info
```
