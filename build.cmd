set DEBUG=true
set GOOS=linux
set EMAIL=ssl@exaple.com
go build -o caddystart main.go

cp caddystart /c/Source/docker/CockroachDockerCluster/bin/
copy caddystart C:\Source\golang\src\sc.tpnfc.us\oyoshi\cockroachcluster\DockerFiles\docker\caddy\