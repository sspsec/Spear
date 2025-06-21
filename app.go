package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"

	wailsRuntime "github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp 创建新的 App 应用
func NewApp() *App {
	return &App{}
}

// startup 在应用启动时调用
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Config 配置结构体
type Config struct {
	JavaPath struct {
		Java8  string `yaml:"Java8"`
		Java11 string `yaml:"Java11"`
		Java17 string `yaml:"Java17"`
		Open   string `yaml:"Open"`
	} `yaml:"javapath"`
}

// Tool 工具结构体
type Tool struct {
	Name        string `yaml:"ToolName"`
	Path        string `yaml:"PATH"`
	FileName    string `yaml:"FileName"`
	Value       string `yaml:"VALUE"`
	Command     string `yaml:"COMMAND"`
	Optional    string `yaml:"Optional"`
	Description string `yaml:"Description,omitempty"` // 添加描述字段
	OpenCount   int    `yaml:"OpenCount,omitempty"`   // 添加打开次数字段
}

// Category 分类结构体
type Category struct {
	Name string `yaml:"CategoryName"`
	Tool []Tool `yaml:"Tools"`
}

// Categories 分类列表结构体
type Categories struct {
	Category []Category `yaml:"Categories"`
}

// GetCategories 获取所有工具分类
func (a *App) GetCategories() (Categories, error) {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")
	var categories Categories
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return categories, fmt.Errorf("读取配置文件失败: %v", err)
	}

	if err := yaml.Unmarshal(data, &categories); err != nil {
		return categories, fmt.Errorf("解析配置文件失败: %v", err)
	}

	return categories, nil
}

// ExecuteCommand 执行工具命令
func (a *App) ExecuteCommand(path, optional, value, filename string) error {
	// 读取配置文件
	var config Config
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析配置文件失败: %v", err)
	}

	// 获取当前工作目录
	currentDir := a.getResourcePath()
	fmt.Printf("当前工作目录: %s\n", currentDir)

	var execCommand string
	switch value {
	case "Java8", "Java11", "Java17":
		// 构建Java可执行文件路径
		var javaExec string
		if value == "Java8" {
			javaExec = filepath.Join(currentDir, config.JavaPath.Java8)
		} else if value == "Java11" {
			javaExec = filepath.Join(currentDir, config.JavaPath.Java11)
		} else if value == "Java17" {
			javaExec = filepath.Join(currentDir, config.JavaPath.Java17)
		}

		// 构建工具路径
		toolPath := filepath.Join(currentDir, path)
		jarPath := filepath.Join(toolPath, filename)

		fmt.Printf("Java可执行文件: %s\n", javaExec)
		fmt.Printf("工具目录: %s\n", toolPath)
		fmt.Printf("JAR文件: %s\n", jarPath)
		fmt.Printf("可选参数: %s\n", optional)

		// 检查Java可执行文件是否存在
		if _, err := os.Stat(javaExec); err != nil {
			return fmt.Errorf("Java可执行文件不存在: %s", javaExec)
		}

		// 检查JAR文件是否存在
		if _, err := os.Stat(jarPath); err != nil {
			return fmt.Errorf("JAR文件不存在: %s", jarPath)
		}

		// 构建执行命令
		execCommand = fmt.Sprintf("cd '%s' && '%s' %s -jar '%s'",
			toolPath, javaExec, optional, filename)

	case "Open":
		toolPath := filepath.Join(currentDir, path)
		execCommand = fmt.Sprintf("cd '%s' && open '%s'",
			toolPath, filename)
	case "openterm":
		toolPath := filepath.Join(currentDir, path)
		return openTerminal(toolPath)
	default:
		return fmt.Errorf("不支持的命令类型: %s", value)
	}

	fmt.Println("执行命令:", execCommand)

	cmd := exec.Command("sh", "-c", execCommand)

	// 设置标准输出和错误输出
	var stdout, stderr bytes.Buffer
	cmd.Stdout = io.MultiWriter(os.Stdout, &stdout)
	cmd.Stderr = io.MultiWriter(os.Stderr, &stderr)

	// 执行命令
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("执行命令失败: %v\n标准输出: %s\n错误输出: %s",
			err, stdout.String(), stderr.String())
	}

	// 等待命令完成
	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("命令执行出错: %v\n标准输出: %s\n错误输出: %s",
			err, stdout.String(), stderr.String())
	}

	// 输出执行结果
	fmt.Printf("命令执行完成\n标准输出: %s\n错误输出: %s\n",
		stdout.String(), stderr.String())

	return nil
}

// openTerminal 打开终端的辅助函数
func openTerminal(dir string) error {
	switch runtime.GOOS {
	case "darwin":
		itermPath := "/Applications/iTerm.app"
		if _, err := os.Stat(itermPath); err == nil {
			script := fmt.Sprintf(`tell application "iTerm"
				create window with default profile
				tell current session of current window
					write text "cd %s; ls --color=always"
				end tell
			end tell`, dir)
			cmd := exec.Command("osascript", "-e", script)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				return fmt.Errorf("打开终端失败: %v", err)
			}
		} else {
			script := fmt.Sprintf(`tell application "Terminal"
				do script "cd %s; ls --color=always"
			end tell`, dir)
			cmd := exec.Command("osascript", "-e", script)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Start()
			if err != nil {
				return fmt.Errorf("打开终端失败: %v", err)
			}
		}
	default:
		return fmt.Errorf("不支持的平台")
	}
	return nil
}

// 添加常量定义
const defaultConfig = `javapath:
    Java8: resources/java8/bin/java
    Java11: resources/java11/bin/java
    Java17: resources/java17/bin/java
    Open: ""
`

// AddTool 添加新工具
func (a *App) AddTool(tool Tool, categoryName string) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")

	// 读取原始YAML内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析Categories和Config
	var categories Categories
	var config Config
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析Config失败: %v", err)
	}

	// 检查工具名称是否已存在
	for _, category := range categories.Category {
		for _, existingTool := range category.Tool {
			if existingTool.Name == tool.Name {
				return fmt.Errorf("工具名称 '%s' 已存在", tool.Name)
			}
		}
	}

	// 查找分类并添加工具
	categoryFound := false
	for i, category := range categories.Category {
		if category.Name == categoryName {
			categories.Category[i].Tool = append(categories.Category[i].Tool, tool)
			categoryFound = true
			break
		}
	}

	if !categoryFound {
		// 如果分类不存在，创建新分类
		newCategory := Category{
			Name: categoryName,
			Tool: []Tool{tool},
		}
		categories.Category = append(categories.Category, newCategory)
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释和默认配置
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# Java 17
# 路径：resources/java17/bin/java
# 这个路径指向Java 17的可执行文件，适用于需要Java 17环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

` + defaultConfig

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Categories数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	// 发送更新成功事件
	wailsRuntime.EventsEmit(a.ctx, "tool-added", true)
	return nil
}

func (a *App) getResourcePath() string {
	if execPath, err := os.Executable(); err == nil {
		// 对于 .app 包，可执行文件在 Contents/MacOS 目录下
		// 资源文件在 Contents/Resources 目录下
		if strings.HasSuffix(execPath, "/Contents/MacOS/Spear") {
			return filepath.Join(filepath.Dir(execPath), "../Resources")
		}
		return filepath.Dir(execPath)
	}
	return "."
}

// DeleteTool 删除工具
func (a *App) DeleteTool(toolName, categoryName string) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")
	var categories Categories
	var config Config

	// 读取原始YAML内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析Categories和Config
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析Config失败: %v", err)
	}

	// 查找并删除工具
	for i, category := range categories.Category {
		if category.Name == categoryName {
			for j, tool := range category.Tool {
				if tool.Name == toolName {
					// 删除工具
					categories.Category[i].Tool = append(
						categories.Category[i].Tool[:j],
						categories.Category[i].Tool[j+1:]...,
					)
					break
				}
			}
		}
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

`
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Config数据
	if newData, err := yaml.Marshal(config); err != nil {
		return fmt.Errorf("序列化Config失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Config数据失败: %v", err)
	}

	// 写入Categories数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	return nil
}

// OpenToolDirectory 打开工具所在目录
func (a *App) OpenToolDirectory(path string) error {
	fullPath := filepath.Join(a.getResourcePath(), path)
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", fullPath)
	case "windows":
		cmd = exec.Command("explorer", fullPath)
	default:
		cmd = exec.Command("xdg-open", fullPath)
	}

	return cmd.Run()
}

// GetToolTypes 获取支持的工具类型
func (a *App) GetToolTypes() []string {
	return []string{"Java8", "Java11", "Open"}
}

// GetFilePath 获取文件的完整路径
func (a *App) GetFilePath(fileName string) (string, error) {
	// 这里可以根据需要处理文件路径
	// 例如，如果文件在特定目录下：
	return filepath.Abs(fileName)
}

// GetFileInfo 获取文件信息
func (a *App) GetFileInfo(filePath string) (map[string]string, error) {
	absPath, err := filepath.Abs(filePath)
	if err != nil {
		return nil, fmt.Errorf("无法获取绝对路径: %v", err)
	}

	dir := filepath.Dir(absPath)
	fileName := filepath.Base(absPath)

	return map[string]string{
		"path":     dir,
		"fileName": fileName,
		"fullPath": absPath,
	}, nil
}

// OpenFileDialog 打开文件选择对话框
func (a *App) OpenFileDialog() (map[string]string, error) {
	filePath, err := wailsRuntime.OpenFileDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		DefaultDirectory: "",
		DefaultFilename:  "",
		Title:            "选择工具文件",
		Filters: []wailsRuntime.FileFilter{
			{
				DisplayName: "所有文件",
				Pattern:     "*.*",
			},
			{
				DisplayName: "Java 文件",
				Pattern:     "*.jar",
			},
			{
				DisplayName: "可执行文件",
				Pattern:     "*.exe;*.sh;*.bat;*.cmd",
			},
		},
	})

	if err != nil {
		return nil, fmt.Errorf("选择文件失败: %v", err)
	}

	if filePath == "" {
		return nil, nil
	}

	return a.GetFileInfo(filePath)
}

// OpenDirectoryDialog 打开目录选择对话框
func (a *App) OpenDirectoryDialog() (string, error) {
	dirPath, err := wailsRuntime.OpenDirectoryDialog(a.ctx, wailsRuntime.OpenDialogOptions{
		DefaultDirectory: "",
		Title:            "选择工具目录",
	})

	if err != nil {
		return "", fmt.Errorf("选择目录失败: %v", err)
	}

	return dirPath, nil
}

// UpdateTool 更新工具信息
func (a *App) UpdateTool(originalName, categoryName string, tool Tool) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")

	// 读取原始YAML内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析Categories和Config
	var categories Categories
	var config Config
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析Config失败: %v", err)
	}

	// 查找并更新工具
	found := false
	for i, category := range categories.Category {
		for j, t := range category.Tool {
			if t.Name == originalName {
				if categoryName == category.Name {
					// 如果在同一分类中，直接更新工具
					found = true
					categories.Category[i].Tool[j] = tool
					break
				} else {
					// 如果在不同分类中，从原分类删除
					categories.Category[i].Tool = append(
						categories.Category[i].Tool[:j],
						categories.Category[i].Tool[j+1:]...)
				}
			}
		}
		if found {
			break
		}
	}

	// 如果没有在原分类中找到或需要移动到新分类，则添加到目标分类
	if !found {
		for i, category := range categories.Category {
			if category.Name == categoryName {
				categories.Category[i].Tool = append(categories.Category[i].Tool, tool)
				break
			}
		}
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释和默认配置
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# Java 17
# 路径：resources/java17/bin/java
# 这个路径指向Java 17的可执行文件，适用于需要Java 17环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

`
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Categories和Config数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	if newData, err := yaml.Marshal(config); err != nil {
		return fmt.Errorf("序列化Config失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Config数据失败: %v", err)
	}

	// 发送更新成功事件
	wailsRuntime.EventsEmit(a.ctx, "tool-updated", true)
	return nil
}

// DeleteCategory 删除分类及其下的所有工具
func (a *App) DeleteCategory(categoryName string) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")

	// 读取原始YAML内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析Categories
	var categories Categories
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}

	// 查找并删除分类
	found := false
	for i, category := range categories.Category {
		if category.Name == categoryName {
			// 删除分类
			categories.Category = append(
				categories.Category[:i],
				categories.Category[i+1:]...,
			)
			found = true
			break
		}
	}

	if !found {
		return fmt.Errorf("分类 '%s' 不存在", categoryName)
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释和默认配置
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# Java 17
# 路径：resources/java17/bin/java
# 这个路径指向Java 17的可执行文件，适用于需要Java 17环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

` + defaultConfig

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Categories数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	return nil
}

// UpdateCategoryTools 批量更新分类下工具顺序
func (a *App) UpdateCategoryTools(categoryName string, tools []Tool) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")
	var categories Categories
	var config Config

	// 读取配置文件
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 同时解析Categories和Config结构
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}
	if err := yaml.Unmarshal(data, &config); err != nil {
		return fmt.Errorf("解析Config失败: %v", err)
	}

	// 找到对应分类并更新工具顺序
	for i, category := range categories.Category {
		if category.Name == categoryName {
			categories.Category[i].Tool = tools
			break
		}
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释和默认配置
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# Java 17
# 路径：resources/java17/bin/java
# 这个路径指向Java 17的可执行文件，适用于需要Java 17环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

`
	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Config数据（包含Java路径）
	if newData, err := yaml.Marshal(config); err != nil {
		return fmt.Errorf("序列化Config失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Config数据失败: %v", err)
	}

	// 写入Categories数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	// 发送更新成功事件
	wailsRuntime.EventsEmit(a.ctx, "tool-updated", true)
	return nil
}

// UpdateToolDescription 更新工具描述
func (a *App) UpdateToolDescription(toolName, categoryName, description string) error {
	configPath := filepath.Join(a.getResourcePath(), "tool.yml")

	// 读取原始YAML内容
	data, err := ioutil.ReadFile(configPath)
	if err != nil {
		return fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析Categories
	var categories Categories
	if err := yaml.Unmarshal(data, &categories); err != nil {
		return fmt.Errorf("解析Categories失败: %v", err)
	}

	// 查找并更新工具描述
	toolFound := false
	for i, category := range categories.Category {
		if category.Name == categoryName {
			for j, tool := range category.Tool {
				if tool.Name == toolName {
					// 只更新描述字段
					categories.Category[i].Tool[j].Description = description
					toolFound = true
					break
				}
			}
			if toolFound {
				break
			}
		}
	}

	if !toolFound {
		return fmt.Errorf("未找到工具: %s", toolName)
	}

	// 创建或打开文件
	file, err := os.OpenFile(configPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("打开配置文件失败: %v", err)
	}
	defer file.Close()

	// 写入注释和默认配置
	content := `# Java 8
# 路径：resources/java8/bin/java
# 这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
# Java 11
# 路径：resources/java11/bin/java
# 这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
# Java 17
# 路径：resources/java17/bin/java
# 这个路径指向Java 17的可执行文件，适用于需要Java 17环境的应用。
# 打开方式
# 命令：open
# 该命令用于打开或执行文件，具体依赖于操作系统的配置。

` + defaultConfig

	if _, err := file.WriteString(content); err != nil {
		return fmt.Errorf("写入配置失败: %v", err)
	}

	// 写入Categories数据
	if newData, err := yaml.Marshal(categories); err != nil {
		return fmt.Errorf("序列化Categories失败: %v", err)
	} else if _, err := file.Write(newData); err != nil {
		return fmt.Errorf("写入Categories数据失败: %v", err)
	}

	// 发送更新成功事件
	wailsRuntime.EventsEmit(a.ctx, "tool-updated", true)
	return nil
}
