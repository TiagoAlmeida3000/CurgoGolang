package repo

import (
	//"github.com/go-sql-driver/mysql" é usado pela aplicação
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

//Db armazena a conexao com o banco de dados
var Db *sqlx.DB

//AbreConexaoComBancoDeDados funcao que abre a conexao com banco
func AbreConexaoComBancoDeDados() (err error) {
	err = nil
	Db, err = sqlx.Open("mysql", "root:12345@/cursogo2")
	if err != nil {
		return
	}
	err = Db.Ping()
	return
}
