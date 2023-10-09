
build back:
	goreleaser release --snapshot --clean

build front:
    export NODE_OPTIONS=--openssl-legacy-provider
    npm run build
