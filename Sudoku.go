package main
import (
	"fmt"
)
func main(){
	board:=[][]int{
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
func checkBox(board [][]int) bool {
	row:=3
	col:=3
	a:=make([][]int , row)
	for i:=0;i<row;i++{
		a[i]=make([]int,col) 
		for j:=0;j<col;j++{
			a[i][j]=board[i][j]
		}
	}
	if checkRowandCol(a){
		return true  
	}else{
		return false
	}
	return true
}

func checkRowandCol(a [][]int) bool{
	r,c :=len(a) , len(a[0])
	for i:=0;i<r;i++{
		for j:=0;j<c;j++{
			for k:=j+1;k<r;k++{
				if a[i][k]==a[i][j]{
					return false
				}
			}
			for k:=i+1;k<r;k++{
				if a[k][j]==a[i][j]{
					return false
				}
			}
		}
	}
	return true
}
func checkBoard(board [][]int) bool{ 
	if checkRowandCol(board) && checkBox(board){
		return true
	}else{
		return false
	}
	return true
}
