package Test

import (
	"fmt"
	"os"
)

func FPrint() {
	file, err := os.Create("D:\\WorkSpace\\WorkspaceOfGo\\Mytest_Go\\Test\\A.txt")

	if err != nil {
		fmt.Printf("发生错误：%s", err)
	}

	fmt.Fprintln(file, "W")
	fmt.Fprintln(file, "Y")
	fmt.Fprintln(file, "D")

}
