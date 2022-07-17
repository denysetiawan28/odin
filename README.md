# Odin

Odin is a service build with go language, handling all service for data reconcile, billing generation (Invoice & Debit Note), generate proforma invoice,

### Function Of This Service
- Data Reconcile
- Invoice Generation
- Debit Note Generation
- Pro Forma Invoice

### Requirements
- Golang 1.17

### Installation
#### Local
```shell
//clone project
git clone git@bitbucket.org:titipaja/odin.git
//change directory to project
cd thor
//downloading go module dependency
go mod tidy
```
#### Docker
```shell
//1ST TIME
git clone git@bitbucket.org:titipaja/odin.git
cd odin
sudo su
docker build -t titipaja/odin-server:1.0.0 .
docker run -d -p 9090:9090 --net bridge --name odin-service titipaja/odin-server:1.0.0

//2ND TIME & Continously
cd odin
git pull
sudo su
docker build -t titipaja/odin-server:1.0.1 .
docker kill odin-service
docker rm odin-service
docker run -d -p 9090:9090 --net bridge --name odin-service titipaja/odin-server:1.0.1
```
### Migration Files
- Not implemented yet

### Environment Variables
 ```
APPS_NAME=odin
APPS_VERSION=1.0.2
APPS_HTTP_PORT=9090

DB_MASTER_HOST=asgard-titip-626e.aivencloud.com
DB_MASTER_USERNAME=app_thor
DB_MASTER_PASSWORD=AVNS_6gEsWAwGG7KJpsG
DB_MASTER_NAME=defaultdb
DB_MASTER_SHCEMA=thor
DB_MASTER_PORT=18882
DB_MASTER_CONNECTION_TIMEOUT= 20s
DB_MASTER_IDLE_CONNECTION=5
DB_MASTER_MAX_OPEN_CONNECTION=10
DB_MASTER_DEBUG_MODE=true

WRAPPER_NETSUITE_URL=https://5318962.restlets.api.netsuite.com/app/site/hosting/restlet.nl
WRAPPER_NETSUITE_CONSUMER_KEY=47e81852dedd41513a8031c29ed3fb1edeac507ebde8e4b0f8a8c366f07922a8
WRAPPER_NETSUITE_CONSUMER_SECRET=d313cc0ed7638e879b654286c3715c114ddb7f79ba73c9900ca3e1ef7c2c9392
WRAPPER_NETSUITE_TOKEN_ID=d79d8e6b48fb5627093e1bb5f2979e2c817fbcbe425bcc7d17bc9f3def29b9e8
WRAPPER_NETSUITE_TOKEN_SECRET=57f0a213b0b04dcf31c916bc1bc6834764c7600cf6aa0bf7d61898a4e2c8810e
WRAPPER_NETSUITE_REALM=5318962
 ```