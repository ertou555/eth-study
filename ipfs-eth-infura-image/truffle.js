// truffle migrate --network ropsten
var HDWalletProvider = require("truffle-hdwallet-provider");

var infura_apikey = "5fb32fed479043a5b9d25816983b6a46";
var mnemonic = "infant shock chair century audit addict tape cup chapter empower lab trust"

module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*" // Match any network id
    },
    ropsten: {
      provider: new HDWalletProvider(mnemonic, "https://ropsten.infura.io/"+infura_apikey),
      network_id: 3
    }
  }
};
