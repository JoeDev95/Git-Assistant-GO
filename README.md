# Git Assistant

Git Assistant é uma ferramenta de linha de comando escrita em Go para simplificar operações comuns do Git.

## Instalação

Certifique-se de ter o Go instalado em sua máquina. Você pode clonar este repositório e compilar o código usando o seguinte comando:

```bash
go build -o git-assistant main.go


./git-assistant init
./git-assistant list-branches
./git-assistant status
./git-assistant add-remote-branch --remote origin --branch NovoBranch
./git-assistant rebase --branch BranchDestino
./git-assistant commit "Mensagem do Commit"
./git-assistant push
