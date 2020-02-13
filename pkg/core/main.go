package core
import(
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)
type QueryError struct {
Query string
Err error
}
type DbError struct {
	Err error
}

func (receiver *DbError) Error() string {
	return fmt.Sprintf("cant handle db operatin:%v",receiver.Err.Error() )
}

func dbError(err error) *DbError {
	return &DbError{Err: err}
}
type BankMachine struct{
	Id int64
	Name string
}
type BankAccounts struct {
	Id int64
	Name string
	BankAccount int64
	Balance int64
}

func (receiver QueryError) Error() string {
	return fmt.Sprintf("can't execute query %s:%s",loginSQL,receiver.Err.Error())
}

/////
var ErrInvalidPass=errors.New("Invalid password")
func queryError(query string,err error) *QueryError{
	return &QueryError{Query: query, Err: err}
}
func Init(db *sql.DB) (err error) {
	ddls := []string{clientsDDL,bankMachinesDDL}
	for _, ddl := range ddls {
		_, err = db.Exec(ddl)
		if err != nil {
			return err
		}}
	initialData:=[]string{clientsInitialData,bankMachinesInitialData}
	for _,info:=range initialData{
		_,err=db.Exec(info)
		if err !=nil{
			return err
		}}
	return nil
}
func Login(login string,password int,db *sql.DB)(bool,error){
var dbLogin string
var dbPassword int
err:=db.QueryRow(
	loginSQL,login).Scan(&dbLogin,&dbPassword)
	if err != nil {
		if err ==sql.ErrNoRows{
			return false, nil
		}
		return false, queryError(loginSQL,err)
	}
	if dbPassword!=password{
		return false, ErrInvalidPass
	}
	return true, nil
}
func Machine(db *sql.DB)(machines []BankMachine, err error){
	rows,err:=db.Query(showBankMachine)
	if err!=nil{
		return nil,queryError(showBankMachine,err)
	}
defer func() {
	if innerErr := rows.Close(); innerErr != nil {
		machines, err = nil, dbError(innerErr)
	}
}()
	for rows.Next(){
		machine:=BankMachine{}
		err=rows.Scan(&machine.Id,&machine.Name)
		if err != nil {
			return nil, dbError(err)
		}
		machines=append(machines,machine)
	}
	if rows.Err()!= nil{
		return nil,dbError(rows.Err())
	}

	return machines, nil
}
func Account(db *sql.DB)(accounts []BankAccounts, err error){
	rows,err:=db.Query(bankAccount)
	if err!=nil{
		return nil,queryError(bankAccount,err)
	}
	defer func() {
		if innerErr := rows.Close(); innerErr != nil {
			accounts, err = nil, dbError(innerErr)
		}
	}()
	for rows.Next(){
		account:=BankAccounts{}
		err=rows.Scan(&account.Id,&account.Name,&account.BankAccount,&account.Balance)
		if err != nil {
			return nil, dbError(err)
		}
		accounts=append(accounts,account)
	}
	if rows.Err()!= nil{
		return nil,dbError(rows.Err())
	}

	return accounts, nil
}
func main() {
	
}
