## demo4
### 读取consul配置、api代理、pub/sub、ratelimit、hystrix、logrus、prometheus监控日志等


1. build 执行文件
```
cd user-service
make build

cd api
make build

cd restaurant-service
make build
```
2. 启动容器
```
docker-compose up
```
3. 访问
```
curl localhost:8080/user/info
```

### TODO
1. log内容整合进ELK
