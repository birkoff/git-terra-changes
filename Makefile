linux:
	env GOOS=linux GOARCH=amd64 GOARM=7 go build -o git-terra-changes

macos:
	env GOOS=darwin GOARCH=amd64 go build -o git-terra-changes

.PHONY: linux macos
