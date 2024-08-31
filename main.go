package main

import (
	"fmt" // 导入fmt包，用于格式化字符串
	"fyne.io/fyne/v2" // 导入Fyne GUI框架相关包
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gopkg.in/yaml.v3" // 导入yaml包，用于解析YAML文件
	"image/color" // 导入color包，用于颜色定义
	"io/ioutil" // 导入ioutil包，用于文件读取
	"log" // 导入log包，用于日志记录
	"os" // 导入os包，用于文件操作和获取系统信息
	"os/exec" // 导入exec包，用于执行外部命令
	"path/filepath" // 导入filepath包，用于路径操作
	"runtime" // 导入runtime包，用于获取运行时信息
	"strings" // 导入strings包，用于字符串操作
	"syscall" // 导入syscall包，用于系统调用
	"time" // 导入time包，用于时间操作
)

// 定义一个自定义主题结构体
type customTheme struct {
	fyne.Theme // 嵌入fyne.Theme接口，用于实现自定义主题
	isDarkMode bool // 是否是暗黑模式
}

// 实现Color方法，用于返回自定义颜色
func (t customTheme) Color(name fyne.ThemeColorName, variant fyne.ThemeVariant) color.Color {
	switch name {
	case theme.ColorNameBackground: // 背景颜色
		if t.isDarkMode {
			return color.RGBA{R: 28, G: 28, B: 30, A: 255} // 暗黑模式下的背景颜色
		}
		return color.White // 亮色模式下的背景颜色
	case theme.ColorNameButton: // 按钮颜色
		if t.isDarkMode {
			return color.RGBA{R: 72, G: 72, B: 74, A: 255} // 暗黑模式下的按钮颜色
		}
		return color.RGBA{R: 212, G: 227, B: 250, A: 255} // 亮色模式下的按钮颜色
	case theme.ColorNameDisabledButton: // 禁用按钮颜色
		if t.isDarkMode {
			return color.RGBA{R: 58, G: 58, B: 60, A: 255} // 暗黑模式下的禁用按钮颜色
		}
		return color.RGBA{R: 174, G: 174, B: 178, A: 255} // 亮色模式下的禁用按钮颜色
	case theme.ColorNameForeground: // 前景颜色
		if t.isDarkMode {
			return color.White // 暗黑模式下的前景颜色
		}
		return color.Black // 亮色模式下的前景颜色
	case theme.ColorNameHover: // 悬停颜色
		if t.isDarkMode {
			return color.RGBA{R: 44, G: 44, B: 46, A: 255} // 暗黑模式下的悬停颜色
		}
		return color.RGBA{R: 230, G: 230, B: 230, A: 255} // 亮色模式下的悬停颜色
	case theme.ColorNameInputBackground: // 输入框背景颜色
		if t.isDarkMode {
			return color.RGBA{R: 44, G: 44, B: 46, A: 255} // 暗黑模式下的输入框背景颜色
		}
		return color.White // 亮色模式下的输入框背景颜色
	default:
		return t.Theme.Color(name, variant) // 默认返回系统主题颜色
	}
}

// 定义配置结构体，用于解析YAML配置文件
type Config struct {
	Path struct { 
		Java8   string `yaml:"Java8"` // Java8路径
		Java11  string `yaml:"Java11"` // Java11路径
		Python3 string `yaml:"Python3"` // Python3路径
		Open    string `yaml:"Open"` // Open命令路径
	} `yaml:"path"`
}

// 定义工具结构体，用于表示每个工具的信息
type Tool struct {
	Name     string `yaml:"ToolName"` // 工具名称
	Path     string `yaml:"PATH"` // 工具路径
	FileName string `yaml:"FileName"` // 文件名
	Value    string `yaml:"VALUE"` // 值（运行方式）
	Command  string `yaml:"COMMAND"` // 命令
	Optional string `yaml:"Optional"` // 可选参数
}

// 定义类别结构体，用于表示工具分类
type Category struct {
	Name string `yaml:"CategoryName"` // 类别名称
	Tool []Tool `yaml:"Tools"` // 工具列表
}

// 定义所有类别的结构体，用于表示整体工具结构
type Categories struct {
	Category []Category `yaml:"Categories"` // 所有类别列表
}

var categories Categories // 全局变量，存储解析后的类别数据
var config Config // 全局变量，存储解析后的配置数据
var cachedYAMLData []byte // 缓存的YAML数据，用于减少文件读取

// 实现Font方法，返回自定义字体资源
func (t customTheme) Font(style fyne.TextStyle) fyne.Resource {
	if style.Bold {
		return resourceMesloLGSNFRegularTtf // 粗体
	}
	if style.Italic {
		return resourceMesloLGSNFRegularTtf // 斜体
	}
	if style.Monospace {
		return resourceMesloLGSNFRegularTtf // 等宽字体
	}
	return resourceMesloLGSNFRegularTtf // 默认字体
}

// 主函数
func main() {
	myApp := app.NewWithID("com.sspsec.Spear") // 创建带有ID的应用实例
	currentTheme := &customTheme{Theme: theme.LightTheme(), isDarkMode: false} // 初始化自定义主题

	if fyne.CurrentApp().Settings().ThemeVariant() == theme.VariantDark {
		currentTheme.isDarkMode = true
		currentTheme.Theme = theme.DarkTheme() // 设置为暗黑主题
	} else {
		currentTheme.isDarkMode = false
		currentTheme.Theme = theme.LightTheme() // 设置为亮色主题
	}
	myApp.Settings().SetTheme(currentTheme) // 应用自定义主题

	myWindow := myApp.NewWindow("SSP渗透集成工具箱V3_by_Spe4r 公众号:SSP安全研究") // 创建主窗口
	myWindow.SetMaster() // 设置窗口为主窗口

	outputLabel := widget.NewLabel("Output will be shown here") // 创建输出标签，用于显示运行结果
	outputLabel.Wrapping = fyne.TextWrapBreak // 设置标签文本自动换行

	yamldata, err := ReadYAMLFile() // 读取YAML文件数据
	if err != nil {
		log.Fatal(err) // 如果读取出错，记录错误并退出程序
	}
	yaml.Unmarshal(yamldata, &categories) // 解析YAML数据到categories结构体
	yaml.Unmarshal(yamldata, &config) // 解析YAML数据到config结构体

	var categoryContainers []fyne.CanvasObject // 存储分类容器的切片
	var allContainers []*fyne.Container // 存储所有工具容器的切片

	updateToolContainers(outputLabel, &categoryContainers, &allContainers) // 更新工具容器内容

	mainContent := container.NewVBox(categoryContainers...) // 创建垂直布局的主内容
	scrollableContent := container.NewScroll(mainContent) // 创建可滚动的主内容容器
	scrollableContent.SetMinSize(fyne.NewSize(850, 650)) // 设置最小尺寸

	searchEntry := widget.NewEntry() // 创建搜索框
	searchEntry.SetPlaceHolder("搜索工具...") // 设置搜索框的占位符
	searchEntry.Resize(fyne.NewSize(600, 40)) // 调整搜索框尺寸

	// 搜索框内容变更时的处理逻辑
	searchEntry.OnChanged = func(s string) {
		if s == "" {
			updateToolContainers(outputLabel, &categoryContainers, &allContainers) // 更新工具容器
			scrollableContent.Content = container.NewVBox(categoryContainers...) // 重置内容
		} else {
			filteredObjects := []fyne.CanvasObject{}
			s = strings.ToLower(s) // 转换搜索字符串为小写
			for _, container := range allContainers {
				btn := container.Objects[0].(*widget.Button) // 获取容器中的按钮
				if strings.Contains(strings.ToLower(btn.Text), s) {
					filteredObjects = append(filteredObjects, container) // 如果匹配则加入结果集
				}
			}
			if len(filteredObjects) > 0 {
				scrollableContent.Content = container.NewVBox(filteredObjects...) // 更新显示过滤后的工具
			} else {
				scrollableContent.Content = container.NewVBox() // 无匹配结果时清空显示
			}
		}
		scrollableContent.Refresh() // 刷新显示内容
	}

	// 取消搜索按钮
	clearButton := widget.NewButton("取消搜索", func() {
		searchEntry.SetText("") // 清空搜索框内容
		updateToolContainers(outputLabel, &categoryContainers, &allContainers) // 更新工具容器
		scrollableContent.Content = container.NewVBox(categoryContainers...) // 重置内容
		scrollableContent.Refresh() // 刷新显示内容
	})

	// 添加工具按钮
	addButton := widget.NewButton("添加工具", func() {
		nameEntry := widget.NewEntry() // 创建输入框，用于输入工具名称
		nameEntry.SetPlaceHolder("工具名称") // 设置输入框占位符
		pathEntry := widget.NewEntry() // 创建输入框，用于输入工具路径
		pathEntry.SetPlaceHolder("需要在app包内resources文件夹内") // 设置输入框占位符
		filenameEntry := widget.NewEntry() // 创建输入框，用于输入文件名
		filenameEntry.SetPlaceHolder("文件名") // 设置输入框占位符
		valueEntry := widget.NewEntry() // 创建输入框，用于输入值
		valueEntry.SetPlaceHolder("Java8/Java11/Python3/Open/openterm") // 设置输入框占位符
		commandEntry := widget.NewEntry() // 创建输入框，用于输入命令
		commandEntry.SetPlaceHolder("命令") // 设置输入框占位符
		optionalEntry := widget.NewEntry() // 创建输入框，用于输入可选参数
		optionalEntry.SetPlaceHolder("可选参数") // 设置输入框占位符

		categoryNames := make([]string, len(categories.Category)) // 创建类别名称的切片
		for i, category := range categories.Category {
			categoryNames[i] = category.Name // 填充类别名称
		}
		categorySelect := widget.NewSelect(categoryNames, nil) // 创建下拉框用于选择类别

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

		var dialogWindow dialog.Dialog // 定义对话框窗口
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
						categories.Category[i].Tool = append(categories.Category[i].Tool, newTool) // 将新工具加入到选定的类别
						break
					}
				}
				updateYAMLFile() // 更新YAML文件
				updateToolContainers(outputLabel, &categoryContainers, &allContainers) // 更新工具容器
				scrollableContent.Content = container.NewVBox(categoryContainers...) // 更新显示内容
				scrollableContent.Refresh() // 刷新显示
			}
		}, myWindow)

		dialogWindow.Resize(fyne.NewSize(400, 400)) // 设置对话框尺寸
		dialogWindow.Show() // 显示对话框
	})

	// 删除工具按钮
	removeButton := widget.NewButton("删除工具", func() {
		categoryNames := make([]string, len(categories.Category)) // 创建类别名称切片
		for i, category := range categories.Category {
			categoryNames[i] = category.Name // 填充类别名称
		}
		categorySelect := widget.NewSelect(categoryNames, nil) // 创建类别选择下拉框
		toolSelect := widget.NewSelect([]string{}, nil) // 创建工具选择下拉框

		categorySelect.OnChanged = func(s string) {
			var toolNames []string
			for _, category := range categories.Category {
				if category.Name == s {
					for _, tool := range category.Tool {
						toolNames = append(toolNames, tool.Name) // 根据选定的类别更新工具名称列表
					}
					break
				}
			}
			toolSelect.Options = toolNames // 更新工具选择下拉框的选项
			toolSelect.Refresh() // 刷新下拉框
		}

		form := container.NewVBox(
			widget.NewForm(
				widget.NewFormItem("类别", categorySelect),
				widget.NewFormItem("工具名称", toolSelect),
			),
		)

		var dialogWindow dialog.Dialog // 定义对话框窗口
		dialogWindow = dialog.NewCustomConfirm("删除工具", "提交", "取消", form, func(confirm bool) {
			if confirm {
				toolName := toolSelect.Selected // 获取选中的工具名称
				var found bool
				for i, category := range categories.Category {
					if category.Name == categorySelect.Selected {
						for j, tool := range category.Tool {
							if tool.Name == toolName {
								categories.Category[i].Tool = append(categories.Category[i].Tool[:j], categories.Category[i].Tool[j+1:]...) // 从工具列表中删除工具
								found = true
								break
							}
						}
						break
					}
				}
				if found {
					updateYAMLFile() // 更新YAML文件
					updateToolContainers(outputLabel, &categoryContainers, &allContainers) // 更新工具容器
					scrollableContent.Content = container.NewVBox(categoryContainers...) // 更新显示内容
					scrollableContent.Refresh() // 刷新显示内容
					outputLabel.SetText("Removed: " + toolName) // 更新输出标签文本
				} else {
					outputLabel.SetText("Tool not found: " + toolName) // 如果工具未找到，则显示提示信息
				}
			}
		}, myWindow)

		dialogWindow.Resize(fyne.NewSize(400, 200)) // 设置对话框尺寸
		dialogWindow.Show() // 显示对话框
	})

	// 创建搜索和添加工具的容器
	searchAndAddContainer := container.NewBorder(nil, nil, nil, container.NewHBox(clearButton, addButton, removeButton), searchEntry)

	// 设置背景
	var background *canvas.RadialGradient
	if currentTheme.isDarkMode {
		background = canvas.NewRadialGradient(color.RGBA{R: 25, G: 25, B: 25, A: 255}, color.RGBA{R: 0, G: 0, B: 0, A: 255}) // 暗黑模式下的背景
	} else {
		background = canvas.NewRadialGradient(color.RGBA{R: 255, G: 255, B: 255, A: 255}, color.RGBA{R: 255, G: 255, B: 255, A: 255}) // 亮色模式下的背景
	}

	// 创建主界面的内容容器
	content := container.NewMax(background, container.NewBorder(searchAndAddContainer, outputLabel, nil, nil, scrollableContent))

	myWindow.SetContent(content) // 设置窗口内容
	myWindow.Resize(fyne.NewSize(800, 750)) // 设置窗口大小
	myWindow.CenterOnScreen() // 将窗口居中显示

	myWindow.ShowAndRun() // 显示窗口并运行应用
}

// 创建工具容器，用于每个工具按钮的显示
func createToolContainer(toolbase Tool, outputLabel *widget.Label) *fyne.Container {
	button := widget.NewButton(toolbase.Name, func() {
		err := ExecuteCommand(toolbase.Path, toolbase.Optional, toolbase.Value, toolbase.FileName) // 执行工具的命令
		if err != nil {
			outputLabel.SetText("Error: " + err.Error()) // 如果执行出错，显示错误信息
		} else {
			outputLabel.SetText("Running: " + toolbase.Name) // 显示工具正在运行的信息
		}
	})

	toolContainer := container.NewVBox(button) // 创建垂直布局容器，包含按钮

	return toolContainer
}

// 更新工具容器，生成各个分类的工具按钮
func updateToolContainers(outputLabel *widget.Label, categoryContainers *[]fyne.CanvasObject, allContainers *[]*fyne.Container) {
	*categoryContainers = nil // 清空分类容器
	*allContainers = nil // 清空所有容器

	for _, category := range categories.Category {
		label := widget.NewLabel(category.Name) // 创建分类标签
		labelContainer := container.NewMax(label) // 创建最大化布局容器，包含标签
		var buttons []fyne.CanvasObject

		for _, toolbase := range category.Tool {
			toolbase := toolbase
			container := createToolContainer(toolbase, outputLabel) // 为每个工具创建容器
			buttons = append(buttons, container) // 添加工具容器到按钮列表
			*allContainers = append(*allContainers, container) // 添加工具容器到所有容器列表
		}

		catContainer := container```go
.NewVBox(labelContainer) // 创建垂直布局容器，包含分类标签
		gridContainer := container.NewGridWrap(fyne.NewSize(200, 40), buttons...) // 创建网格布局容器，包含工具按钮
		catContainer.Add(gridContainer) // 将按钮网格容器添加到分类容器中
		*categoryContainers = append(*categoryContainers, catContainer) // 将分类容器添加到分类容器列表中
	}
}

// 读取YAML文件内容
func ReadYAMLFile() ([]byte, error) {
	if cachedYAMLData != nil { 
		return cachedYAMLData, nil // 如果有缓存数据，直接返回
	}
	filename := "tool.yml" // 定义YAML文件名
	data, err := ioutil.ReadFile(filename) // 读取文件内容
	if err != nil {
		WriteErrorToFile("Error reading file", filename, err) // 如果读取文件出错，记录错误日志
		return nil, fmt.Errorf("读取文件出错: %v", err) // 返回错误信息
	}
	cachedYAMLData = data // 缓存读取到的数据
	return data, nil // 返回文件内容
}

// 更新YAML文件内容
func updateYAMLFile() {
	filename := "tool.yml" // 定义YAML文件名
	data, err := yaml.Marshal(&struct { 
		Categories []Category `yaml:"Categories"` // 更新的类别数据
		Path       struct {
			Java8   string `yaml:"Java8"` // Java8路径
			Java11  string `yaml:"Java11"` // Java11路径
			Python3 string `yaml:"Python3"` // Python3路径
			Open    string `yaml:"Open"` // Open命令路径
		} `yaml:"path"`
	}{
		Categories: categories.Category, // 更新的类别列表
		Path:       config.Path, // 更新的路径配置
	})
	if err != nil {
		WriteErrorToFile("YAML Marshal error", filename, err) // 如果序列化出错，记录错误日志
		return
	}

	// 定义文件内容模板，包含路径说明
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

	err = ioutil.WriteFile(filename, []byte(content), 0644) // 将数据写入文件
	if err != nil {
		WriteErrorToFile("WriteFile error", filename, err) // 如果写入出错，记录错误日志
	}
	cachedYAMLData = nil // 使缓存无效，强制下次读取文件
}

// 将错误信息写入日志文件
func WriteErrorToFile(msg, filename string, err error) {
	logFile, fileErr := os.OpenFile("errors.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644) // 打开或创建错误日志文件
	if fileErr != nil {
		fmt.Println("Failed to open error log file:", fileErr) // 如果文件打开失败，打印错误信息
		return
	}
	defer logFile.Close() // 确保文件关闭

	timestamp := time.Now().Format("2006-01-02 15:04:05") // 获取当前时间戳
	logMsg := fmt.Sprintf("%s | ERROR | %s | %s: %v\n", timestamp, filename, msg, err) // 格式化日志信息
	_, fileErr = logFile.WriteString(logMsg) // 将日志信息写入文件
	if fileErr != nil {
		fmt.Println("Failed to write to error log file:", fileErr) // 如果写入失败，打印错误信息
	}
}

// 执行工具命令
func ExecuteCommand(path, optional, value, filename string) error {
	yamldata, err := ReadYAMLFile() // 读取YAML文件内容
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(yamldata, &config); err != nil { // 解析YAML文件内容
		WriteErrorToFile("YAML Unmarshal error", "config", err) // 如果解析出错，记录错误日志
		return err
	}

	currentDir, err := os.Getwd() // 获取当前工作目录
	if err != nil {
		WriteErrorToFile("Get working directory error", "currentDir", err) // 如果获取目录出错，记录错误日志
		return err
	}

	// 构建完整路径
	fullPath := filepath.Join(currentDir, path)

	var execCommand string
	var cmd *exec.Cmd

	// 根据不同的值选择不同的执行方式
	switch value {
	case "Java8", "Java11":
		javaPath := config.Path.Java8 // 获取Java8路径
		if value == "Java11" {
			javaPath = config.Path.Java11 // 如果是Java11，则获取Java11路径
		}
		execCommand = fmt.Sprintf("cd %s && %s %s -jar %s", fullPath, currentDir+javaPath, optional, filename) // 构建Java执行命令
		cmd = exec.Command("cmd", "/C", execCommand) // 创建命令
	case "Open":
		execCommand = fmt.Sprintf("cd %s && %s %s", fullPath, filename, optional) // 构建打开命令
		cmd = exec.Command("cmd", "/C", execCommand) // 创建命令
	case "Python3":
		pythonPath := config.Path.Python3 // 获取Python3路径
		batchFileContent := fmt.Sprintf(`
@echo off
set PYTHONPATH=%s
cd /d "%s"
python %s
echo python %s %s
pause
`, currentDir+pythonPath, fullPath, filename, filename, optional) // 创建批处理文件内容

		execCommand = filepath.Join(os.TempDir(), "run_python.bat") // 创建批处理文件路径
		ioutil.WriteFile(execCommand, []byte(batchFileContent), 0644) // 将内容写入批处理文件

		// 通过 "start" 命令打开新的 cmd 窗口并执行命令
		cmd = exec.Command("cmd", "/C", "start", execCommand) // 创建命令

	case "openterm":
		openTerminal(path) // 打开终端
		return nil
	default:
		return fmt.Errorf("unsupported value: %s", value) // 如果值不支持，返回错误
	}

	// 记录执行的命令到日志
	WriteErrorToFile("Executing command: "+execCommand, "command.log", nil)

	// 设置隐藏窗口的属性
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏启动时的黑框
	}

	if err := cmd.Start(); err != nil { // 执行命令
		WriteErrorToFile("Command execution failed", "command.log", err) // 如果执行失败，记录错误日志
		return err
	}
	return nil
}

// 打开终端窗口
func openTerminal(dir string) error {
	cmd := exec.Command("cmd", "/C", "start", "cmd", "/K", "cd", dir) // 创建命令，打开新终端并进入指定目录
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true} // 隐藏启动时的黑框

	cmd.Stdout = os.Stdout // 设置标准输出
	cmd.Stderr = os.Stderr // 设置标准错误输出
	err := cmd.Start() // 执行命令
	if err != nil {
		WriteErrorToFile("Failed to open terminal", dir, err) // 如果执行失败，记录错误日志
		return fmt.Errorf("failed to open terminal: %w", err) // 返回错误信息
	}
	return nil
}
