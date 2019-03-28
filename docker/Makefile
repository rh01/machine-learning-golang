all: compile docker push clean

compile:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o model

docker:
	docker build --force-rm=true -t rh02/gotrain:single .

push:
	docker push rh02/gotrain:single

clean:
	rm model
