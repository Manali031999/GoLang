package main
import "fmt" 

func transpose(matrix [][]int,r,c int) [][]int{
	a:=len(matrix[0])
	b:=len(matrix)
	transpose := make([][]int , a)
	for i:=0; i<a; i++{
		transpose[i] = make([]int, b)
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ { 
			transpose[j][i] = matrix[i][j] 
		} 
	}
	fmt.Println("Transpose of matrix: ") 
	return transpose
}
func main() {
	matrix :=[][]int{
		{1,2,3},
		{4,5,6},
	} 
	fmt.Println(transpose(matrix,len(matrix),len(matrix[0])))
}

			 
