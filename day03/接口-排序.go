package main
import ("fmt"
		"sort")

type numbers []int

func (n numbers)Len()int{
	return len(n)
}

func (n numbers)Less(i,j int) bool{
	return n[i]<n[j]
}

func (n numbers)Swap(i,j int){
	n[i],n[j] = n[j],n[i]
}


func main(){
	nums := numbers{10,8,6,9,3,7,2,4,1,5}
	sort.Sort(nums)
	for _,v := range nums {
		fmt.Printf("%v\t",v)
	}
}