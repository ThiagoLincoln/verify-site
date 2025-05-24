package checksite

import (
	"fmt"
	"net/http"
	"time"
)

func CheckSite(url string) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("❌ %s está offline. Erro: %s\n", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		fmt.Printf("✅ %s está online.\n", url)
	} else {
		fmt.Printf("⚠️ %s retornou status %d\n", url, resp.StatusCode)
	}
}