package main

import (
	"context"
	"syscall"

	"github.com/lxn/win"
)

// App struct
type App struct {
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	hwnd := win.FindWindow(nil, syscall.StringToUTF16Ptr("Your App Title"))
    win.SetWindowLong(hwnd, win.GWL_EXSTYLE, win.GetWindowLong(hwnd, win.GWL_EXSTYLE)|win.WS_EX_LAYERED)}

// Greet returns a greeting for the given name

