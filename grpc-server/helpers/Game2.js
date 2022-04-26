const { pairPlayers } = require('./help');

function getPairWinnerFlipit(pair){
    let flipResult = Math.round(Math.random()*1+1);
    if(pair.p1.selection == flipResult){
        return pair.p1.player
    }
    return pair.p2.player;
}

function processPairsFlipit(pairArray){
    let winners = [];
    pairArray.forEach((pair) => {
        let winner = getPairWinnerFlipit(pair);
        winners.push(winner);
    });
    if(winners.length == 1) return winners[0];
    return processPairsFlipit(pairPlayers(winners, 2));
}

const flipit = (players) => {
    console.log("Jugando Toss para ",players, " jugadores");
    let playerList = [];
    for(let i=1;i<=players;i++){
        playerList.push(i);
    }
    let list = pairPlayers(playerList, 2);
    let winner = processPairsFlipit(list);
    console.log("Ganador: ",winner);
    return winner;
}

module.exports = {
    flipit
}