// +build wasm

package main

import "github.com/maxence-charriere/go-app/v7/pkg/app"

type hello struct {
	app.Compo

	Name string

	UpdateAvailable bool
}

func (h *hello) OnAppUpdate(ctx app.Context) {
	h.UpdateAvailable = ctx.AppUpdateAvailable
	h.Update()
}

func (h *hello) Render() app.UI {
	return app.Div().Body(
		app.H1().Body(
			app.Text("Hello WASM!!!"),
			app.Text(h.Name),
		),
		app.Input().
			Value(h.Name).
			OnChange(h.OnInputChange),
		app.If(h.UpdateAvailable,
			app.Button().
				Text("Update to the latest version?").
				OnClick(h.onUpdateClick),
		),
	)
}

func (h *hello) OnInputChange(ctx app.Context, e app.Event) {
	h.Name = ctx.JSSrc.Get("value").String()
	h.Update()
}

func (h *hello) onUpdateClick(ctx app.Context, e app.Event) {
	app.Reload()
}

func main() {
	app.Route("/", &hello{})
	app.Run()
}
