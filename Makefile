build:
	GOOS=linux GOARCH=amd64 go build -v -ldflags '-d -s -w' -a -tags netgo -installsuffix netgo -o build/bin/app .

init:
	terraform -chdir="./infrastructure" init

plan:
	terraform -chdir="./infrastructure" apply

apply:
	terraform -chdir="./infrastructure" apply --auto-approve

destroy:
	terraform -chdir="./infrastructure" destroy --auto-approve