package main
import "fmt"
func main(){
	m1:=[][]int{
		{1,2},
		{3,4},
		{5,6},
	}
	m2:=[][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}
	mul(m1,m2)
}
func mul(m1 [][]int , m2 [][]int){
	m3:=[3][3]int{}
	fmt.Println("Answer is: ")
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			for k:=0;k<2;k++{
				m3[i][j]=m3[i][j]+(m1[i][k]*m2[k][j])
			}
		}
	}
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			fmt.Printf("%d\t",m3[i][j])
		}
		fmt.Println()
	}
}
	
