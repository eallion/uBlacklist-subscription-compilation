package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// 通过两重循环过滤重复元素
func RemoveRepeatedElement(arr []string) []string {
	// 存放结果，空初始化比make定义要快
	result := []string{}
	// 外层循环准备添加到结果的切片
	for i := 0; i < len(arr); i++ {
		// 初始定义该元素不存在，很奇怪，初始在循环里面比先声明在外面之后再赋值要快
		exist := false
		// 这里根据当前切片的长度进行循环，直接使用 len 比初始一个 count变量 记数要快
		for j := 0; j < len(result); j++ {
			// 如果遇到重复提前退出
			if result[j] == arr[i] {
				// 并且说明已存在
				exist = true
				break
			}
		}
		// 如果在 result切片都没有遍历到此元素
		if !exist {
			// 那么就追加到 result
			result = append(result, arr[i])
		}
	}
	return result
}

// 读取文件，返回字符串
func ReadFile(Path string) (result string, err error) {
	// 只读方式打开文件
	f, err1 := os.Open(Path)
	if err1 != nil {
		fmt.Println("os.Open err", err1)
		err = err1
		return
	}
	defer f.Close()

	// 读取内容保存到这个切片
	var content []byte
	buf := make([]byte, 4096)
	for {
		n, err2 := f.Read(buf)
		if err2 != nil && err2 != io.EOF {
			fmt.Println("f.Read err:", err2)
			err = err2
			return
		}
		if err2 == io.EOF {
			break
		}
		// 合并切片
		content = append(content, buf[:n]...)
	}
	// 字符切片转字符串
	result = string(content)
	return
}

// 创建文件并写入数据
func CreateFile(path, data string) (err error) {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer f.Close()

	// 第一个参数返回数据长度，第二个，错误信息
	_, err = f.WriteString(data)
	if err != nil {
		fmt.Println("f.WriteString err:", err)
		return
	}
	return
}

// 备份文件，避免丢失
func BackupFile(dst, src string) (err error) {
	// 源文件
	fSrc, err := os.Open(src)
	if err != nil {
		fmt.Println("os.Open err:", err)
		return
	}
	defer fSrc.Close()

	// 目标文件
	fDst, err := os.Create(dst)
	if err != nil {
		fmt.Println("os.Create err:", err)
		return
	}
	defer fDst.Close()

	buf := make([]byte, 4096)
	for {
		n, err := fSrc.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("f.Read err", err)
			return  err
		}
		if err == io.EOF {
			break
		}

		// 读多少写多少
		if _, err = fDst.Write(buf[:n]); err != nil {
			return err
		}
	}
	return
}

func main() {
	// 要更新去重的列表
	path := `uBlacklist.txt`

	// 备份路径
	Backup := `uBlacklist_backup.txt`

	// 执行前先备份一份
	if err := BackupFile(Backup, path); err != nil {
		fmt.Println("BackupFile err", err)
		return
	}

	str, err := ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	// 统一换行符
	str = strings.Replace(str, "\r\n", "\n", -1)

	// 将字符串转成字符串切片
	strSlice := strings.Split(str, "\n")

	// 去重
	strSlice = RemoveRepeatedElement(strSlice)

	// 排序
	sort.Strings(strSlice)

	// 字符串切片转字符串
	str = strings.Join(strSlice, "\n")

	err = CreateFile(path, str)
	if err != nil {
		fmt.Println("CreateFile err:", err)
		return
	}

	fmt.Println("更新去重成功！")
}
