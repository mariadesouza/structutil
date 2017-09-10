# structutil

# Get Package

go get github.com/mariadesouza/utilities/structutil

#

* CopyToStringStruct

<pre>
func CopyToStringStruct(src interface{}, dest interface{})
</pre>
CopyToStringStruct : Copies src struct fields to dest struct Both src and dest structs have elements with same names but different types

* StructToMap

<pre>
func StructToMap(v interface{}) *map[string]interface{}
</pre>
StructToMap : returns a map of names to values for struct passed .


# Usage
<pre>

type srcStruct struct {
	Id     int
	Name   string
	Email  string
	Salary float64
}

type destStruct struct {
	Id     string
	Name   string
	Email  string
	Salary string
}

func main() {

	src := srcStruct{1024, "Ethan", "emdesouza@gmail.com", 200100.23}
	dest := destStruct{}
	structutil.CopyToStringStruct(&src, &dest)
	m := structutil.StructToMap(&dest)
	for key, val := range *m {
		fmt.Println(key, val)
	}
}
</pre>
