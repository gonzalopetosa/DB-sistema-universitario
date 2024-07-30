package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = ""
    initialdb   = "postgres"
    targetdb = "herrera_lambrecht_petosa_sotelo_db1"

)

func main() {

    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Println("Menu:")
        fmt.Println("1. Crear base de datos")
		fmt.Println("2. Crear tablas")
        fmt.Println("3. Crear constraints")
        fmt.Println("4. Crear stored procedures")
        fmt.Println("5. Crear triggers")
        fmt.Println("6. Eliminar constraints")
        fmt.Println("7. Cargar tablas")
        fmt.Println("8. Mas opciones->")
        fmt.Println("9. Ejecutar test")
        fmt.Println("0. Salir")
        fmt.Print("Seleccione una opción: ")

        var choice int
        fmt.Scanln(&choice)

        switch choice {
        case 1: //Crear base de datos
            eliminarBaseDeDatos("../scripts/eliminar_db.sql", initialdb)
            crearBaseDeDatos("../scripts/crear_db.sql",initialdb)
        case 2: //Crear tablas
            crearTablas("../scripts/crear_tablas.sql",targetdb)
        case 3: // Crear constraints
            agregarConstraints("../scripts/agregar_constraints.sql", targetdb)
        case 4: // Crear stored procedures
            agregarStoredProcedure("../scripts/stored_procedures/apertura_inscripcion.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/inscripcion_a_materia.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/baja_inscripcion.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/cierre_inscripcion.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/aplicacion_de_cupos.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/ingreso_nota_cursada.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/cierre_de_cursada.sql", targetdb)
            agregarStoredProcedure("../scripts/stored_procedures/test.sql", targetdb)
        case 5: // Crear triggers
            agregarTrigger("../scripts/triggers/trg_email_inscripcion.sql", targetdb)
            agregarTrigger("../scripts/triggers/trg_actualizar_estado_en_espera.sql", targetdb)
            agregarTrigger("../scripts/triggers/trg_cierre_cursada2.sql", targetdb)
            agregarTrigger("../scripts/triggers/trg_email_baja_inscripcion.sql", targetdb)
            agregarTrigger("../scripts/triggers/trg_email_cupo_aplicado.sql", targetdb)
            agregarTrigger("../scripts/triggers/trg_email_alumne_aceptade.sql", targetdb)
        case 6: // Eliminar constraints
            eliminarConstraints("../scripts/eliminar_constraints.sql", targetdb)
        case 7: // Cargar tablas
            cargarTablas("../scripts/cargar_tablas.sql", targetdb)
        case 8: // Submenu
            Submenu:
				for{
					fmt.Println("Mas opciones:")
					fmt.Println("\t1.Apertura de inscripcion")
					fmt.Println("\t2.Inscripcion a materia")
					fmt.Println("\t3.Baja de inscripcion")
					fmt.Println("\t4.Cierre de inscripcion")
					fmt.Println("\t5.Aplicacion de cupos")
					fmt.Println("\t6.Ingreso de nota de cursada")
					fmt.Println("\t7.Cierre de cursada")
					fmt.Println("\t0.Volver al menu principal")
					fmt.Println("")
					fmt.Println("\tSeleccione una opcion:")
					
					var subChoice int
					fmt.Scanln(&subChoice)

					switch subChoice {
					case 1: 
                        aperturaInscripcion()
					case 2: 
						inscripcionMateria()
					case 3: 
						bajaInscripcion()
					case 4: 
						cierreInscripcion()
					case 5: 
						aplicacionCupos()
					case 6: 
						ingresoNota()
					case 7: 
                        cierreCursada()
					case 0: 
						fmt.Println("Volviendo al menu principal..")
						break Submenu
					default:
						fmt.Println("Opción no válida en el submenu")
								}
					}
      
        case 9: // Ejecutar test                
            ejecutarTest()
		case 0: //Salir
            fmt.Println("Saliendo..")
            return
        default:
            fmt.Println("Opción no válida")
        }

        esperarTeclaOprimida(reader)

    }
}

func crearTablas(fileName string, dbName string){
    fmt.Println("Creando tablas...")

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Tablas creadas.")
}

func cargarTablas(fileName string, dbName string){

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Tablas cargadas.")
}

func agregarConstraints(fileName string, dbName string){

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Constraints agregadas.")
}

func agregarStoredProcedure(fileName string, dbName string){

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Stored Prodecure agregada.")
}

func agregarTrigger(fileName string, dbName string){

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Trigger agregado.")
}

func eliminarConstraints(fileName string, dbName string){
    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Constraints eliminadas.")
}


func crearBaseDeDatos(fileName string, dbName string){
    fmt.Println("Creando base de datos...")

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }
    
    fmt.Println("Base de datos creada.") 
}

func eliminarBaseDeDatos(fileName string, dbName string){
    fmt.Println("Eliminando si existe la base de datos...")

    err := ejecutarScriptSQL(fileName, dbName)

    if err != nil {
        log.Printf("%v", err)
        return
    }

}

func esperarTeclaOprimida(reader *bufio.Reader) {
    fmt.Println("Presione Enter para continuar...")
    
    for {
        input, _ := reader.ReadString('\n')
        if input == "\n" {
            break
        } else {
            fmt.Println("Presione Enter para continuar...")
        }
    }

}

func ejecutarScriptSQL(filePath string, dbName string) error {

    db, err := conexionDB(dbName)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()
    
    sqlFile, err := os.ReadFile(filePath)
    if err != nil {
        return fmt.Errorf("error al leer el archivo SQL: %v", err)
    }

    _, err = db.Exec(string(sqlFile))
    if err != nil {
        return fmt.Errorf("error al ejecutar el archivo SQL: %v", err)
    }
    
    return nil
}

func conexionDB( dbName string) (*sql.DB , error){
    psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable",
    host, user, dbName)

    db, err := sql.Open("postgres", psqlInfo)
    if err != nil {
        return nil,fmt.Errorf("error al conectar con la base de datos: %v", err)
    }
    return db,nil
}


// ---- Stored Procedures

func aperturaInscripcion() {
    var año int
    var semestre int
    fmt.Print("Ingrese el año: ")
    fmt.Scanln(&año)
    fmt.Print("Ingrese el número de semestre: ")
    fmt.Scanln(&semestre)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT apertura_inscripcion($1, $2)", año, semestre).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado de apertura de inscripción:", result)
    }
}

func inscripcionMateria() {
    var alumne int
    var materia int
    var comision int
    fmt.Print("Ingrese el id del alumne: ")
    fmt.Scanln(&alumne)
    fmt.Print("Ingrese el id de la materia: ")
    fmt.Scanln(&materia)
    fmt.Print("Ingrese el id de la comisión: ")
    fmt.Scanln(&comision)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT inscripcion_a_materia($1, $2, $3)", alumne, materia, comision).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado de la inscripción:", result)
    }
}

func bajaInscripcion(){
    var alumne int
    var materia int
    fmt.Print("Ingrese el id del alumne: ")
    fmt.Scanln(&alumne)
    fmt.Print("Ingrese el id de la materia: ")
    fmt.Scanln(&materia)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT baja_inscripcion($1, $2)", alumne, materia).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado de la baja de inscripción:", result)
    }
}

func cierreInscripcion(){
    var año int
    var semestre int
    fmt.Print("Ingrese el año: ")
    fmt.Scanln(&año)
    fmt.Print("Ingrese el número de semestre: ")
    fmt.Scanln(&semestre)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT cerrar_inscripcion($1, $2)", año, semestre).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado del cierre de inscripción:", result)
    } 
}

func aplicacionCupos(){
    var año int
    var semestre int
    fmt.Print("Ingrese el año: ")
    fmt.Scanln(&año)
    fmt.Print("Ingrese el número de semestre: ")
    fmt.Scanln(&semestre)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT aplicar_cupos($1, $2)", año, semestre).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado de la aplicación de cupos:", result)
    } 
}

func ingresoNota(){
    var alumne int
    var materia int
    var comision int
    var nota int
    fmt.Print("Ingrese el id del alumne: ")
    fmt.Scanln(&alumne)
    fmt.Print("Ingrese el id de la materia: ")
    fmt.Scanln(&materia)
    fmt.Print("Ingrese el id de la comisión: ")
    fmt.Scanln(&comision)
    fmt.Print("Ingrese la nota: ")
    fmt.Scanln(&nota)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT ingresar_nota_cursada($1, $2, $3, $4)", alumne, materia, comision, nota).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado del ingreso de la nota de cursada:", result)
    }

}

func cierreCursada(){
    var materia int
    var comision int
    fmt.Print("Ingrese el id de la materia: ")
    fmt.Scanln(&materia)
    fmt.Print("Ingrese el id de la comisión: ")
    fmt.Scanln(&comision)

    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    var result bool
    err = db.QueryRow("SELECT cerrar_cursada($1, $2)", materia, comision).Scan(&result)
    if err != nil {
        fmt.Println("Error al ejecutar la stored procedure:", err)
    } else {
        fmt.Println("Resultado del cierre de la nota de cursada:", result)
    }  
}

func ejecutarTest(){
    
    db, err := conexionDB(targetdb)
    if err != nil {
        log.Fatalf("No se pudo conectar a la base de datos: %v", err)
    }
    defer db.Close()

    db.QueryRow("select testear()")
   
}
