


# 仅适用于本机的test发布。 不具有泛用性
git pull

GO15VENDOREXPERIMENT=1 go build src/devops/main.go
if [ ! -d "../deploy" ]; then
  mkdir ../deploy
fi

mv main ../deploy/
yes | cp -rf src/conf  ../deploy/
cd ../deploy
./main -mode=test
