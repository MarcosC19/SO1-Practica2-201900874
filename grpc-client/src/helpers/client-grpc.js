require('dotenv').config()

var PROTO_PATH = __dirname + '/./protos/client.proto';

var parseArgs = require('minimist');
var grpc = require('@grpc/grpc-js');
var protoLoader = require('@grpc/proto-loader');
var packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {keepCase: true,
     longs: String,
     enums: String,
     defaults: true,
     oneofs: true
    });
var fase2_proto = grpc.loadPackageDefinition(packageDefinition).practica2;

var argv = parseArgs(process.argv.slice(2), {
  string: 'target'
});
var target;
if (argv.target) {
  target = argv.target;
} else {
  target = `${process.env.IP_SERVER}:50051`;
}
var client = new fase2_proto.PlayGame(target, grpc.credentials.createInsecure());

module.exports = client