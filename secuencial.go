package main

//importar paquetes necesarios
import (
	"fmt"
	"time"
)

// Funcion para ejecutar el metodo secuencial
func EjecutarSecuencial(n, umbral int, archivoSalida string) (time.Duration, string, interface{}) {
	startTime := time.Now()
	fmt.Println("=== INICIANDO EJECUCION SECUENCIAL ===")

	// Calcular traza de producto de matrices
	fmt.Println("Evaluando condicion (calculando traza de matrices)...")
	traza := CalcularTrazaDeProductoDeMatrices(n)

	// Decidir rama a ejecutar entre A y B
	var ramaEjecutada string
	var resultado interface{}

	fmt.Printf("Condicion evaluada - Traza: %d, Umbral: %d\n", traza, umbral)

	if traza > umbral {
		ramaEjecutada = "A"
		fmt.Println("Ejecutando Rama A (Proof-of-Work)...")
		hash, nonce := SimularProofOfWork("blockdata", 2)
		resultado = fmt.Sprintf("Hash: %s, Nonce: %d", hash, nonce)
	} else {
		ramaEjecutada = "B"
		fmt.Println("Ejecutando Rama B (Busqueda de primos)...")
		primos := EncontrarPrimos(10000)
		resultado = fmt.Sprintf("Encontrados %d numeros primos", len(primos))
	}

	tiempoTotal := time.Since(startTime)

	// Imprimir resultados y retornar tiempo total, rama ejecutada y resultado
	fmt.Printf("\n=== EJECUCION SECUENCIAL COMPLETADA ===\n")
	fmt.Printf("Tiempo inicio: %s\n", startTime.Format("15:04:05.000"))
	fmt.Printf("Tiempo fin: %s\n", time.Now().Format("15:04:05.000"))
	fmt.Printf("Rama ejecutada: %s\n", ramaEjecutada)
	fmt.Printf("Resultado: %v\n", resultado)
	fmt.Printf("Tiempo total: %v\n", tiempoTotal)
	fmt.Printf("======================================\n\n")

	return tiempoTotal, ramaEjecutada, resultado
}
