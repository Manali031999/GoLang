package main
import (
	"fmt"
)
type address struct{
	area string
	country string
}
type prsnlDets struct{
	id int
	name string
	add address
}
type technology struct{
	tech string
	exp float64
}
type techDets struct{
	id int
	techlogy []technology
}
type email struct{
	emailid string
	phone string
}

type contactDets struct{
	id int
	em email
}
type Dets struct{
	nameof string
	addof address
	techdetails []technology
	mailid string
	phoneno string
}
func main(){
	code := make(map[string]string)
	code["IND"] = "+91"
	code["UK"] = "+41"
	prsnl := []prsnlDets{
		prsnlDets{ 1, "Radha",  address{"bopal", "IND"}},
		prsnlDets{ 2, "Aniket",  address{"maninagar", "UK"}},
	}
	tech := []techDets{
		techDets{ 1, []technology{{"React", 3},{"GoLang",1.5}}},
		techDets{ 2, []technology{{"Vue", 0.9},{"GoLang",1}}},
	}
	con := []contactDets{
		contactDets{ 1, email{"radha.kotecha@bacancy.com","12345 67890"}},
		contactDets{ 2, email{"aniket.amin@bacancy.com","12345 67890"}},
	}
	Details := make(map[int]Dets,0)
	prsnllen:=len(prsnl)
	techlen:=len(tech)
	conlen:=len(con)
	for i:=0;i<prsnllen;i++{
		var temp Dets
		id:=prsnl[i].id
		temp.nameof = prsnl[i].name
		temp.addof= prsnl[i].add
		country:=prsnl[i].add.country
		for j:=0;j<techlen;j++{
			if tech[j].id==id{
				temp.techdetails=tech[j].techlogy
			}
		}
		for j:=0;j<conlen;j++{
			if con[j].id==id{
				temp.mailid=con[j].em.emailid
			if country == "IND"{
				temp.phoneno=code["IND"]+" "+con[j].em.phone
			}else{
				temp.phoneno=code["UK"]+" "+con[j].em.phone
			}
		}
		}
		Details[id] = temp
	}
	for i := range Details {
		fmt.Println(i, ":", Details[i])
	}
}
