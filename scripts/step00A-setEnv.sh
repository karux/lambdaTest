export GOPATH=${PWD}
export GOBIN=${PWD}/bin
mkdir ${GOBIN}
export GOTOOLING=${PWD}/tooling
mkdir ${GOTOOLING}
mkdir -p ${PWD}/src/github.com/karux
export PATH=${PATH}:${GOBIN}:${GOTOOLING}/bin
# export GOROOT =
export GOOS=darwin
export GOARCH=amd64
