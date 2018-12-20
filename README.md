# 这是一个后端是Golang和node的grpc-web demo

* 首先你要有golang的环境 ，node 环境，docker环境。

```bash
$ ./build.sh
$ npm install 
$ npx webpack client.js
$ docker-compose up
$ python -m SimpleHTTPServer 8081
```
然后在浏览器打开 http://localhost:8081
在console 可以看到交互信息