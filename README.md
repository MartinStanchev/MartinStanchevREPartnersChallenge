# MartinStanchevREPartnersChallenge

### Deployment

The app is temporarily deployed on a GCP Cloud Run free trial and can be accessed here:
`https://martinstanchevrepartnerschallenge-snkn5idqfa-ew.a.run.app`


### Build

```
git clone git@github.com:MartinStanchev/MartinStanchevREPartnersChallenge.git
cd MartinStanchevREPartnersChallenge
go build -o packDistributor
```

### Running

The application takes two environment variables:
`APP_PORT` - specifies the port the app will be hosted on. Default ":8080"
`PACK_SIZES` - a comma-separated list of integers that represent the initial sizes that the app will run on. Default is "250,500,1000,2000,5000"

To run locally:

```sh
APP_PORT=":8080" PACK_SIZES="250,500,1000,2000,5000" ./packDistributor
```

### Tests

To run the unit tests use:
```sh
go test ./...
```

### Endpoints and Querying

The app exposes two endpoints 

POST `/order` - Accepts a json with an integer "order" field, based on which it calculates and returns the pack distribution

examples:
```sh
curl -X POST --data '{"order": 12001}' https://martinstanchevrepartnerschallenge-snkn5idqfa-ew.a.run.app/order`
```

```sh
curl -X POST --data '{"order": 1}' localhost:8080/order
```

PATCH `/update/sizes` - Updates the configured pack sizes dynamically. Accepts a json array "packSize" that contains integers

examples:
```sh
curl -X PATCH --data '{"packSizes": [250, 1000]}' https://martinstanchevrepartnerschallenge-snkn5idqfa-ew.a.run.app/update/sizes
```

```sh
curl -X PATCH --data '{"packSizes": [150, 300, 600]}' localhost:8080/update/sizes
```

