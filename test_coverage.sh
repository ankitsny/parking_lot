go test ./... -coverprofile cover.out
go tool cover -html=cover.out -o cover.html

# install libgnome first
nohup gnome-open ./cover.html