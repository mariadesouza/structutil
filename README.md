# structutil

# Get Package

go get github.com/mariadesouza/utilities/structutil

# Usage
<pre>
type testStruct struct {
	Name       string
	Email      []string
	Occupation string
}
</pre>

<pre>
	testStruct := testStruct{"Ethan", []string{"emdesouza@gmail.com"}, "engineer"}
        m := structutil.StructToMap(&testStruct)
	for key, val := range *m {
		fmt.Println(key, val)
	}
</pre>
