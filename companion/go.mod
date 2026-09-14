module github.com/jhunthrop/foreversixty/companion

go 1.25.11

require (
	github.com/jedisct1/go-minisign v0.0.0-20260527172527-a09352b57a22
	github.com/jhunthrop/foreversixty/logs v0.0.0
	github.com/klauspost/compress v1.20.0
	github.com/zalando/go-keyring v0.2.8
)

require (
	github.com/danieljoos/wincred v1.2.3 // indirect
	github.com/godbus/dbus/v5 v5.2.2 // indirect
	golang.org/x/crypto v0.52.0 // indirect
	golang.org/x/sys v0.45.0 // indirect
)

replace github.com/jhunthrop/foreversixty/logs => ../logs
