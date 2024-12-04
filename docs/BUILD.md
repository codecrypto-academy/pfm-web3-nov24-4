## Compilar programa
Para compilar el API se deben cumplir los siguientes requisitos:
- Tener Golang 1.23 [instalado](https://go.dev/doc/install)
- Ejecutar ```go mod tidy``` en la raiz del proyecto para descargar las dependencias
- (Opcional) Se proporciona el fichero .air.toml para ejecutar el programa con live reload utilizando [air](https://github.com/air-verse/air)
- Ejecutar ```go build -o pfm-web3-nov24-4 main.go```
- Definir la variable global **KUBECONFIG** en caso de utilizar un fichero de configuración de kubernetes diferente de **config**:
    ```bash
    export KUBECONFIG=$HOME/.kube/config-hlf
    ```
- Ejecutar el programa: ```./pfm-web3-nov24-4```