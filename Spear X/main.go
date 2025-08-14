package main

import (
	"embed"
	"fmt"
	"log"
	"strings"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
//go:embed build/appicon.png
var assets embed.FS
var icon []byte

func main() {
	// 创建一个新的应用实例
	app := NewApp()

	// 创建应用配置
	err := wails.Run(&options.App{
		Title:  "SpearX",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 0},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		// Mac 设置
		Mac: &mac.Options{
			Appearance:           mac.NSAppearanceNameDarkAqua, // 固定使用深色模式
			WebviewIsTransparent: true,
			WindowIsTranslucent:  true,
			TitleBar:             mac.TitleBarHiddenInset(),
			About: &mac.AboutInfo{
				Title:   "SpearX",
				Message: "A modern cross-platform tool manager\n\nDeveloped by Spe4r\n© 2025 Spe4r Development\n\nVersion 2.0.0",
				Icon:    icon, // 使用嵌入的图标
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) Select(selectFolder bool) (string, error) {
	var dialog string
	var err error

	options := runtime.OpenDialogOptions{
		Title:            "选择工具",
		DefaultDirectory: "/Applications/Spear.app/Contents/Resources",
	}

	if selectFolder {
		dialog, err = runtime.OpenDirectoryDialog(a.ctx, options)
	} else {
		options.Filters = []runtime.FileFilter{
			{
				DisplayName: "所有文件 (*)",
				Pattern:     "*",
			},
		}
		dialog, err = runtime.OpenFileDialog(a.ctx, options)
	}

	if err != nil {
		return "", err
	}

	// 验证路径
	if !strings.Contains(dialog, "Contents/Resources") {
		return "", fmt.Errorf("无效的工具路径：必须位于 App包内 resources 目录下")
	}

	// 提取相对路径
	parts := strings.Split(dialog, "Contents/Resources/")
	if len(parts) != 2 {
		return "", fmt.Errorf("无法解析工具路径")
	}

	return parts[1], nil
}
