package DataBase

import (
	"database/sql"
	"log"
	"errors"
	"fmt"
	"strings"
	"github.com/lib/pq"
)

type DataStore interface {
	Create(client ...interface{})
	Name() string
}

type DataStoreFactory func(conf map[string]string)(DataStore,error)

type PostgreSQLDataStore struct {
	DSN string
	DB *sql.DB
}

var datastoreEactories = make(map[string]DataStoreFactory)


/**
	HELPERS
 */

//VerifyCon checks if connection can be made
func VerifyCon(db *sql.DB) (error,bool) {
	if err := db.Ping(); err != nil {
		return err,false
	}

	return nil,true
}

//Register a DataStore
func Register(name string,factory DataStoreFactory)  {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist",name)
	}
	_,registered := datastoreEactories[name]
	if registered {
		fmt.Errorf("Datastore factory %s alredy registered.Ignoring.",name)
	}
	datastoreEactories[name] = factory
}

func init()  {
	Register("postgres",NewPostgreSQLDataStore)
}

//Create instance of the DataStore interface
func CreateDataStore(conf map[string]string)(DataStore,error)  {

	engineName := conf["DATASTORE"];

	engineFactory, ok := datastoreEactories[engineName]
	if !ok {

		availableDatastores := make([]string,len(datastoreEactories))
		for k,_ := range datastoreEactories{
			availableDatastores = append(availableDatastores,k)
		}
		return nil,errors.New(fmt.Sprintf("Invalid Datastore name.Must be one of: %s",strings.Join(availableDatastores,",")))
	}

	return engineFactory(conf)

}


/**
	PGSQL
 */
func (pds *PostgreSQLDataStore) Name() string {
	return "PostgreSQLDataStore"
}

func NewPostgreSQLDataStore(conf map[string]string)(DataStore,error)  {

	dsn := conf["DATASTORE_POSTGRES_DSN"]
	if dsn == ""{
		return nil, errors.New(fmt.Sprintf("%s is required for the postgres datastore", "DATASTORE_POSTGRES_DSN"))
	}

	db, _ := sql.Open("postgres",dsn)

	if err,ok := VerifyCon(db);!ok{
		log.Panicf("Failed to connect to datastore: %s", err.Error())
		return nil, err
	}

	return &PostgreSQLDataStore{
		DSN:dsn,
		DB:db,
	},nil

}

//Create a client and return the last inserted id
func (pds *PostgreSQLDataStore) Create(client...interface{})  {
	txn, err := pds.DB.Begin()
	if err != nil {
		log.Fatal(err)
	}
	stmt, err := txn.Prepare(pq.CopyIn("clientes","numero","nombres","pais"))
	if err != nil {
		log.Fatal(err)
	}
	_,err = stmt.Exec(client)
	if err != nil {
		log.Fatal(err)
	}

	_,err = stmt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	err = stmt.Close()
	if err != nil {
		log.Fatal(err)
	}
	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Sprintf("***Client has been created***")
}

