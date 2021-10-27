package main
import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	
)

type postJSON struct {
	Data string
}

var ctx = context.Background()



func main() {
	r := gin.Default()

	
	r.POST("/go/sha256", func(c *gin.Context) {
		println("yes1")
		jsonData, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "something went wrong with database",
			})
		}
		println("yes2")
		var data postJSON
		json.Unmarshal(jsonData, &data)
		fmt.Println(data.Data)
		println("yes3")
		if len([]rune(data.Data)) < 8 {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "Your message lenght must be more than 8 characters!",
			})
			println("yes4")
		}else{
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "yess",
			})
			println("yes5")
		}
	})
	r.Run(":80") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
