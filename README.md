### RabbitMQ Golang Example

Create one RabbitMQ server on the same machine as rec.go.
Then run:
  go run rec.go


On a client run:
  go run sendfile.go -file a.xml

### Install Two RHEL7/CentOS7
yum -y install epel-release
yum -y install rabbitmq-server git golang

systemctl enable rabbitmq-server
firewall-cmd --permanent --add-port=5672/tcp
firewall-cmd --reload
setsebool -P nis_enabled 1
systemctl start rabbitmq-server
systemctl status rabbitmq-server

### Config Golang
vi /etc/profile.d/go.sh
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin


Opnieuw inloggen en directory structuur aanmaken

go get github.com/streadway/amqp

 rabbitmqctl list_queues




Todo:
- receive example save to file
- error checkin on both .go files
