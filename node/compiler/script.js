function execFn(fileName, funcName , args){
    fileName = require(`./src/${fileName}`);
    exResult = fileName[funcName].apply(this, args);
    return exResult;
}

fileName = process.argv[2];
funcName = process.argv[3];

args = process.argv[4].split(',')

result = execFn(fileName, funcName, args);
console.log(result);