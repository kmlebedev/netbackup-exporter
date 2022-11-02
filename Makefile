build:
	CGO_ENABLED=0 go build -ldflags "-extldflags -static"

dev:
	docker build --no-cache -t registry.tochka-tech.com/devexp/oci/netbackup-exporter:dev -f Dockerfile .

codegen: NBU_VERSION = 8.3
codegen: NBU_API = admin
codegen:
	mkdir -p ./nbu-$(NBU_API)-api
	curl -Ls "https://sort.veritas.com/public/documents/nbu/$(NBU_VERSION)/windowsandunix/productguides/html/$(NBU_API)-api/admin.yaml" -o ./swagger-nbu-admin-api.yaml
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate --input-spec /local/swagger-nbu-admin-api.yaml --lang go --output /local/nbu-$(NBU_API)-api
