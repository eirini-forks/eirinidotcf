#!/bin/bash

kubectl apply -f ./serviceaccount.yml

useExistingSecret=false
if helm status pheed-mysql; then
    useExistingSecret=true
fi

helm init --client-only
helm repo update
helm upgrade pheed-mysql stable/mysql --install  \
  --namespace pheed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=pheed,service.type=NodePort,existingSecret="$useExistingSecret"
