# Comandos utéis 

inicia projeto 
`go run github.com/99designs/gqlgen init`  

cria a estrutura do projeto 
`go run github.com/99designs/gqlgen generate`

inicia servidor 
`go run cmd/server/server.go`

baixa as dependencias 
`go mod tidy`

Problema CGO_ENABLED ao usar sqlite:  

`go env CGO_ENABLED`

`go env -w CGO_ENABLED=<0 or1>`

# Exemplo de client usando graphQL

```

# Para criar/atualizar objetos usamos mutation

mutation creteCategory {
createCategory(input: { name : "Tecnologia", description: "foo" }) {
id
name
description

}
}

mutation creteCourse {
createCourse(input: { name : "Course a ", description: "boo" , category: "2f3b494c-8938-4828-90eb-f82fc1ee92f0"}) {
id
name
description

}
}

query queryCategories {
categories{
id
name

}
}

# Para obter dados de relacionamentos
# 1 - Dentro de graph/model definimos modelos separados para categoria e curso

# 2 - em seguida executei o comento para gerar os Resolvers novamente (go run github.com/99designs/gqlgen generate)

# 3 - Dentro de cada modelo interno (internal/database) foi definido metodos responsáveis por
buscar os dados no bd  (FindByCourseID/FindByCategoryID)

# 4 - No arquivo schema.resolvers.go  foi implementado os resolvers que basicamente são metodos que serão chamados caso o usuário solicite alguma informação de relacionamento , como courses na query a seguir

query queryCategoriesWithCourses {
categories{
id
name
courses {
id
name
}

}
}

query queryCourses {
courses{
id
name

}
}

query queryCoursesWithCategory {
courses{
id
name
category {
id
name
}
}
}
```

# Referencias 

https://gqlgen.com/ 

Problema CGO_ENABLED ao usar sqlite 
Binary was compiled with 'CGO_ENABLED=0'
https://forum.golangbridge.org/t/getting-errors-when-running-go-project-that-uses-github-com-mattn-go-sqlite3-library/31800
