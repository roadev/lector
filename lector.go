package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	//"code.google.com/p/x-go-binding/ui/x11"
)

func main() {
	palabras := []string{}
	tablas := []string{}
	datos := []string{}
	tipos := []string{}
	primarias := []string{}
	foraneas := []string{}
	foraneas2 := []string{}
	foraneas3 := []string{}
	restricciones := []string{}
	dobles := [][]string{}
	triples := [][][]string{}
	palabra := ""

	//-------Variables de control--------
	//  lenguaje := "PHP"  //v 0.1
	//lenguaje := "Laravel" //v 0.9
	//  lenguaje := "Django"  //v 0.3
	lenguaje := "Rails" //v 0.1

	dbHost := "localhost"
	dbUser := "usuario"
	dbPass := "contraseña"
	dbName := "mydb"
	autor := "David Vanegas"
	//-------FIN - Variables de control--------

	content, err := ioutil.ReadFile("lectura3.txt")
	if err != nil {
		fmt.Printf("%v\n", "error de lectura")
	}
	lines := strings.Split(string(content), "\n")

	for v := range lines {
		line := strings.Split(string(lines[v]), " ")
		for l := range line {

			lin := strings.Split(string(line[l]), ",")
			if len(lin) > 1 {
				palabras = append(palabras, lin[0])
				palabras = append(palabras, ",")
			} else if lin[0] != "" {
				palabras = append(palabras, lin[0])
			}

		} //for l
	} //for v

	cantidad := len(palabras)
	d := 0
	for d < cantidad {

		if palabras[d] == "CREATE" {
			tablas = append(tablas, palabras[d+2])
			d = d + 3
		}
		if palabras[d] == "(" {

			datos = append([]string{}, palabras[d+1])
			if palabras[d+2] == "character" {
				tipos = append([]string{}, palabras[d+2]+"_"+palabras[d+3])
				d = d + 4
				palabra = ""
				for palabras[d] != "," {
					palabra = palabra + palabras[d] + "_"
					d = d + 1
				}
				restricciones = append(restricciones, palabra)

			} else if palabras[d+2] == "timestamp" {
				tipos = append([]string{}, palabras[d+2]+"_"+palabras[d+3]+"_"+palabras[d+4]+"_"+palabras[d+5])
				d = d + 6
				palabra = ""
				for palabras[d] != "," {
					palabra = palabra + palabras[d] + "_"
					d = d + 1
				}
				restricciones = append(restricciones, palabra)

			} else {
				tipos = append([]string{}, palabras[d+2])
				d = d + 3
				palabra = ""
				for palabras[d] != "," {
					palabra = palabra + palabras[d] + "_"
					d = d + 1
				}
				restricciones = append(restricciones, palabra)
			}
			d = d - 1

		}
		if palabras[d] == "," {
			if palabras[d+1] != "CONSTRAINT" {

				datos = append(datos, palabras[d+1])
				if palabras[d+2] == "character" {
					tipos = append(tipos, palabras[d+2]+"_"+palabras[d+3])
					d = d + 4
					palabra = ""
					for palabras[d] != "," {
						palabra = palabra + palabras[d] + "_"
						d = d + 1
					}
					restricciones = append(restricciones, palabra)

				} else if palabras[d+2] == "timestamp" {
					tipos = append(tipos, palabras[d+2]+"_"+palabras[d+3]+"_"+palabras[d+4]+"_"+palabras[d+5])
					d = d + 6
					palabra = ""
					for palabras[d] != "," {
						palabra = palabra + palabras[d] + "_"
						d = d + 1
					}
					restricciones = append(restricciones, palabra)

				} else {
					tipos = append(tipos, palabras[d+2])
					d = d + 3
					palabra = ""
					for palabras[d] != "," {
						palabra = palabra + palabras[d] + "_"
						d = d + 1
					}
					restricciones = append(restricciones, palabra)

				}
				d = d - 1
			}
		}

		if palabras[d] == "PRIMARY" {

			d = d + 2
			primarias = append([]string{})
			palabra = ""

			if strings.Contains(palabras[d], "(") && strings.Contains(palabras[d], ")") {
				primarias = append(primarias, (palabras[d])[1:len((palabras[d]))-1])
			} else if strings.Contains(palabras[d], "(") {
				primarias = append(primarias, (palabras[d])[1:len((palabras[d]))])
				d = d + 1
				for !strings.Contains(palabras[d], ")") {
					if palabras[d] != "," {
						primarias = append(primarias, palabras[d])
					}
					d = d + 1
				}
				primarias = append(primarias, (palabras[d])[0:len((palabras[d]))-1])
			}

		}

		if palabras[d] == "FOREIGN" {
			d = d + 2
			foraneas = append(foraneas, (palabras[d])[1:len(palabras[d])-1])
			foraneas2 = append(foraneas2, palabras[d+2])
			foraneas3 = append(foraneas3, (palabras[d+3])[1:len(palabras[d+3])-1])
			d = d + 4
		}

		if palabras[d] == ")" {
			dobles = append(dobles, datos)
			dobles = append(dobles, tipos)
			dobles = append(dobles, restricciones)
			dobles = append(dobles, primarias)
			dobles = append(dobles, foraneas)
			dobles = append(dobles, foraneas2)
			dobles = append(dobles, foraneas3)
			triples = append(triples, dobles)

			dobles = append([][]string{})
			restricciones = append([]string{})
			primarias = append([]string{})
			foraneas = append([]string{})
			foraneas2 = append([]string{})
			foraneas3 = append([]string{})

		}
		d = d + 1
	}

	for x := range triples {
		fmt.Printf("tabla %v \n", tablas[x])
		fmt.Printf(" %v \n\n", triples[x])
	}

	if lenguaje == "Rails" {
		os.Mkdir("."+string(filepath.Separator)+"Rails", 0777)
		f := createFile("Rails/commands.txt")
		for y := range tablas {
			writeFile(f, ""+getCommandsRails(triples, tablas, y))
		}
		writeFile(f, " ")
		defer closeFile(f)

	}

	if lenguaje == "Django" {
		os.Mkdir("."+string(filepath.Separator)+"Django", 0777)
		f := createFile("Django/models.py")
		writeFile(f, "from django.db import models")
		writeFile(f, "from django.utils import timezone")
		for y := range tablas {
			writeFile(f, "")
			writeFile(f, "class "+getClase(tablas[y])+"(models.Model):")
			writeFile(f, "    "+getVariablesModeloDjango(triples, tablas, y))
		}
		writeFile(f, "")
		writeFile(f, "")
		defer closeFile(f)

		f = createFile("Django/forms.py")
		writeFile(f, "from django import forms")
		for x := range tablas {
			writeFile(f, "from models import "+getClase(tablas[x]))
		}
		writeFile(f, "")
		for y := range tablas {
			writeFile(f, "")
			writeFile(f, "class "+getClase(tablas[y])+"Form(forms.ModelForm):")
			writeFile(f, "    class Meta:")
			writeFile(f, "        model = "+getClase(tablas[y]))
		}
		writeFile(f, "")
		defer closeFile(f)

		f = createFile("Django/views.py")
		writeFile(f, "from django.http import HttpResponseRedirect")
		writeFile(f, "from django.core.context_processors import crsf")
		for x := range tablas {
			writeFile(f, "from forms import "+getClase(tablas[x])+"Form")
		}
		writeFile(f, "")
		for y := range tablas {
			writeFile(f, "")
			writeFile(f, "def new_"+(tablas[y])+"(request):")
			writeFile(f, "    if request-method=='POST':")
			writeFile(f, "        form = "+getClase(tablas[y])+"Form(request.POST, request.FILES)")
			writeFile(f, "        if form.is_valid():")
			writeFile(f, "            form.save()")
			writeFile(f, "            return HttpResponseRedirect('/"+tablas[y]+"')")
			writeFile(f, "    else:")
			writeFile(f, "        form = "+getClase(tablas[y])+"Form()")
			writeFile(f, "    return render_to_response('"+tablas[y]+"form.html', {'form':form}, context_instance=RequestContext(request))")
		}
		writeFile(f, "")
		defer closeFile(f)

		os.Mkdir("."+string(filepath.Separator)+"Django/views", 0777)
		for x := range tablas {
			f = createFile("Django/views/" + tablas[x] + "form.html")
			writeFile(f, "{% extends 'base.html' %}")
			writeFile(f, "{% block content %}")
			writeFile(f, "    <form id='form' method='post' enctype='multipart/form-data' action=''>{% csrf_token %}")
			writeFile(f, "        {{ form.as_p }}")
			writeFile(f, "        <p><input type='submit' value='confirm' /></p>")
			writeFile(f, "    </form>")
			defer closeFile(f)
		}

	}

	for x := range tablas {
		//writeFile(f, tablas[x])

		if lenguaje == "Golangweb" {
			os.Mkdir("."+string(filepath.Separator)+"Golang", 0777)
			f := createFile("Golang/db_abstract_model.php")
			writeFile(f, "package main")
			writeFile(f, "")
			writeFile(f, "import (")
			writeFile(f, "    \"database/sql\"")
			writeFile(f, "    \"fmt\"")
			writeFile(f, "    _ \"github.com/lib/pq\"")
			writeFile(f, "    \"time\"")
			writeFile(f, "    )")
			writeFile(f, "")
			writeFile(f, "const (")
			writeFile(f, "    dbUser     = \""+dbUser+"\"")
			writeFile(f, "    dbPassWORD = \""+dbPass+"\"")
			writeFile(f, "    dbName     = \""+dbName+"\"")
			writeFile(f, ")")
			writeFile(f, "")
			writeFile(f, "func main() {")
			writeFile(f, "    dbinfo := fmt.Sprintf(\"user=%s password=%s dbname=%s sslmode=disable\", dbUser, dbPassWORD, dbName)")
			writeFile(f, "    db, err := sql.Open(\"postgres\", dbinfo)")
			writeFile(f, "    checkErr(err)")
			writeFile(f, "    defer db.Close()")
			writeFile(f, "")
			writeFile(f, "    ")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")

			defer closeFile(f)

		}

		if lenguaje == "PHP" {
			os.Mkdir("."+string(filepath.Separator)+"BD", 0777)
			f := createFile("BD/db_abstract_model.php")
			writeFile(f, "<?php")
			writeFile(f, "#Hecho por "+autor)
			writeFile(f, "abstract class DBAbstractModel {")
			writeFile(f, "  private static $dbHost = '"+dbHost+"';")
			writeFile(f, "  private static $dbUser = '"+dbUser+"';")
			writeFile(f, "  private static $dbPass = '"+dbPass+"';")
			writeFile(f, "  protected $dbName = '"+dbName+"';")
			writeFile(f, "  protected $query;")
			writeFile(f, "  protected $rows = array();")
			writeFile(f, "  private $conn;")
			writeFile(f, "")
			writeFile(f, "  # métodos abstractos para ABM de clases que hereden")
			writeFile(f, "  abstract protected function get();")
			writeFile(f, "  abstract protected function set();")
			writeFile(f, "  abstract protected function edit();")
			writeFile(f, "  abstract protected function delete();")
			writeFile(f, "")
			writeFile(f, "  # los siguientes métodos pueden definirse con exactitud")
			writeFile(f, "  # y no son abstractos")
			writeFile(f, "  # Conectar a la base de datos")
			writeFile(f, "  private function open_connection() {")
			writeFile(f, "    $this->conn = new mysqli(self::$dbHost, self::$dbUser, self::$dbPass, $this->dbName);")
			writeFile(f, "  }")
			writeFile(f, "")
			writeFile(f, "  # Desconectar la base de datos")
			writeFile(f, "  private function close_connection() {")
			writeFile(f, "    $this->conn->close();")
			writeFile(f, "  }")
			writeFile(f, "")
			writeFile(f, "  # Ejecutar un query simple del tipo INSERT, DELETE, UPDATE")
			writeFile(f, "  protected function execute_single_query() {")
			writeFile(f, "    $this->open_connection();")
			writeFile(f, "    $this->conn->query($this->query);")
			writeFile(f, "    $this->close_connection();")
			writeFile(f, "   }")
			writeFile(f, "")
			writeFile(f, "  # Traer resultados de una consulta en un Array")
			writeFile(f, "  protected function get_results_from_query() {")
			writeFile(f, "    $this->open_connection();")
			writeFile(f, "    $result = $this->conn->query($this->query);")
			writeFile(f, "    while ($this->rows[] = $result->fetch_assoc());")
			writeFile(f, "    $result->close();")
			writeFile(f, "    $this->close_connection();")
			writeFile(f, "    array_pop($this->rows);")
			writeFile(f, "    }")
			writeFile(f, "}")
			writeFile(f, "?>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")
			defer closeFile(f)

			f = createFile("BD/" + getClase(tablas[x]) + ".php")
			writeFile(f, "<?php")
			writeFile(f, "require_once('db_abstract_model.php');")
			writeFile(f, "")
			writeFile(f, "#Hecho por "+autor)
			writeFile(f, "class "+getClase(tablas[x])+" extends DBAbstractModel {")
			writeFile(f, "    "+getVariablesModeloPHP(triples, x))
			writeFile(f, "    function __construct() {")
			writeFile(f, "      $this->dbName = '"+dbName+"';")
			writeFile(f, "    }")
			writeFile(f, "")
			writeFile(f, "    public function get($user_email='') {")
			writeFile(f, "      if($user_email != ''):")
			writeFile(f, "        $this->query = \"")
			writeFile(f, "          SELECT")
			writeFile(f, "          "+getColumnasPHP(triples, x, 0))
			writeFile(f, "          FROM")
			writeFile(f, "          usuarios")
			writeFile(f, "          WHERE")
			writeFile(f, "          email = '$user_email'")
			writeFile(f, "          \";")
			writeFile(f, "        $this->get_results_from_query();")
			writeFile(f, "        endif;")
			writeFile(f, "        if(count($this->rows) == 1):")
			writeFile(f, "          foreach ($this->rows[0] as $propiedad=>$valor):")
			writeFile(f, "            $this->$propiedad = $valor;")
			writeFile(f, "          endforeach;")
			writeFile(f, "        endif;")
			writeFile(f, "      }")
			writeFile(f, "")
			writeFile(f, "      public function set($user_data=array()) {")
			writeFile(f, "      if(array_key_exists('email', $user_data)):")
			writeFile(f, "        $this->get($user_data['email']);")
			writeFile(f, "        if($user_data['email'] != $this->email):")
			writeFile(f, "         foreach ($user_data as $campo=>$valor):")
			writeFile(f, "           $$campo = $valor;")
			writeFile(f, "         endforeach;")
			writeFile(f, "      $this->query = \"")
			writeFile(f, "      INSERT INTO")
			writeFile(f, "      usuarios")
			writeFile(f, "      ("+getColumnasEspecialesPHP(triples, x, 1, 1)+")")
			writeFile(f, "      VALUES")
			writeFile(f, "      ('$nombre', '$apellido', '$email', '$clave')")
			writeFile(f, "      \";")
			writeFile(f, "      $this->execute_single_query();")
			writeFile(f, "      endif;")
			writeFile(f, "      endif;")
			writeFile(f, "      }")
			writeFile(f, "")
			writeFile(f, "      public function edit($user_data=array()) {")
			writeFile(f, "        foreach ($user_data as $campo=>$valor):")
			writeFile(f, "           $$campo = $valor;")
			writeFile(f, "        endforeach;")
			writeFile(f, "      $this->query = \"")
			writeFile(f, "        UPDATE")
			writeFile(f, "        usuarios")
			writeFile(f, "        SET")
			writeFile(f, "        nombre='$nombre',")
			writeFile(f, "        apellido='$apellido',")
			writeFile(f, "        clave='$clave'")
			writeFile(f, "        WHERE")
			writeFile(f, "        email = '$email'")
			writeFile(f, "        \";")
			writeFile(f, "      $this->execute_single_query();")
			writeFile(f, "      }")
			writeFile(f, "")
			writeFile(f, "      public function delete($user_email='') {")
			writeFile(f, "        $this->query = \"")
			writeFile(f, "        DELETE FROM")
			writeFile(f, "        usuarios")
			writeFile(f, "        WHERE")
			writeFile(f, "        email = '$user_email'")
			writeFile(f, "        \";")
			writeFile(f, "      $this->execute_single_query();")
			writeFile(f, "      }")
			writeFile(f, "")
			writeFile(f, "      function __destruct() {")
			writeFile(f, "        unset($this);")
			writeFile(f, "      }")
			writeFile(f, "}")
			writeFile(f, "?>")

			defer closeFile(f)

		}
		if lenguaje == "Lumen" {
			os.Mkdir("."+string(filepath.Separator)+"models", 0777)
			os.Mkdir("."+string(filepath.Separator)+"controllers", 0777)
			os.Mkdir("."+string(filepath.Separator)+"views", 0777)
			os.Mkdir("."+string(filepath.Separator)+"views/"+getClase(tablas[x]), 0777)

			f := createFile("models/" + getClase(tablas[x]) + ".php")
			writeFile(f, "<?php namespace App;")
			writeFile(f, "")
			writeFile(f, "use Illuminate\\Database\\Eloquent\\Model;")
			writeFile(f, "")
			writeFile(f, "class "+getClase(tablas[x])+" extends Model")
			writeFile(f, "{")
			writeFile(f, "")
			writeFile(f, "    //protected $primaryKey = '"+getPrimaryKeys(triples, x)+"';")
			writeFile(f, "    //protected $table = '"+tablas[x]+"';")
			writeFile(f, "")
			writeFile(f, "    protected $fillable = ["+getColumnas(triples, x)+"];")
			writeFile(f, "")
			writeFile(f, "}")
			defer closeFile(f)

			if len(triples[x][3]) == 1 {

				f = createFile("controllers/" + getClase(tablas[x]) + "Controller.php")
				writeFile(f, "<?php")
				writeFile(f, "")
				writeFile(f, "namespace App\\Http\\Controllers;")
				writeFile(f, "")
				writeFile(f, "use App\\"+getClase(tablas[x])+";")
				writeFile(f, "use App\\Http\\Controllers\\Controller;")
				writeFile(f, "use Illuminate\\Http\\Request;")
				writeFile(f, "")
				writeFile(f, "class "+getClase(tablas[x])+"Controller extends Controller")
				writeFile(f, "{")
				writeFile(f, "")
				writeFile(f, "    public function index()")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+"s = "+getClase(tablas[x])+"::all();")
				writeFile(f, "        return response()->json($"+tablas[x]+"s);")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "    public function get"+getClase(tablas[x])+"($id)")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        return response()->json($"+tablas[x]+");")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "    public function save"+getClase(tablas[x])+"(Request $request)")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::create($request->all());")
				writeFile(f, "        return response()->json($"+tablas[x]+");")
				writeFile(f, "    }") //ESTE CREATE ESTA EN DUDA GRANDEMENTE
				writeFile(f, "")
				writeFile(f, "    public function delete"+getClase(tablas[x])+"($id)")
				writeFile(f, "  	{")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        $"+tablas[x]+"->delete();")
				writeFile(f, "        return response()->json('success');")
				writeFile(f, "  	}")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "    public function update"+getClase(tablas[x])+"(Request $request, $id)")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        "+getVariablesUpdateLumen(triples, tablas, x))
				writeFile(f, "        $"+tablas[x]+"->save();")
				writeFile(f, "")
				writeFile(f, "        return response()->json($"+tablas[x]+");")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "}")
				defer closeFile(f)

				os.Mkdir("."+string(filepath.Separator)+"views/"+getClase(tablas[x]), 0777)
				f = createFile("views/" + getClase(tablas[x]) + "/form.blade.php")
				writeFile(f, "@extends ('template')")
				writeFile(f, "")
				writeFile(f, "<?php")
				writeFile(f, "if ($"+tablas[x]+"->exists):")
				writeFile(f, "    $form_data = array('route' => array('"+tablas[x]+".update', $"+tablas[x]+"->"+triples[x][3][0]+"), 'method' => 'PATCH', 'files'=> true);")
				writeFile(f, "    $action    = 'Editar';")
				writeFile(f, "else:")
				writeFile(f, "    $form_data = array('route' => '"+tablas[x]+".store', 'method' => 'POST', 'files'=> true);")
				writeFile(f, "    $action    = 'Crear';")
				writeFile(f, "endif;")
				writeFile(f, "?>")
				writeFile(f, "")
				writeFile(f, "@section ('title') {{ $action }} "+getClase(tablas[x])+" @stop")
				writeFile(f, "")
				writeFile(f, "@section ('contenido')")
				writeFile(f, "<center>")
				writeFile(f, "    <h1>")
				writeFile(f, "        {{ $action }} "+getClase(tablas[x]))
				writeFile(f, "    </h1>")
				writeFile(f, "</center>")
				writeFile(f, "")
				writeFile(f, "<p>")
				writeFile(f, "    <a href=\"{{ route('"+tablas[x]+".index') }}\" class=\"btn btn-info\">Lista de "+getClase(tablas[x])+"s</a>")
				writeFile(f, "</p>")
				writeFile(f, "<br>")
				writeFile(f, "")
				writeFile(f, "{{ Form::model($"+tablas[x]+", $form_data, array('role' => 'form')) }}")
				writeFile(f, "")
				writeFile(f, "@include ('errors', array('errors' => $errors))")
				writeFile(f, "")
				writeFile(f, "    "+getVariablesVistaFormLaravel(triples, tablas, x))
				writeFile(f, "")
				writeFile(f, "    {{ Form::button($action . ' "+tablas[x]+"', array('type' => 'submit', 'class' => 'btn btn-primary')) }}")
				writeFile(f, "")
				writeFile(f, "    {{ Form::close() }}")
				writeFile(f, "@stop")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				defer closeFile(f)
			}

			f = createFile("views/" + getClase(tablas[x]) + "/list.blade.php")
			writeFile(f, "@extends ('template')")
			writeFile(f, "")
			writeFile(f, "@section ('title') Lista de "+getClase(tablas[x])+" @stop")
			writeFile(f, "")
			writeFile(f, "@section ('contenido')")
			writeFile(f, "<center>")
			writeFile(f, "    <h1>")
			writeFile(f, "        Lista de "+getClase(tablas[x]))
			writeFile(f, "    </h1>")
			writeFile(f, "</center>")
			writeFile(f, "")
			writeFile(f, "<p>")
			writeFile(f, "    <a href=\"{{ route('"+tablas[x]+".create') }}\" class=\"btn btn-info\">Crear "+getClase(tablas[x])+"</a>")
			writeFile(f, "</p>")
			writeFile(f, "<br>")
			writeFile(f, "")
			writeFile(f, "<div class=\"row\">")
			writeFile(f, "    @foreach ($"+tablas[x]+"s as $"+tablas[x]+")")
			writeFile(f, "        <a href=\"curso/{{ $curso->id_curso }}\">")
			writeFile(f, "            <div class=\"col-md-4 col-sm-6 col-xs-12\">")
			writeFile(f, "                "+getVariablesVistaListaLaravel(triples, tablas, x, "div"))
			writeFile(f, "            </div>")
			writeFile(f, "        </a>")
			writeFile(f, "    @endforeach")
			writeFile(f, "</div>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "<table class=\"row\">")
			writeFile(f, "    @foreach ($"+tablas[x]+"s as $"+tablas[x]+")")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "tabla"))
			writeFile(f, "    @endforeach")
			writeFile(f, "</table>")
			writeFile(f, "@stop")
			writeFile(f, "")
			defer closeFile(f)

			f = createFile("views/" + getClase(tablas[x]) + "/view.blade.php")
			writeFile(f, "@extends ('template')")
			writeFile(f, "")
			writeFile(f, "@section ('title') Ver "+getClase(tablas[x])+" @stop")
			writeFile(f, "")
			writeFile(f, "@section ('content')")
			writeFile(f, "<br>")
			writeFile(f, "")
			writeFile(f, "<div class=\"row\">")
			writeFile(f, "    <div class=\"col-md-4 col-sm-6 col-xs-12\">")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "div"))
			writeFile(f, "    </div>")
			writeFile(f, "</div>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "<table class=\"row\">")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "table-view"))
			writeFile(f, "</table>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, ""+getVariablesForaneosVistaLaravel(triples, tablas, x))
			writeFile(f, "@stop")
			writeFile(f, "")
			defer closeFile(f)
		}

		if lenguaje == "Laravel" {

			os.Mkdir("."+string(filepath.Separator)+"models", 0777)
			os.Mkdir("."+string(filepath.Separator)+"controllers", 0777)
			os.Mkdir("."+string(filepath.Separator)+"views", 0777)
			os.Mkdir("."+string(filepath.Separator)+"views/"+getClase(tablas[x]), 0777)

			f := createFile("models/" + getClase(tablas[x]) + ".php")
			writeFile(f, "<?php")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "use Illuminate\\Auth\\UserTrait;")
			writeFile(f, "use Illuminate\\Auth\\UserInterface;")
			writeFile(f, "use Illuminate\\Auth\\Reminders\\RemindableTrait;")
			writeFile(f, "use Illuminate\\Auth\\Reminders\\RemindableInterface;")
			writeFile(f, "")
			writeFile(f, "class "+getClase(tablas[x])+" extends Eloquent implements UserInterface, RemindableInterface")
			writeFile(f, "{")
			writeFile(f, "")
			writeFile(f, "    use UserTrait, RemindableTrait;")
			writeFile(f, "")
			writeFile(f, "    public $errors;")
			writeFile(f, "    protected $primaryKey = '"+getPrimaryKeys(triples, x)+"';")
			writeFile(f, "")
			writeFile(f, "    protected $table = '"+tablas[x]+"';")
			writeFile(f, "")
			writeFile(f, "    protected $fillable = array("+getColumnas(triples, x)+");")
			writeFile(f, "")
			writeFile(f, "    public $timestamps = false;")
			writeFile(f, "")
			writeFile(f, "    public function isValid($data)")
			writeFile(f, "    {")
			writeFile(f, "        $rules = array(")
			writeFile(f, "        "+getReglasLaravel(triples, x))
			writeFile(f, "        );")
			writeFile(f, "")
			writeFile(f, "        $validator = Validator::make($data, $rules);")
			writeFile(f, "        if ($validator->passes())")
			writeFile(f, "        {")
			writeFile(f, "            return true;")
			writeFile(f, "        }")
			writeFile(f, "")
			writeFile(f, "        $this->errors = $validator->errors();")
			writeFile(f, "        return false;")
			writeFile(f, "    }")
			writeFile(f, "")
			writeFile(f, "    public function validAndSave($data){")
			writeFile(f, "        if ($this->isValid($data)){")
			writeFile(f, "            $this->fill($data);")
			writeFile(f, "            $this->save();")
			writeFile(f, "            return true;")
			writeFile(f, "         }")
			writeFile(f, "        return false;")
			writeFile(f, "    }")
			writeFile(f, "")
			writeFile(f, "    "+getForaneos(triples, tablas, x))
			writeFile(f, "")
			writeFile(f, "}")
			defer closeFile(f)

			if len(triples[x][3]) == 1 {

				f = createFile("controllers/" + getClase(tablas[x]) + "Controller.php")
				writeFile(f, "<?php")
				writeFile(f, "")
				writeFile(f, "class "+getClase(tablas[x])+"Controller extends BaseController")
				writeFile(f, "{")
				writeFile(f, "")
				writeFile(f, "  /**")
				writeFile(f, "   * Display a listing of the resource.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "")
				writeFile(f, "    public function index()")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+"s = "+getClase(tablas[x])+"::all();")
				writeFile(f, "        return View::make('"+getClase(tablas[x])+"/lista')->with('"+tablas[x]+"s', $"+tablas[x]+"s);")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "    public function lista()")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+"s = "+getClase(tablas[x])+"::all();")
				writeFile(f, "        return View::make('"+getClase(tablas[x])+"/lista')->with('"+tablas[x]+"s', $"+tablas[x]+"s);")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Show the form for creating a new resource.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "    public function create()")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = new "+getClase(tablas[x])+";")
				writeFile(f, "        "+getForaneosControlador(triples, tablas, x))
				writeFile(f, "        return View::make('"+getClase(tablas[x])+"/form')->with('"+tablas[x]+"', $"+tablas[x]+")"+getForaneosRetornadosControlador(triples, tablas, x))
				writeFile(f, "    }") //ESTE CREATE ESTA EN DUDA GRANDEMENTE
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Store a newly created resource in storage.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "    public function store()")
				writeFile(f, "  	{")
				writeFile(f, "        $"+tablas[x]+" = new "+getClase(tablas[x])+";")
				writeFile(f, "  		  $data = Input::all();")
				writeFile(f, "")
				writeFile(f, "  		  if ($"+tablas[x]+"->isValid($data))")
				writeFile(f, "    		{")
				writeFile(f, "    			  $"+tablas[x]+"->fill($data);")
				writeFile(f, "    			  $"+tablas[x]+"->save();")
				writeFile(f, "    			  return Redirect::route('"+tablas[x]+".show', array($"+tablas[x]+"->"+triples[x][3][0]+"));")
				writeFile(f, "    		}")
				writeFile(f, "    		else")
				writeFile(f, "    		{")
				writeFile(f, "  			    return Redirect::route('"+tablas[x]+".create')->withInput()->withErrors($"+tablas[x]+"->errors);")
				writeFile(f, "    		}")
				writeFile(f, "  	}")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Display the specified resource.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @param  int  $id")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "    public function show($id)")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        return View::make('"+getClase(tablas[x])+"/view')->with('"+tablas[x]+"', $"+tablas[x]+");")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Show the form for editing the specified resource.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @param  int  $id")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "  	 public function edit($id)")
				writeFile(f, "  	 {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        if (is_null ($"+tablas[x]+"))")
				writeFile(f, "        {")
				writeFile(f, "            App::abort(404);")
				writeFile(f, "        }")
				writeFile(f, "")
				writeFile(f, "        $form_data = array('route' => array('$"+tablas[x]+".update', $"+tablas[x]+"->"+triples[x][3][0]+"), 'method' => 'PATCH');")
				writeFile(f, "        $action = 'Editar';")
				writeFile(f, "")
				writeFile(f, "        "+getForaneosControlador(triples, tablas, x))
				writeFile(f, "        return View::make('Curso/form', compact('$"+tablas[x]+"', 'form_data', 'action'))"+getForaneosRetornadosControlador(triples, tablas, x))
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Update the specified resource in storage.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @param  int  $id")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "    public function update($id)")
				writeFile(f, "    {")
				writeFile(f, "        $"+tablas[x]+" = "+getClase(tablas[x])+"::find($id);")
				writeFile(f, "        $data = Input::all();")
				writeFile(f, "        if ($"+tablas[x]+"->validAndSave($data))")
				writeFile(f, "        {")
				writeFile(f, "            return Redirect::route('"+tablas[x]+".show', array($"+tablas[x]+"->"+triples[x][3][0]+"));")
				writeFile(f, "        }")
				writeFile(f, "        else")
				writeFile(f, "        {")
				writeFile(f, "            return Redirect::route('"+tablas[x]+".edit', $"+tablas[x]+"->"+triples[x][3][0]+")->withInput()->withErrors($curso->errors);")
				writeFile(f, "        }")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "	/**")
				writeFile(f, "	 * Remove the specified resource from storage.")
				writeFile(f, "	 *")
				writeFile(f, "	 * @param  int  $id")
				writeFile(f, "	 * @return Response")
				writeFile(f, "	 */")
				writeFile(f, "    public function destroy($id)")
				writeFile(f, "    {")
				writeFile(f, "")
				writeFile(f, "    }")
				writeFile(f, "")
				writeFile(f, "}")
				defer closeFile(f)

				os.Mkdir("."+string(filepath.Separator)+"views/"+getClase(tablas[x]), 0777)
				f = createFile("views/" + getClase(tablas[x]) + "/form.blade.php")
				writeFile(f, "@extends ('template')")
				writeFile(f, "")
				writeFile(f, "<?php")
				writeFile(f, "if ($"+tablas[x]+"->exists):")
				writeFile(f, "    $form_data = array('route' => array('"+tablas[x]+".update', $"+tablas[x]+"->"+triples[x][3][0]+"), 'method' => 'PATCH', 'files'=> true);")
				writeFile(f, "    $action    = 'Editar';")
				writeFile(f, "else:")
				writeFile(f, "    $form_data = array('route' => '"+tablas[x]+".store', 'method' => 'POST', 'files'=> true);")
				writeFile(f, "    $action    = 'Crear';")
				writeFile(f, "endif;")
				writeFile(f, "?>")
				writeFile(f, "")
				writeFile(f, "@section ('title') {{ $action }} "+getClase(tablas[x])+" @stop")
				writeFile(f, "")
				writeFile(f, "@section ('contenido')")
				writeFile(f, "<center>")
				writeFile(f, "    <h1>")
				writeFile(f, "        {{ $action }} "+getClase(tablas[x]))
				writeFile(f, "    </h1>")
				writeFile(f, "</center>")
				writeFile(f, "")
				writeFile(f, "<p>")
				writeFile(f, "    <a href=\"{{ route('"+tablas[x]+".index') }}\" class=\"btn btn-info\">Lista de "+getClase(tablas[x])+"s</a>")
				writeFile(f, "</p>")
				writeFile(f, "<br>")
				writeFile(f, "")
				writeFile(f, "{{ Form::model($"+tablas[x]+", $form_data, array('role' => 'form')) }}")
				writeFile(f, "")
				writeFile(f, "@include ('errors', array('errors' => $errors))")
				writeFile(f, "")
				writeFile(f, "    "+getVariablesVistaFormLaravel(triples, tablas, x))
				writeFile(f, "")
				writeFile(f, "    {{ Form::button($action . ' "+tablas[x]+"', array('type' => 'submit', 'class' => 'btn btn-primary')) }}")
				writeFile(f, "")
				writeFile(f, "    {{ Form::close() }}")
				writeFile(f, "@stop")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				writeFile(f, "")
				defer closeFile(f)
			}

			f = createFile("views/" + getClase(tablas[x]) + "/list.blade.php")
			writeFile(f, "@extends ('template')")
			writeFile(f, "")
			writeFile(f, "@section ('title') Lista de "+getClase(tablas[x])+" @stop")
			writeFile(f, "")
			writeFile(f, "@section ('contenido')")
			writeFile(f, "<center>")
			writeFile(f, "    <h1>")
			writeFile(f, "        Lista de "+getClase(tablas[x]))
			writeFile(f, "    </h1>")
			writeFile(f, "</center>")
			writeFile(f, "")
			writeFile(f, "<p>")
			writeFile(f, "    <a href=\"{{ route('"+tablas[x]+".create') }}\" class=\"btn btn-info\">Crear "+getClase(tablas[x])+"</a>")
			writeFile(f, "</p>")
			writeFile(f, "<br>")
			writeFile(f, "")
			writeFile(f, "<div class=\"row\">")
			writeFile(f, "    @foreach ($"+tablas[x]+"s as $"+tablas[x]+")")
			writeFile(f, "        <a href=\"curso/{{ $curso->id_curso }}\">")
			writeFile(f, "            <div class=\"col-md-4 col-sm-6 col-xs-12\">")
			writeFile(f, "                "+getVariablesVistaListaLaravel(triples, tablas, x, "div"))
			writeFile(f, "            </div>")
			writeFile(f, "        </a>")
			writeFile(f, "    @endforeach")
			writeFile(f, "</div>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "<table class=\"row\">")
			writeFile(f, "    @foreach ($"+tablas[x]+"s as $"+tablas[x]+")")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "tabla"))
			writeFile(f, "    @endforeach")
			writeFile(f, "</table>")
			writeFile(f, "@stop")
			writeFile(f, "")
			defer closeFile(f)

			f = createFile("views/" + getClase(tablas[x]) + "/view.blade.php")
			writeFile(f, "@extends ('template')")
			writeFile(f, "")
			writeFile(f, "@section ('title') Ver "+getClase(tablas[x])+" @stop")
			writeFile(f, "")
			writeFile(f, "@section ('content')")
			writeFile(f, "<br>")
			writeFile(f, "")
			writeFile(f, "<div class=\"row\">")
			writeFile(f, "    <div class=\"col-md-4 col-sm-6 col-xs-12\">")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "div"))
			writeFile(f, "    </div>")
			writeFile(f, "</div>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, "<table class=\"row\">")
			writeFile(f, "        "+getVariablesVistaListaLaravel(triples, tablas, x, "table-view"))
			writeFile(f, "</table>")
			writeFile(f, "")
			writeFile(f, "")
			writeFile(f, ""+getVariablesForaneosVistaLaravel(triples, tablas, x))
			writeFile(f, "@stop")
			writeFile(f, "")
			defer closeFile(f)

		}

	} //FIN del for range tablas
} //Fin del main

//funciones GO para creacion de archivos
func createFile(p string) *os.File {
	fmt.Println("crear")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File, val string) {
	//fmt.Println("escribir")
	fmt.Fprintln(f, val)

}

func closeFile(f *os.File) {
	fmt.Println("cerrar")
	f.Close()
}

func getClase(nombre string) string {
	//clase := strings.ToUpper((nombre)[0:1])+(nombre)[1:len(nombre)]
	lineas := strings.Split(string(nombre), "_")
	respuesta := ""
	k := 0
	for k < len(lineas) {
		respuesta = respuesta + strings.ToUpper((lineas[k])[0:1]) + (lineas[k])[1:len(lineas[k])]
		k = k + 1
	}
	return respuesta
}

//Este metodo hace referencia a los foraneos en LARAVEL llamados en el view
func getVariablesForaneosVistaLaravel(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	cantidad := len(triples[element][4])
	cantidadForaneos := len(triples)
	k := 0
	for k < cantidad {
		resultado = resultado + "{{ $" + tablas[element] + "->get" + getClase(triples[element][5][k]) + "() }} " + "\n\t\t"
		k = k + 1
	}
	k = 0
	for k < cantidadForaneos {
		for x := range triples[k][5] {
			if triples[k][5][x] == tablas[element] {
				resultado = resultado + "@foreach( $" + tablas[element] + "->get" + getClase(tablas[k]) + "() as $x_" + tablas[k] + " ) " + "\n\t\t\t"
				for y := range triples[k][0] {
					resultado = resultado + "{{ $x_" + tablas[k] + "->" + triples[k][0][y] + " }} " + "\n\t\t\t"
				}
				resultado = resultado + "@endforeach " + "\n\t\t"
			}
		}
		k = k + 1
	}
	return resultado
}

//Metodo que me retorna las variables de la vista (views list y view view) de los modelos
//TENER MUY EN CUENTA PARA FUTUROS FORMS DE OTROS TIPOS
func getVariablesVistaListaLaravel(triples [][][]string, tablas []string, element int, tipoResultado string) string {
	resultado := ""
	cantidad := len(triples[element][0])
	k := 0
	if tipoResultado == "tabla" {
		resultado = resultado + "<tr>" + "\n\t\t\t"
	}
	for k < cantidad {
		if tipoResultado == "tabla" {
			resultado = resultado + "<td>{{ $" + tablas[element] + "->" + triples[element][0][k] + " }}</td>" + "\n\t\t\t\t"
		} else if tipoResultado == "table-view" {
			resultado = resultado + "<tr>" + "\n\t\t\t"
			resultado = resultado + "<th>" + getClase(triples[element][0][k]) + "</th>" + "\n\t\t\t"
			resultado = resultado + "<td>{{ $" + tablas[element] + "->" + triples[element][0][k] + " }}</td>" + "\n\t\t"
			resultado = resultado + "</tr>" + "\n\t\t"
		} else {
			resultado = resultado + "{{ $" + tablas[element] + "->" + triples[element][0][k] + " }}" + "\n\t\t\t"
		}
		k = k + 1
	}
	if tipoResultado == "tabla" {
		resultado = resultado + "</tr>" + "\n\t\t\t"
	}
	return resultado
}

//Metodo que me retorna las variables de la vista (views forms) de los modelos
//TENER MUY EN CUENTA PARA FUTUROS FORMS DE OTROS TIPOS
func getVariablesVistaFormLaravel(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	cantidad := len(triples[element][0])
	k := 0
	for k < cantidad {
		if triples[element][1][k] == "serial" {

		} else if stringInSlice(triples[element][0][k], triples[element][3]) && stringInSlice(triples[element][0][k], triples[element][4]) {
			resultado = resultado + "<div class=\"row\">" + "\n\t\t"
			resultado = resultado + "<div class=\"form-group col-md-5\">" + "\n\t\t\t"
			resultado = resultado + "{{ Form::label('" + triples[element][0][k] + "', '" + getClase(triples[element][0][k]) + " del " + getClase(tablas[element]) + "') }}" + "\n\t\t\t"
			resultado = resultado + "{{ Form::hidden('" + triples[element][0][k] + "', $" + getClaseForaneo(triples[element][0][k], triples, element) + ")) }}" + "\n\t\t"
			resultado = resultado + "</div>" + "\n\t"
			resultado = resultado + "</div>" + "\n\t"
		} else if stringInSlice(triples[element][0][k], triples[element][4]) {
			resultado = resultado + "<div class=\"row\">" + "\n\t\t"
			resultado = resultado + "<div class=\"form-group col-md-5\">" + "\n\t\t\t"
			resultado = resultado + "{{ Form::label('" + triples[element][0][k] + "', '" + getClase(triples[element][0][k]) + " del " + getClase(tablas[element]) + "') }}" + "\n\t\t\t"
			resultado = resultado + "{{ Form::select('" + triples[element][0][k] + "', $" + getClaseForaneo(triples[element][0][k], triples, element) + "s, null, array('placeholder' => 'Escoje', 'class' => 'form-control' )) }}" + "\n\t\t"
			resultado = resultado + "</div>" + "\n\t"
			resultado = resultado + "</div>" + "\n\t"
		} else {
			resultado = resultado + "<div class=\"row\">" + "\n\t\t"
			resultado = resultado + "<div class=\"form-group col-md-5\">" + "\n\t\t\t"
			resultado = resultado + "{{ Form::label('" + triples[element][0][k] + "', '" + getClase(triples[element][0][k]) + " del " + getClase(tablas[element]) + "') }}" + "\n\t\t\t"
			resultado = resultado + "{{ Form::text('" + triples[element][0][k] + "', null, array('Introduce el " + getClase(triples[element][0][k]) + " de " + getClase(tablas[element]) + "', 'class'=>'form-control')) }}" + "\n\t\t"
			resultado = resultado + "</div>" + "\n\t"
			resultado = resultado + "</div>" + "\n\t"
		}
		k = k + 1
	}
	return resultado

}

//Metodo encargado de devolver las primary keys en array con comas
func getPrimaryKeys(triples [][][]string, element int) string {
	resultado := triples[element][3][0] + ""
	cantidad := len(triples[element][3])
	k := 1
	if cantidad == 1 {
		return resultado
	}
	for k < cantidad {
		resultado = resultado + ", " + triples[element][3][k]
		k = k + 1
	}
	return "array(" + resultado + ")"
}

//Metodo que retorna los elementos de un

//Metodo que me retorna las columnas de una tabla, no solo para laravel, pero le pone comillas
func getColumnas(triples [][][]string, element int) string {
	resultado := "'" + triples[element][0][0] + "'"
	cantidad := len(triples[element][0])
	k := 1
	for k < cantidad {
		resultado = resultado + ", '" + triples[element][0][k] + "'"
		k = k + 1
	}
	return resultado
}

//Metodo que me retorna las reglas de validacion del modelo de una tabla en LARAVEL
func getReglasLaravel(triples [][][]string, element int) string {
	resultado := ""
	cantidad := len(triples[element][0])
	k := 0
	for k < cantidad {
		if triples[element][1][k] == "integer" {
			resultado = resultado + "'" + triples[element][0][k] + "' => 'required|numeric'," + "\n\t\t\t"
		} else if triples[element][1][k] == "serial" {

		} else {
			resultado = resultado + "'" + triples[element][0][k] + "' => 'required'," + "\n\t\t\t"
		}
		k = k + 1
	}
	return resultado
}

//Este metodo hace referencia a los foraneos en LARAVEL
func getForaneos(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	cantidad := len(triples[element][4])
	cantidadForaneos := len(triples)
	k := 0
	for k < cantidad {
		resultado = resultado + "public function get" + getClase(triples[element][5][k]) + "(){ " + "\n\t\t\t"
		resultado = resultado + "$" + (triples[element][5][k]) + " = " + getClase(triples[element][5][k]) + "::find($this->" + (triples[element][4][k]) + ");\n\t\t\t"
		resultado = resultado + "return $" + (triples[element][5][k]) + ";\n\t\t}\n\t\t"
		k = k + 1
	}
	k = 0
	for k < cantidadForaneos {
		for x := range triples[k][5] {
			if triples[k][5][x] == tablas[element] {
				resultado = resultado + "public function get" + getClase(tablas[k]) + "Foraneo(){ " + "\n\t\t\t"
				resultado = resultado + "$" + (tablas[k]) + "s = " + getClase(tablas[k]) + "::where('" + (triples[k][6][x]) + "',  '=', $this->" + (triples[k][4][x]) + ")->get();\n\t\t\t"
				resultado = resultado + "return $" + (tablas[k]) + "s;\n\t\t}\n\t\t"
			}
		}
		k = k + 1
	}
	return resultado
}

//Este metodo hace referencia a los foraneos en el create de LARAVEL
func getForaneosControlador(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	cantidad := len(triples[element][4])
	k := 0
	for k < cantidad {
		resultado = resultado + "$" + triples[element][5][k] + "s = " + getClase(triples[element][5][k]) + "::lists('', '" + triples[element][4][k] + "');" + "\n\t\t\t"
		k = k + 1
	}
	return resultado
}

//Este metodo hace referencia a los foraneos en el return de metodos del controlador en LARAVEL
func getForaneosRetornadosControlador(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	cantidad := len(triples[element][4])
	k := 0
	for k < cantidad {
		resultado = resultado + "->with('" + triples[element][5][k] + "', $" + triples[element][5][k] + ")"
		k = k + 1
	}
	resultado = resultado + ";"

	return resultado
}

//
func getVariablesModeloPHP(triples [][][]string, element int) string {
	resultado := ""
	cantidad := len(triples[element][0])
	k := 0
	for k < cantidad {
		if stringInSlice(triples[element][0][k], triples[element][3]) {
			resultado = resultado + "protected $" + triples[element][0][k] + "; " + "\n\t\t\t"
		} else {
			resultado = resultado + "public $" + triples[element][0][k] + "; " + "\n\t\t\t"
		}
		k = k + 1
	}
	return resultado
}

//Metodo que retorna las columnas de una tabla
// tipo: 0 sin nada, 1 como variables php, 2 como variable php con =''
func getColumnasPHP(triples [][][]string, element int, tipo int) string {
	resultado := "" + triples[element][0][0] + ""
	cantidad := len(triples[element][0])
	k := 1
	if tipo == 1 { //como variable PHP
		resultado = "$" + resultado
	}
	if tipo == 2 { //como variable PHP con =''
		resultado = "$" + resultado + "=''"
	}
	for k < cantidad {
		if tipo == 0 { //como variable PHP
			resultado = resultado + ", " + triples[element][0][k] + ""
		}
		if tipo == 1 { //como variable PHP
			resultado = resultado + ", $" + triples[element][0][k] + ""
		}
		if tipo == 2 { //como variable PHP
			resultado = resultado + ", $" + triples[element][0][k] + "=''"
		}
		k = k + 1
	}
	return resultado
}

//Metodo que retorna las llaves primarias o todas aquellas que no hacen parte de la llave primaria
// tipoDatos: 0 las llaves primarias, 1 las llaves que no son primarias
func getColumnasEspecialesPHP(triples [][][]string, element int, tipoDatos int, tipoRespuestas int) string {
	resultado := ""
	cantidad := 0
	k := 1
	suf, pref := "", ""

	if tipoDatos == 1 { //las columnas sobrantes que no son llaves primarias
		cantidad = len(triples[element][0])
		for k < cantidad {
			if !stringInSlice(triples[element][0][k], triples[element][3]) {
				resultado = resultado + pref + triples[element][0][k] + suf + ", "
			}
			k = k + 1
		}
		resultado = (resultado)[0 : len(resultado)-2]
	}
	if tipoDatos == 2 { //todas
		cantidad = len(triples[element][0])
		for k < cantidad {
			resultado = resultado + pref + triples[element][0][k] + suf + ", "
			k = k + 1
		}
		resultado = (resultado)[0 : len(resultado)-2]
	} else { //las llaves primarias
		resultado = "" + triples[element][3][0] + ""
		cantidad = len(triples[element][3])
		for k < cantidad {
			resultado = resultado + ", " + pref + triples[element][3][k] + suf + ""
			k = k + 1
		}
	}

	return resultado
}

//
//METODOS LUMEN
//
func getVariablesUpdateLumen(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	for x := range triples[element][0] {
		resultado = resultado + "$" + tablas[element] + "->" + triples[element][0][x] + "= $request->input('" + triples[element][0][x] + "');\n\t\t"
	}
	return resultado
}

//
//METODOS RAILS
//
func getCommandsRails(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	k := 0
	resultado = resultado + "rails generate scaffold " + tablas[element] + " "

	for k = range triples[element][0] {

		if triples[element][1][k] == "integer" {
			resultado = resultado + triples[element][0][k] + ":integer"
		} else if triples[element][1][k] == "serial" {
			resultado = resultado + triples[element][0][k] + ":primary_key"
		} else if triples[element][1][k] == "double" {
			resultado = resultado + triples[element][0][k] + ":decimal"
		} else if triples[element][1][k] == "date" {
			resultado = resultado + triples[element][0][k] + ":date"
		} else if ((triples[element][1][k])[0:9]) == "character" {
			resultado = resultado + triples[element][0][k] + ":string"
		} else if ((triples[element][1][k])[0:9]) == "timestamp" {
			resultado = resultado + triples[element][0][k] + ":timestamp"
		} else {
			resultado = resultado + triples[element][0][k] + ":string"
		}

		//if stringInSlice(triples[element][0][k], triples[element][3]) && stringInSlice(triples[element][0][k], triples[element][4]) {
		//	resultado = resultado + triples[element][0][k] + "= models.ForeignKey(primary_key=True)\n    "
		//}
		if len(triples[element][4]) > 0 {
			if stringInSlice(triples[element][0][k], triples[element][4]) {
				resultado = resultado + ":" + getClaseForaneo(triples[element][0][k], triples, element) + " "
			} else {
				resultado = resultado + " "
			}
		} else if len(triples[element][3]) > 0 {
			if stringInSlice(triples[element][0][k], triples[element][3]) && triples[element][1][k] == "serial" {
				//resultado = resultado + triples[element][0][k] + "= models.AutoField(primary_key=True)\n    "
				resultado = resultado + ":" + getClaseForaneo(triples[element][0][k], triples, element) + " "
			} else {
				resultado = resultado + " "

			}
		} else {
			resultado = resultado + " "
		}

		/* else if stringInSlice(triples[element][0][k], triples[element][3]) {
			resultado = resultado + triples[element][0][k] + "= models.CharField(primary_key=True, max_length=5000)\n    "
		}
		*/
	}

	return resultado
}

//
//METODOS DJANGO
//
func getVariablesModeloDjango(triples [][][]string, tablas []string, element int) string {
	resultado := ""
	//tipo := ""
	k := 0
	for k = range triples[element][0] {
		if stringInSlice(triples[element][0][k], triples[element][3]) && stringInSlice(triples[element][0][k], triples[element][4]) {
			resultado = resultado + triples[element][0][k] + "= models.ForeignKey(primary_key=True)\n    "
		} else if stringInSlice(triples[element][0][k], triples[element][4]) {
			resultado = resultado + triples[element][0][k] + "= models.ForeignKey(" + getClase(getClaseForaneo(triples[element][0][k], triples, element)) + ")\n    "
		} else if stringInSlice(triples[element][0][k], triples[element][3]) && triples[element][1][k] == "serial" {
			resultado = resultado + triples[element][0][k] + "= models.AutoField(primary_key=True)\n    "
		} else if stringInSlice(triples[element][0][k], triples[element][3]) {
			resultado = resultado + triples[element][0][k] + "= models.CharField(primary_key=True, max_length=5000)\n    "
		} else if triples[element][1][k] == "integer" {
			resultado = resultado + triples[element][0][k] + "= models.IntegerField()\n    "
		} else if triples[element][1][k] == "date" {
			resultado = resultado + triples[element][0][k] + "= models.DateField()\n    "
		} else if ((triples[element][1][k])[0:9]) == "timestamp" {
			resultado = resultado + triples[element][0][k] + "= models.DateTimeField(auto_now_add=True)\n    "
		} else if len(triples[element][1][k]) > 17 {
			resultado = resultado + triples[element][0][k] + "= models.CharField(max_length=" + (triples[element][1][k])[18:len(triples[element][1][k])] + "\n    "
		} else {
			resultado = resultado + triples[element][0][k] + "= models.CharField(max_length=5000)\n    "
		}

	}
	return resultado
}

//
//  METODOS AUXILIARES
//
//Metodo para descubrir si una palabra hace parte de un slice o no
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//metodo que retorna la clase de la cual es foranea
func getClaseForaneo(a string, triples [][][]string, element int) string {
	for c, b := range triples[element][4] {
		if b == a {
			return triples[element][5][c]
		}
	}
	return ""
}
