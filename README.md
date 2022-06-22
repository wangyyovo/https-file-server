# 简易的https文件服务器

```shell
openssl genrsa -out server.key 2048
 
openssl req -new -x509 -key server.key -out server.crt -days 365
```
