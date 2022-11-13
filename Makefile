build: GOOS =?
build: GOARCH =?
build:
	CGO_ENABLED=0 GOOS=$(GOOS) GOARCH=$(GOARCH) go build -ldflags "-extldflags -static"

dev:
	docker build --no-cache -t kmlebedev/netbackup-exporter:dev -f Dockerfile .

codegen: NBU_VERSION = 8.3
codegen: NBU_API = admin
codegen:
	mkdir -p ./nbu-$(NBU_API)-api
	curl -Ls "https://sort.veritas.com/public/documents/nbu/$(NBU_VERSION)/windowsandunix/productguides/html/$(NBU_API)-api/admin.yaml" -o ./swagger-nbu-admin-api.yaml
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate --input-spec /local/swagger-nbu-admin-api.yaml --lang go --output /local/nbu-$(NBU_API)-api

gh_upload: GOOS = linux
gh_upload: GOARCH = amd64
gh_upload: RELEASE = v0.0.7
gh_upload: REPO = netbackup-exporter
gh_upload: build
	tar -czvf $(REPO)-$(RELEASE)-$(GOOS)-$(GOARCH).tar.gz $(REPO)
	gh release upload $(RELEASE) $(REPO)-$(RELEASE)-$(GOOS)-$(GOARCH).tar.gz --repo kmlebedev/$(REPO)
	rm $(REPO)-$(RELEASE)-$(GOOS)-$(GOARCH).tar.gz