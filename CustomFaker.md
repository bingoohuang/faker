## Custom Generator Provider

You can also add your own generator function to your own defined tags. See example below

```go
type Gondoruwo struct {
	Name       string
	Locatadata int
}

type Sample struct {
	ID                 int64     `faker:"customIdFaker"`
	Gondoruwo          Gondoruwo `faker:"gondoruwo"`
	Danger             string    `faker:"danger"`
	Bingoo             string    `faker:"bingoo,size=10"`
}

func CustomGenerator() {
	faker.AddProvider("customIdFaker", func(v reflect.Value) (interface{}, error) {
		 return int64(43), nil
	})
	faker.AddProvider("danger", func(v reflect.Value) (interface{}, error) {
		return "danger-ranger", nil
	})

	faker.AddProvider("gondoruwo", func(v reflect.Value) (interface{}, error) {
		obj := Gondoruwo{
			Name:       "Power",
			Locatadata: 324,
		}
		return obj, nil
	})
	
	faker.AddProviderV2("bingoo", func(v reflect.Value, tag Tag) (interface{}, error) {
    		return "size:" + tag.Opts["size"], nil
    })
}

func main() { 
	CustomGenerator()
	var sample Sample
	faker.Fake(&sample)
	fmt.Printf("%+v", sample)
}
```

Results:
```
{ID:43 Gondoruwo:{Name:Power Locatadata:324} Danger:danger-ranger}
```
