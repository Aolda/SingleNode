const readline = require("readline");
 
const rl = readline.createInterface({
    input: process.stdin,
    output: process.stdout,
});
 
let input = []
let sum = 0;

rl.on("line", (line) => {
    if(line == "q") { rl.close(); }
    input.push(parseInt(line));
});
 
rl.on('close', () => {
    input.forEach(el => {
        sum += el;
    })
    console.log(sum);
    process.exit();
})

function add() {
    
}