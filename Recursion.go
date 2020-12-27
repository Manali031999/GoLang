package main
import "fmt"
func addNum(n int) int{
	defer count(n)
	if n==0{
		return 0
	}
	if n==1 {
		return 1
	}
	return n + addNum(n-1)
}

func count(a int){
		fmt.Println("iteration number: ",a)
	
}
func main(){
	var n int
	fmt.Println("Enter number:")
	fmt.Scan(&n)
	fmt.Println(addNum(n))
}