# 根据 GPT-3 实现小工具  


## 环境依赖

```linux
Node >= V16.0
Npm >= 7.x
Go >= 17.1
```

## 创建配置文件并填写

```sh
cp config.yaml.example config.yaml
```

## 运行

1.
```sh
cd resource && npm install && npm run build
```

2. 
```sh
cd ../ && go run cmd/chat.go
```

## docker 运行

1.
```
docker build -t chat .
```

2.
```
docker run -it -p 127.0.0.1:10090:10060  -v /Users/config.yaml:/usr/local/bin/config.yaml chat:latest
```
