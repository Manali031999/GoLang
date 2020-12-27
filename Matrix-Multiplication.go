package main
import "fmt"
func main(){
	m1:=[3][2]int{
		[2]int{3,4},
		[2]int{7,2},
		[2]int{5,9},
	}
	m2:=[2][3]int{
		[3]int{3,1,5},
		[3]int{6,9,7},
	}
	m3:=[3][3]int{}
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