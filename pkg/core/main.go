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
	return fmt.Sprintf("cant handle db operation:%v",receiver.Err.Error() )
}

func dbError(err error) *DbError {
	return &DbError{Err: err}
}
type BankMachine struct{
	Id int64
	Name string
	Street string
}
type Client struct {
	Id int64
	Name string
	Login string
	Password int64
	BankAccount uint64
	PhoneNumber int64
	Balance uint64
}
type BankAccounts struct {
	Id int64
	Name string
	BankAccount int64
	Balance int64
}
type Services struct{
	Id int64
	Name string
	Price uint64
}
func (receiver *QueryError) Error() string {
	return fmt.Sprintf("can't execute query %s:%s",loginSQLlog,receiver.Err.Error())
}

/////
var ErrInvalidPass=errors.New("Invalid password")
func queryError(query string,err error) *QueryError{
	return &QueryError{Query: query, Err: err}
}
func Init(db *sql.DB) (err error) {
	ddls := []string{clientsDDL,bankMachinesDDL,servicesDDL}
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
func Login(login string,password int,db *sql.DB)(int64,bool,error){
var dbLogin string
var dbPassword int
var dbId int64
err:=db.QueryRow(
	loginSQLlog,login).Scan(&dbId,&dbLogin,&dbPassword)
	if err != nil {
		if err ==sql.ErrNoRows{
			return 0,false, nil
		}
		return 0,false, queryError(loginSQLlog,err)
	}
	if dbPassword!=password{
		return 0,false, ErrInvalidPass
	}
	return dbId,true, nil
}
func Machine(db *sql.DB,userId int64)(machines []BankMachine, err error){
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
		err=rows.Scan(&machine.Id,&machine.Name,&machine.Street)
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
func Account(db *sql.DB,userId int64)(accounts []BankAccounts, err error){
	rows,err:=db.Query(bankAccount,userId)
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
//////
func AddClients(client Client,db *sql.DB)(err error)  {
	_,err=db.Exec(
		insertClients,
		sql.Named("name",client.Name),
  		sql.Named("login",client.Login),
		sql.Named("password",client.Password),
		sql.Named("bankAccount",client.BankAccount),
		sql.Named("phoneNumber",client.PhoneNumber),
		sql.Named("balance",client.Balance))
	if err != nil {
		return err
	}
	return nil
}
func AddBankMachine(machine BankMachine,db *sql.DB)(err error){
	_,err=db.Exec(
		insertBankMachine,
		sql.Named("name",machine.Name),
		sql.Named("street",machine.Street),
		)
	if err != nil {
		return err
	}
	return nil
}
func AddServices(service Services,db *sql.DB)(err error){
	_,err=db.Exec(
		insertServices,
		sql.Named("name",service.Name),
		sql.Named("price",service.Price),
		)
	if err != nil {
		return err
	}
	return nil
}
func UpdateBalance(update Client,db *sql.DB)(err error){
	_,err=db.Exec(
		updateBalanceSQL,
		sql.Named("id",update.Id),
		sql.Named("balance",update.Balance),
		)
	if err != nil {
		return err
	}
	return nil
}
func TransferPlusByPhoneNumber(phoneNumbers int64,balance uint64,db *sql.DB)(err error){
	_,err=db.Exec(
	updateTransferByPhoneNumberPlus,
		sql.Named("phoneNumber",phoneNumbers),
		sql.Named("balance",balance),
		)
	if err != nil {
		return err
	}
	return nil
}
func TransferMinusByPhoneNumber(transfer Client,db *sql.DB)(err error){
	_,err=db.Exec(
		updateTransferByPhoneNumberMinus,
		sql.Named("phoneNumber",transfer.PhoneNumber),
		sql.Named("balance",transfer.Balance),
		)
	if err != nil {
		return err
	}
	return nil
}
func TransferPlusByBankAccount(bankAccount uint64,balance uint64,db *sql.DB)(err error){
	_,err=db.Exec(
		updateTransferByBankAccountPlus,
		sql.Named("bankAccount",bankAccount),
		sql.Named("balance",balance),
	)
	if err != nil {
		return err
	}
	return nil
}
func TransferMinusByBankAccount(transfer Client,db *sql.DB)(err error){
	_,err=db.Exec(
		updateTransferBYBankAccountMinus,
		sql.Named("bankAccount",transfer.BankAccount),
		sql.Named("balance",transfer.Balance),
	)
	if err != nil {
		return err
	}
	return nil
}
func PayServicesMinus(pay Client,db *sql.DB) (err error) {
	_, err = db.Exec(
		payServicesMinus,
		sql.Named("id",pay.Id),
		sql.Named("balance",pay.Balance),
		)
	if err != nil {
		return err
	}
	return nil
}
func PayServicesPlus(id int64,price uint64,db *sql.DB)(err error)  {
		_,err=db.Exec(
			payServicesplus,
			sql.Named("id",id),
			sql.Named("price",price),
		)
		if err != nil {
			return err
		}
		return nil
	}
func main() {}
