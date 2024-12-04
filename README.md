## Objetivo

El objetivo es desplegar una red de blockchain utilizando la tecnología Hyperledger Fabric 2.5 en un entorno Kubernetes (KinD). Para ello se dispone de esta API desarrollada en Go 1.23 utilizando el framework [Gin](https://gin-gonic.com/)

## Funcionalidades
En esta versión inicial del API tenemos disponibles las siguientes funcionalidades:

- Desplegar organizaciones (MSP) tipo Orderer
  - Desplegar entidades certificadoras (CA)
  - Desplegar nodos (Orderer)
  - Eliminar organización
- Desplegar organizaciones (MSP) tipo Peer
  - Desplegar entidades certificadoras (CA)
  - Desplegar nodos (Peer)
  - Eliminar organización

## TODO
En posteriores versiones se añadirá las siguientes funcionalidades:

- Preparar cadena de conexión para el orderer
- Crear canales (channel)
- Unir peer a un canal
- Preparar cadena de conexión para el peer
- Instalar chaincode
- Aprobar chaincode

## Intrucciones
- Revisa los requisitos para desplegar un cluster de Kubernetes con KinD, Istio y una configuración mínima para el DNS del cluster [Requisitos previos](docs/REQUISITES.md)
- Aquí encontrarás algunos ejemplo con el comando Curl para hacer peticiones al API [Instrucciones de uso del API](docs/USAGE.md)
- Para compilar el programa seguir las [instrucciones](docs/BUILD.md) 
