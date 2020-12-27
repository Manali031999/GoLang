package main
import (
	"fmt"
	"strconv"
)
func odd(number string){
	var countzero int
	for i:=range number{
		if number[i]==48{
			countzero++
			if countzero>=3{
				break
			}
		}else{
			countzero=0
		}
	}
	if countzero>=3{
		fmt.Println("000-ODD")
	}else if countzero<3{
		fmt.Println(">3-0s ODD")
	}
}
func even(number string){
	var countone int
	for i:=range number{
		if number[i]==49{
			countone++
			if countone<=5{
				break
			}
		}else{
			countone=0
		}
	}
	if countone<=5{
		fmt.Println(">5 1s- EVEN")
	}

}
func main(){

	var num int64
	fmt.Println("Enter Number: ")
	fmt.Scan(&num)
	number:=strconv.FormatInt(num,2)
	fmt.Println(number)
	switch {
	case num%2==0:
		even(number)
	case num%2!=0:
		odd(number)
	default : fmt.Println("end!")
	}
}
