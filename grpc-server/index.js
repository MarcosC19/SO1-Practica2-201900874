var PROTO_PATH = "./protos/server.proto";

const { rps } = require('./helpers/Game1');
const { flipit } = require('./helpers/Game2');
const { bigBrother } = require('./helpers/Game3');
const { smallBrother } = require('./helpers/Game4');
const { roulette } = require('./helpers/Game5');

const amqp = require('amqplib/callback_api');
const grpc = require('@grpc/grpc-js');
const protoLoader = require('@grpc/proto-loader');
const { delay } = require('lodash');
const packageDefinition = protoLoader.loadSync(
    PROTO_PATH,
    {
        keepCase: true,
        longs: String,
        enums: String,
        defaults: true,
        oneofs: true
    }
);

const practica2_proto = grpc.loadPackageDefinition(packageDefinition).practica2;

const IPRabbit = process.env.HOSTNAME_RABBIT || "localhost";
const UserRabbit = "rabbit";
const PassRabbit = "sopes1";
const ColaRabbit = "GamesQueue";
var CanalRabbit = undefined;

function Playing(call, callback) {
    let gameName;
    let winner;

    switch (call.request.gameid) {
        case 1:
            gameName = "Piedra, Papel o Tijeras";
            winner = rps(call.request.players);
            break;
        case 2:
            gameName = "Cara o Cruz";
            winner = flipit(call.request.players);
            break;
        case 3:
            gameName = "Numero mayor";
            winner = bigBrother(call.request.players);
            break;
        case 4:
            gameName = "Numero menor";
            winner = smallBrother(call.request.players);
            break;
        case 5:
            gameName = "Ruleta";
            winner = roulette(call.request.players);
            break;
        default:
            gameName = "Piedra, Papel o Tijeras";
            winner = rps(call.request.players);
            break;
    }

    let log = {
        game_id: parseInt(call.request.gameid),
        players: parseInt(call.request.players),
        game_name: gameName,
        winner: parseInt(winner),
        queue: "RabbitMQ"
    }

    CanalRabbit.sendToQueue(ColaRabbit, Buffer.from(JSON.stringify(log)));

    callback(null, {
        status: 1
    });
}

function main() {
    var server = new grpc.Server();
    server.addService(practica2_proto.PlayGame.service, {
        Playing: Playing
    });
    server.bindAsync(
        `0.0.0.0:50051`,
        grpc.ServerCredentials.createInsecure(),
        () => {
            server.start();
        }
    );
}

const rabbitConnect = () => {
    amqp.connect(`amqp://${UserRabbit}:${PassRabbit}@${IPRabbit}`, async (error, connection) => {
        if (error) {
            console.log("Error with connection:", error);
            return;
        }
        console.log("Succesfully Rabbit connection");
        connection.createChannel((error, channel) => {
            if (error) {
                console.log("Error with channel creation:", error);
                return;
            }

            CanalRabbit = channel;
            CanalRabbit.assertQueue(ColaRabbit, {
                durable: false
            });
            console.log("Successfully created queue");
        });
    });
}

delay(rabbitConnect, 50000);
delay(main, 2500);