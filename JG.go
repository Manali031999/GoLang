package main
import "fmt"

func main(){
	path:=[]int{2,4,3,2,2,0,1}
	//path:=[]int{0}
	count:=jump(path)
	if count==0{
		fmt.Println("Cannot reach")
	}else{
		fmt.Println("Reached at",count,"counts")
	}
}

func jump(path []int) int{
	length:=len(path)
	var count int 
	i:=0
	for i<length{
		if path[i]==0{
			return count
		}
		if i== length-1{
			return 1
		}else{
			value:=path[i]
			last:=i+value
				if last >= length-1{
					count++
					break
				}
			for j:=1;j<=value;j++{
				if path[i+1]>=path[i]{
					i=i+1
					count++
					break
				}else{
					i=i+1
					continue
				}
			}
		}
	}
	return count
}
