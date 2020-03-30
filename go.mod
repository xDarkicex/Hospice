module github.com/xDarkicex/cchha_new_server

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/go-chi/chi/v4 v4.0.0-rc1
	github.com/labstack/gommon v0.3.0
	github.com/mattn/go-colorable v0.1.6 // indirect
	github.com/mgutz/ansi v0.0.0-20170206155736-9520e82c474b
	github.com/scorredoira/email v0.0.0-20191107070024-dc7b732c55da
	github.com/valyala/fasttemplate v1.1.0
	golang.org/x/net v0.0.0-20200226121028-0de0cce0169b // indirect
	golang.org/x/text v0.3.2 // indirect
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace github.com/xDarkicex/cchha_new_server/server => ./Hospice/server

replace github.com/xDarkicex/cchha_new_server/app/controllers => ./Hospice/app/controllers
