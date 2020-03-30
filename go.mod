module github.com/xDarkicex/cchha_new_server

go 1.14

require (
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751
	github.com/go-chi/chi v4.0.4+incompatible
	github.com/go-chi/chi/v4 v4.0.0-rc1
	github.com/labstack/gommon v0.3.0
	github.com/scorredoira/email v0.0.0-20191107070024-dc7b732c55da
	github.com/valyala/fasttemplate v1.0.1
	golang.org/x/xerrors v0.0.0-20191204190536-9bdfabe68543
)

replace github.com/xDarkicex/cchha_new_server/server => ./Hospice/server

replace github.com/xDarkicex/cchha_new_server/app/controllers => ./Hospice/app/controllers
