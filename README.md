# cau_go
cau 中国农业大学 校园网认证 cli golang版本
```
go install github.com/NingYuanLin/cau_go@latest
```
> 请确保$GOPATH/bin在PATH环境变量中

[python版本看这里，使用pip安装更简单](https://github.com/NingYuanLin/cau_auth)

## 使用方法
### 1. 创建配置文件
```
cau_go config -c
```
### 2. 登录
```
cau_go login
或
cau_go i
```
> 如果没有创建配置文件，可以通过`cau_go login -u 学号 -p 密码`来登录
### 3. 登出
```
cau_go logout
或
cau_go o
```
### 4. 检查登录状态
```
cau_go status
或
cau_go s
```

## Shell autocomplete 
shell autocomplete for your application (bash, zsh, fish, powershell)  
[使用方法](https://github.com/spf13/cobra/blob/main/shell_completions.md)
<img width="1341" alt="图片" src="https://user-images.githubusercontent.com/57001533/192970538-65aace5f-2668-49bb-b313-6e60cfe99490.png">


## 与python版本的区别
1. 新的命令格式，即`command + flag`,这与`git`等工具类似。
2. 配置文件格式现在为`.yaml`。
3. 支持shell自动补全。
