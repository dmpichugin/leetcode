fmt:
	gosimports -w ./

test:
	richgo test -count=1 ./...

# check code before commit
cc: fmt test