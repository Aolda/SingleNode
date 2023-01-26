function execFn(funcName, fileName, args){
    fileName = require(`./src/${fileName}`);
    exResult = fileName[funcName].apply(this, args);
    return exResult;
}

fileName = process.argv[2];
funcName = process.argv[3];

i = 4;
args = [];

while(process.argv[i] != undefined) {
    args.push(process.argv[i]);
    i++;
}

result = execFn(funcName, fileName, args);
console.log(result);