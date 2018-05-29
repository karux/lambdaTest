mkdir -p ./tooling/bin
cd tooling
export GOPATH=${PWD}
export GOBIN=${GOPATH}/bin

# step 1 - install the tool
# step 2 - download the source for safe keeping

# dependency manager with vendoring 
go get -u github.com/golang/dep/cmd/dep
#go get github.com/golang/dep/cmd/dep

# convert raw json into a go structure
go get -u github.com/ChimeraCoder/gojson/gojson
#go get github.com/ChimeraCoder/gojson/gojson
