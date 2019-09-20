# saas golang sdk

## Usage
### Quick start
```go
package main

import (
	"encoding/json"
	"fmt"

	sdk "github.com/nbltrust/jadepool-saas-sdk-go"
)

func main() {
	app := sdk.NewApp("appkey", "appsecret")
	result, err := app.GetAssets()
    if err != nil {
        fmt.Printf("execute error: %v", err)
        return
    }
    
	fmt.Println("code:", result.Code)
	fmt.Println("message:", result.Message)
	fmt.Println("data:")
	printMap(result.Data)
}
    
func printMap(m map[string]interface{}) {
	b, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
    }
	fmt.Print(string(b))
}
```

## CLI
Usage:
`ctl <key> <secret> <action> [<params>...] [-a <host>]`

e.g.
```bash
$ go run cmd/ctl/main.go "appkey" "appsecret" "CreateAddress" "ETH"
```