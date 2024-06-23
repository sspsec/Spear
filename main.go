package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	_ "github.com/lengzhao/font/autoload"
	"gopkg.in/yaml.v3"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

type customTheme struct {
	fyne.Theme
}

type Config struct {
	Java_path struct {
		Java8  string `yaml:"Java8"`
		Java11 string `yaml:"Java11"`
		Open   string `yaml:"Open"`
	} `yaml:"javapath"`
}

type Categories struct {
	Category []struct {
		Name string `yaml:"CategoryName"`
		Tool []struct {
			Name     string `yaml:"ToolName"`
			PATH     string `yaml:"PATH"`
			FileName string `yaml:"FileName"`
			VALUE    string `yaml:"VALUE"`
			COMMAND  string `yaml:"COMMAND"`
			Optional string `yaml:"Optional"`
		} `yaml:"Tools"`
	} `yaml:"Categories"`
}

func (t customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameForeground:
		return color.Black // 设置文字颜色为纯黑以增强对比度和可读性
	case theme.ColorNameButton:
		return color.White
	case theme.ColorNameBackground:
		return color.RGBA{R: 173, G: 216, B: 230, A: 255} // 清新一点的颜色
	default:
		return t.Theme.Color(name, variant)
	}
}

func main() {
	myApp := app.NewWithID("com.sspsec.Spear")

	myApp.Settings().SetTheme(customTheme{theme.LightTheme()})
	myWindow := myApp.NewWindow("SSP渗透集成工具箱V2_by_Spe4r 公众号:SSP安全研究")

	background := canvas.NewRectangle(color.RGBA{R: 173, G: 216, B: 230, A: 255}) // 设置清新背景颜色

	outputLabel := widget.NewLabel("Output will be shown here")
	outputLabel.Wrapping = fyne.TextWrapBreak

	var categories Categories
	yamldata, err := ReadYAMLFile()
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yamldata, &categories)

	var canvasObjects []fyne.CanvasObject
	var allButtons []*widget.Button
	var categoryContainers []fyne.CanvasObject

	for _, category := range categories.Category {
		label := widget.NewLabel(category.Name)
		labelContainer := container.NewMax(label)
		var buttons []fyne.CanvasObject

		for _, toolbase := range category.Tool {
			toolbase := toolbase // 在循环内创建当前迭代的副本
			btn := widget.NewButton(toolbase.Name, func() {
				err := ExecuteCommand(yamldata, toolbase.PATH, toolbase.Optional, toolbase.VALUE, toolbase.COMMAND, toolbase.FileName)
				if err != nil {
					outputLabel.SetText("Error: " + err.Error())
				} else {
					outputLabel.SetText("Running: " + toolbase.Name)
				}
			})
			buttons = append(buttons, btn)
			allButtons = append(allButtons, btn) // 保证 allButtons 包含所有按钮的引用
		}

		catContainer := container.NewVBox(labelContainer)
		catContainer.Add(container.NewGridWrap(fyne.NewSize(160, 30), buttons...))
		canvasObjects = append(canvasObjects, catContainer)
		categoryContainers = append(categoryContainers, catContainer)
	}

	mainContent := container.NewVBox(canvasObjects...)
	scrollableContent := container.NewScroll(mainContent)
	scrollableContent.SetMinSize(fyne.NewSize(850, 650))

	backgroundWithContent := container.NewMax(background, scrollableContent)
	// 创建搜索框
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("搜索工具...")

	// 搜索框内容变化时过滤显示的按钮
	searchEntry.OnChanged = func(s string) {
		if s == "" {
			scrollableContent.Content = container.NewVBox(categoryContainers...)
		} else {
			filteredObjects := []fyne.CanvasObject{}
			s = strings.ToLower(s)
			for _, btn := range allButtons {
				if strings.Contains(strings.ToLower(btn.Text), s) {
					filteredObjects = append(filteredObjects, btn)
				}
			}
			if len(filteredObjects) > 0 {
				scrollableContent.Content = container.NewVBox(filteredObjects...)
			} else {
				scrollableContent.Content = container.NewVBox()
			}
		}
		scrollableContent.Refresh()
	}

	// 创建取消搜索按钮
	clearButton := widget.NewButton("取消搜索", func() {
		searchEntry.SetText("")
		scrollableContent.Content = container.NewVBox(categoryContainers...)
		scrollableContent.Refresh()
	})

	searchContainer := container.NewBorder(nil, nil, nil, clearButton, searchEntry)

	content := container.NewBorder(searchContainer, outputLabel, nil, nil, backgroundWithContent)

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 600))
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}

func ReadYAMLFile() ([]byte, error) {
	currentDir, err := os.Getwd()
	filename := "tool.yml"
	data, err := ioutil.ReadFile(filename)
	WriteErrorToFile(filename)
	WriteErrorToFile(currentDir)
	if err != nil {
		WriteErrorToFile(err.Error())
		return nil, fmt.Errorf("读取文件出错: %v", err)
	}
	return data, nil
}

func WriteErrorToFile(errMsg string) {
	file, err := os.OpenFile("errorsss.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("打开错误日志文件失败:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString(errMsg + "\n")
	if err != nil {
		fmt.Println("写入错误日志文件失败:", err)
		return
	}
}

func ExecuteCommand(yamldata []byte, path, optional, command, value string, filename string) error {
	var config Config
	if err := yaml.Unmarshal(yamldata, &config); err != nil {
		WriteErrorToFile("YAML Unmarshal error: " + err.Error())
		return err
	}

	// 获取应用的当前工作目录
	currentDir, err := os.Getwd()
	if err != nil {
		WriteErrorToFile("Get working directory error: " + err.Error())
		return err
	}

	var execCommand string
	// 根据 command 参数决定使用 Java8 还是 Java11 路径，或者是执行 Open 命令
	switch command {
	case "Java8", "Java11":
		javaPath := filepath.Join(currentDir, config.Java_path.Java8)
		if command == "Java11" {
			javaPath = filepath.Join(currentDir, config.Java_path.Java11)
		}
		execCommand = fmt.Sprintf("cd %s && %s %s %s %s", path, javaPath, optional, value, filename)
	case "Open":
		execCommand = fmt.Sprintf("cd %s && open %s", path, filename)
	case "openterm":
		return openTerminal(path)
	default:
		return fmt.Errorf("unsupported command: %s", command)
	}

	fmt.Println("Executing command:", execCommand)

	var cmd *exec.Cmd

	cmd = exec.Command("sh", "-c", execCommand)
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setpgid: true, // 在 UNIX-like 系统中创建一个新的进程组
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		WriteErrorToFile("Command execution failed: " + err.Error())
		return err
	}
	return nil
}

func openTerminal(dir string) error {
	var cmd *exec.Cmd
	currentDir, err := os.Getwd()
	if err != nil {
		return err
	}
	switch runtime.GOOS {
	case "darwin":
		full_path := filepath.Join(currentDir, dir)
		script := fmt.Sprintf("tell application \"Terminal\" to do script \"cd %s; ls --color=always\"", full_path)
		cmd = exec.Command("osascript", "-e", script)
	//case "windows":
	//	cmd = exec.Command("cmd", "/C", "start", "cmd", "/k", "cd", dir)
	default:
		return fmt.Errorf("unsupported platform")
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return fmt.Errorf("failed to open terminal: %w", err)
	}
	return nil
}
