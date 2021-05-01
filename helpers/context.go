package helpers

import (
	"reflect"

	"github.com/gin-gonic/gin"
)


func GetUserDataFromContext(c *gin.Context) (bool,map[string]interface{}){
	if userDataAsInterface, exist := c.Get("user"); exist{
		var arbitaryData = make(map[string]interface{}) // make the empty interface first	
		v := reflect.ValueOf(userDataAsInterface)
		if v.Kind() == reflect.Map {
			for _, key := range v.MapKeys() {
				strct := v.MapIndex(key)
				//fmt.Println(key.String(), strct.Interface())
				arbitaryData[key.String()] = strct.Interface()
			}
		}
		return true,arbitaryData
	}else{
		return false,nil
	}
}