package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	for {
		fmt.Println("Menú:")
		fmt.Println("1. Ejecutar boltDBcrearAlumne")
		fmt.Println("2. Ejecutar boltDBcrearMateria")
		fmt.Println("3. Ejecutar boltDBcrearComision")
		fmt.Println("4. Ejecutar boltDBcrearPeriodo")
		fmt.Println("5. Ejecutar boltDBcrearHistoriaAcademica")
		fmt.Println("6. Salir")

		var choice int
		fmt.Print("Elija una opción: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Ejecutando boltDBcrearAlumne.go...")
			cmd := exec.Command("go", "run", "/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/boltDBcrearAlumne.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error al ejecutar el archivo: %v\n", err)
			}
		case 2:
			fmt.Println("Ejecutando boltDBcrearMateria.go...")
			cmd := exec.Command("go", "run", "/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/boltDBcrearMateria.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error al ejecutar el archivo: %v\n", err)
			}
		case 3:
			fmt.Println("Ejecutando boltDBcrearComision.go...")
			cmd := exec.Command("go", "run", "/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/boltDBcrearComision.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error al ejecutar el archivo: %v\n", err)
			}
		case 4:
			fmt.Println("Ejecutando boltDBcrearPeriodo.go...")
			cmd := exec.Command("go", "run", "/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/boltDBcrearPeriodo.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error al ejecutar el archivo: %v\n", err)
			}
		case 5:
			fmt.Println("Ejecutando boltDBcrearHistoriaAcad.go...")
			cmd := exec.Command("go", "run", "/home/lilo/herrera-lambrecht-petosa-sotelo-db1/boltDB/boltDBcrearHistoriaAcad.go")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Printf("Error al ejecutar el archivo: %v\n", err)
			}				
		case 6:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción no válida, por favor elija de nuevo.")
		}
	}
}

