const { pairPlayers } = require('./help');

const PAPER = 1; 
const ROCK = 2;
const SCISSORS = 3;

function getPairWinnerRPS(pair){
    if(
        pair.p1.selection == PAPER && pair.p2.selection == ROCK
    ){
        return pair.p1.player;
    }else if(
        pair.p1.selection == PAPER && pair.p2.selection == SCISSORS
    ){
        return pair.p2.player;
    }else if(
        pair.p1.selection == PAPER && pair.p2.selection == PAPER
    ){
        return pair.p1.player > pair.p2.player ? (pair.p1.player):(pair.p2.player);
    }else if(
        pair.p1.selection == ROCK && pair.p2.selection == ROCK
    ){
        return pair.p1.player > pair.p2.player ? (pair.p1.player):(pair.p2.player);
    }else if(
        pair.p1.selection == ROCK && pair.p2.selection == SCISSORS
    ){
        return pair.p1.player;
    }else if(
        pair.p1.selection == ROCK && pair.p2.selection == PAPER
    ){
        return pair.p2.player;
    }else if(
        pair.p1.selection == SCISSORS && pair.p2.selection == ROCK
    ){
        return pair.p2.player
    }else if(
        pair.p1.selection == SCISSORS && pair.p2.selection == SCISSORS
    ){
        return pair.p1.player > pair.p2.player ? (pair.p1.player):(pair.p2.player);
    }else if(
        pair.p1.selection == SCISSORS && pair.p2.selection == PAPER
    ){
        return pair.p1.player;
    }
    return -1;
}

function processPairsRPS(pairArray) {
  let winners = [];
  pairArray.forEach((pair) => {
      let winner = getPairWinnerRPS(pair);
      winners.push(winner);
  });
  if(winners.length == 1) return winners[0];
  return processPairsRPS(pairPlayers(winners, 3));
}


const rps = (players) => {
    console.log("Jugando RPS para ",players, " jugadores");
    let playerList = [];
    for(let i=1;i<=players;i++){
        playerList.push(i);
    }
    let list = pairPlayers(playerList, 3);
    let winner = processPairsRPS(list);
    console.log("Ganador: ",winner);
    return winner;
}

module.exports = {
    rps
}