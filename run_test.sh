


# 仅适用于本机的test发布。 不具有泛用性
git pull
go build src/main.go
if [ ! -d "../deploy" ]; then
  mkdir ../deploy
fi

mv main ../deploy/
yes | cp -rf src/conf  ../deploy/
cd ../deploy
./main -mode=test
