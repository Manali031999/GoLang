package main
import "fmt"
func main(){
	m1:=[3][2]int{}
	m2:=[2][3]int{}
	fmt.Println("Enter 3x2 matrix:")
	for i:=0;i<3;i++{
		for j:=0;j<2;j++{
			fmt.Scan(&m1[i][j])
		}
	}
	fmt.Println("Enter 2x3 matrix:")
	for i:=0;i<2;i++{
		for j:=0;j<3;j++{
			fmt.Scan(&m2[i][j])
		}
	}
	fmt.Println()
	mul(m1,m2)
}
func mul(m1 [3][2]int , m2 [2][3]int){
	m3:=[3][3]int{}
	fmt.Println("Transpose is: ")
	for i:=0;i<3;i++{
		for j:=0;j<3;j++{
			m3[i][j]=0
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
	
