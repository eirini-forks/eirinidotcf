#!/bin/bash

kubectl apply -f ./serviceaccount.yml

useExistingSecret="nil"
if helm status pheed-mysql; then
    useExistingSecret="pheed-mysql"
fi

helm init --client-only
helm repo update
helm upgrade pheed-mysql stable/mysql --install  \
  --namespace pheed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=pheed,service.type=NodePort,existingSecret="$useExistingSecret"
