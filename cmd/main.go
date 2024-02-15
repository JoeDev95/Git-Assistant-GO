// Package main é o ponto de entrada para o programa.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// runGitCommand é uma função utilitária para executar comandos Git.
func runGitCommand(args ...string) error {
	cmd := exec.Command("git", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// rootCmd é o comando raiz para a CLI.
var rootCmd = &cobra.Command{
	Use:   "git-assistant",
	Short: "Um Assistente de Linha de Comando para o Git",
	Long:  "Um assistente de linha de comando para comandos avançados do Git.",
	Run: func(cmd *cobra.Command, args []string) {
		// Exibe mensagem de ajuda se nenhum comando for fornecido
		cmd.Help()
	},
}

// configureCmd é um comando para configurar o nome e email do usuário.
var configureCmd = &cobra.Command{
	Use:   "configure",
	Short: "Configurar nome e email do usuário",
	RunE: func(cmd *cobra.Command, args []string) error {
		name, _ := cmd.Flags().GetString("name")
		email, _ := cmd.Flags().GetString("email")

		if name == "" || email == "" {
			return fmt.Errorf("o nome e o email do usuário são obrigatórios")
		}

		if err := runGitCommand("config", "--global", "user.name", name); err != nil {
			return err
		}

		return runGitCommand("config", "--global", "user.email", email)
	},
}

// listBranchCmd é um comando para listar todos os branches no repositório.
var listBranchCmd = &cobra.Command{
	Use:   "list-branches",
	Short: "Listar todos os branches no repositório",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGitCommand("branch")
	},
}

// statusCmd é um comando para exibir o status do repositório.
var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Exibir o status do repositório",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runGitCommand("status")
	},
}

// addRemoteBranchCmd é um comando para criar um novo branch remoto.
var addRemoteBranchCmd = &cobra.Command{
	Use:   "add-remote-branch",
	Short: "Criar um novo branch remoto",
	RunE: func(cmd *cobra.Command, args []string) error {
		remoteName, _ := cmd.Flags().GetString("remote")
		branchName, _ := cmd.Flags().GetString("branch")

		if remoteName == "" || branchName == "" {
			return fmt.Errorf("o nome do repositório remoto e do branch são obrigatórios")
		}

		return runGitCommand("push", "--set-upstream", remoteName, branchName)
	},
}

// rebaseCmd é um comando para reorganizar commits com rebase.
var rebaseCmd = &cobra.Command{
	Use:   "rebase",
	Short: "Reorganizar commits com rebase",
	RunE: func(cmd *cobra.Command, args []string) error {
		branchName, _ := cmd.Flags().GetString("branch")

		if branchName == "" {
			return fmt.Errorf("o nome do branch é obrigatório")
		}

		return runGitCommand("rebase", branchName)
	},
}

// init é um comando para inicializar um novo repositório Git.
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Inicializar um novo repositório Git",
	RunE: func(cmd *cobra.Command, args []string) error {
		gitCmd := exec.Command("git", "init")
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		return gitCmd.Run()
	},
}

// branchCmd é um comando para criar um novo branch.
var branchCmd = &cobra.Command{
	Use:   "branch",
	Short: "Criar um novo branch",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return fmt.Errorf("por favor, forneça um nome para o branch")
		}

		branchName := args[0]
		gitCmd := exec.Command("git", "branch", branchName)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		return gitCmd.Run()
	},
}

// commitCmd é um comando para fazer commit de alterações.
var commitCmd = &cobra.Command{
	Use:   "commit",
	Short: "Commit de alterações",
	RunE: func(cmd *cobra.Command, args []string) error {
		message := ""
		if len(args) > 0 {
			message = args[0]
		}

		gitCmd := exec.Command("git", "commit", "-m", message)
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		return gitCmd.Run()
	},
}

// pushCmd é um comando para enviar alterações para um repositório remoto.
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "Enviar alterações para um repositório remoto",
	RunE: func(cmd *cobra.Command, args []string) error {
		gitCmd := exec.Command("git", "push")
		gitCmd.Stdout = os.Stdout
		gitCmd.Stderr = os.Stderr
		return gitCmd.Run()
	},
}

// main é a função principal que configura a CLI e executa comandos.
func main() {
	// Configurar Viper para persistência de dados de configuração
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	// Carregar configurações se existirem
	if err := viper.ReadInConfig(); err != nil {
		log.Println("Configurações não encontradas. Será criado um novo arquivo de configuração.")
	}

	// Executar o comando raiz
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	// Salvar configurações após a execução bem-sucedida
	if err := viper.WriteConfig(); err != nil {
		log.Println("Erro ao salvar as configurações:", err)
	}
}
