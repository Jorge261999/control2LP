package main

//importar paquetes necesarios
import (
	"flag"
	"fmt"
	"os"
	"time"
)

// definir flags y funciones principales
func main() {
	n := flag.Int("n", 80, "Dimension de las matrices para calcular traza")
	umbral := flag.Int("umbral", 100000, "Umbral para decidir rama ganadora")
	archivo := flag.String("archivo", "resultados/metricas.txt", "Archivo para guardar metricas")
	modo := flag.String("modo", "especulativo", "Modo de ejecucion: especulativo|secuencial|comparar")
	iteraciones := flag.Int("iter", 30, "Iteraciones para modo comparar")

	flag.Parse()
	// Crear directorio para resultados si no existe
	os.MkdirAll("resultados", 0755)

	fmt.Printf("Ejecutando en modo: %s\n", *modo)
	fmt.Printf("Parametros: n=%d, umbral=%d\n\n", *n, *umbral)

	// Seleccionar modo de ejecucion
	switch *modo {
	case "especulativo":
		EjecutarEspeculativo(*n, *umbral, *archivo)
	case "secuencial":
		EjecutarSecuencial(*n, *umbral, *archivo)
	case "comparar":
		compararRendimiento(*n, *umbral, *archivo, *iteraciones)
	default:
		fmt.Println("Modo no reconocido")
	}
}

// Funcion para comparar rendimiento entre ambos metodos especulativo y secuencial
func compararRendimiento(n, umbral int, archivo string, iteraciones int) {
	fmt.Printf("=== COMPARACION DE RENDIMIENTO (%d iteraciones) ===\n\n", iteraciones)

	var tiemposEspeculativo, tiemposSecuencial []time.Duration

	for i := 0; i < iteraciones; i++ {
		fmt.Printf("--- Iteracion %d/%d ---\n", i+1, iteraciones)

		fmt.Println("-> Ejecucion ESPECULATIVA:")
		tiempoEsp, ramaEsp, _ := EjecutarEspeculativo(n, umbral, archivo)
		tiemposEspeculativo = append(tiemposEspeculativo, tiempoEsp)

		fmt.Println("-> Ejecucion SECUENCIAL:")
		tiempoSec, ramaSec, _ := EjecutarSecuencial(n, umbral, archivo)
		tiemposSecuencial = append(tiemposSecuencial, tiempoSec)

		fmt.Printf("Resumen iteracion %d: Especulativo=%v (rama %s), Secuencial=%v (rama %s)\n\n",
			i+1, tiempoEsp, ramaEsp, tiempoSec, ramaSec)
	}

	// Calcular promedios entre ambas ejecuciones, speedup y guardar metricas
	var promEsp, promSec time.Duration
	for i := 0; i < iteraciones; i++ {
		promEsp += tiemposEspeculativo[i]
		promSec += tiemposSecuencial[i]
	}
	promEsp /= time.Duration(iteraciones)
	promSec /= time.Duration(iteraciones)

	speedup := float64(promSec) / float64(promEsp)

	fmt.Println("=== RESULTADOS FINALES ===")
	fmt.Printf("Tiempo promedio ESPECULATIVO: %v\n", promEsp)
	fmt.Printf("Tiempo promedio SECUENCIAL: %v\n", promSec)
	fmt.Printf("SPEEDUP: %.2fx\n", speedup)

	guardarMetricas(archivo, promEsp, promSec, speedup, iteraciones)
}

// Funcion para guardar metricas en archivo
func guardarMetricas(archivo string, tiempoEsp, tiempoSec time.Duration, speedup float64, iteraciones int) {
	content := fmt.Sprintf("Metricas (%d iteraciones)\n", iteraciones)
	content += fmt.Sprintf("Especulativo: %v\n", tiempoEsp)
	content += fmt.Sprintf("Secuencial: %v\n", tiempoSec)
	content += fmt.Sprintf("Speedup: %.2fx\n", speedup)

	os.WriteFile(archivo, []byte(content), 0644)
	fmt.Printf("Metricas guardadas en: %s\n", archivo)
}
