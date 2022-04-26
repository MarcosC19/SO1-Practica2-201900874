const { pairPlayers } = require('./help');

const bigBrother = (players) => {
    console.log("Jugando RumbleBiggest para ",players, " jugadores");
    let playerList = [];
    for(let i=1;i<=players;i++){
        playerList.push(i);
    }
    let list = pairPlayers(playerList, 100);

    let maximum = {
        player: 0,
        selection: 0
    };
    list.forEach((pair) => {
        if(pair.p1.selection > maximum.selection){
            maximum = pair.p1;
        }

        if(pair.p2.selection > maximum.selection){
            maximum = pair.p2;
        }
    });

    console.log(`Ganador: ${maximum.player}, (${maximum.selection})`)
    return maximum.player;
}

module.exports = {
    bigBrother
}