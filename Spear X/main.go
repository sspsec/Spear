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
		Title:  "Spear X",
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
			TitleBar:             mac.TitleBarDefault(),
			About: &mac.AboutInfo{
				Title:   "Spear X",
				Message: "© 2025 微信公众号：SSP安全研究 https://github.com/sspsec/Spear",
				Icon:    icon, // 使用嵌入的图标
			},
		},
	})

	if err != nil {
		log.Fatal(err)
	}
}

func (a *App) SelectFile() (string, error) {
	dialog, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title:            "选择工具文件",
		DefaultDirectory: "/Applications/Spear.app/Contents/Resources",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "所有文件 (*)",
				Pattern:     "*",
			},
			{
				DisplayName: "Mac可执行文件 (无扩展名, *.command)",
				Pattern:     "*;*.command;*.app;*.*",
			},
			{
				DisplayName: "Shell脚本 (*.sh)",
				Pattern:     "*.sh",
			},
			{
				DisplayName: "脚本文件 (*.py, *.php, *.js, *.jsp, *.asp)",
				Pattern:     "*.py;*.php;*.js;*.jsp;*.asp;*.aspx",
			},
			{
				DisplayName: "Java文件 (*.jar, *.class, *.java)",
				Pattern:     "*.jar;*.class;*.java",
			},
			{
				DisplayName: "文本文件 (*.txt, *.log, *.conf)",
				Pattern:     "*.txt;*.log;*.conf;*.ini;*.yaml;*.json",
			},
		},
	})

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
