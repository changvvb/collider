## collider IM信令服务配置
注意: 第三条有一个配置IP的地方,另外有需要翻墙的地方
### 1.请先安装golang,以Ubuntu为例
```bash
$ sudo apt-get update
$ sudo apt-get install golang
```
### 2.创建文件夹，并将该文件夹路径设置为$GOPATH
```bash
$ cd
$ mkdir go
$ echo "export GOPATH=~/go" >> ~/.bashrc
$ echo "export PATH=$PATH:~/go/bin" >> ~/.bashrc
$ source ~/.bashrc
$ mkdir -p ~/go/bin ~/go/src ~/go/pkg
```
### 3.配置和部署
```bash
$ cd ~/go/src
$ git clone https://github.com/changvvb/collidersrc.git
$ mv collidersrc/collider* ./
$ go get github.com/go-sql-driver/mysql
$ go get golang.org/x/net/websocket
$ go install collidermain
$ collidermain -room-server=http://IP:6060
```
<a href="https://github.com/changvvb/IM">BACK</a> 
