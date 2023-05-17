async function execFn(fileName, funcName , args) {
    fileName = require(`../src/${fileName}`);
    exResult = await fileName[funcName].apply(this, args);
    return exResult;
}
fileName = process.argv[2];
funcName = process.argv[3];

i = 4;
args = [];

while(process.argv[i] != null) {
    args.push(process.argv[i]);
    i++;
}

async function print() {
    result = await execFn(fileName, funcName, args);
    console.log(result);
}
print();