package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
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
	"time"
)

type customTheme struct {
	fyne.Theme
	isDarkMode bool
}

func (t customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground:
		if t.isDarkMode {
			return color.RGBA{R: 28, G: 28, B: 30, A: 255}
		}
		return color.White
	case theme.ColorNameButton:
		if t.isDarkMode {
			return color.RGBA{R: 72, G: 72, B: 74, A: 255}
		}
		return color.RGBA{R: 212, G: 227, B: 250, A: 255}
	case theme.ColorNameDisabledButton:
		if t.isDarkMode {
			return color.RGBA{R: 58, G: 58, B: 60, A: 255}
		}
		return color.RGBA{R: 174, G: 174, B: 178, A: 255}
	case theme.ColorNameForeground:
		if t.isDarkMode {
			return color.White
		}
		return color.Black
	case theme.ColorNameHover:
		if t.isDarkMode {
			return color.RGBA{R: 44, G: 44, B: 46, A: 255}
		}
		return color.RGBA{R: 230, G: 230, B: 230, A: 255}
	case theme.ColorNameInputBackground:
		if t.isDarkMode {
			return color.RGBA{R: 44, G: 44, B: 46, A: 255}
		}
		return color.White
	default:
		return t.Theme.Color(name, variant)
	}
}

type Config struct {
	JavaPath struct {
		Java8  string `yaml:"Java8"`
		Java11 string `yaml:"Java11"`
		Open   string `yaml:"Open"`
	} `yaml:"javapath"`
}

type Tool struct {
	Name     string `yaml:"ToolName"`
	Path     string `yaml:"PATH"`
	FileName string `yaml:"FileName"`
	Value    string `yaml:"VALUE"`
	Command  string `yaml:"COMMAND"`
	Optional string `yaml:"Optional"`
}

type Category struct {
	Name string `yaml:"CategoryName"`
	Tool []Tool `yaml:"Tools"`
}

type Categories struct {
	Category []Category `yaml:"Categories"`
}

var categories Categories
var config Config
var cachedYAMLData []byte

var myApp = app.NewWithID("com.sspsec.Spear")
var myWindow = myApp.NewWindow("SSP渗透集成工具箱V4_by_Spe4r 公众号:SSP安全研究")
var scrollableContents *container.Scroll

var found bool
var filename = "tool.yml"
var cmd *exec.Cmd

type rightButton struct {
	widget.BaseWidget
	tool              Tool
	label             *widget.Label
	OnTapped          func()
	OnTappedSecondary func(*fyne.PointEvent)
	OnMouseIn         func(*fyne.PointEvent)
	OnMouseOut        func(*fyne.PointEvent)
}

func (w *rightButton) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewWithoutLayout(w.label))
}

func (w *rightButton) Tapped(_ *fyne.PointEvent) {
	if w.OnTapped != nil {
		w.OnTapped()
	}
}

func (w *rightButton) TappedSecondary(e *fyne.PointEvent) {
	if w.OnTappedSecondary != nil {
		w.OnTappedSecondary(e)
	}
}

func newrightButton(name string) *rightButton {
	w := &rightButton{
		label: widget.NewLabel(name),
	}
	w.ExtendBaseWidget(w)
	return w
}

func main() {
	currentTheme := &customTheme{Theme: theme.LightTheme(), isDarkMode: false}
	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		currentTheme.isDarkMode = true
		currentTheme.Theme = theme.DarkTheme()
	} else {
		currentTheme.isDarkMode = false
		currentTheme.Theme = theme.LightTheme()
	}
	myApp.Settings().SetTheme(currentTheme)

	myWindow.SetMaster()

	outputLabel := widget.NewLabel("Output will be shown here")
	outputLabel.Wrapping = fyne.TextWrapBreak

	yamldata, err := ReadYAMLFile()
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yamldata, &categories)

	yaml.Unmarshal(yamldata, &config)

	var categoryContainers []fyne.CanvasObject
	var allContainers []*fyne.Container

	updateToolContainers(outputLabel, &categoryContainers, &allContainers)
	updateToolContainers(outputLabel, &categoryContainers, &allContainers)

	mainContent := container.NewVBox(categoryContainers...)
	scrollableContent := container.NewScroll(mainContent)
	scrollableContent.SetMinSize(fyne.NewSize(850, 650))
	scrollableContents = scrollableContent
	searchEntry := widget.NewEntry()
	searchEntry.SetPlaceHolder("搜索工具...")
	searchEntry.Resize(fyne.NewSize(600, 40))

	searchEntry.OnChanged = func(s string) {
		if s == "" {

			updateToolContainers(outputLabel, &categoryContainers, &allContainers)
			scrollableContent.Content = container.NewVBox(categoryContainers...)
		} else {
			filteredObjects := []fyne.CanvasObject{}
			s = strings.ToLower(s)
			for _, container := range allContainers {
				btn := container.Objects[0].(*widget.Button)

				if strings.Contains(strings.ToLower(btn.Text), s) {
					filteredObjects = append(filteredObjects, container)
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

	clearButton := widget.NewButton("取消搜索", func() {
		searchEntry.SetText("")
		updateToolContainers(outputLabel, &categoryContainers, &allContainers)
		scrollableContent.Content = container.NewVBox(categoryContainers...)
		scrollableContent.Refresh()
	})

	addButton := widget.NewButton("添加工具", func() {
		nameEntry := widget.NewEntry()
		nameEntry.SetPlaceHolder("工具名称")
		pathEntry := widget.NewEntry()
		pathEntry.SetPlaceHolder("需要在app包内resources文件夹内")
		filenameEntry := widget.NewEntry()
		filenameEntry.SetPlaceHolder("文件名")
		valueEntry := widget.NewEntry()
		valueEntry.SetPlaceHolder("Java8/Java11/Open/openterm")
		commandEntry := widget.NewEntry()
		commandEntry.SetPlaceHolder("命令")
		optionalEntry := widget.NewEntry()
		optionalEntry.SetPlaceHolder("可选参数")

		categoryNames := make([]string, len(categories.Category))
		for i, category := range categories.Category {
			categoryNames[i] = category.Name
		}
		categorySelect := widget.NewSelect(categoryNames, nil)

		form := container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("工具名称", nameEntry),
				widget.NewFormItem("工具路径", pathEntry),
				widget.NewFormItem("执行文件名", filenameEntry),
				widget.NewFormItem("运行方式", valueEntry),
				widget.NewFormItem("命令", commandEntry),
				widget.NewFormItem("可选参数", optionalEntry),
				widget.NewFormItem("类别", categorySelect),
			),
		)

		var dialogWindow dialog.Dialog
		dialogWindow = dialog.NewCustomConfirm("添加新工具", "提交", "取消", form, func(confirm bool) {
			if confirm {
				newTool := Tool{
					Name:     nameEntry.Text,
					Path:     pathEntry.Text,
					FileName: filenameEntry.Text,
					Value:    valueEntry.Text,
					Command:  commandEntry.Text,
					Optional: optionalEntry.Text,
				}
				for i, category := range categories.Category {
					if category.Name == categorySelect.Selected {
						categories.Category[i].Tool = append(categories.Category[i].Tool, newTool)
						break
					}
				}
				insertYAMLFile()
				updateToolContainers(outputLabel, &categoryContainers, &allContainers)
				scrollableContent.Content = container.NewVBox(categoryContainers...)
				scrollableContent.Refresh()
			}
		}, myWindow)

		dialogWindow.Resize(fyne.NewSize(400, 400))
		dialogWindow.Show()
	})

	searchAndAddContainer := container.NewBorder(nil, nil, nil, container.NewHBox(clearButton, addButton), searchEntry)

	var background *canvas.RadialGradient
	if currentTheme.isDarkMode {
		background = canvas.NewRadialGradient(color.RGBA{R: 25, G: 25, B: 25, A: 255}, color.RGBA{R: 0, G: 0, B: 0, A: 255})
	} else {
		background = canvas.NewRadialGradient(color.RGBA{R: 255, G: 255, B: 255, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	}

	content := container.NewMax(background, container.NewBorder(searchAndAddContainer, outputLabel, nil, nil, scrollableContent))

	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(800, 750))
	myWindow.CenterOnScreen()

	myWindow.ShowAndRun()
}

func createToolContainer(toolbase Tool, outputLabel *widget.Label, categoryContainers *[]fyne.CanvasObject, allContainers *[]*fyne.Container) *fyne.Container {

	button := widget.NewButton(toolbase.Name, func() {

	})
	menu := fyne.NewMenu("",
		fyne.NewMenuItem("删除", func() {
			dialog.NewConfirm("确认操作", fmt.Sprintf("确定要删除 %s？", toolbase.Name), func(confirmed bool) {
				if confirmed {
					removeYAMLFile(toolbase.Name, outputLabel, categoryContainers, allContainers)
				}
			}, myWindow).Show()
		}),
		fyne.NewMenuItem("修改", func() {
			updateToolConfig(toolbase, outputLabel, categoryContainers, allContainers)
		}),
		fyne.NewMenuItem("打开目录", func() {
			currentDir := getFullPath(toolbase.Path)
			if err := openToolDirectory(currentDir); err != nil {
				WriteErrorToFile("打开路径错误", currentDir, err)
			} else {
				fmt.Println("目录已打开:", currentDir)
			}
		}),
	)

	buttonTop := newrightButton("")
	buttonTop.tool = toolbase

	buttonTop.OnTapped = func() {
		err := ExecuteCommand(toolbase.Path, toolbase.Optional, toolbase.Value, toolbase.FileName)
		if err != nil {
			outputLabel.SetText("Error: " + err.Error())
		} else {
			outputLabel.SetText("Running: " + toolbase.Name)
		}
	}

	buttonTop.OnTappedSecondary = func(e *fyne.PointEvent) {
		outputLabel.SetText("右键点击: " + toolbase.Name)
		if e == nil {
			println(e)
			return
		}
		menus := widget.NewPopUpMenu(menu, myWindow.Canvas())

		menus.ShowAtPosition(fyne.NewPos(e.AbsolutePosition.X, e.AbsolutePosition.Y))
	}
	toolContainer := container.NewMax(button, buttonTop)
	return toolContainer
}

func updateToolConfig(toolbase Tool, outputLabel *widget.Label, categoryContainers *[]fyne.CanvasObject, allContainers *[]*fyne.Container) {
	nameEntry := widget.NewEntry()
	nameEntry.Text = toolbase.Name

	pathEntry := widget.NewEntry()
	pathEntry.Text = toolbase.Path

	filenameEntry := widget.NewEntry()
	filenameEntry.Text = toolbase.FileName

	valueEntry := widget.NewEntry()
	valueEntry.Text = toolbase.Value

	commandEntry := widget.NewEntry()
	commandEntry.Text = toolbase.Command

	optionalEntry := widget.NewEntry()
	optionalEntry.Text = toolbase.Optional

	categoryNames := make([]string, len(categories.Category))
	for i, category := range categories.Category {
		categoryNames[i] = category.Name
	}

	categorySelect := widget.NewSelect(categoryNames, nil)

	for i, category := range categories.Category {
		for _, tool := range category.Tool {
			if tool.Name == toolbase.Name {
				categorySelect.Selected = categories.Category[i].Name
				break
			}
		}
	}

	form := container.NewVBox(
		widget.NewForm(
			widget.NewFormItem("工具名称", nameEntry),
			widget.NewFormItem("工具路径", pathEntry),
			widget.NewFormItem("执行文件名", filenameEntry),
			widget.NewFormItem("运行方式", valueEntry),
			widget.NewFormItem("命令", commandEntry),
			widget.NewFormItem("可选参数", optionalEntry),
			widget.NewFormItem("类别", categorySelect),
		),
	)

	var dialogWindow dialog.Dialog
	dialogWindow = dialog.NewCustomConfirm("修改工具", "提交", "取消", form, func(confirm bool) {
		if confirm {

			newTool := Tool{
				Name:     nameEntry.Text,
				Path:     pathEntry.Text,
				FileName: filenameEntry.Text,
				Value:    valueEntry.Text,
				Command:  commandEntry.Text,
				Optional: optionalEntry.Text,
			}
			for i, category := range categories.Category {
				if category.Name == categorySelect.Selected {
					categories.Category[i].Tool = append(categories.Category[i].Tool, newTool)
					break
				}
			}
			fileName := toolbase.Name
			updateYAMLFile(newTool, fileName, categorySelect.Selected)
			updateToolContainers(outputLabel, categoryContainers, allContainers)
			scrollableContents.Content = container.NewVBox(*categoryContainers...)
			scrollableContents.Refresh()
		}
	}, myWindow)
	dialogWindow.Resize(fyne.NewSize(400, 400))
	dialogWindow.Show()
}

func updateToolContainers(outputLabel *widget.Label, categoryContainers *[]fyne.CanvasObject, allContainers *[]*fyne.Container) {

	*categoryContainers = nil
	*allContainers = nil

	for _, category := range categories.Category {
		label := widget.NewLabel(category.Name)
		labelContainer := container.NewMax(label)
		var buttons []fyne.CanvasObject

		for _, toolbase := range category.Tool {
			toolbase := toolbase
			container := createToolContainer(toolbase, outputLabel, categoryContainers, allContainers)
			buttons = append(buttons, container)
			*allContainers = append(*allContainers, container)
		}

		catContainer := container.NewVBox(labelContainer)
		gridContainer := container.NewGridWrap(fyne.NewSize(200, 35), buttons...)
		catContainer.Add(gridContainer)
		*categoryContainers = append(*categoryContainers, catContainer)
	}
}

func ReadYAMLFile() ([]byte, error) {
	if cachedYAMLData != nil {
		return cachedYAMLData, nil
	}

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		WriteErrorToFile("Error reading file", filename, err)
		return nil, fmt.Errorf("读取文件出错: %v", err)
	}
	cachedYAMLData = data
	return data, nil
}

func insertYAMLFile() {
	data, err := yaml.Marshal(&struct {
		Categories []Category `yaml:"Categories"`
		JavaPath   struct {
			Java8  string `yaml:"Java8"`
			Java11 string `yaml:"Java11"`
			Open   string `yaml:"Open"`
		} `yaml:"javapath"`
	}{
		Categories: categories.Category,
		JavaPath:   config.JavaPath,
	})
	if err != nil {
		WriteErrorToFile("YAML Marshal error", filename, err)
		return
	}

	content := "# Java 8\n" +
		"# 路径：resources/java8/bin/java\n" +
		"# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。\n" +
		"# Java 11\n" +
		"# 路径：resources/java11/bin/java\n" +
		"# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。\n" +
		"# 打开方式\n" +
		"# 命令：open\n" +
		"# 该命令用于打开或执行文件，具体依赖于操作系统的配置。\n" +
		string(data)

	err = ioutil.WriteFile(filename, []byte(content), 0644)
	if err != nil {
		WriteErrorToFile("WriteFile error", filename, err)
	}
	cachedYAMLData = nil // Invalidate the cache
}

func updateYAMLFile(toolbase Tool, fileName, fileSourceName string) {

	yamldata, err := ReadYAMLFile()
	if err != nil {
		log.Fatal(err)
	}
	yaml.Unmarshal(yamldata, &categories)

	yaml.Unmarshal(yamldata, &config)

	found := false

	for i, categorie := range categories.Category {
		for j, t := range categorie.Tool {
			if t.Name == fileName {
				if fileSourceName == categories.Category[i].Name {
					found = true
					categories.Category[i].Tool[j] = toolbase
					break
				} else {
					categories.Category[i].Tool = append(categories.Category[i].Tool[:j], categories.Category[i].Tool[j+1:]...)
				}
			}
		}
		if found {
			break
		}
	}
	if !found {
		for i, category := range categories.Category {
			if category.Name == fileSourceName {
				categories.Category[i].Tool = append(categories.Category[i].Tool, toolbase)
				break
			}
		}
	}

	file, err := os.OpenFile("tool.yml", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return
	}
	content := "# Java 8\n" +
		"# 路径：resources/java8/bin/java\n" +
		"# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。\n" +
		"# Java 11\n" +
		"# 路径：resources/java11/bin/java\n" +
		"# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。\n" +
		"# 打开方式\n" +
		"# 命令：open\n" +
		"# 该命令用于打开或执行文件，具体依赖于操作系统的配置。\n"
	_, _ = file.WriteString(content)
	newData, _ := yaml.Marshal(categories)
	_, _ = file.Write(newData)
	newData, _ = yaml.Marshal(config)
	_, _ = file.Write(newData)
	cachedYAMLData = nil
}

func removeYAMLFile(toolName string, outputLabel *widget.Label, categoryContainers *[]fyne.CanvasObject, allContainers *[]*fyne.Container) {

	for i, category := range categories.Category {
		for j, tool := range category.Tool {
			if tool.Name == toolName {
				categories.Category[i].Tool = append(categories.Category[i].Tool[:j], categories.Category[i].Tool[j+1:]...)
				found = true
				break
			}
		}
	}
	if found {
		insertYAMLFile()
		updateToolContainers(outputLabel, categoryContainers, allContainers)

		scrollableContents.Content = container.NewVBox(*categoryContainers...)
		scrollableContents.Refresh()
		outputLabel.SetText("Removed: " + toolName)
	} else {
		outputLabel.SetText("Tool not found: " + toolName)
	}
}

func WriteErrorToFile(msg, filename string, err error) {
	logFile, fileErr := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if fileErr != nil {
		fmt.Println("Failed to open error log file:", fileErr)
		return
	}
	defer logFile.Close()

	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logMsg := fmt.Sprintf("%s | ERROR | %s | %s: %v\n", timestamp, filename, msg, err)
	_, fileErr = logFile.WriteString(logMsg)
	if fileErr != nil {
		fmt.Println("Failed to write to error log file:", fileErr)
	}
}

func ExecuteCommand(path, optional, value, filename string) error {
	yamldata, err := ReadYAMLFile()
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamldata, &config); err != nil {
		WriteErrorToFile("YAML Unmarshal error", "config", err)
		return err
	}

	currentDir, err := os.Getwd()
	if err != nil {
		WriteErrorToFile("Get working directory error", "currentDir", err)
		return err
	}

	var execCommand string
	switch value {
	case "Java8", "Java11":
		javaPath := filepath.Join(currentDir, config.JavaPath.Java8)
		if value == "Java11" {
			javaPath = filepath.Join(currentDir, config.JavaPath.Java11)
		}
		execCommand = fmt.Sprintf("cd %s && %s %s -jar %s", path, javaPath, optional, filename)
	case "Open":
		execCommand = fmt.Sprintf("cd %s && open %s", path, filename)
	case "openterm":
		return openTerminal(path)
	default:
		return fmt.Errorf("unsupported value: %s", value)
	}

	fmt.Println("Executing command:", execCommand)

	cmd := exec.Command("sh", "-c", execCommand)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		WriteErrorToFile("Command execution failed", execCommand, err)
		return err
	}
	return nil
}

func getFullPath(relativePath string) string {
	currentDir, _ := os.Getwd()
	return filepath.Join(currentDir, relativePath)
}

func openTerminal(dir string) error {
	switch runtime.GOOS {
	case "darwin":
		fullPath := getFullPath(dir)
		itermPath := "/Applications/iTerm.app"

		if _, err := os.Stat(itermPath); err == nil {
			script := fmt.Sprintf(`tell application "iTerm"
                create window with default profile
                tell current session of current window
                    write text "cd %s; ls --color=always"
                end tell
            end tell`, fullPath)
			cmd = exec.Command("osascript", "-e", script)
		} else {
			script := fmt.Sprintf(`tell application "Terminal"
                do script "cd %s; ls --color=always"
            end tell`, fullPath)
			cmd = exec.Command("osascript", "-e", script)
		}
	default:
		err := fmt.Errorf("unsupported platform")
		WriteErrorToFile("Open terminal error", dir, err)
		return err
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	if err != nil {
		WriteErrorToFile("Failed to open terminal", dir, err)
		return fmt.Errorf("failed to open terminal: %w", err)
	}
	return nil
}

func openToolDirectory(dir string) error {

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", dir)
	default:
		return fmt.Errorf("unsupported platform")
	}
	return cmd.Start()
}
