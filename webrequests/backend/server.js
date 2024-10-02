
const express = require('express')

const PORT = 5000

const app = express()

app.use(express.json())
app.use(express.urlencoded({extended : true}))

app.get('/',(req,res)=>{
    res.send(`<center> <h1> Express as Backend Server for Go Lang </h1> </center>`)
})

app.get('/get',(req,res)=>{
    res.json({message : "Hello from backend server This is Get Request"})
})

app.post('/post',(req,res)=>{
    const data = req.body
    console.log(data);
    res.json(data)
})

app.post('/postform',(req,res)=>{
    const data = req.body
    console.log(data);
    res.json(data)
})

app.listen(PORT,()=>{
    console.log(`Server running on port ${PORT}`);
})