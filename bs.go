// package bs

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func GetSlash(c *gin.Context) {

// var fooObj myType

// name := "Adam"

// fooObj = myType{
// 	Foo: "bar",
// 	Age: 69,
// 	Name: &name,
// }

// if fooObj.Name != nil {

// 	fmt.Println(*fooObj.Name)
// 	fooObj.SetName("Warfa")
// 	fmt.Println(*fooObj.Name)
// }

// fooObj.Fish = "Tuna"
// fooObj.AgeInMonths = fooObj.GetAgeInMonths()

// 	c.JSON(http.StatusOK, fooObj)

// }

// type struct2 struct {
//     Fish string `json:"fish"`
// }

// type myType struct{
//     Foo string `json:"foo"`
//     Age int `json:"age"`
// 	AgeInMonths string `json:"ageInMonths"`
// 	Name *string `json:"name"`
//     struct2
// }

// func (this *myType) GetAgeInMonths() string {
// 	return fmt.Sprintf("The value of age is %v", this.Age * 12)
// }

// func (m *myType) SetName(name string) {
// 	m.Name = &name
// }