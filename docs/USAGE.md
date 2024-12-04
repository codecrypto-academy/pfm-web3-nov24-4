El procedimiento de ejemplo para desplegar los nodos que participarán de la red de Hyperledger Fabric 2.5 será el siguiente:
- Org0MSP (Orderer)
    - Desplegar CA para la organización Org0MSP de tipo orderer
    - Registrar el usuario **orderer** de tipo orderer
    - Deplegar los nodos Orderer (2 nodos) de la organización.
- Org1MSP (Peer)
    - Desplegar CA para la organización Org1MSP de tipo peer
    - Registrar el usuario **peer** de tipo peer
    - Deplegar los nodos Peer (2 nodos) de la organización.

## Crear CA para la Org0MSP
### ca-org0
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-ca \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "image": "hyperledger/fabric-ca",
  "version": "1.5.6",
  "capacity": "1Gi",
  "storageclass": "standard",
  "name": "ca-org0",
  "enroll_id": "enroll",
  "enroll_pw": "enrollpw",
  "hosts": [
    "ca-org0.homelab.local"
  ],
  "istio_port": 443
}'
```

## Comprobar estado del Pod ca-org0
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-ca-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ca-org0",
  "namespace": "default"
}'
```

## Registrar el usuario orderer
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/register-fabric-ca \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ca-org0",
  "user": "orderer",
  "secret": "ordererpw",
  "type": "orderer",
  "enrollId": "enroll",
  "enrollSecret": "enrollpw",
  "mspid": "Org0MSP",
  "caurl": "https://ca-org0.homelab.local"
}'
```

## Desplegar los nodos Orderer
### ord0-org0
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-orderer \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "ordererImage": "hyperledger/fabric-orderer",
  "ordererVersion": "2.4.6",
  "scName": "standard",
  "mspid": "Org0MSP",
  "enrollId": "enroll",
  "enrollPw": "enrollpw",
  "capacity": "2Gi",
  "name": "ord0-org0",
  "caName": "ca-org0.default",
  "hosts": "ord0-org0.homelab.local",
  "istioPort": 443
}'
```

#### Comprobar estado del Pod ord0-org0
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-orderer-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ord0-org0",
  "namespace": "default"
}'
```

### ord1-org0
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-orderer \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "ordererImage": "hyperledger/fabric-orderer",
  "ordererVersion": "2.4.6",
  "scName": "standard",
  "enrollId": "enroll",
  "mspid": "Org0MSP",
  "enrollPw": "enrollpw",
  "capacity": "2Gi",
  "name": "ord1-org0",
  "caName": "ca-org0.default",
  "hosts": "ord1-org0.homelab.local",
  "istioPort": 443
}'
```

#### Comprobar estado del Pod ord1-org0
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-orderer-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ord1-org0",
  "namespace": "default"
}'
```

### ord2-org0
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-orderer \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "ordererImage": "hyperledger/fabric-orderer",
  "ordererVersion": "2.4.6",
  "scName": "standard",
  "enrollId": "enroll",
  "mspid": "Org0MSP",
  "enrollPw": "enrollpw",
  "capacity": "2Gi",
  "name": "ord2-org0",
  "caName": "ca-org0.default",
  "hosts": "ord2-org0.homelab.local",
  "istioPort": 443
}'
```

#### Comprobar estado del Pod ord2-org0
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-orderer-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ord2-org0",
  "namespace": "default"
}'
```

## Borrar organización Org0MSP
```bash
curl --request DELETE \
  --url http://127.0.0.1:8080/api/v1/delete-fabric-orderer-org \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "caname": "ca-org0",
  "orderers": [
      "ord0-org0", 
      "ord1-org0",
      "ord2-org0"
  ],
  "namespace": "default"
}'
```

## Crear CA para la Org1MSP
### ca-org1
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-ca \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "image": "hyperledger/fabric-ca",
  "version": "1.5.6",
  "capacity": "1Gi",
  "storageclass": "standard",
  "name": "ca-org1",
  "enroll_id": "enroll",
  "enroll_pw": "enrollpw",
  "hosts": [
    "ca-org1.homelab.local"
  ],
  "istio_port": 443
}'
```

#### Comprobar estado del Pod ca-org1
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-ca-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ca-org1",
  "namespace": "default"
}'
```

## Registrar usuario **peer** para Org1MSP
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/register-fabric-ca \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "ca-org1",
  "user": "peer",
  "secret": "peerpw",
  "type": "peer",
  "enrollId": "enroll",
  "enrollSecret": "enrollpw",
  "mspid": "Org1MSP",
  "caurl": "https://ca-org1.homelab.local"
}'
```

## Desplegar los nodos Peer para Org1MSP
### peer0-org1
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-peer \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "statedb": "couchdb",
  "peerImage": "hyperledger/fabric-peer",
  "peerVersion": "2.4.6",
  "scName": "standard",
  "enrollId": "enroll",
  "mspid": "Org1MSP",
  "enrollPw": "enrollpw",
  "capacity": "5Gi",
  "name": "peer0-org1",
  "caName": "ca-org1.default",
  "hosts": "peer0-org1.homelab.local",
  "istioPort": 443
}'
```

#### Comprobar el estado del Pod peer0-org1
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-peer-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "peer0-org1",
  "namespace": "default"
}'
```

### peer1-org1
```bash
curl --request POST \
  --url http://127.0.0.1:8080/api/v1/create-fabric-peer \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "statedb": "couchdb",
  "peerImage": "hyperledger/fabric-peer",
  "peerVersion": "2.4.6",
  "scName": "standard",
  "enrollId": "enroll",
  "mspid": "Org1MSP",
  "enrollPw": "enrollpw",
  "capacity": "5Gi",
  "name": "peer1-org1",
  "caName": "ca-org1.default",
  "hosts": "peer1-org1.homelab.local",
  "istioPort": 443
}'
```

#### Comprobar estado del Pod peer1-org1
```bash
curl --request GET \
  --url http://127.0.0.1:8080/api/v1/check-fabric-peer-status \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "name": "peer1-org1",
  "namespace": "default"
}'
```

## Borra la organización org1MSP
```bash
curl --request DELETE \
  --url http://127.0.0.1:8080/api/v1/delete-fabric-peer-org \
  --header 'Authorization: Basic YWRtaW46cGFzc3dvcmQ=' \
  --header 'content-type: application/json' \
  --data '{
  "ca_name": "ca-org1",
  "peers": [
      "peer0-org1", 
      "peer1-org1"
  ],
  "namespace": "default"
}'
```






