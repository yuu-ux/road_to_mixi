package main

import (
	"fmt"
	"net/http"
    "os"
)

func test(testCases map[string]int) {
    for url, expected := range testCases {
        resp, err := http.Get(url)
        if err != nil {
            fmt.Printf("[ERROR] %s => %v\n", url, err)
            continue
        }
        defer resp.Body.Close()

        if resp.StatusCode == expected {
            fmt.Printf("[OK]   %s => %d\n", url, resp.StatusCode)
        } else {
            fmt.Printf("[FAIL] %s => %d (expected %d)\n", url, resp.StatusCode, expected)
        }
    }
}

func main() {
	testCases := map[string]int{
		"http://localhost:8090": 200,
		"http://localhost:8090/hoge": 404,
		"http://localhost:8090/img/image1.png": 200,
		"http://localhost:8090/img/image2.png": 200,
		"http://localhost:8090/test": 200,
        "http://localhost:8090/app": 200,
	}

	mainteTestCases := map[string]int{
		"http://localhost:8090": 503,
		"http://localhost:8090/hoge": 503,
		"http://localhost:8090/img/image1.png": 503,
		"http://localhost:8090/img/image2.png": 503,
		"http://localhost:8090/test": 503,
        "http://localhost:8090/app": 503,
	}

    if len(os.Args) == 1 {
        test(testCases)
    } else {
        fmt.Println("メンテナンス中")
        test(mainteTestCases)
    }
}
