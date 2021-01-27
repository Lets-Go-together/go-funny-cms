package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"gocms/app/http/admin/validates"
	"gocms/pkg/config"
	"net/http"
	"testing"
)

var _ = func() bool {
	testing.Init()
	return true
}()

func BenchmarkNormalHandleFunc(b *testing.B) {

	router := gin.New()

	router.POST("api/register", func(ctx *gin.Context) {
		p := validates.RegisterParams{}
		validator := validates.RegisterAction{}
		if !validator.Validate(ctx, &p) {
			return
		}
		if len(p.Account) == 0 {
			fmt.Println("error")
		}
	})
	config.Router = router
	go func() {
		_ = router.Run(":8081")
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testPost("8081")
	}
}

func BenchmarkReflectHandleFunc(b *testing.B) {

	router := gin.New()

	group("api",
		post("/register", func(ctx *gin.Context, params *validates.RegisterParams) {
			if len(params.Account) == 0 {
				fmt.Println("error")
			}
		}),
	).setup(router)
	config.Router = router
	go func() {
		_ = router.Run(":8082")
	}()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		testPost("8082")
	}
}

func testPost(port string) {
	params := struct {
		Account  string
		Password string
		Email    string
		Captcha  string
	}{
		Account:  "account",
		Password: "1231ljasd",
		Email:    "email@exmpale.com",
		Captcha:  "12345",
	}
	paramsByte, _ := json.Marshal(params)
	r, _ := http.Post("http://127.0.0.1:"+port+"/api/register", "application/json", bytes.NewReader(paramsByte))

	if r.StatusCode != 200 {
		fmt.Println("errr")
	}
}
