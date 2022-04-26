const { pairPlayers } = require('./help');

const smallBrother = (players) => {
    console.log("Jugando RumbleSmallets para ",players, " jugadores");
    let playerList = [];
    for(let i=1;i<=players;i++){
        playerList.push(i);
    }
    let list = pairPlayers(playerList, 100);

    let minimum = {
        player: 0,
        selection: 101
    };
    
    list.forEach((pair) => {
        if(pair.p1.selection < minimum.selection){
            minimum = pair.p1;
        }

        if(pair.p2.selection < minimum.selection){
            minimum = pair.p2;
        }
    });

    console.log(`Ganador: ${minimum.player}, (${minimum.selection})`)
    return minimum.player;
}

module.exports = {
    smallBrother
}