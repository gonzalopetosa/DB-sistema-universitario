package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"io/ioutil"
	_ "github.com/lib/pq"
	"github.com/boltdb/bolt"
)

type Materia struct {
    ID     int    `json:"id_materia"`
    Nombre string `json:"nombre"`
}

func main()  {
	guardarEnBoltDBMateria()
	mostrarEnBoltDBMateria()
}

func guardarEnBoltDBMateria(){
	//Conexion a BD
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=herrera_lambrecht_petosa_sotelo_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creando la bd en bolt
	boltDB, err := bolt.Open("Materia.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// Crea el bucket para materia
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("materia"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//Leer el json
	jsonString, err := ioutil.ReadFile("/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/json/materias.json")
	if err != nil {
		log.Fatal(err)
	}

	var materias []Materia
	err = json.Unmarshal([]byte(jsonString), &materias)
	if err != nil {
		log.Fatal(err)
	}

	//Guardar materias en la BD
	err = boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("materia"))
		for _, c := range materias {
			jsonString, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}

			err = bucket.Put([]byte(fmt.Sprintf("%d", c.ID)), jsonString)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de datos materia y datos guardados correctamente.")
}

func mostrarEnBoltDBMateria() {
	boltDB, err := bolt.Open("Materia.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	err = boltDB.View(func(tx *bolt.Tx) error {
		// Traer bucket "materia"
		materiaBucket := tx.Bucket([]byte("materia"))
		if materiaBucket == nil {
			log.Fatal(err)
		}

		//Iterar bucket y mostrar por pantalla
		err := materiaBucket.ForEach(func(k, v []byte) error {
			fmt.Printf("Clave: %s\n", k)
			fmt.Printf("Valor: %s\n", v)
			fmt.Println("--------------------")
			return nil
		})
		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}
}
