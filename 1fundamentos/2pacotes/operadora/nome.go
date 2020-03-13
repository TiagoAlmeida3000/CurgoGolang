package operadora

import (
	"strconv"

	"github.com/TiagoAlmeida3000/cursogo2/1fundamentos/2pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora que o usuario esta usando
var NomeOperadora = "XPTO Telecon"

//PrefixoDaCapitalOperadora  prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
