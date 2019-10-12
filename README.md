# EIRINI-DOT-CF - The Eirini Web Page 

This repository contains the resources for `https://eirini.cf`. 

## API

1. Navigate to the `api/` directory
1. Run `cf push`
1. Wait for the app to get deployed
1. Navigate back to the root directory of `eirinidotcf`

## Deploy the database

1. Navigate to the `db/` directory
2. Run `./helm-install-db.sh`
3. Export the `MYSQL_ROOT_PASSWORD` to an environment variable:
```command
export MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace pheed-db pheed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
```
4. Get the node port of the database:
```command
kubectl get svc feed-db -o yaml | grep nodePort
```
5. Navigate back to the root directory of `feedelphia`

## Bind the api to the MySql Database

1. Navigate back to the `api` directory
2. Export env variables `$DB_IP`(the node IP) and `$DB_PORT` (the node port of the database)
3. Create a User-Provided-Service instance:
```command
cf cups feed-db -p "$(./db-cups.json.sh)"
```
4. Bind the `api` app to the `feed-db` service instance:
```command
cf bind-service pheed-api feed-db
```
5. Restage the `api` app: 
```
cf restage pheed-api
```

## Frontend

1. Set the `api_url` in `web/src/config.json` to `http://pheed-api.<your-cf-domain>`
2. Build the static content using `yarn`:
```command
npm install
yarn run build
cf push
```

The frontend app should be avialable now at `http://web.<your-cf-domain>`

If you have a domain registered to your cluster node IP and you want to use it for this app, run the following commands:
```command
cf create-domain <org> <domain>
cf map-route web <domain>
```
