


# 仅适用于测试机192.168.49.20的test发布。 不具有泛用性
git pull
export GOROOT=/usr/local/go
export PATH=$PATH:$GOROOT/bin
export GOPATH=/home/devops/devops
GO15VENDOREXPERIMENT=1 go build src/devops/main.go
if [ ! -d "../deploy" ]; then
  mkdir ../deploy
fi

mv main ../deploy/devops
yes | cp -rf src/devops/conf  ../deploy/
cd ../deploy
nohup ./devops -mode=test > stdout.log 2>&1 &
