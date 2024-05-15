package main
import(
	"net/http"
	"errors"
	"github.com/gin-gonic/gin"
)
type book struct{
/*ID string 'json:"id"'
Title string 'json:"title"'
Author string 'json:"author"'
Quantity int 'json:"quantity"'*/
ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}
var books = []book{
	{ID: "1", Title:"RANDOM BOOK 1",Author:"RA 1",Quantity:1},
	{ID: "2", Title:"RANDOM BOOK 2",Author:"RA 2",Quantity:1},
	{ID: "3", Title:"RANDOM BOOK 3",Author:"RA 3",Quantity:1},
}
func getBooks(c*gin.Context){
	c.IndentedJSON(http.StatusOK,books)
}
func createBook(c*gin.Context){
	var newbook books
if err := c.BindJSON(&newBook);err != nil{
	return
}
books = append(books,newBook)
c.IndentedJSON(http.StatusCreated,newBook)
}
func main(){
	router := gin.Default()
	router.GET("/books",getBooks)
	router.POST("/books",createBooks)
	router.Run("localhost:8080")
}