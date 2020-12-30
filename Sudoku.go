package main
import (
	"fmt"
)
func main(){
	board:=[9][9]int{
		{5,3,4,6,7,8,9,1,2},
		{6,7,2,1,9,5,3,4,8},
		{1,9,8,3,4,2,5,6,7},
		{8,5,9,7,6,1,4,2,3},
		{4,2,6,8,5,3,7,9,1},
		{7,1,3,9,2,4,8,5,6},
		{9,6,1,5,3,7,2,8,4},
		{2,8,7,4,1,9,6,3,5},
		{3,4,5,2,8,6,1,7,9},
	}
	if checkBoard(board){
		fmt.Println("correct")
	}else{
		fmt.Println("Invalid")
	}
}
func checkRow(board [9][9]int) bool{
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			for k:=j+1;k<9;k++{
				if board[i][k]==board[i][j]{
					return false
				}
			}
		}
	}
	return true
}
func checkCol(board [9][9]int) bool{
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			for k:=i+1;k<9;k++{
				if board[k][j]==board[i][j]{
					return false
				}
			}
		}
	}
	return true
} 

func checkBox(board [9][9]int) bool{ 
	for i:=0;i<9;i++{
		for j:=0;j<9;j++{
			r:=(i/3)*3
			c:=(i/3)*3
			for i:=r ; i<r+3 ; i++{
				for j:=c ; j<c+3; j++{
					for k:=j+1 ; k<9 ; k++{
						if board[i][k]==board[i][j]{
							return false
						}
					}
					for k:=i+1;k<9 ; k++{
						if board[k][j]==board[i][j]{
							return false
						}
					}
				}
			}
		}
	}
	return true
}
func checkBoard(board [9][9]int) bool{ 
	if checkRow(board) && checkCol(board) && checkBox(board){
		return true
	}else{
		return false
	}
	return true
}
