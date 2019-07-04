## DEMO2
### 如何使用api的方式作为api gateway代理后端的rpc服务

1. go run main.go
2. micro --registry=consul api --handler=api --namespace=anakin.sun.api
3. GET localhost:8080/user/get