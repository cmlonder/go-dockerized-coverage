# Integrate Go Project with SonarQube

Read detailed description in [this blog post](http://cemalonder.com/post/go-sonar-integration-using-docker-part-1/)

Run tests
```
go test ./... -v -coverprofile=coverage.out -covermode=count -coverpkg=./...
```

Start SonarQube
```
docker run --rm -p 9000:9000 sonarqube
```

Analyse using SonarScannerCli (pay attention to IP which is local one)
```
 docker run --rm -e SONAR_HOST_URL=http://192.168.0.11:9000 -v ${PWD}:/usr/src sonarsource/sonar-scanner-cli
```