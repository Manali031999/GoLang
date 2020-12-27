package main
import "fmt"

func main(){
	a:=[]int{1,3,2,4,5}
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
		if i==len(a)-1 || a[i]==0{
			return count
		}else{
			b:=a[i]
			for j:=1;j<=b;j++{
				if a[i+1]>a[i]{
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