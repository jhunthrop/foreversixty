module github.com/jhunthrop/foreversixty/companion

go 1.25.11

require (
	github.com/jhunthrop/foreversixty/logs v0.0.0
	github.com/klauspost/compress v1.20.0
	github.com/zalando/go-keyring v0.2.8
)

require (
	github.com/danieljoos/wincred v1.2.3 // indirect
	github.com/godbus/dbus/v5 v5.2.2 // indirect
	golang.org/x/sys v0.38.0 // indirect
)

replace github.com/jhunthrop/foreversixty/logs => ../logs
