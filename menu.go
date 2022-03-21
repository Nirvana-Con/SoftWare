package main

import "fmt"

func main(){
	choice := ""
	for{
		fmt.Println("----------MENU-----------")
		fmt.Println("a:上一页")
		fmt.Println("b:下一页")
		fmt.Println("c:起始页")
		fmt.Println("d:退出")
		fmt.Println("请输入选项")
		fmt.Println("------------------------")
		fmt.Scan(&choice)
		switch choice {
		case "a":
			fmt.Println("这是上一页")
		case "b":
			fmt.Println("这是下一页")
		case "c":	
			fmt.Println("这是起始页")
		case "d":
			fmt.Println("退出成功")
			return 
		default:
			fmt.Println("请输入正确选项")
		}
	}
}

