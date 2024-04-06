#! /usr/bin/env node

const { exec } = require("child_process");
const path = require("path");

console.log("===Instalando dependências===");

exec(`./hello`, (error, stdout, stderr) => {
  if (error) {
    console.log(`error: ${error.message}`);
    return;
  }
  if (stderr) {
    console.log(`stderr: ${stderr}`);
  }
  console.log("==Tudo pronto para começar==");
  console.log("==stdout==", stdout);
});
