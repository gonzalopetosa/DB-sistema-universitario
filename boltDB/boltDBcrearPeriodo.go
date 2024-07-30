package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	_ "github.com/lib/pq"
	"github.com/boltdb/bolt"
)

type Periodo struct {
    
    Semestre    string `json:"semestre"`
    Estado      string `json:"estado"`
}

func main() {
	guardarEnBoltDBperiodo()
	mostrarEnBoltDBPeriodo()
}

func guardarEnBoltDBperiodo(){
	//Conexion a BD
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=herrera_lambrecht_petosa_sotelo_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creando la bd en bolt
	boltDB, err := bolt.Open("Periodo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// Crea el bucket para periodo
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("periodo"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//Leer el json
	jsonString, err := ioutil.ReadFile("/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/json/periodos.json")
	if err != nil {
		log.Fatal(err)
	}

	var periodos []Periodo
	err = json.Unmarshal([]byte(jsonString), &periodos)
	if err != nil {
		log.Fatal(err)
	}

	//Guardar periodos en la BD
	err = boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("periodo"))
		for _, c := range periodos {
			jsonString, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}

			err = bucket.Put([]byte(c.Semestre), jsonString)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de datos periodo y datos guardados correctamente.")
}

func mostrarEnBoltDBPeriodo() {
	boltDB, err := bolt.Open("Periodo.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	err = boltDB.View(func(tx *bolt.Tx) error {
		// Traer bucket "periodo"
		periodoBucket := tx.Bucket([]byte("periodo"))
		if periodoBucket == nil {
			log.Fatal(err)
		}

		//Iterar bucket y mostrar por pantalla
		err := periodoBucket.ForEach(func(k, v []byte) error {
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
