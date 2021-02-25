[![Join us](https://img.shields.io/static/v1?label=slack&message=Join%20Us&color=blue)](https://grafadruid.slack.com)

# go-druid
A Golang client for Druid.
Now supports Query API and Common API.

### Development

#### Testing
`go-druid` uses mage to run tests locally.
    Install Mage: 
```    
    git clone https://github.com/magefile/mage
    cd mage
    go run bootstrap.go
```
`mage -l` provides a list of targets that can be run. Default is `Check`

```
Targets:
  build            runs go mod download and then installs the binary.
  check*           run linters and tests
  fmt              run gofmt linter
  lint             run golint linter https://github.com/golang/lint
  testCoverHTML    generates test coverage report
  testRace         run tests with race detector
  vet              run go vet linter

* default target
```