package main
import "fmt" 
func main() {
	var i, j, r, c int 
	var matrix [10][10]int
	var transpose [10][10]int
	   	fmt.Print("Enter the number of rows the matrix: ")
		fmt.Scanln(&r) 
		fmt.Print("Enter the number of columns th matrix: ") 
		fmt.Scanln(&c) 
		fmt.Println("Enter the matrix elements") 
		for i = 0; i < r; i++ { 
			for j = 0; j < c; j++ { 
				fmt.Scan(&matrix[i][j]) 
				} 
			} 
			for i = 0; i < r; i++ {
				for j = 0; j < c; j++ { 
					 transpose[j][i] = matrix[i][j] 
				} 
			} 
			fmt.Println("Transpose of matrix: ") 
			for i = 0; i < r; i++ { 
				for j = 0; j < c; j++ { 
					fmt.Print(transpose[i][j], "\t") 
				} 
				fmt.Println() 
			} 
}