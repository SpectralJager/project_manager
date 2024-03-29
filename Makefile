server:
	go run cmd/server/server.go

sqlc:
	sqlc generate

templ:
	templ generate ./templates/templates.tmpl

tw:
	./tailwing build -i input.css -o static/main.css --watch
