const roulette = (players) => {
    console.log("Jugando Roulette para ",players, " jugadores");
    let playerList = [];
    for(let i=1;i<=players;i++){
        playerList.push(i);
    }

    while(playerList.length > 1){
        let randomIndex = Math.round(Math.random()*(playerList.length-1));
        playerList.splice(randomIndex, 1);
    }

    return playerList[0];
}

module.exports = {
    roulette
}