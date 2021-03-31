package exercises

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func Fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "https://") {
			url = fmt.Sprintf("https://%v",url)
		}
		fmt.Println(url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		b, err := io.Copy(os.Stderr, resp.Body)
		statusCode := resp.Status
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "reading %v: %v\n",url, err)
			os.Exit(1)
		}
		fmt.Printf("%s\n\n%v\n", strconv.FormatInt(b, 10), statusCode)
	}
}
