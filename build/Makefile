#  __  __   _   _  _____
# |  \/  | /_\ | |/ / __|
# | |\/| |/ _ \| ' <| _|
# |_|  |_/_/ \_\_|\_\___|
#

build-all: build-linux build-windows build-android build-unix zip-all

#   ___ _    ___   _   _  _
#  / __| |  | __| /_\ | \| |
# | (__| |__| _| / _ \| .` |
#  \___|____|___/_/ \_\_|\_|
#

clean:
	rm pluto-*

#  _    ___ _  _ _   ___  __
# | |  |_ _| \| | | | \ \/ /
# | |__ | || .` | |_| |>  <
# |____|___|_|\_|\___//_/\_\
#

build-linux:
	GOOS=linux GOARCH=386 		go build -o pluto-linux-i386 	../pluto.go
	GOOS=linux GOARCH=amd64 	go build -o pluto-linux-amd64 	../pluto.go

# __      _____ _  _ ___   _____      _____
# \ \    / /_ _| \| |   \ / _ \ \    / / __|
#  \ \/\/ / | || .` | |) | (_) \ \/\/ /\__ \
#   \_/\_/ |___|_|\_|___/ \___/ \_/\_/ |___/
#

build-windows:
	GOOS=windows GOARCH=386 	go build -o pluto-windows-i386.exe 	../pluto.go
	GOOS=windows GOARCH=amd64 	go build -o pluto-windows-amd64.exe ../pluto.go

#    _   _  _ ___  ___  ___ ___ ___
#   /_\ | \| |   \| _ \/ _ \_ _|   \
#  / _ \| .` | |) |   / (_) | || |) |
# /_/ \_\_|\_|___/|_|_\\___/___|___/
#


build-android:
	GOOS=android GOARCH=arm64 	go build -o pluto-android-arm64 ../pluto.go

#  _   _ _  _ _____  __
# | | | | \| |_ _\ \/ /
# | |_| | .` || | >  <
#  \___/|_|\_|___/_/\_\
#

build-unix:
	GOOS=freebsd GOARCH=386 	go build -o pluto-freebsd-i386 	../pluto.go
	GOOS=freebsd GOARCH=amd64 	go build -o pluto-freebsd-amd64 ../pluto.go
	GOOS=openbsd GOARCH=386 	go build -o pluto-openbsd-i386 	../pluto.go
	GOOS=openbsd GOARCH=amd64 	go build -o pluto-openbsd-amd64 ../pluto.go
	GOOS=netbsd GOARCH=386 		go build -o pluto-netbsd-i386 	../pluto.go
	GOOS=netbsd GOARCH=amd64 	go build -o pluto-netbsd-amd64 	../pluto.go

#  _______ ___
# |_  /_ _| _ \
#  / / | ||  _/
# /___|___|_|
#

zip-all:
	xz -9 pluto-*