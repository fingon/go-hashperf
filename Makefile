#
# Author: Markus Stenberg <fingon@iki.fi>
#
# Copyright (c) 2018 Markus Stenberg
#
# Created:       Fri Nov 30 10:12:07 2018 mstenber
# Last modified: Fri Nov 30 10:29:11 2018 mstenber
# Edit time:     7 min
#
#

build_binaries: native.test linux.arm.test linux.mips.test
	GOOS=linux GOARCH=mips GOMIPS=softfloat go test -c -o linux.mips.softfloat.test

native.test:
	go test -c -o native.test


linux.%.test: $(wildcard *.go)
	GOOS=linux GOARCH=$* go test -c -o linux.$*.test

clean:
	rm -f *.test
