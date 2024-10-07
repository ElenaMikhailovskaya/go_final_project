# запускает тесты и формирует итоговый отчет по покрытию кода тестами
test-coverage:
	go test -v ./... -coverprofile cover.out.tmp | tee report.out && cat cover.out.tmp | grep -v "_mock_test.go" | grep -v "/app/" | grep -v "/interfaces/" | grep -v "/transport/" |  grep -v "/cmd/" | grep -v "/domain/" | grep -v "/models/" | grep -v "/docker/" | grep -v "/docs/" > cover.out && rm -f cover.out.tmp
# запускает в браузере окно с просмотром кусков кода, который покрыты и не покрыты тестами
coverage-view:
	go tool cover -html cover.out -o coverage.html