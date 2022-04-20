git clone git@github.com:nbltrust/hashkey-custody-sdk-go.git  

git checkout business  

go mod tidy  

prepare pri_hashkey-hub.pem and pub_xpert_238.pem  

go run cmd/ctl/main.go hashkey-hub pri_hashkey-hub.pem BusinessAssetsGet -a "https://develop-saas.nbltrust.com/saas-business" -p pub_xpert_238.pem  

`
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

go run cmd/ctl/main.go hashkey-hub pri_hashkey-hub.pem BusinessClientGet 235 -a "https://develop-saas.nbltrust.com/saas-business" -p pub_xpert_238.pem

`
{
  "email": "cob63176@tuofs.com",
  "id": 235,
  "kycLevel": 1,
  "name": "Christina",
  "phone": "+86-13817572905"
}
`
