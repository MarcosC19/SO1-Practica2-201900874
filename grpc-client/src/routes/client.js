const express = require('express')
const router = express.Router()
const client = require('../helpers/client-grpc') // importing the grpc client

router.get('/', (req, res) => {
    return res.json({ 'response': 'Server running' })
})

router.post('/runGame', (req, res) => {
    const { game_id, players } = req.body

    client.Playing({ gameId: game_id, players: players }, function (err, response) {
        if (err){
            res.status(400).json({'status': err})
        }

        let result = {
            "game_id": response.gameId,
            "players": response.players,
            "game_name": response.gameName,
            "winner": response.winner
        }

        res.status(200).json(result)
    });
})

module.exports = router