mv commands.go commands.go.old
curl -LO https://github.com/antirez/redis-doc/raw/master/commands.json
echo "package main\n\nvar redisCommandsJSON = \`" > commands.go
cat commands.json >> commands.go
echo "\`" >> commands.go
