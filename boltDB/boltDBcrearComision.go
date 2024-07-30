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

type Comision struct {
    IDMateria int `json:"id_materia"`
    IDComision int `json:"id_comision"`
    Cupo int `json:"cupo"`
}

func main(){
guardarEnBoltDBComision()
mostrarEnBoltDBComision()	
}


func guardarEnBoltDBComision(){
	//Conexion a BD
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=herrera_lambrecht_petosa_sotelo_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creando la bd en bolt
	boltDB, err := bolt.Open("Comision.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// Crea el bucket para comision
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("comision"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//Leer el json
	jsonString, err := ioutil.ReadFile("/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/json/comisiones.json")
	if err != nil {
		log.Fatal(err)
	}

	var comisiones []Comision
	err = json.Unmarshal([]byte(jsonString), &comisiones)
	if err != nil {
		log.Fatal(err)
	}

	//Guardar comisiones en la BD
	err = boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("comision"))
		for _, c := range comisiones {
			jsonString, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}

			err = bucket.Put([]byte(fmt.Sprintf("%d", c.IDComision)), jsonString)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de datos comision y datos guardados correctamente.")
}

func mostrarEnBoltDBComision() {
	boltDB, err := bolt.Open("Comision.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	err = boltDB.View(func(tx *bolt.Tx) error {
		// Traer bucket "comision"
		comisionBucket := tx.Bucket([]byte("comision"))
		if comisionBucket == nil {
			log.Fatal(err)
		}

		//Iterar bucket y mostrar por pantalla
		err := comisionBucket.ForEach(func(k, v []byte) error {
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
