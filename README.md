### RUN UNIT TEST ALL FOLDER
```go
go test ./...
```

### RUN UNIT TEST ONE FOLDER
```
cd folder_path
```
```go
go test
```

### RUN UNIT TEST WITH DETAIL
```go
go test -v
```

### RUN UNIT TEST ONE UNIT
```go
go test -v -run=NameFunctionTest
```

### RUN UNIT SUB TEST JUST ONE
```go
go test -v -run=NameFunctionTest/NameSubTest
```