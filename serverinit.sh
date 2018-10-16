#!/usr/bin/env bash
sudo add-apt-repository ppa:webupd8team/java
sudo apt-get update
sudo apt-get install oracle-java8-installer


sudo apt install openjdk-8-jre-headless 
sudo wget -q -O - https://pkg.jenkins.io/debian/jenkins-ci.org.key | sudo apt-key add -
sudo sh -c 'echo deb http://pkg.jenkins.io/debian-stable binary/ > /etc/apt/sources.list.d/jenkins.list'
sudo apt-get update
sudo apt-get install jenkins
sudo systemctl start jenkins
sudo apt-get install golang


sudo chmod 777 -R /usr/local/bin
export GOPATH=~/go
sudo snap install go --classic
sudo apt-get install gcc
curl https://raw.githubusercontent.com/struCoder/pmgo/master/install.sh | sh
go get -u github.com/BurntSushi/toml
go get -u github.com/Sirupsen/logrus
go get -u github.com/fatih/color
go get -u github.com/olekukonko/tablewriter
go get -u github.com/sevlyar/go-daemon
go get -u github.com/struCoder/pidusage
go get -u gopkg.in/alecthomas/kingpin.v2

go get -u github.com/labstack/echo/...
go get -u github.com/go-sql-driver/mysql
go get -u golang.org/x/net/websocket
go get -u github.com/googollee/go-socket.io
go get -u github.com/rs/cors
go get -u github.com/graarh/golang-socketio

# 라이브러리 다 설치하고 실행 라이브러리 다설치해야 돌아감 pmgo
# pmgo안되면 ~/.pmgo지우고 하면됨