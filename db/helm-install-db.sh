#!/bin/bash

kubectl apply -f ./serviceaccount.yml

MYSQL_ROOT_PASSWORD=""
if helm status pheed-mysql; then
    MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace pheed-db pheed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
fi

helm init --client-only
helm repo update
helm upgrade pheed-mysql stable/mysql --install  \
  --namespace pheed-db \
  --set persistence.storageClass=standard,mysqlDatabase=pheed,service.type=NodePort,mysqlRootPassword="$MYSQL_ROOT_PASSWORD"
