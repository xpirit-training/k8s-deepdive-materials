const fs = require('fs');
const express = require('express');
const app = express();

let backendUrl;
if (fs.existsSync('/etc/config/config.properties')) {
    const configProperties = fs.readFileSync('/etc/config/config.properties', 'utf8');
    backendUrl = configProperties.split('=')[1].trim();
} else {
    backendUrl = process.env.BACKEND_URL || 'http://localhost:8080';
}

app.get('/env.js', (req, res) => {
    const myEnvVar = process.env.APP_NAME || '';
    res.send(`const APP_NAME = "${myEnvVar}";`);
});


app.get('/config.js', (req, res) => {
    res.send(`const backendUrl = "${backendUrl}";`);
});

app.use(express.static('public'));

const PORT = process.env.PORT || 3000;
app.listen(PORT, () => {
    console.log(`Server is running on http://localhost:${PORT}`);
});
