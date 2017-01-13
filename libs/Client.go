package libs

import (
	"log"
	"github.com/joaquinicolas/FormAndroidMigrator/DataBase"
)


type Client struct {
	Numero string
	Nombres string
	Pais string
}

func NewClient(numero,nombres,pais string) *Client {
	return &Client{
		Numero:numero,
		Nombres:nombres,
		Pais:pais,
	}
}

//
func (c *Client) CreateFromFormatter(input Formatter)(func(DataBase.DataStore)) {
	data := input.Read()
	clients,err := input.Parser(data)
	checkErr(err)
	return func(store DataBase.DataStore) {
		for _,i:= range clients{
			c.Nombres = i[1]
			c.Numero = i[0]
			c.Pais = i[2]
			store.Create(c.Numero,c.Nombres,c.Pais)
		}
	}
}


func checkErr(err error)  {
	if err != nil {
		log.Fatal(err)
	}
}


/**
 id      | integer                | not null  | plain    |              |
 numero  | character varying(10)  |           | extended |              |
 nombres | character varying(100) |           | extended |              |
 pais    | character varying(2)   |           | extended |
 */