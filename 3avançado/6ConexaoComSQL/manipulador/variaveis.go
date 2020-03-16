package manipulador

import "html/template"

//Modelos armazena o modelo de pagina ola
var Modelos = template.Must(template.ParseFiles("6ConexaoComSQL/html/ola.html"))

//ModeloLocal armazena o modelo de pagina local
var ModeloLocal = template.Must(template.ParseFiles("6ConexaoComSQL/html/local.html"))
