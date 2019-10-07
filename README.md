# CLI

# 服务计算命令行开发

## 17343140 杨泽涛

## 实验目的

使用 golang 开发 [开发 Linux 命令行实用程序](https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html) 中的 **selpg**

**！！！！ 插入图片好像有错，请TA看一下pdf版报告，麻烦ta了！！！！！！**

提示：

1. 请按文档 **使用 selpg** 章节要求测试你的程序
2. 请使用 pflag 替代 goflag 以满足 Unix 命令行规范， 参考：[Golang之使用Flag和Pflag](https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/)
3. golang 文件读写、读环境变量，请自己查 os 包
4. “-dXXX” 实现，请自己查 `os/exec` 库，例如案例 [Command](https://godoc.org/os/exec#example-Command)，管理子进程的标准输入和输出通常使用 `io.Pipe`，具体案例见 [Pipe](https://godoc.org/io#Pipe)

## 实验内容

### 处理参数部分

根据网站要求，使用pflag来处理参数，相关代码如下

<img src="/Users/yang/Library/Application Support/typora-user-images/image-20191007154438409.png" alt="image-20191007154438409" style="zoom:50%;" />

![image-20191007154420611](/Users/yang/Library/Application Support/typora-user-images/image-20191007154420611.png)

根据selpg文档要求，我们设置参数报错程序如下

![image-20191007154549699](/Users/yang/Library/Application Support/typora-user-images/image-20191007154549699.png)

### 读写创建部分

我们根据参数的类型来为我们的输入输出创建正确的对象：

![image-20191007154846739](/Users/yang/Library/Application Support/typora-user-images/image-20191007154846739.png)

### 程序读写部分

我们根据定长和不定长来进行不同类型的读写操作

![image-20191007154949521](/Users/yang/Library/Application Support/typora-user-images/image-20191007154949521.png)

## 相关测试

1.`selpg -s1 -e1 input_file`

![image-20191007155044202](/Users/yang/Library/Application Support/typora-user-images/image-20191007155044202.png)

2.`selpg -s1 -e1 < input_file`

![image-20191007155134923](/Users/yang/Library/Application Support/typora-user-images/image-20191007155134923.png)

3.`other_command | selpg -s10 -e20`

改成了-s1 -e1 因为小文件方便修改

![image-20191007155350122](/Users/yang/Library/Application Support/typora-user-images/image-20191007155350122.png)

4.`selpg -s10 -e20 input_file >output_file`

![image-20191007155434432](/Users/yang/Library/Application Support/typora-user-images/image-20191007155434432.png)

![image-20191007155453632](/Users/yang/Library/Application Support/typora-user-images/image-20191007155453632.png)

5.`$ selpg -s10 -e20 input_file 2>error_file`

![image-20191007155910201](/Users/yang/Library/Application Support/typora-user-images/image-20191007155910201.png)

![image-20191007155925664](/Users/yang/Library/Application Support/typora-user-images/image-20191007155925664.png)

6.`selpg -s10 -e20 input_file >output_file 2>error_file`

![image-20191007155722014](/Users/yang/Library/Application Support/typora-user-images/image-20191007155722014.png)

![image-20191007155642502](/Users/yang/Library/Application Support/typora-user-images/image-20191007155642502.png)

7.`selpg -s10 -e20 input_file >output_file 2>/dev/null`

![image-20191007160048116](/Users/yang/Library/Application Support/typora-user-images/image-20191007160048116.png)

8.`$ selpg -s10 -e20 input_file >/dev/null`

![image-20191007160134656](/Users/yang/Library/Application Support/typora-user-images/image-20191007160134656.png)

9.`selpg -s10 -e20 input_file | other_command`

![image-20191007161543907](/Users/yang/Library/Application Support/typora-user-images/image-20191007161543907.png)

10`selpg -s10 -e20 input_file 2>error_file | other_command`

效果和之前类似，不再重复。

11.`selpg -s10 -e20 input_file > output_file 2>error_file &`

![image-20191007161738362](/Users/yang/Library/Application Support/typora-user-images/image-20191007161738362.png)

测试中关于-f和-d的命令我没有测试，因为没有终端和有换页符的文件，测试中大多使用的是小文件，大文件检测在下面

![image-20191007170139917](/Users/yang/Library/Application Support/typora-user-images/image-20191007170139917.png)

in.txt行数

![image-20191007170209904](/Users/yang/Library/Application Support/typora-user-images/image-20191007170209904.png)

out.txt行数

![image-20191007170235750](/Users/yang/Library/Application Support/typora-user-images/image-20191007170235750.png)

指定一页12行后

![image-20191007170301538](/Users/yang/Library/Application Support/typora-user-images/image-20191007170301538.png)

![image-20191007170314945](/Users/yang/Library/Application Support/typora-user-images/image-20191007170314945.png)