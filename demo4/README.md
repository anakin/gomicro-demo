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
4. consul UI
http://localhost:8500

5. jaeger UI
http://localhost:16686

6. prometheus UI
http://localhost:9090

### TODO
1. log内容整合进ELK
