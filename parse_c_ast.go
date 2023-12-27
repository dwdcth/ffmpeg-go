package main

import (
	"fmt"
	"log"

	"modernc.org/cc/v4"
)

func main() {
	// C 语言源代码
	src := `
    #include <stdio.h>

    int main() {
        printf("Hello, World!");
        return 0;
    }
    `

	// 创建解析器实例
	ast, err := cc.Parse("example.c", src, cc.AllowCompatibleExtensions())
	if err != nil {
		log.Fatal(err)
	}

	// 遍历 AST
	ast.Walk(func(node interface{}) bool {
		fmt.Printf("%T\n", node)
		return true
	})
}
