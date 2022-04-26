function randomSel(player, max) {
    let result = {
        player: player,
        selection: Math.round(
            (Math.random() * (max - 1) + 1), 0
        )
    }
    return result;
}

function pairPlayers(players, max) {
    let pairlist = [];
    for (let i = 0; i < players.length; i = i + 2) {
        if (i + 1 < players.length) {
            let pair = {
                p1: randomSel(players[i], max),
                p2: randomSel(players[i + 1], max)
            }
            pairlist.push(pair);
        } else {
            let pair = {
                p1: randomSel(i + 1, max),
                p2: randomSel(i + 1, max)
            }
            pairlist.push(pair);
        }
    }
    return pairlist;
}

module.exports = {
    pairPlayers
}