#instalacion de pmgo 
cd /home/ubuntu

git clone https://github.com/struCoder/pmgo.git

cd pmgo
go build pmgo.go
mv pmgo /usr/local/bin

# proceso de crear la carpeta que ejecutara el pmgo
cd /
mkdir src && src
mkdir consumer

# aqui va el repositorio donde se tiene el codigo de golang

git clone https://github.com/AxlMax/PlcGoConsumer.git

pmgo start consumer/PlcGoConsumer consumer
