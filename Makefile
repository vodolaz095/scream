# Set this values according to your application

export pckName=scream
export semver=0.0.1
export myGPGKey=994C6375

# Do not change after this string
export arch=$(shell uname)-$(shell uname -m)
export gittip=$(shell git log --format='%h' -n 1)
export ver=$(semver).$(gittip).$(arch)
export subver=$(pckName) made via $(shell hostname) on $(shell date)
export archiv=build/scream-$(arch)-v$(semver)


all: build

engrave:
	rm -f ver.go
	touch ver.go
	echo "package $(pckName)" >> ver.go
	echo "" >>ver.go
	echo "//VERSION constant is engraved on build process" >>ver.go
	echo "const VERSION = \"$(ver)\"" >> ver.go
	echo "" >>ver.go
	echo "//SUBVERSION constant is engraved on build process" >>ver.go
	echo "const SUBVERSION = \"$(subver)\"" >>ver.go
	echo "" >>ver.go

deps:
	go get -u github.com/tools/godep
	go get -u github.com/golang/lint/golint
	godep restore

check: deps
	gofmt  -w=true -s=true -l=true ./..
	golint ./...
	go vet
	go test -v


build: clean engrave deps check
	go build -o "build/$(pckName)" app/$(pckName).go
	git checkout ver.go

dist: build
	zip $(archiv).zip  build/$(pckName) README.md README_RU.md CHANGELOG.md homedir/ contrib/ -r
	tar -czvf $(archiv).tar.gz  build/$(pckName) README.md README_RU.md CHANGELOG.md homedir/ contrib/
	tar -cjvf $(archiv).tar.bz2 build/$(pckName) README.md README_RU.md CHANGELOG.md homedir/ contrib/


sign:
	rm build/*.txt -f
	rm build/*.txt.sig -f
	find build/ -name $(pckName)-* -exec md5sum {} + > build/md5sum.txt
	gpg2 -a --output build/md5sum.txt.sig  --detach-sig build/md5sum.txt
	gpg2 --verify build/md5sum.txt.sig build/md5sum.txt
	find build/ -name $(pckName)-* -exec sha1sum {} + > build/sha1sum.txt
	gpg2 -a --output build/sha1sum.txt.sig --detach-sig build/sha1sum.txt
	gpg2 --verify build/sha1sum.txt.sig build/sha1sum.txt
	@echo ""
	@echo ""
	@echo "MD5 hashes"
	@echo "========================"
	@cat build/md5sum.txt
	@echo ""
	@echo ""
	@echo "SHA1 hashes"
	@echo "========================"
	@cat build/sha1sum.txt
	@echo ""
	@echo ""
	@echo "*.sig files are signed with my GPG key of \`$(myGPGKey)\`"

clean:
	git checkout ver.go
	rm -rf build/$(pckName)
	rm -rf build/*.zip
	rm -rf build/*.tar.gz
	rm -rf build/*.tar.bz2
	rm -rf build/*.txt
	rm -rf build/*.txt.sig

test: check

install: build
	su -c 'cp -f build/$(pckName) /usr/bin/'

uninstall:
	su -c 'rm -rf /usr/bin/$(pckName)'
