#Build stage

FROM golang:latest AS BuildStage 
WORKDIR /usr/src/app
COPY /app . 
#RUN CGO_ENABLED=0 GOOS=linux go build -o hello hello.go
RUN go build -o hello hello.go


#Deploy stage 

# O segredo para reduzir o máximo o tamanho da imagem final é usar como imagem base a strach
# strach é uma imagem que contem o minimo necessário    
FROM scratch
WORKDIR /app 
COPY --from=BuildStage /usr/src/app /app 


CMD ["./hello"] 
