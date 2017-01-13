package libs

import (
	"io/ioutil"
	"strings"
	"errors"
	"fmt"
	"log"
)


type Formatter interface {
	Read() string
	Parser(string) ([][]string,error)
}
var formattersFactories = make(map[string]FormattereFactory)
type FormattereFactory func(conf map[string]string)(Formatter,error)
type CSVInput struct {
	LOCATION string
	SEPARATOR string
}



func (csv *CSVInput) Read() string{

	dat, err := ioutil.ReadFile(csv.LOCATION)
	checkErr(err)
	data := string(dat)
	return data
}



func (csv *CSVInput) Parser(data string) ([][]string,error) {

	if data == "" {

		return nil, errors.New("CSV data is empty")
	}

	row := strings.Split(data,"\n")
	d := make([][]string,len(row))
	for i:=0;i < len(row);i++{

		d[i] = strings.Split(row[i],csv.SEPARATOR)
	}

	fmt.Println(d[5])
	return d,nil

}

//Create Formatter from configuration map, this param specified the FORMAT,
// LOCATION AND SEPARATOR
func CreateFormatter(conf map[string]string) (Formatter,error) {
	format := conf["FORMAT"]
	formatterFactory,_ := formattersFactories[format]
	return formatterFactory(conf)
}

func NewCSVFormatter(conf map[string]string)(Formatter,error)  {
	location := conf["LOCATION"]
	separator := conf["SEPARATOR"]
	if location == "" {
		return nil,errors.New(fmt.Sprintf("The location of csv is required."))
	}

	if separator == "" {
		return  nil,errors.New(fmt.Sprintf("" +
			"The separator is required. It can be a comma(,) or semicolon(;)" +
			""))
	}
	return &CSVInput{
		SEPARATOR:separator,
		LOCATION:location,
	},nil

}

func init()  {
	Register("CSV",NewCSVFormatter)
}

func Register(name string,factory FormattereFactory)  {
	if factory == nil {
		log.Panicf("Formatter factory %s does not exist",name)
	}

	_,registered := formattersFactories[name]
	if registered {
		fmt.Errorf("Formatter factory %s alredy registered.Ignoring.",name)
	}
	formattersFactories[name] = factory
}

