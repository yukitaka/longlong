# LongLong

## Usage
### Development execution
#### Prepared
```
docker compose up -d
```
```
go run cmd/llctl/llctl.go help
```

### Generate
#### All
```
go generate ./...
```
#### Stringer
```
go install golang.org/x/tools/cmd/stringer@latest
```

### Execute tests
#### All tests
```
go test ./...
```
