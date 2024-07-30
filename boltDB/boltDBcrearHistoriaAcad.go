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

type HistoriaAcad struct {
    IDAlumne        int    `json:"id_alumne"`
    Semestre        string `json:"semestre"`
    IDMateria       int    `json:"id_materia"`
    IDComision      int    `json:"id_comision"`
    Estado          string `json:"estado"`
    NotaRegular     int    `json:"nota_regular"`
    NotaFinal       int    `json:"nota_final"`
}

func main(){
	guardarEnBoltDBHistoriaAcad()
	mostrarEnBoltDBHistoriaAcad()
}

func guardarEnBoltDBHistoriaAcad(){
	//Conexion a BD
	db, err := sql.Open("postgres", "user=postgres host=localhost dbname=herrera_lambrecht_petosa_sotelo_db1 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//Creando la bd en bolt
	boltDB, err := bolt.Open("HistoriaAcad.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	// Crea el bucket para historia academica
	err = boltDB.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("historiaacademica"))
		return err
	})
	if err != nil {
		log.Fatal(err)
	}

	//Leer el json
	jsonString, err := ioutil.ReadFile("/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/json/historia_academica.json")
	if err != nil {
		log.Fatal(err)
	}

	var historia_academica []HistoriaAcad
	err = json.Unmarshal([]byte(jsonString), &historia_academica)
	if err != nil {
		log.Fatal(err)
	}

	//Guardar historia academica en la BD
	err = boltDB.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("historiaacademica"))
		for _, c := range historia_academica {
			jsonString, err := json.Marshal(c)
			if err != nil {
				log.Fatal(err)
			}

			err = bucket.Put([]byte(fmt.Sprintf("%d-%s-%d", c.IDAlumne, c.Semestre ,c.IDMateria)), jsonString)
			if err != nil {
				log.Fatal(err)
			}
		}
		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Base de datos historiaacademica y datos guardados correctamente.")
}

func mostrarEnBoltDBHistoriaAcad() {
	boltDB, err := bolt.Open("HistoriaAcad.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer boltDB.Close()

	err = boltDB.View(func(tx *bolt.Tx) error {
		// Traer bucket "historiaacademica"
		historiaacademicaBucket := tx.Bucket([]byte("historiaacademica"))
		if historiaacademicaBucket == nil {
			log.Fatal(err)
		}

		//Iterar bucket y mostrar por pantalla
		err := historiaacademicaBucket.ForEach(func(k, v []byte) error {
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
