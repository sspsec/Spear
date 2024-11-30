## 前言

在 **Spear 工具箱 V2 版本** 中，引入了 **YAML 文件** 来管理工具并动态加载 GUI 页面。虽然这种方式可以实现工具管理，但修改配置文件的方式不够便捷和直观。因此在 **V3 版本** 中，新增了两个按钮：**添加工具** 和 **删除工具**。这两个按钮本质上仍然是对 `tool.yml` 文件进行操作，添加工具时会将执行信息添加到 YAML 文件中，删除工具时则会从 YAML 文件中删除对应的工具配置。

此外，用户反馈希望能够在打开终端时避免使用系统自带的终端，而是使用 **iTerm**。为此，我添加了一个判断条件：如果电脑中存在 `/Applications/iTerm.app`，则使用 **iTerm**，否则使用 macOS 自带的终端。

最后，朋友们指出页面颜色不够好看，因此我更换了主题颜色，使用了简洁的 **白色** 和 **黑色**，并且支持自适应系统的浅色模式和深色模式。

------

## V4版本改动

### 新增右键功能：

1. **删除**：右键点击工具名称，删除按钮及工具配置。
2. **修改**：右键点击工具名称，修改对应工具的配置信息。
3. **打开目录**：右键点击工具名称，在 **访达** 中打开工具所在的目录。

------

## 演示

**首次运行提示验证签名失败**，执行以下命令可以正常启动：

```
xattr -rd com.apple.quarantine Spear.app
```

对于 macOS Sequoia 测试版系统，执行：

```
codesign --sign - Spear.app
```

------

### 主页面展示

#### 浅色主题

![浅色主题](https://github.com/user-attachments/assets/2fd6f65f-d7dc-49dd-bdd1-89598e870981)

#### 深色主题

![深色主题](https://github.com/user-attachments/assets/6158c3bd-31a8-4501-bb41-dbeffeba60b8)

------

## 添加工具功能

### 添加工具流程：

1. 将工具目录（例如 `JYso-1.3.1.jar`）复制到 **app 包内 `resources` 文件夹** 中。
2. 点击 **添加工具** 按钮，填写工具名称及路径，路径应以 `resources` 开头（这与代码中的拼接方式相关）。
3. 执行文件名为 `Jyso-1.3.1.jar`，因为该工具没有 GUI 界面，运行时会通过终端启动。选择 **openterm** 方式在终端中打开工具目录。

### 不同运行方式：

- **Java 应用**：可以选择 `Java8` 或 `Java11` 执行，例如：`-jar`。

- **打开 APP**：例如打开 `OSS-Browser.app`，使用 `open` 运行方式。

  **示例**：打开终端并切换到工具目录：

  ![打开工具所在路径的终端](https://github.com/user-attachments/assets/2b8814bc-da65-4859-b30a-a44c4090b3da)

------

## 删除工具

右键点击工具名称，可以删除对应工具和 `tool.yml` 配置文件中的相关内容。

![删除工具](https://github.com/user-attachments/assets/92624560-5c36-43a4-ab26-4e78513c46fb)

------

## 编译

### 编译命令：

1. M系列芯片编译

   ：

   ```
   go install fyne.io/fyne/v2/cmd/fyne@latest
   fyne package -os darwin -icon Icon.png
   ```

2. Intel芯片编译

   ：

   ```
   export GOOS=darwin
   export GOARCH=amd64
   fyne package -os darwin -icon Icon.png
   ```

3. Windows编译

   ：

   ```
   GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o Spear.exe main.go
   ```

------

## 蚁剑资源文件夹

```
Spear.app/Contents/Resources/resources/webshell/AntSword/antSword-2.1.15
```

------

## BurpSuite 激活

```
/Applications/Spear.app/Contents/Resources/resources/pentest/BurpSuite/BurpSuite.app/Contents/Resources/app/BurpSuiteLoader.jar
```

------

## Windows 版本

与 Mac 版本基本相同，只是增加了 Python 环境及一些常用 Python 工具。支持自定义添加 Python、Java、GUI 程序等功能，用户可以根据自己的需求选择工具管理方式。

------

## 公众号

![公众号二维码](https://github.com/user-attachments/assets/8d233519-0f1e-49bc-9b2a-c46ded91bbf9)

------
