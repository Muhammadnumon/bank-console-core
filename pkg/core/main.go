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

func (receiver QueryError) Error() string {
	return fmt.Sprintf("can't execute query %s:%s",loginSQL,receiver.Err.Error())
}

/////
var ErrInvalidPass=errors.New("Invalid password")
func queryError(query string,err error) *QueryError{
	return &QueryError{Query: query, Err: err}
}
func Login(login string,password int,db sql.DB)(bool,error){
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

func main() {
	
}
