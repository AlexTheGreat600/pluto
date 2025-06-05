#  __  __   _   _  _____
# |  \/  | /_\ | |/ / __|
# | |\/| |/ _ \| ' <| _|
# |_|  |_/_/ \_\_|\_\___|
#

build-all: build-linux build-windows build-android build-unix

build-linux:
	GOOS=linux GOARCH=386 go build -o pluto-linux-i386 ../pluto.go
	GOOS=linux GOARCH=amd64 go build -o pluto-linux-amd64 ../pluto.go

build-windows:
	GOOS=windows GOARCH=386 go build -o pluto-windows-i386.exe ../pluto.go
	GOOS=windows GOARCH=amd64 go build -o pluto-windows-amd64.exe ../pluto.go

build-android:
	GOOS=android GOARCH=arm64 go build -o pluto-android-arm64 ../pluto.go

build-unix:
	GOOS=freebsd GOARCH=386 	go build -o pluto-freebsd-i386 ../pluto.go
	GOOS=freebsd GOARCH=amd64 	go build -o pluto-freebsd-amd64 ../pluto.go
	GOOS=openbsd GOARCH=386 	go build -o pluto-openbsd-i386 ../pluto.go
	GOOS=openbsd GOARCH=amd64 	go build -o pluto-openbsd-amd64 ../pluto.go
	GOOS=netbsd GOARCH=386 		go build -o pluto-netbsd-i386 ../pluto.go
	GOOS=netbsd GOARCH=amd64 	go build -o pluto-netbsd-amd64 ../pluto.go