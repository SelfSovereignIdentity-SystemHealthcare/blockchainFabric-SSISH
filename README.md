# Blockchain Hyperledger Fabric - Self Sovereign Identity System for Healthcare

## Estrutura de Diretórios

- `/fabric`: Rede Fabric v2.2 usada como ambiente de teste
- `/chaincode`: arquivos relacionados ao chaincode
- `/ccapi`: API REST do chaincode em projeto Golang

## Desenvolvimento

A biblioteca `cc-tools` foi testada nas redes Fabric v1.4, v2.2 e v2.4.

Dependências para chaincode e API do chaincode:

- Go 1.14 ou superior

Dependências para ambiente de teste:

- Docker 20.10.5 ou superior
- Docker Compose 1.28.5 ou superior

Instalação da API Chaincode Go:

```bash
$ cd chaincode; go mod vendor; cd ..
$ cd ccapi; go mod vendor; cd ..
```

## Implementando ambiente de teste na v2.2
Após a instalação, use o script ./startDev.sh na pasta raiz para iniciar o ambiente de desenvolvimento. Ele irá iniciar todos os componentes do projeto com 3 organizações.

Se você quiser implementar com 1 organização, execute o comando ./startDev.sh -n 1.

Para aplicar mudanças no chaincode, execute $ ./upgradeCC.sh <versão> <sequência> com uma versão superior à atual (começa com 0.1). Acrescente -n 1 ao comando se estiver executando com 1 organização.

Para aplicar mudanças na API CC, execute $ ./reloadCCAPI.sh.

## Teste automatizado e tentativa
Para testar transações após iniciar todos os componentes, execute $ ./tryout.sh.

Para testar transações usando a ferramenta godog, execute $ ./godog.sh.

## Mais
Você pode entrar em contato com os desenvolvedores da GoLedger e mantenedores do cc-tools em nosso Discord - Junte-se a nós!

Mais documentação e detalhes sobre cc-tools podem ser encontrados em https://goledger-cc-tools.readthedocs.io/en/latest/

Para implantação em produção, considere usar o GoFabric - https://gofabric.io

<hr/>
<hr/>
<hr/>
<hr/>
<hr/>

## Assets:
**VerifiableCredential**
- issuerHash: sha256 (CK)
- receiverHash: sha256 (CK)
- credentialType: string
- credentialData: string
- status: string

**Wallet**
- holderHash: sha256 (PK)
- credentials: []->verifiableCredential
- verifications: []->credentialVerification

**DID**
- walletHash: sha256 (CK)
- userController: sha256 (CK)
- publicKey: string
- authenticationMethods: string: string
- services: string
- timestamp: datetime
- status: string

**CredentialVerification**
- verifierHash: sha256 (PK)
- credentialHash: sha256
- timestamp: datetime
- result: string

**SSISHEvent**
- walletHash: sha256 (CK)
- eventType: string (CK)
- timestamp: datetime
- eventDetails: string

**Exam**:
- patientWalletHolderHash: sha256 (CK)
- doctorWalletHolderHash: sha256 (CK)
- timestamp: datetime
- name: string
- urlExamDocument: string
- diagnosisHash: sha256
- treatmentHash: sha256

**Diagnosis**:
- exam: ->exam (PK)
- timestamp: datetime
- name: string
- obs: string

**Treatment**:
- diagnosis: ->diagnosis (PK)
- timestamp: datetime
- name: string
- obs: string

