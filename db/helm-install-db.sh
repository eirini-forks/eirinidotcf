#!/bin/bash

kubectl apply -f ./serviceaccount.yml

helm init --client-only
helm repo update
helm upgrade pheed-mysql stable/mysql --install  \
  --namespace pheed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=pheed,service.type=NodePort
