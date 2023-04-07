# Arquitetura Hexagonal (Ports And Adapters)

- A _Arquitetura Hexagonal_, de forma genérica, consiste em isolar a aplicação e definir como acessar pontos externos e como outros pontos externos irão acessar a aplicação.

- Esta é uma aplicação de _Produtos_, em que é possível ativar, desativar, criar ou acessar um _Produto_.

- Como adaptadores responsáveis por acessar a aplicação, temos um _Command-line Interface_ (_CLI_) e um _Servidor Web_.

- E como adaptador que poder ser acessado pela aplicação está um banco de dados _Sqlite_.

- O Domínio da aplicação fica isolado dentro do diretório **application**. Já os adaptadores ficam localizados sob o diretório **adapters**.

- _ProductPersistenceInterface_ é uma interface que atua como porta para as operações de leitura e escrita a serem utilizadas por um mecanismo de persistência qualquer. Essa porta pode ser utilizada tanto pelos adaptadores de _CLI_ quanto de _Servidor Web_.

- Para testar a aplicação, executar:

    - Através do adaptador de _CLI_, sob o diretório _src_:

        - Para criar um novo produto:

            - `go run main.go cli -a=create -n=Product CLI -p=25.0`

        - E para acessar o produto recém criado:

            - `go run main.go cli -a=get --id=<ID produto recém criado>`

    - Através do adaptador de _Servidor Web_, sob o diretório _src_:

        - Para subir o servidor:

            - `go run main.go http`

        - E para acessar o produto recém criado, em um novo terminal:

            - `curl http://localhost:8080/product/<ID produto recém criado>`

> N.T.: Linguagem de programação para esta aplicação: Golang.
