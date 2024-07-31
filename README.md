**1**►**前言**

  在 Spear 工具箱 V2 版本，引入了 yaml 文件来管理工具，动态进行加载 GUI 页面，但是这样需要修改配置文件不够便捷优雅和直观，因此在 V3 版本中添加了两个按钮，一个是添加工具按钮，一个是删除工具按钮，本质上这两个按钮还是对 tool.yml文件进行操作，添加的时候将执行信息添加到 yml 文件中，在删除的时候从 yml 文件中进行删除。

  除此之外，用户反馈想要打开终端的时候不打开系统自带的终端，而打开 iterm，这个功能我添加了一个判断，如果电脑中存在/Applications/iTerm.app 就会使用 iterm 进行打开，否则将使用 mac 自带的终端。

  最后，朋友说页面颜色太丑了，所以换成了简单的白色与黑色，并且自适应系统的浅色模式和深色模式。



**2**►**演示**

第一次运行会提示验证签名失败 执行如下命令 即可正常执行

```
xattr -rd com.apple.quarantine Spear.app
```

如果是 macos sequoia 测试版系统则需要执行

```
codesign --sign - Spear.app
```

主页面 浅色主题

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGpdXd6S8Zf90fWzSxRvc783gjDcU6mBcNcVVuI8g19UfYf6YJJXQyGhw/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

深色主题

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGprpgoItYUT6f5QZv9YkpYl7rGzH9Go6Y7dK20iaV1TaGBjTwKvb6bF2g/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)

添加工具功能

如添加JYso-1.3.1.jar 这个工具 我们需要将这个工具目录复制到

app 包内 resources 文件夹内（其实也可以做其他文件夹的，但是方便统一管理就没加）

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGp84zTRLk852n2k7FHUCAP43Z8ougZIoibNuBppvW9AI4QNKgLxwvrsTg/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)



之后点击按钮进行添加

填写工具名称

工具路径为，以 resources 开头的路径，因为在代码中进行拼接，以及 Mac 上 app 的资源目录在此，所以用了相对路径

执行文件名为 Jyso-1.3.1.jar 因为其没有 GUI 界面，只能通过打开所在文件的终端进行运行，所以也可不写



运行方式一共有四种，如打开 fofaviewer 等带有 GUI 页面的工具时，可使用 Java8、Java11 进行打开 命令为-jar 可选参数是留给 CS、冰蝎、哥斯拉等需要 java 的其他参数时用到的，因JYso没有 GUI 页面所以，这里使用 openterm 进行打开终端，选择类别进行提交即可添加成功。

还有一类运行方式为 open，需要添加执行 app 程序可以使用该运行方式，如打开 yakit.app

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGpR4IlIsK0s3emVpUEqUDmoAKDMB48e97SAqB1bcslSvGOIsVqBmJTFw/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)


![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGp2Tp0icpYI7EuicERUdibVqlib2TfdurX5Hhqnkicib0yQRibkLpvdGVAuLIGQ/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)


打开后为打开工具所在的路径的终端

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGpdKBibX5Uq0ibPBmQlibhjdUicdOlftHPgnHyKfec4icfmZWn9J8fAY0HkPg/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)


删除工具只需要选择类别，工具名称即可删除

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGpl4LdaOzaicCF4PsVK6YsoRAYEeCv4exOatqG8p2lQ5HeZqaJXicI3J7g/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)


**3**►**其他**

编译

注意：在Intel芯片的Mac上运行时 要将资源文件中的java版本更换成amd版本的 自带的为arm架构的

```
#M系列芯片编译go install fyne.io/fyne/v2/cmd/fyne@latestfyne package -os darwin -icon Icon.png
#Intel芯片编译go install fyne.io/fyne/v2/cmd/fyne@latestexport GOOS=darwin export GOARCH=amd64fyne package darwin-icon Icon.png
#windows编译GOOS=windows GOARCH=amd64 go build -ldflags -H=windowsgui -o Spear.exe main.go
```

蚁剑资源文件夹

```
Spear.app/Contents/Resources/resources/webshell/AntSword/antSword-2.1.15
```

BurpSuite激活

```
/Applications/Spear.app/Contents/Resources/resources/pentest/BurpSuite/BurpSuite.app/Contents/Resources/app/BurpSuiteLoader.jar
```

参考国光的BurpSuite激活

https://www.sqlsec.com/2023/07/ventura.html#Burp-Suite

关于 Windows 版本

还有些小问题目前还在优化中～ 尽请期待 也可自行对代码进行修改编译。

有什么问题欢迎与我沟通。

![图片](https://mmbiz.qpic.cn/mmbiz_png/Oum0kexPoVojLE0ZYA4h4Wm9m7TJXeGp37mSwJUmmia2FJSTvicSMg9fbtDhEzggUvKyvEcib9NMRz2C4Qv1trPfA/640?wx_fmt=png&from=appmsg&tp=webp&wxfrom=5&wx_lazy=1&wx_co=1)
