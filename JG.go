package main
import "fmt"

func main(){
	a:=[]int{2,4,3,2,2,0,1}
	//a:=[]int{1}
	count:=jump(a)
	if count==0{
		fmt.Println("Cannot reach")
	}else{
		fmt.Println("Reached",count)
	}
}

func jump(a []int) int{
	var count int 
	i:=0
	for i<len(a){
		if a[i]==0{
			return count
		}
		if i== len(a)-1{
			return 1
		}else{
			b:=a[i]
			c:=i+b
				if c >= len(a)-1{
					count++
					break
				}
			for j:=1;j<=b;j++{
				if a[i+1]>=a[i]{
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
