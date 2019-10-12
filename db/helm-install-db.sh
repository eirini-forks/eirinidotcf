#!/bin/bash

helm upgrade pheed-mysql stable/mysql --install  \
  --namespace pheed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=pheed,service.type=NodePort
