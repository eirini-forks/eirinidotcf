# EIRINI-DOT-CF - The Eirini Web Page 

This repository contains the resources for `https://eirini.cf`. 

## Deploy the database

1. Navigate to the `db/` directory
2. Run `./helm-install-db.sh`
3. Export the `MYSQL_ROOT_PASSWORD` to an environment variable:
```command
export MYSQL_ROOT_PASSWORD=$(kubectl get secret --namespace pheed-db pheed-mysql -o jsonpath="{.data.mysql-root-password}" | base64 --decode; echo)
```

4. Get the node port of the database and export it as `$DB_PORT`:
```command
export DB_PORT=$(kubectl get service -npheed-db pheed-mysql -ojsonpath='{.spec.ports[?(.name=="mysql")].nodePort}')
```

5. Get a Node IP and export it as `$DB_IP`

## API

1. Navigate to the `api/` directory
2. Run `cf push`
3. Wait for the app to get deployed (it will crash, that's fine)
4. Create the Custom Service for the Database

```command
cf cups pheed-db -p "$(./db-cups.json.sh)"
```

5. Bind the service with the API CF app:

```command
cf bind-service eirinidotcf-api pheed-db
```

6. Restage the API:

```command
cf restage eirinidotcf-api
```

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

1. Set the `api_url` in `web/src/config.json` to `http://eirinidotcf.<your-cf-domain>`
2. Run `yarn install` inside the `/web`
3. Build the static content using `yarn`:
```command
yarn install
yarn run build
```
4. CF push the web frontend with `$ cf push`

The frontend app should be avialable now at `http://web.<your-cf-domain>`

If you have a domain registered to your cluster node IP and you want to use it for this app, run the following commands:
```command
cf create-domain <org> <domain>
cf map-route web <domain>
```
