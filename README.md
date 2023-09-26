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