package main
import "fmt" 
func main() {
	matrix :=[][]int{
		{1,2,3},
		{4,5,6},
	} 
	fmt.Println("Transpose of matrix: ",transpose(matrix,len(matrix),len(matrix[0])))
}
func transpose(matrix [][]int,row,col int) [][]int{
	transpose := make([][]int , col)
	for i:=0; i<col; i++{
		transpose[i] = make([]int, row)
	}
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ { 
			transpose[j][i] = matrix[i][j] 
		} 
	}
	return transpose
}

			 
