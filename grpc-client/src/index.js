const express = require('express')
const app = express()

const cors = require('cors')

// SETTINGS
app.set('port', process.env.PORT || 8080)
app.use(express.json())
app.use(cors())

// ROUTES
app.use(require('./routes/client'))


// STARTING SERVER
app.listen(app.get('port'), () => {
    console.log(`Server on port ${app.get('port')}`)
})