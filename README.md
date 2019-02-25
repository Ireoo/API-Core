# API Core with golang
> version: 1.0.0

## Golang 最新版本
```bash
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go
```

## dep 包管理安装
```bash
curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh
```


## 数据库操作
https://www.jb51.net/article/143208.htm


## 提交数据形式
```bash
curl \
  -X POST \
  http://localhost:2019/users/insert \
  -H 'Content-Type: application/json' \
  -d '{"where":{"id":1},"data":"2", "other":"3"}'
```

## 链接说明
```bash
http://localhost:2019/:table/:mode

// table -> 你要选择的表名
// mode -> 对该表的操作
// insert, update, findOne, findAll, delete
```
