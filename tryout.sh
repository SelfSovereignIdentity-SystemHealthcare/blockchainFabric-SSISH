#!/usr/bin/env bash

if [ $# -lt 1 ] ; then
  HOST="localhost"
else
  HOST=$1
fi

CCAPI_ORG1_PORT=80
CCAPI_ORG2_PORT=980
CCAPI_ORG3_PORT=1080

if [[ $(docker container ls -f name=ccapi.org --format '{{.Names}}') == "ccapi.org.example.com" ]]
then
  CCAPI_ORG2_PORT=80
  CCAPI_ORG3_PORT=80
fi

printf "Sending requests to localhost"

# Ignore commands that don't take part on the main flow
if false
then

printf '\n\nGet Header\n';
curl -k \
  "http://${HOST}:${CCAPI_ORG1_PORT}/api/gateway/query/getHeader" \
  -H 'cache-control: no-cache'

printf '\n\nGet Transactions\n';
curl -k \
  "http://${HOST}:${CCAPI_ORG1_PORT}/api/gateway/query/getTx" \
  -H 'cache-control: no-cache'

printf '\n\nGet CreateAsset definition\n';
curl -k -X POST \
  "http://${HOST}:${CCAPI_ORG1_PORT}/api/gateway/query/getTx" \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
        "txName": "createAsset"
      }'

printf '\n\nGet Asset Types\n';
curl -k \
  "http://${HOST}:${CCAPI_ORG1_PORT}/api/gateway/query/getSchema" \
  -H 'cache-control: no-cache'

printf '\n\nGet noticia schema\n';
curl -k -X POST \
  "http://${HOST}:${CCAPI_ORG1_PORT}/api/gateway/query/getSchema" \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{
        "assetType": "noticia"
      }'

fi;

printf '\n\nCreate Fact-checker 1\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/criarFactChecker" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache' \
     -d '{
            "sub":"1",
            "cnpj":"0001",
            "nome":"Fact-checker 1",
            "username":"factchecker1",
            "password":"123",
            "userType": "factChecker"
         }'

printf '\n\nCreate Fact-checker 2\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/criarFactChecker" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache' \
     -d '{
            "sub":"2",
            "cnpj":"0002",
            "nome":"Fact-checker 2",
            "username":"factchecker2",
            "password":"123",
            "userType": "factChecker"
         }'

printf '\n\nCreate Fact-checker 3\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/criarFactChecker" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache' \
     -d '{
            "sub":"3",
            "cnpj":"0003",
            "nome":"Fact-checker 3",
            "username":"factchecker3",
            "password":"123",
            "userType": "factChecker"
         }'

printf '\n\nCreate notícia 1\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/criarNoticia" \
  -H "Content-Type: multipart/form-data" \
  -F "titulo=Título de notícia 1" \
  -F "origem=Jornal 1" \
  -F "tipo=0" \
  -F "dataCriacao=2024-01-01T03:00:00.000Z" \
  -F "politicaDeRotulo=1" \
  -F "file=@./test-data/noticia-1.pdf"

printf '\n\nCreate notícia 2\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/criarNoticia" \
  -H "Content-Type: multipart/form-data" \
  -F "titulo=Título de notícia 2" \
  -F "origem=Jornal 2" \
  -F "tipo=0" \
  -F "dataCriacao=2024-01-01T03:00:00.000Z" \
  -F "politicaDeRotulo=1" \
  -F "file=@./test-data/noticia-2.pdf"

printf '\n\nVotar autenticidade notícia 1 com fact-checker 1\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:84cf1d13-86d8-5af0-8b8e-607ea95fb42b"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:996ad860-2a9a-504f-8861-aeafd0b2ae29"
           },
           "decisao": 1,
           "justificativa": "J1"
         }'

printf '\n\nVotar autenticidade notícia 2 com fact-checker 1\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:8eee75ba-83dc-58fb-8354-66c066be569c"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:996ad860-2a9a-504f-8861-aeafd0b2ae29"
           },
           "decisao": 0,
           "justificativa": "J1"
         }'

printf '\n\nVotar autenticidade notícia 1 com fact-checker 2\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:84cf1d13-86d8-5af0-8b8e-607ea95fb42b"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:59e06cf8-f390-5093-af2e-3685be593a25"
           },
           "decisao": 2,
           "justificativa": "J2"
         }'

printf '\n\nVotar autenticidade notícia 2 com fact-checker 2\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:8eee75ba-83dc-58fb-8354-66c066be569c"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:59e06cf8-f390-5093-af2e-3685be593a25"
           },
           "decisao": 0,
           "justificativa": "J2"
         }'         

printf '\n\nVotar autenticidade notícia 1 com fact-checker 3\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:84cf1d13-86d8-5af0-8b8e-607ea95fb42b"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:391ada15-580c-5baa-b16f-eeb35d9b1122"
           },
           "decisao": 0,
           "justificativa": "J3"
         }'

printf '\n\nVotar autenticidade notícia 2 com fact-checker 3\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/invoke/votarAutenticidadeNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "noticia": {
             "@assetType": "noticia",
             "@key": "noticia:8eee75ba-83dc-58fb-8354-66c066be569c"
           },
           "factChecker": {
             "@assetType": "factChecker",
             "@key": "factChecker:391ada15-580c-5baa-b16f-eeb35d9b1122"
           },
           "decisao": 1,
           "justificativa": "J3"
         }'

printf '\n\nValidar notícia 1\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/validarNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "titulo": "Título de notícia 1",
            "origem": "Jornal 1"
         }'         

printf '\n\nValidar notícia 2\n'
curl -X POST "http://${HOST}:${CCAPI_ORG1_PORT}/api/validarNoticia" \
     -H 'Content-Type: application/json' \
     -H 'cache-control: no-cache'  \
     -d '{
           "titulo": "Título de notícia 2",
            "origem": "Jornal 2"
         }'