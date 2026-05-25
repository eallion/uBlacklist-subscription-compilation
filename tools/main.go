package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func RemoveRepeatedElement(arr []string) []string {
	result := []string{}
	for i := 0; i < len(arr); i++ {
		exist := false
		for j := 0; j < len(result); j++ {
			if result[j] == arr[i] {
				exist = true
				break
			}
		}
		if !exist {
			result = append(result, arr[i])
		}
	}
	return result
}

func main() {
	path := `uBlacklist.txt`
	backup := `uBlacklist_backup.txt`

	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("ReadFile err:", err)
		return
	}

	if err := os.WriteFile(backup, data, 0644); err != nil {
		fmt.Println("BackupFile err:", err)
		return
	}

	str := string(data)
	str = strings.Replace(str, "\r\n", "\n", -1)
	strSlice := strings.Split(str, "\n")
	strSlice = RemoveRepeatedElement(strSlice)
	sort.Strings(strSlice)
	str = strings.Join(strSlice, "\n") + "\n"

	if err := os.WriteFile(path, []byte(str), 0644); err != nil {
		fmt.Println("WriteFile err:", err)
		return
	}

	fmt.Println("更新去重成功！")
}
