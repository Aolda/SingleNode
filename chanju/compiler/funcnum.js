
const start = performance.now();
//node funcnum.js RTFN 하면 함수 갯수 알려줌
const fs = require('fs');
const esprima = require('esprima'); // 모듈 추가 가능????????????????

function ReadtheFuncNum(){
    // Read the JavaScript file
    const file = fs.readFileSync('./script.js', 'utf-8');

    // Use Esprima to parse the file
   const ast = esprima.parseScript(file);

    // Iterate over the functions in the file
    ast.body.forEach(element => {
        if (element.type === 'FunctionDeclaration'){
            if(element.id.name !== "execFn"){
                console.log(`Function name: ${element.id.name}`);
                console.log(`Parameters: ${element.params.map(p => p.name).join(', ')}`);    
            } 
           }
    })
}

function execFn(fnName, context){
    args = Array.prototype.slice.call(arguments,2);
    return context[fnName].apply(context, args);
}

context = {}
context["RTFN"]=ReadtheFuncNum;

fnName = process.argv[2];
inputArgs = process.argv[3];
args = []

result = execFn(fnName,context,...args);

const end = performance.now();
const duration = (end - start).toFixed(2);
console.log(`실행 속도: ${duration}ms`);

