# WDB
Módulo de acesso ao banco de dados.
<br/>
## Iniciando a conexão com o banco de dados
    
```go
wdb.InitDB("postgresql", wdb.Config{
    DBHost:     "localhost",
    DBPort:     "5432",
    DBUser:     "postgres",
    DBPassword: "123456",
    DBName:     "welber",
})
```
<br/>

## Resgatando a conexão com o Banco de Dados

Ao criar uma conexão, ela é armazenada em um array de conexões. A posição 0 é a conexão padrão. Pois ela é a primeira que foi chamada pela wdb.InitDB().    
```go
db := wdb.GetConn(0)
```
<br/>

## Migrate
Faz o migrate da estrutura do modelo.
```go
wdb.Migrate(wdb.GetConn(0), &EstruturaModel{})
```
