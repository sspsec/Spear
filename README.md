

## 前言

​	在 Spear 工具箱 V2 版本，引入了 yaml 文件来管理工具，动态进行加载 GUI 页面，但是这样需要修改配置文件不够便捷优雅和直观，因此在 V3 版本中添加了两个按钮，一个是添加工具按钮，一个是删除工具按钮，本质上这两个按钮还是对 tool.yml 文件进行操作，添加的时候将执行信息添加到 yml 文件中，在删除的时候从 yml 文件中进行删除。除此之外，用户反馈想要打开终端的时候不打开系统自带的终端，而打开 iTerm，这个功能我添加了一个判断，如果电脑中存在 `/Applications/iTerm.app` 就会使用 iTerm 进行打开，否则将使用 mac 自带的终端。最后，朋友说页面颜色太丑了，所以换成了简单的白色与黑色，并且自适应系统的浅色模式和深色模式。

  V4版本改动如下
  新增右键功能：
  删除：会删除button按钮以及tool.yml文件中相关工具的配置
  修改：修改后会对应修改tool.yml中相关工具的配置信息
  打开目录：点击后会在访达中打开工具所在目录

## 演示

第一次运行会提示验证签名失败，执行如下命令即可正常执行：

```sh
xattr -rd com.apple.quarantine Spear.app
```

如果是 macOS Sequoia 测试版系统则需要执行：

```sh
codesign --sign - Spear.app
```

## 主页面

### 浅色主题
<img width="858" alt="image" src="https://github.com/user-attachments/assets/2fd6f65f-d7dc-49dd-bdd1-89598e870981">



### 深色主题
<img width="858" alt="image" src="https://github.com/user-attachments/assets/6158c3bd-31a8-4501-bb41-dbeffeba60b8">


## 添加工具功能

如添加 JYso-1.3.1.jar 这个工具，我们需要将这个工具目录复制到 app 包内 resources 文件夹内（其实也可以做其他文件夹的，但是方便统一管理就没加）。

<img width="858" alt="image" src="https://github.com/user-attachments/assets/76aec28c-c549-4f3e-9b55-7f7ead8a9aaf">


​	之后点击按钮进行添加，填写工具名称，工具路径为以 `resources` 开头的路径，因为在代码中进行拼接，以及 Mac 上 app 的资源目录在此，所以用了相对路径。执行文件名为 `Jyso-1.3.1.jar`，
因为其没有 GUI 界面，只能通过打开所在文件的终端进行运行，所以也可不写。运行方式一共有四种，如打开 fofaviewer 等带有 GUI 页面的工具时，可使用 Java8、Java11 进行打开，命令为 `-jar`。
可选参数是留给 CS、冰蝎、哥斯拉等需要 Java 的其他参数时用到的，因 JYso 没有 GUI 页面，所以这里使用 `openterm` 进行打开终端，选择类别进行提交即可添加成功。


还有一类运行方式为 `open`，需要添加执行 app 程序可以使用该运行方式，如打开 OSS-Browser.app。

<img width="858" alt="image" src="https://github.com/user-attachments/assets/44dfb6d0-f34b-4861-a029-479576b71c53">


打开后为打开工具所在的路径的终端。
<img width="858" alt="image" src="https://github.com/user-attachments/assets/2b8814bc-da65-4859-b30a-a44c4090b3da">

<img width="858" alt="image" src="https://github.com/user-attachments/assets/271afe88-292c-426c-bfc2-495c8fc2a305">


## 删除工具

右键工具名称即可删除。

<img width="858" alt="image" src="https://github.com/user-attachments/assets/92624560-5c36-43a4-ab26-4e78513c46fb">
<img width="858" alt="image" src="https://github.com/user-attachments/assets/a2d9c4a8-ad84-42d2-9f68-66d350d9dd92">


## 编译

注意：在 Intel 芯片的 Mac 上运行时，要将资源文件中的 Java 版本更换成 amd 版本的，自带的为 arm 架构的。

```sh
# M系列芯片编译
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon Icon.png

# Intel芯片编译
go install fyne.io/fyne/v2/cmd/fyne@latest
export GOOS=darwin
export GOARCH=amd64
fyne package -os darwin -icon Icon.png

# Windows编译
GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o Spear.exe main.go
```

## 蚁剑资源文件夹

```bash
Spear.app/Contents/Resources/resources/webshell/AntSword/antSword-2.1.15
```

## BurpSuite激活

```bash
/Applications/Spear.app/Contents/Resources/resources/pentest/BurpSuite/BurpSuite.app/Contents/Resources/app/BurpSuiteLoader.jar
```

参考国光的 BurpSuite 激活：[链接](https://www.sqlsec.com/2023/07/ventura.html#Burp-Suite)

## Windows 版本
由于很多粉丝朋友都问过我关于有没有Windows版本的问题，我就改了下代码编译了一下。

与Mac版本的基本没区别，只是增加了Python的运行环境，以及一些常用的python工具。

支持自定义添加python，java，GUI程序，打开命令行等功能。

我觉得像一个启动器了，把需要执行的命令，参数，等等写进去就会按照这个执行。

还是那个初衷，因为工具这个东西，每个人都有每个人顺手的工具，本项目呢只是提供了一个方便搜索，添加，删除，管理工具的一个框架，更好的管理常用的工具，当然这么一说，它的应用范围就广了，不仅仅局限于一个渗透工具的启动器了，也可以是其他任何的工具，看大家发挥了。有什么问题私信留言我，谢谢～
对于配置文件修改的问题可以去看看往期的Mac Spear工具箱V3的文章，这里就不再赘述。

自适应深色浅色主题

<img width="847" alt="image" src="https://github.com/user-attachments/assets/f9a27634-9d16-40e2-a4b3-3f8c9453d1f5">

<img width="846" alt="image" src="https://github.com/user-attachments/assets/e2914e49-0a42-4ac4-931f-0d8e29947c7f">

添加python工具

![](https://s2.loli.net/2024/08/24/FXkoxwDRUHS2rEi.png#errorMessage=unknown%20error&id=sy9Vu&originHeight=1536&originWidth=1700&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)

添加java工具

<img width="852" alt="image" src="https://github.com/user-attachments/assets/f95852eb-095a-4a6f-a86f-904cf173f45f">

对于exe这样的需要命令行调用的 直接打开工具所在目录，可只写路径名和工具名，运行方式为openterm

![](https://s2.loli.net/2024/08/24/N7e9jQzbU4q3BXl.png#errorMessage=unknown%20error&id=hGgoB&originHeight=1552&originWidth=1708&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)

添加像蚁剑这样点击exe就可以出现图形化界面的可以使用Open运行方式

![](https://s2.loli.net/2024/08/24/tzs17wReOQVrjoB.png#errorMessage=unknown%20error&id=u0phZ&originHeight=1550&originWidth=1706&originalType=binary&ratio=1&rotation=0&showTitle=false&status=error&style=none)

## 公众号
![image](https://github.com/user-attachments/assets/8d233519-0f1e-49bc-9b2a-c46ded91bbf9)

