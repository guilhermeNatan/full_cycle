# Desafio 2

**Comandos:** 
Para execua a aplicação execute 
`docker-compose up -d`

A aplicação irá subir na porta 80 podendo ser acessada por 
http://localhost


**Papel de cada imagem** 


Mysql
 * Banco de dados mysql , dessa forma os dados inseridos pela aplicação podem ser persistidos 
 * Dentro do diretorio` mysql/initdb` contêm o script que será executado quando o banco inicializar 

Backend
* Servidor node da aplicação,  dentro do arquivo index.js é definido as rotas da api. Foram definidas duas rotas principais api/inserir e api/listar que inserem e lista os dados persistidos no banco de dados 
Foi feito o mapeamento de volumes da pasta node , dessa forma é possível criar arquivos 'compartilhados' entre a máquina host e o container.  

Nginx:

* Proxy reverso do nginx , responsável por servir o frontend da aplicação e fazer o proxy com a imagem do backend. Foram definidas duas configurações de location  dentro do arquivo nginx.conf  barra (/) onde será retornado o frontend da aplicação e /api  que é um proxi para a backend da aplicação que esta definido na imagem backend, ou seja, todo chamada para /api será direcionada para o backend 

**Dependencias entre os containes**

Para garantir a ordem de inicialização dos containers foi utilizado o **dockerize** 
os containers devem seguir a seguinte ordem ao ser inicializados 
1. db
1. backend
1. nginx 

**Otimizações** 
* Definir previamente as dependencias dentro do node/package.json, dessa forma apenas um comando RUN npm install dentro da imagem já instala todas as dependencias 
* Para imagem do node usei a versão 18-bullseye-slim  que  é uma versão minimalista do node dessa forma a imagem final esta com menos de 250MB 
* Scripts de inicialização do banco foram colocados dentro do diretorio `mysql/initdb `

**Referencias:** 

Problema com a imagem do mysql que não inicializava : 
https://stackoverflow.com/questions/65245078/docker-image-build-context-canceled-error-on-windows-10 

Para fazer o compartilhamento de volume entre máquina host e container foi necessário explicitar o caminho da pasta node_modules la lista de volumes 
https://stackoverflow.com/questions/30043872/docker-compose-node-modules-not-present-in-a-volume-after-npm-install-succeeds


