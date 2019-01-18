package manipulador

import "html/template"

//modelos armazena os modelos de pagina que serão executados pelos manipuladores
var ModeloOla = template.Must(template.ParseFiles("html/ola.html"))

var ModeloLocal = template.Must(template.ParseFiles("html/local.html"))
