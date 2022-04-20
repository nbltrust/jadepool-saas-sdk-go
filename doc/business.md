git clone git@github.com:nbltrust/hashkey-custody-sdk-go.git  

git checkout business  

go mod tidy  

prepare pri_hashkey-hub.pem and pub_xpert_238.pem  

go run cmd/ctl/main.go hashkey-hub pri_hashkey-hub.pem BusinessAssetsGet -a "https://develop-saas.nbltrust.com/business" -p pub_xpert_238.pem  

should see something like below 

`
code: 0
message: success
sign: true
data:
{
  "assets": [
    {
      "decimal": 18,
      "id": 1,
      "name": "ETH",
      "switch": true
    },
    {
      "decimal": 8,
      "id": 2,
      "name": "BTC",
      "switch": true
    }
  ]
}
`
