```
1.    Crie uma API usando Gin rodando na porta :8080
2.    Crie um endpoint GET /ping que retorne: { "message": "pong" }
3.    Crie um endpoint GET /users/:id que retorne o id recebido na URL
4.    Crie um endpoint GET /search?q=valor que retorne o valor da query q
5.    Crie um endpoint POST /users que receba um JSON com name e email e retorne os mesmos dados
6.    Adicione validação no endpoint acima:
    •    name obrigatório
    •    email obrigatório
```

1. Não pode criar duas contas com o mesmo email (crie um setup para que o email seja definido como index e unique
2. Get users para pegar todos os usuarios 
3. Na criação e edição verificar se é um email valido
4. Instalar o air para não ter que ficar reiniciando o servidor
5. Instalar o golint e lintar o projeto inteiro (copiar o lint yaml dos nossos projetos)
6. Ao deletar, ter que digitar a senha para verificar se o usuario pode ser deletado