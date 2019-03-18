                                             # ETH+IPFS Dapp development process
1 Client browser: install metamask wallet , keep mnemonic, switch Ropsten Network;

2 Project code: install truffle, npm; truffle unbox react;

    1) in directory contracts, write new contracts; it compile success, in directory migrations can generate *.js；
        truffle compile; truffle migrate --network development
        0x50f1c8814bec311812ee8bbfbd42e809d8ef0a5d generate new contract.
    2) login infura.io get apikey, modify truffle.js(need infura_apikey, mnemonic);
    3) Website or Port can be specified  (www.website.com) (80) in file scripts/start.js
        var DEFAULT_PORT = process.env.PORT || 80;
        function run(port) {
          var protocol = process.env.HTTPS === 'true' ? "https" : "http";
          var host = process.env.HOST || 'www.website.com';
          //var host = process.env.HOST || 'localhost';
          setupCompiler(host, port, protocol);
          runDevServer(host, port, protocol);
        }
     4) in file src/App.js: add const simpleStorage = contract(SimpleStorageContract) ,
                           simpleStorage.at('0x50f1c8814bec311812ee8bbfbd42e809d8ef0a5d').then((contract)
        
     5) npm install; npm start
     
 3 Client browser: visit url: http://www.website.com/
   阿里云购买网址： http://www.website.com/但未备案，太麻烦了；
   使用国外免费域名解析：https://www.cloudflare.com;
   使用免费infura测试节点： https://infura.io/
