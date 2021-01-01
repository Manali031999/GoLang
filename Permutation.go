package main
import (
	"fmt"
	"strconv"
)
func main() {
	var n , k int
	var str string
	fmt.Println("Enter number:")
	fmt.Scan(&n)
	fmt.Println("Enter Position: ")
	fmt.Scan(&k)
	for i:=1;i<=n;i++{
		str=str+strconv.Itoa(i)
	}
	result:=permutations(str)
	fmt.Println(result[k-1])
}

func permutations(str string) []string {
	le:=len(str)
	if le == 1 {
		return []string{str}
	}
	current := str[0:1]
	rem := str[1:]
	perms := permutations(rem)
	result := make([]string, 0) 
	for _, perm := range perms {
		permLen:=len(perm)
		for i := 0; i <= permLen; i++ {
			str := swap(i, current, perm)
			result = append(result, str)
		}
	}
	return result
}
func swap(i int, current string, perm string) string {
	start := perm[0:i]
	end := perm[i:len(perm)]
	return start + current + end
}
