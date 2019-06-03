const fs = require('fs')
const yaml = require('js-yaml')
const uuid = require('uuid/v4')
const { execSync } = require('child_process')
const express = require('express')
const bodyParser = require('body-parser')
const cors = require('cors')

const app = express()
const port = 3000

app.use(bodyParser.json())
app.use(cors())

app.post('/trigger', (req, res) => {
    const gitUrl = req.body.repository.git_url
    const taskID = uuid()

    execSync(`mkdir -p tmp`)
    execSync(`git clone --depth=1 ${gitUrl} tmp/${taskID}`)
    execSync(`rm -rf tmp/${taskID}/.git`)
    setTimeout(() => {
        const st = runPipeline(`./tmp/${taskID}/.foxy-ci.yml`, taskID)
        console.log(st)
        execSync(`rm -rf tmp/${taskID}`)
    })
    res.send()
})

app.post('/configure/:name', (req, res) => {
    execSync(`mkdir -p pipelines/${req.params.name}`)
    const config = getConfig(req.params.name)
    console.log(config)
    res.send()
})

app.listen(port, () => console.log(`Example app listening on port ${port}!`))


const getConfig = (name) => {
    try {
        const config = require(`pipelines/${name}/config.json`)
        return config
    } catch (error) {
        fs.writeFileSync(`pipelines/${name}/config.json`, '{}')
        return {}
    }
}


const runPipeline = (file, taskID) => {
    const config = yaml.safeLoad(fs.readFileSync(file, 'utf8'))
    const output = {}

    for (const actionKey in config.actions) {
        const action = config.actions[actionKey]
        output[actionKey] = runAction(action, taskID)
    }
    return output
}

const runAction = (action, taskID) => {
    const output = []
    execSync(`docker run --name ${taskID} -t -d ${action.image}`)

    for (const step of action.steps) {
        const stdout = execSync(`docker exec ${taskID} ${step}`)
        output.push(
            {
                step,
                output: stdout.toString().trim()
            })
    }

    execSync(`docker stop ${taskID}`)
    execSync(`docker rm ${taskID}`)
    return output
}

