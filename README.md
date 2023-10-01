# MartinStanchevREPartnersChallenge


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

```
APP_PORT=":8080" PACK_SIZES="250,500,1000,2000,5000" ./packDistributor
```

### Tests

To run the unit tests use:
```
go test ./...
```