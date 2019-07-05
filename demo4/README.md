## demo4
### 读取consul配置，api代理等
1. consul地址 127.0.0.1：8500
2. mysql地址 127.0.0.1：3306
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
./micro api --handler=api namespace=chope.co.api
```
7. 访问
```
curl localhost:8080/user/info
```
