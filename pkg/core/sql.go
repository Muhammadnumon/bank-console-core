package core

const clientsDDL  =`create table if not exists clients(
    id integer primary key autoincrement ,
    name text not null,
    login text not null unique ,
    password integer not null ,
    bankAccount integer not null 
    )`
const loginSQL=`select login,password from clients where login ?`