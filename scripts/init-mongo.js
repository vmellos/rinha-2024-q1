// init-mongo.js

db = db.getSiblingDB('rinha');

// Insira seus registros aqui
db.clientes.insertMany([
  { user_id: "1", limite: 10000, saldo: 0 },
  { user_id: "2", limite: 80000, saldo: 0 },
  { user_id: "3", limite: 1000000, saldo: 0 },
  { user_id: "4", limite: 100000000, saldo: 0 },
  { user_id: "5", limite: 5000000, saldo: 0 }
]);
