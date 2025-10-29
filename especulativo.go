package main

//importar paquetes necesarios
import (
	"fmt"
	"time"
)

// Resultado estructura para comunicar resultados de ramas
type Resultado struct {
	rama  string
	valor interface{}
}

// EjecutarRamaA ejecuta la primera rama especulativa
func ejecutarRamaA(data string, dificultad int, stop <-chan struct{}, resultadoChan chan Resultado) {
	fmt.Println("  Rama A: Iniciando Proof-of-Work...")
	hash, nonce := SimularProofOfWork(data, dificultad, stop)
	resultado := hash
	if len(hash) >= 9 && hash[:9] == "cancelled_" {
		resultado = fmt.Sprintf("cancelled (nonce %d)", nonce)
	}
	resultadoChan <- Resultado{
		rama:  "A",
		valor: fmt.Sprintf("Hash: %s, Nonce: %d", resultado, nonce),
	}
	fmt.Println("  Rama A: Resultado enviado")
}

// ejecutarRamaB ejecuta la segunda rama especulativa
func ejecutarRamaB(max int, stop <-chan struct{}, resultadoChan chan Resultado) {
	fmt.Println("  Rama B: Buscando numeros primos...")
	primos := EncontrarPrimos(max, stop)
	if len(primos) == 0 {
		resultado := "cancelled or none"
		resultadoChan <- Resultado{
			rama:  "B",
			valor: resultado,
		}
	} else {
		resultadoChan <- Resultado{
			rama:  "B",
			valor: fmt.Sprintf("Encontrados %d numeros primos", len(primos)),
		}
	}
	fmt.Println("  Rama B: Resultado enviado")
}

// EjecutarEspeculativo ejecuta el metodo especulativo
func EjecutarEspeculativo(n, umbral int, archivoSalida string) (time.Duration, string, interface{}) {
	startTime := time.Now()
	fmt.Println("=== INICIANDO EJECUCION ESPECULATIVA ===")

	// Canal para recibir resultados de las ramas
	resultadoChan := make(chan Resultado, 2)

	// Crear canales de cancelacion para cada rama
	stopA := make(chan struct{})
	stopB := make(chan struct{})

	// Iniciar ambas ramas en goroutines, pasando su respectivo canal de stop
	go ejecutarRamaA("blockdata", 2, stopA, resultadoChan) // Dificultad 2 para evitar pruebas largas
	go ejecutarRamaB(10000, stopB, resultadoChan)          // Solo hasta 10,000 para evitar pruebas largas

	// Calcular traza de producto de matrices
	fmt.Println("Evaluando condicion (calculando traza de matrices)...")
	traza := CalcularTrazaDeProductoDeMatrices(n)

	// Decidir rama ganadora entre A y B comparando entre traza y umbral
	var ramaGanadora string
	if traza > umbral {
		ramaGanadora = "A"
		// cancelar B
		close(stopB)
	} else {
		ramaGanadora = "B"
		// cancelar A
		close(stopA)
	}

	fmt.Printf("Condicion evaluada - Traza: %d, Umbral: %d, Rama ganadora: %s\n", traza, umbral, ramaGanadora)

	// Esperar y recibir resultados de ambas ramas
	fmt.Println("Esperando resultado de la rama ganadora...")
	var resultadoGanador Resultado
	var resultadoPerdedor Resultado

	// Recibir ambos resultados y determinar ganador y perdedor
	for i := 0; i < 2; i++ {
		resultado := <-resultadoChan
		fmt.Printf("Resultado recibido de rama: %s\n", resultado.rama)

		if resultado.rama == ramaGanadora {
			resultadoGanador = resultado
		} else {
			resultadoPerdedor = resultado
		}
	}

	// Imprimir resultado descartado
	fmt.Printf("Descartando resultado de rama %s: %v\n", resultadoPerdedor.rama, resultadoPerdedor.valor)

	tiempoTotal := time.Since(startTime)

	// Imprimir resultados y retornar tiempo total, rama ganadora y resultado ganador
	fmt.Printf("\n=== EJECUCION ESPECULATIVA COMPLETADA ===\n")
	fmt.Printf("Tiempo inicio: %s\n", startTime.Format("15:04:05.000"))
	fmt.Printf("Tiempo fin: %s\n", time.Now().Format("15:04:05.000"))
	fmt.Printf("Rama ganadora: %s\n", ramaGanadora)
	fmt.Printf("Resultado utilizado: %v\n", resultadoGanador.valor)
	fmt.Printf("Resultado descartado: %v\n", resultadoPerdedor.valor)
	fmt.Printf("Tiempo total: %v\n", tiempoTotal)
	fmt.Printf("=========================================\n\n")

	return tiempoTotal, ramaGanadora, resultadoGanador.valor
}
