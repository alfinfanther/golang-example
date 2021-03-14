package main
import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ListData struct{
	Id int `json:"id"`
	PropertyName string `json:"property_name"`
	Url string `json:"url"`
}

func main(){
	fmt.Println(ReadJson())
}

func ReadJson() []ListData{
	var list_data []ListData
	var respData []ListData
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &list_data)

	for i := 0; i < len(list_data); i++ {
		lst := ListData{
			Id: list_data[i].Id,
			PropertyName: list_data[i].PropertyName,
			Url: list_data[i].Url,
		}

		respData = append(respData,lst)
	}
	return respData
}