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

type Alumne struct {
    ID        int    `json:"id_alumne"`
    Nombre    string `json:"nombre"`
    Apellido  string `json:"apellido"`
    DNI       int    `json:"dni"`
    FechaNac  string `json:"fecha_nacimiento"`
    Telefono  string `json:"telefono"`
    Email     string `json:"email"`
}

func main()  {
	guardarEnBoltDBAlumne()
	mostrarEnBoltDBAlumne()
}

func guardarEnBoltDBAlumne(){
	//Conexion a BD
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=herrera_lambrecht_petosa_sotelo_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creando la bd en bolt
	boltDB, err := bolt.Open("Alumne.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// Crea el bucket para alumne
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("alumne"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//Leer el json
	jsonString, err := ioutil.ReadFile("/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/json/alumnes.json")
	if err != nil {
		log.Fatal(err)
	}

	var alumnes []Alumne
	err = json.Unmarshal([]byte(jsonString), &alumnes)
	if err != nil {
		log.Fatal(err)
	}

	//Guardar alumnes en la BD
	err = boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("alumne"))
		for _, c := range alumnes {
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

	fmt.Println("Base de datos alumne y datos guardados correctamente.")
}

func mostrarEnBoltDBAlumne() {
	boltDB, err := bolt.Open("Alumne.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	err = boltDB.View(func(tx *bolt.Tx) error {
		// Traer bucket "alumne"
		alumneBucket := tx.Bucket([]byte("alumne"))
		if alumneBucket == nil {
			log.Fatal(err)
		}

		//Iterar bucket y mostrar por pantalla
		err := alumneBucket.ForEach(func(k, v []byte) error {
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
