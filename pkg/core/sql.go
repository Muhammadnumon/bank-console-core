package core

const clientsDDL = `create table if not exists clients(
    id integer primary key autoincrement ,
    name text not null,
    login text not null unique ,
    password integer not null ,
    bankAccount integer not null ,
    phoneNumber integer not null unique,
    balance integer not null check(balance>=0)
    );`
const bankMachinesDDL = `
CREATE TABLE IF NOT EXISTS bankMachine
(
    id    INTEGER PRIMARY KEY AUTOINCREMENT,
    name  TEXT    NOT NULL UNIQUE,
    street  TEXT    NOT NULL UNIQUE
);`
const servicesDDL=`create table if not exists services(
id integer primary key autoincrement,
name text not null,
price integer not null check(price>0)
);`
const bankMachinesInitialData = `INSERT INTO bankMachine(id, name,street)
VALUES (1, 'Alif bankMachine 1 ','Chapayev 24'),       
       (2, 'Alif bankMachine 2 ','Nemat Karaboyev 44'),  
       (3, 'Alif bankMachine 3 ','Yakkatut 26')
ON CONFLICT DO NOTHING;`
const clientsInitialData = `insert into clients
values(1, 'Vasya', 'vasya', 1000, 21060001,934646999,3000),
      (2, 'Petya', 'petya', 2000, 21060002,938520111,2500)
       ON CONFLICT DO NOTHING;`
const loginSQL = `select login,password from clients where login =?`
const loginSQLlog = `select id,login,password from clients where login =?`
const showBankMachine = `select id,name,street from bankMachine `
const bankAccount = `select id,name,bankAccount,balance from clients where id=?`
const insertBankMachine=`insert into bankMachine(name,street)values(:name,:street); `
const insertServices=`insert into services(name,price) values(:name,:price);`
const insertClients=`insert into clients(name,login,password,bankAccount,phoneNumber,balance)values(:name,:login,:password,:bankAccount,:phoneNumber,:balance);`
const updateBalanceSQL=`update clients set balance=balance + :balance where id=:id;`
const updateTransferByPhoneNumberMinus=`update clients set balance =balance - :balance where phoneNumber=:phoneNumber;`
const updateTransferByPhoneNumberPlus=`update clients set balance = balance + :balance where phoneNumber=:phoneNumber;`
const updateTransferBYBankAccountMinus=`update clients set balance=balance - :balance where bankAccount=:bankAccount;`
const updateTransferByBankAccountPlus=`update clients set balance=balance + :balance where bankAccount=:bankAccount;`
const payServicesplus=`update services set price=price + :price where id=:id;`
const payServicesMinus=`update clients set balance=balance - :balance where id=:id;`