## demo4
### 读取consul配置、api代理、jwt、pub/sub、ratelimit、hystrix、logrus、prometheus监控日志等


#### 1. build 执行文件
```
chmod +x make-build.sh
./make-build.sh
```

#### 2. 访问

创建用户：
```
post localhost:8080/user/create 
```
获取用户：
```
curl localhost:8080/user/info
```
获取jwt token(auth)：
```
post localhost:8080/user/auth
```

#### 3. consul UI
http://localhost:8500

#### 4. jaeger UI
http://localhost:16686

#### 5. prometheus UI
http://localhost:9090

### TODO
1. log内容整合进ELK
2. nats版本的问题：
```
go mod edit -replace=github.com/nats-io/nats.go@v1.8.2-0.20190607221125-9f4d16fe7c2d=github.com/nats-io/nats.go@v1.8.1
```
