.PHONY: build_wsl_to_windows
build_wsl_to_windows:
	CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CGO_LDFLAGS=-static-libgcc CC=/usr/bin/x86_64-w64-mingw32-gcc go build -o build/