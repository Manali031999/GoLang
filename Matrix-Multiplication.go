package main
import "fmt"
func main(){
	m1:=[][]int{{1,2},{3,4},{5,6},}
	m2:=[][]int{{1,2,3},{4,5,6},}
	a:=len(m1)
	b:=len(m2[0])
	c:=len(m1[0])
	if a==b {
		fmt.Println(mul(m1,m2,c,a))
	}else{
		fmt.Println("Cannot be Multiplied")
	}
}
func mul(m1  , m2 [][]int , c ,a int ) [][]int{
	m3:=make([][]int,a)
	for i:=0 ; i<a ; i++{
		m3[i]=make([]int,a)
	}
	fmt.Println("Answer is: ")
	for i:=0;i<a;i++{
		for j:=0;j<a;j++{
			for k:=0;k<c;k++{
				m3[i][j]+=(m1[i][k]*m2[k][j])
			}
		}
	}
	return m3
}
	
