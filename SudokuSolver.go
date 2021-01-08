package main
import "fmt"
func main(){
	board := map[int][]int{
		1: []int{5, 0, 4, 0, 7, 8, 9, 1, 2},
		2: []int{0, 7, 2, 0, 9, 5, 0, 4, 0},
		3: []int{1, 0, 8, 0, 4, 0, 5, 6, 7},
		4: []int{8, 5, 0, 7, 6, 1, 4, 2, 3},
		5: []int{0, 2, 6, 8, 5, 3, 0, 0, 1},
		6: []int{7, 0, 3, 9, 2, 4, 8, 5, 0},
		7: []int{9, 6, 1, 0, 3, 0, 2, 8, 4},
		8: []int{0, 0, 0, 4, 1, 9, 6, 0, 5},
		9: []int{3, 4, 5, 2, 8, 6, 1, 7, 0},
	}
	if solveSudoku(board,0,0){
		print(board)
	}else{
		fmt.Println("No solution exists.")
	}
}

func print(board map[int][]int){
	for i:=0;i<9;i++{
		for _,val:=range board[i+1]{
			fmt.Print(val," ")
		}
		fmt.Println()
	}
}
func isValid(board map[int][]int , row , col , num int) bool{
	//same num in similar row
	for i:=0;i<9;i++{
		if board[i+1][col]==num{
			return false
		}
	}
	//same num in similar col
	for i:=0;i<9;i++{
		if board[row+1][i]==num{
			return false
		}
	}
	//3X3 matrix
	startRow:=row-row%3
	startCol:=col-col%3
	for i:=startRow;i<startRow+3;i++{
		for j:=startCol;j<startCol+3;j++{
			if board[i+1][j]==num{
				return false
			}
		}
	}
	return true
}
func solveSudoku(board map[int][]int , row, col  int) bool{
	//if we have reached last row and last col
	if row==8 && col==9{
		return true
	}
	if col==9{
		row++
		col=0
	}
	if board[row+1][col]>0{
		return solveSudoku(board,row,col+1)
	}else{
		for num:=1;num<=9;num++{
			if isValid(board,row,col,num){
				board[row+1][col]=num
				if solveSudoku(board,row,col+1){
					return true
				}
			}
			board[row+1][col]=0
		}
	}
		
	
	
	return false
}