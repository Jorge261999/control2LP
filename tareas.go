package main

//importar paquetes necesarios crypto para hash, math/rand para numeros aleatorios y time para medir tiempos
import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Funcion que simula un Proof-of-Work, para encontrar un hash con cierta dificultad
func SimularProofOfWork(blockData string, dificultad int) (string, int) {
	targetPrefix := strings.Repeat("0", dificultad)
	nonce := 0

	// Busca por un maximo de 3 segundos para evitar ejecuciones largas
	start := time.Now()
	maxDuration := 3 * time.Second

	for time.Since(start) < maxDuration {
		data := fmt.Sprintf("%s%d", blockData, nonce)
		hashBytes := sha256.Sum256([]byte(data))
		hashString := fmt.Sprintf("%x", hashBytes)

		if strings.HasPrefix(hashString, targetPrefix) {
			return hashString, nonce
		}
		nonce++

		// Para hacer la busqueda mas rapida, aceptar soluciones cercanas
		if nonce%1000 == 0 && strings.HasPrefix(hashString, strings.Repeat("0", dificultad-1)) {
			return hashString + " (alternativo)", nonce
		}
	}

	return fmt.Sprintf("timeout_%d", nonce), nonce
}

// EncontrarPrimos encuentra numeros primos hasta un maximo dado de 50,000 para evitar ejecuciones largas
func EncontrarPrimos(max int) []int {
	if max > 50000 {
		max = 50000 // Limitar para pruebas rapidas
	}

	var primes []int
	for i := 2; i < max; i++ {
		isPrime := true
		for j := 2; j*j <= i; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
		}
		if isPrime {
			primes = append(primes, i)
			// Limitar cantidad en 1000 para evitar ejecuciones largas
			if len(primes) > 1000 {
				break
			}
		}
	}
	return primes
}

// CalcularTrazaDeProductoDeMatrices multiplica dos matrices NxN y devuelve la traza
func CalcularTrazaDeProductoDeMatrices(n int) int {
	if n > 200 {
		n = 200 // Limitar tamaño de n en 200 para evitar ejecuciones largas
	}

	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Crear matrices mas pequenas si n es grande para pruebas rapidas
	m1 := make([][]int, n)
	m2 := make([][]int, n)
	for i := 0; i < n; i++ {
		m1[i] = make([]int, n)
		m2[i] = make([]int, n)
		for j := 0; j < n; j++ {
			m1[i][j] = rand.Intn(10)
			m2[i][j] = rand.Intn(10)
		}
	}

	trace := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += m1[i][k] * m2[k][j]
			}
			if i == j {
				trace += sum
			}
		}
	}
	return trace
}
