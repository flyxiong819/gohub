# how to run
## 分别build server和client
```
cd server
go build
# 产出：chat.server可执行文件
```
```
cd client
go build
# 产出：chat.client可执行文件
```
## 启动服务端
```
./chat.server
```
## 启动1个或多个客户端（可多次，多处地方(copy chat.client过去)启动）
```
./chat.client
```
