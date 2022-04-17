build: 
	docker build . -t github.com/raygervais/dfw:latest

run:
	docker run -p 8080:8080 github.com/raygervais/dfw:latest
