start:
	@env go run main.go start
auto:
	@env go run main.go AutoCurd -c .tmp/config.json

ent_m:
	@env go run -mod=mod entgo.io/ent/cmd/ent generate ./ent/schema
	@env go run main.go ent