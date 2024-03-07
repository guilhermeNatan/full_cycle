consul agent -dev

consul members 

curl localhost:8500/v1/catalog/nodes

instalar o dig 
apk -U add bind-tools


dig @localhost -p 8600


dig @localhost -p 8600 consul01.node.consul


docker exec -it consulserver01 sh




mkdir /etc/consul.d
mkdir /var/lib/consul
consul agent -server -bootstrap-expect=3 -node=consulserver01 -bind=172.18.0.4 -data-dir=/var/lib/consul -config-dir=/etc/consul.d


consul agent -server -bootstrap-expect=<quantidade servers> -node=<nome do server> -bind=<ip do server na rede> -data-dir=<diretorio onde o consul guarda seus arquivos >-config-dir=<diretorio onde fica os arquivos de configuração do consul tem que ser criados>

conectar dois servidores 
consul join <ip maquina membro>]



docker exec -it consulclient01 sh
dentro do client
consul agent  -bind=<ip da maquina na rede> -data-dir=/var/lib/consul -config-dir=/etc/consul.d

consul join  <server ip >


consul agent  -bind=172.18.0.5 -data-dir=/var/lib/consul -config-dir=/etc/consul.d

IMPORTANTE:
Registrar o serviço e subir o serviço são coisas diferentes primeiro vc registra o serviço e depois vc
sobe o serviço na maquina
Você deve subir o agente na mesma máquina onde esta o serviço 
exemplo subir um agente que monitora o nginx na mesma maquina onde o nginx será instalado 

para ataualizar o serviços do consul 
consul reload 


dig @localhost -p 8600 nginx.service.consul

curl localhost:8500/v1/catalog/services

consul catalog nodes -service nginx 

consul catalog nodes detailed 


consul agent -bind=<ip> -data-dir<DIR>

consul agent -bind=172.18.0.2 -data-dir=/var/lib/consul -config-dir=/etc/consul.d  -retry-join=172.18.0.4  -retry-join=172.18.0.5

com retry join o client já entra no cluster 


Implementando checks 

apk add nginx:q:q:q

depois que é definido as configurações dentros dos arquivos servers/server0X/server0X.json basta executar o 
seguinte comando apos o servidor subir 

consul agent -config-dir=/etc/consul.d

para gerar chave de criptografia 
consul keygen 
rm -rf /tmp/*
