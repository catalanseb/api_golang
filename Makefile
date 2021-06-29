BUILDPATH=$(CURDIR)
GOPATH=$(shell go env GOPATH)
GO=$(shell which go)
GOINSTALL=$(GO) install
BINARY=api-go


makedir:
	@if [ ! -d $(BUILDPATH)/build/bin ] ; then mkdir -p $(BUILDPATH)/build/bin ; fi

clean:
	@echo "Limpiando binarios..."
	@if [ -d $(BUILDPATH)/build/bin ] ; then rm -rf $(BUILDPATH)/build/ ; fi
	@rm -rf coverfile*

mod:
	@echo "Vendoring..."
	@go mod vendor

build: makedir clean
	@echo "Compilando..."
	@go build -mod vendor -ldflags "-s -w" -o $(BUILDPATH)/build/bin/${BINARY} cmd/server.go
	@echo "Binario generado en build/bin/"${BINARY}