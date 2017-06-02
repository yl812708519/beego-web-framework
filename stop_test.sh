

ps -ef|grep ./devops|grep -v grep|awk '{print $2}' |xargs kill -9