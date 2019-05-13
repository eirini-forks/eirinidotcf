#!/bin/bash

helm install stable/mysql \
  --name pheed-mysql \
  --namespace pheed-db \
  --set persistence.storageClass=hostpath,mysqlDatabase=pheed
