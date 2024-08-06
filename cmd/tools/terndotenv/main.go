package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/joho/godotenv"
)

func main() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		panic(fmt.Sprintf("Error loading .env file: %v", err))
	}

	// Definir o comando e seus argumentos
	cmd := exec.Command(
		"tern",
		"migrate",
		"--migrations",
		"./internal/store/pgstore/migrations",
		"--config",
		"./internal/store/pgstore/migrations/tern.conf",
	)

	// Passar o ambiente atual para o comando, incluindo as variáveis carregadas
	cmd.Env = os.Environ()

	// Capturar a saída padrão e a saída de erro
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Executar o comando
	if err := cmd.Run(); err != nil {
		panic(fmt.Sprintf("Command execution failed: %v", err))
	}
}
