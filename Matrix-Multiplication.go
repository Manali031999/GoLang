package main
import "fmt"
func main(){
	m1:=[][]int{{1,2},{3,4},{5,6},}
	m2:=[][]int{{1,2,3},{4,5,6},}
	rm1:=len(m1)
	cm2:=len(m2[0])
	cm1:=len(m1[0])
	if rm1==cm2 {
		fmt.Println("Answer is :",mul(m1,m2,cm1,rm1))
	}else{
		fmt.Println("Cannot be Multiplied")
	}
}
func mul(m1  , m2 [][]int , cm1 ,rm1 int ) [][]int{
	m3:=make([][]int,rm1)
	for i:=0 ; i<rm1 ; i++{
		m3[i]=make([]int,rm1)
	}
	for i:=0;i<rm1;i++{
		for j:=0;j<rm1;j++{
			for k:=0;k<cm1;k++{
				m3[i][j]+=(m1[i][k]*m2[k][j])
			}
		}
	}
	return m3
}
	
