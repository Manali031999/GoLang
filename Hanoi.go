package main
import "fmt"
func main(){
	var n int
	fmt.Println("Enter Number: ")
	fmt.Scan(&n)
	hanoi(n,"A","C","B")
}
func hanoi(n int, from,to,mid string){
	if n>0{
		hanoi(n-1,from,mid,to)
		fmt.Printf("Move disk %d from %s to %s\n",n,from,to)
		hanoi(n-1,mid,to,from)
	}
}