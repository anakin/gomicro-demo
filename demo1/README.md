## DEMO1
### 如何使用http方式的api gateway

1. go run main.go

2. micro --registry=consul api --handler=http --namespace=anakin.sun.api 

3. curl localhost:8080/user/test