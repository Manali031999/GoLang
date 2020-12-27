package main
import (
	"fmt"
	"strconv"
)
func swap(s string, i, j int) string {
      var result []byte
      for k := 0; k < len(s); k++ {
              if k == i {
                      result = append(result, s[j])
              } else if k == j {
                      result = append(result, s[i])
              } else {
                      result = append(result, s[k])
              }
      }
      return string(result)
}
func permutations(str string, i, n,m int) {
	var j int  =-1
	if i == n-1  {
			  println(str)
              return
	  }
      for j = i; j < n; j++ {
             str = swap(str, i, j)
			permutations(str, i+1, n,m)
	  }
	}
func main() {
	var n ,m  int
	fmt.Println("Enter Number: ")
	fmt.Scan(&n)
	fmt.Println("Enter Iteration Number: ")
	fmt.Scan(&m)
	var str string
	i:=1
	for i<=n{
		str = str+strconv.Itoa(i)
		i++
	}
    permutations(str, 0, len(str),m)
}