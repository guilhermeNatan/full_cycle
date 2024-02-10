const express = require('express')
const app = express()
const port = 3000
const mysql = require('mysql')

const config = {
    host: 'db',
    user: 'root',
    password: 'root',
    database:'nodedb'
};
const connection = mysql.createConnection(config)

app.use(express.urlencoded({extended:true}))
app.use(express.json())

app.get('/', (req,res) => {
    res.send('<h1>Full Cycle</h1>')
})


app.post('/api/inserir', (req,res) => {
    console.log(req.body)
    const {nomePessoa} = req.body;
  
    const sql = `INSERT INTO user(name) values('${nomePessoa}')`

    connection.query(sql)

    res.json({nomePessoa})
})

app.get('/api/listar', (req,res) => {

 const sql = 'SELECT * FROM user';
 connection.query(sql, (erro, resultados) => {
    if (erro) {
      console.error('Erro ao obter dados da tabela:', erro);
      res.status(500).send('Erro interno do servidor');
    } else {
      res.json(resultados);
    }
  });
})


app.listen(port, ()=> {
    console.log('Rodando na porta ' + port)
})