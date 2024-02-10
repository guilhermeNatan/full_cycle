# Desafio 1 

Comandos para construir e executar a imagem 

`docker build -t capela/fc-golang-app-d1 .`

`docker run -it --rm --name my-running-app capela/fc-golang-app-d1`

Para otimizar o tamanho da imagem final , esta foi construida a partir da imagem
scratch , que é uma imagem minimalista  que contem apenas o binario  e o necessario para executar a apicação


**Link da imagem no docker hub** 

https://hub.docker.com/repository/docker/capela/fc-golang-app-d1/general 
