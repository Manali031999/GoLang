
package main
import (
	"fmt"
	"strconv"
)
var count int = 0

func swap(s string,i,j int) string{
	var result []byte
	for k:=0;k<len(s);k++{
		if k==i{
			result = append(result,s[j])
		}else if k==j{
			result = append(result,s[i])
		}else{
			result = append(result,s[k])
		}
	}
	
	return string(result)
}

func permutations(str string, i,n ,x int){
	if i == n-1 {
		count ++
		if count == x{
			println(str)
		}
		
		return
	}
	
	for j:=i;j<n;j++{
		str = swap(str,i,j)
		permutations(str,i+1,n,x)
	}

}

func main() {
	var(
		n , x int
		str string
	) 
	fmt.Println("enter Number: ")
	fmt.Scan(&n)
	fmt.Println("Enter Position: ")
	fmt.Scan(&x)
	for i:=1;i<=n ; i++{
		str= str+strconv.Itoa(i)
	}
	permutations(str,0,len(str),x)
	
} 
