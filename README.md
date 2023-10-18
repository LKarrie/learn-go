# learn-go
learn-go

## 安装
```markdown
# 安装完成 查看版本
go version

# 配置环境
go env -w GO111MODULE=on 
go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

# 常用命令
# 参考 https://golang.org/doc/cmd
build: 编译包和依赖
clean: 移除对象文件
doc: 显示包或者符号的文档
env: 打印环境信息
bug: 启动错误报告
fix：运行 go tool fix
fmt：运行 gofmt 进行格式化
generate：从 processing source 生成go文件
get：下载并安装包和依赖
install：编译并安装包和依赖
list：列出包
run：编译并运行go程序
test：运行测试
tool：运行go提供的工具
vet：运行go tool vet

# 快捷键
行注释 ctrl + /
块注释 shift+alt+a (可以修改为ctrl+shift+/)
ctrl+a 全选
ctrl+shift+k 删除行
ctrl+e 查找文件
ctrl+shift+p 打开设置命令行

# 代码块快捷键
pkgm  main包+main主函数
ff  fmt.Printf("", var)
for for i := 0; i < count; i++ {}
forr for _, v := range v {}
fmain func main() {}
a.print! fmt.Printf("a: %v\n", a)
```

## Golang包

包可以区分命令空间（一个文件夹中不能有两个同名文件），也可以更好管理项目。go中创建一个包，一般是创建一个文件夹，在该文件夹里的go文件中，使用package声明包名称，通常文件夹名称和包名称相同。并且，同一个文件下只有一个包

注意：

* 一个文件夹下只能有一个package
  * import后其实是GOPATH开始的相对目录路径，包括最后一段。但是由于一个目录下只能有一个package，所以import一个路径就等于是import了这个路径下的包
  * 如果有子目录，子目录和父目录是两个包
* 比如实现了一个计算package，名为calc，位于calc目录下，但是又希望有个范例，可以在calc下建立一个example子目录（calc/example/），这个子目录有个example.go（calc/example/example.go）。此时这里的example.go可以是main包，里面还可以有个main函数
* 一个package的文件不能在多个文件夹下
  * 如果多个文件下有重名的package，它们其实是彼此无关的
  * 如果一个go文件需要同时使用不同目录下的同名package，需要在import这些目录时为每个目录指定一个package别名

## Golang包管理go module

go modules 是 go 1.11 新特性，用来管理模块中包的依赖关系

### go mod使用

```markdown
# 初始化模块
go mod init <项目名称>
# 依赖关系处理 根据go.mod文件 
go mod tity
# 将依赖包复制到项目下的 vendor 目录
# 如果包被墙了 可以使用这个命令 随后使用 go build -mod=vendor编译
go mod vendor
# 显示依赖关系
go list -m all
# 显示详细依赖关系 
go list -m -json all
# 下载依赖 
# [path@version] 非必写
go mod download [path@version]
```

