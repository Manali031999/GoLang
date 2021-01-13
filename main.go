package main
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"os"
)
type address struct{
	Area string `json:"area"`
	Country string `json:"country"`
}
type prsnlDets struct{
	Id int `json:"id"`
	Name string `json:"name"`
	Add address `json:"add"`
}
type prsnlArray struct{
	Prsnl []prsnlDets `json:"prsnl"`
}
type technology struct{
	Tech string `json:"tech"`
	Exp float64 `json:"exp"`
}
type techDets struct{
	Id int `json:"id"`
	Techlogy []technology `json:"techlogy"`
}
type techArray struct{
	Tech []techDets `json:"tech"`
}
type email struct{
	Emailid string `json:"email"`
	Phone string `json:"phone"`
}
type contactDets struct{
	Id int `json:"id"`
	Em email `json:"em"`
}
type contactArray struct{
	Contact []contactDets `json:"contact"`
}
type Dets struct{
	nameof string
	addof address
	techdetails []technology
	mailid string
	phoneno string
}
func main(){
	prsnljson , err := os.Open("prsnl.json")
	if err != nil{
		fmt.Println(err)
	}
	prsnlbyteValue,_ := ioutil.ReadAll(prsnljson)
	var prsnl prsnlArray 
	json.Unmarshal(prsnlbyteValue,&prsnl)
	techjson , err := os.Open("tech.json")
	if err != nil{
		fmt.Println(err)
	}
	techbyteValue,_ := ioutil.ReadAll(techjson)
	var tech techArray 
	json.Unmarshal(techbyteValue,&tech)
	contactjson , err := os.Open("contact.json")
	if err != nil{
		fmt.Println(err)
	}
	contactbyteValue,_ := ioutil.ReadAll(contactjson)
	var contact contactArray 
	json.Unmarshal(contactbyteValue,&contact)
	code := make(map[string]string)
	code["IND"] = "+91"
	code["UK"] = "+41"
	Details := make(map[int]Dets)
	prsnllen:=len(prsnl.Prsnl)
	techlen:=len(tech.Tech)
	conlen:=len(contact.Contact)
	for i := 0 ; i<prsnllen ; i++{
		var temp Dets
		id:=prsnl.Prsnl[i].Id
		temp.nameof = prsnl.Prsnl[i].Name
		temp.addof= prsnl.Prsnl[i].Add
		country:=prsnl.Prsnl[i].Add.Country
		for j:=0;j<techlen;j++{
			if tech.Tech[j].Id==id{
				temp.techdetails=tech.Tech[j].Techlogy
			}
		}
		for j:=0;j<conlen;j++{
			if contact.Contact[j].Id==id{
				temp.mailid=contact.Contact[j].Em.Emailid
			if country == "IND"{
				temp.phoneno=code["IND"]+" "+contact.Contact[j].Em.Phone
			}else{
				temp.phoneno=code["UK"]+" "+contact.Contact[j].Em.Phone
			}
		}
		}
		Details[id] = temp
	}
	for i := range Details {
		fmt.Println(i, ":", Details[i])
	}
}
