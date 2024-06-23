# Spear
基于GO的Mac渗透工具箱框架

  引入了yaml 文件， 通过读取yaml 配置文件动态生成GUI页面且不用修改源代码， 方便师傅们对工具的删除与添加。在此感谢EDI团队的penguin师傅提出的宝贵建议，以及提供的部分参考源代码。在GUI页面中添加了一个实时搜索框，随着工具的不断增加，也是比较难找到按钮，所以就加了一个搜索框，方便查找工具。在第一个版本中，只支持了带有GUI 图形化界面的工具， 在这个版本中，我们增加了 点击没有图形化界面的工具按钮，直接打开终端进入所在的目录。这里我们只增加了一个这样的功能，例如dirsearch可以使用python或brew进行下载，能够使用这种方式安装的工具有很多，所以并未自带许多这种工具，这个接口的本意是提供给诸如GitHack等这种无法使用python或brew直接安装到Mac中的脚本工具，师傅们可以根据需要自行添加。

  下面以Spear.app为例显示包内容resources文件夹下是存放工具的地方
![image](https://github.com/sspsec/Spear/assets/142762749/be722db6-e1fe-48c7-ae89-1a8931e0a1f5)


  tool.yml 为GUI配置文，为核心文件，缺少它程序会启动不起来
  ![image](https://github.com/sspsec/Spear/assets/142762749/bb0dea4c-6f6a-484b-ab97-30a4cf824e08)

```

Java 8
路径：resources/java8/bin/java
这个路径指向Java 8的可执行文件，适用于需要Java 8环境的应用。
Java 11
路径：resources/java11/bin/java
这个路径指向Java 11的可执行文件，适用于需要Java 11环境的应用。
打开方式
命令：open
该命令用于打开或执行文件，具体依赖于操作系统的配置。

  - CategoryName: 信息收集
    Tools: #信息手机下包含的信息收集工具
      ToolName: 工具名称 即显示的按钮名字
        PATH: 工具所在的文件夹，相对路径 都在resources文件夹下
        FileName: 需要启动的Jar文件名称
        VALUE: 使用Java哪个版本执行 java8或java11
        COMMAND: -jar
        Optional: 添加参数 例如-Xdock:icon=godzilla.icns -Dfile.encoding=UTF-8
        
        ToolName: 工具名称 即显示的按钮名字
        PATH: 工具所在的文件夹，相对路径 都在resources文件夹下
        FileName: 需要启动的Jar文件名称
        VALUE: 使用Java哪个版本执行 java8或java11
        COMMAND: -jar
        Optional: 添加参数 例如-Xdock:icon=godzilla.icns -Dfile.encoding=UTF-8
   - CategoryName: 渗透利器
    Tools:
      - ToolName: BurpSuite
        PATH: resources/pentest/BurpSuite
        FileName: BurpSuite.app
        VALUE: Open
        COMMAND:
        Optional:
```

对应关系图如下
![image](https://github.com/sspsec/Spear/assets/142762749/c9379ea8-4e09-45ad-800c-8b0b704438fc)

可以根据需求自行修改tool.yml文件
第一次运行会提示验证签名失败 执行如下命令 即可正常执行

```
xattr -rd com.apple.quarantine Spear.app
```

搜索功能 快捷方便的实时搜索
![image](https://github.com/sspsec/Spear/assets/142762749/447548a9-778e-4040-9dd3-21afc17c291e)

非GUI程序终端打开所在路径
![image](https://github.com/sspsec/Spear/assets/142762749/72b03d61-ea2c-4053-a731-b174eb57e4e9)
编译

注意：在Intel芯片的Mac上运行时 要将资源文件中的java版本更换成amd版本的 自带的为arm架构的
```
#M系列芯片编译
go install fyne.io/fyne/v2/cmd/fyne@latest
fyne package -os darwin -icon Icon.png

#Intel芯片编译
go install fyne.io/fyne/v2/cmd/fyne@latest
export GOOS=darwin 
export GOARCH=amd64
fyne package darwin-icon Icon.png
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
