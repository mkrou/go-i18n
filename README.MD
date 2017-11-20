# go-i18n
go-i18n is a package that helps you translate go programs.

### Installation
###### go get
```bash
go get github.com/mkrou/go-i18n
```
###### dep
```bash
dep ensure -u github.com/mkrou/go-i18n
```
### Usage
```golang
package main

import (
	"github.com/mkrou/go-i18n"
	"log"
	"fmt"
)

var translate = i18n.Map{
	"key1": "val1",
	"subcat1": i18n.Map{
		"key2": "val2",
		"subcat2": i18n.Map{
			//Supported native go formatting
			"key3": "string: %s; Int: %d.",
		},
	},
}

func main() {
    //you have to add languages before usage
	if err := i18n.AddLanguage("en", translate); err != nil {
		log.Fatal(err)
	}
	//You have to choose default language.
	//You can choose few languages.
	//go-i18n will search right key in each language until find it
	if err := i18n.Current([]string{"en","ru"}); err != nil {
		log.Fatal(err)
	}
	fmt.Println(i18n.T("key1"))
	fmt.Println(i18n.T("subcat1.key2"))
	fmt.Println(i18n.T("subcat1.subcat2.key3", "hello world", 10))
}

```
###### Output
```bash
val1
val2
string: hello world; Int: 10.
```