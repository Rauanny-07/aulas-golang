package main
   "fmt"
   "strings"
   "sort"

func main(){
    greeting := "hello my frinds!"
	fmt.Println(strings.Contains(greeting,"dogs")) 
	fmt.Println(strings.ReplaceAll(greeting,"my", "mine")) 
	fmt.Println(strings.ToUpper(greeting)) 
	fmt.Println(strings.Index(greeting,"my")) 
	fmt.Println(strings.Split(greeting,"")) 
	ages:= []int{50,80,10} 
	sort.Ints(ages) 
	fmt.Println(ages) 
	Index:= sort.SearchInts(ages,50)  
    fmt.Println(index) 
    names :=[]string{"Alice", "marco", "Diego"} 
    sort.Strings(names) 
    fmt.Println(names) 
    fmt.Println(sort.SearchStrings(nomes,"Alice"))

	x <5

	for i :=0; i <len(names); i++{
		fmt.Println(names[i])
	}

	for index, value := range names {
		fmt.Println("O Indice é", index, "e o valor", Value)
	
	}

		for index, Value := range ages {
			fmt.Println("O Indice é", index, "é o valor", Value)
		}