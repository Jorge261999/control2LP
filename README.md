# Control 2: Ejecución Especulativa en Go

## Descripción del Proyecto
Implementación del patrón de **ejecución especulativa** utilizando Goroutines y Channels en Go. El sistema ejecuta múltiples tareas computacionalmente intensivas en paralelo mientras evalúa una condición, seleccionando posteriormente el resultado correcto y descartando el innecesario.

## Objetivo
Demostrar las ventajas de la ejecución especulativa frente al enfoque secuencial tradicional mediante la comparación de tiempos de ejecución y cálculo del speedup.

## Arquitectura del Proyecto
### Estructura de Archivos
control2/
├── main.go # Punto de entrada y coordinador
├── especulativo.go # Implementación concurrente
├── secuencial.go # Implementación secuencial
├── tareas.go # Funciones de cómputo intensivo
├── go.mod # Módulo de Go
└── resultados/
└── metricas.txt # Reportes de métricas


### 🔧 Funcionalidades por Archivo

#### **main.go** - Director de Orquesta
- Manejo de parámetros por línea de comandos
- Coordinación de ejecuciones
- Análisis comparativo de rendimiento
- Generación de reportes

#### **especulativo.go** - Cerebro Concurrente
- Implementación con goroutines y channels
- Ejecución paralela de ramas especulativas
- Selección y descarte inteligente de resultados
- Patrón de ejecución especulativa

#### **secuencial.go** - Contraste Tradicional
- Implementación secuencial de referencia
- Ejecución lineal sin concurrencia
- Base para comparación de rendimiento

#### **tareas.go** - Motor de Cómputo
- `SimularProofOfWork()`: Minado blockchain simulado
- `EncontrarPrimos()`: Búsqueda de números primos
- `CalcularTrazaDeProductoDeMatrices()`: Evaluación de condición

## ⚙️ Instalación y Configuración

### Prerrequisitos
- Go 1.16 o superior
- Git para control de versiones

### Instalación
```bash
# Clonar repositorio
git clone <url-del-repositorio>
cd control2

# Inicializar módulo Go
go mod init control2

# Compilar proyecto
go build -o control2.exe