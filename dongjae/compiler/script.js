const fs = require('fs'); // 모듈 추가해야되는데 이게 맞는건지 잘 모르겠음
const esprima = require('esprima'); // 동일

function add(a,b){
    a = Number(a)
    b = Number(b)
    return a+b
}

function sub(a,b){
    a = Number(a)
    b = Number(b)
    return a-b;
}

function mul(a,b){
    a = Number(a)
    b = Number(b)
    return a*b;
}

function div(a,b){
    a = Number(a)
    b = Number(b)
    return a/b;
}

function mod(a,b){
    a = Number(a)
    b = Number(b)
    return a%b;
}

function addthree(a,b,c){
    a = Number(a)
    b = Number(b)
    c = Number(c)
    return a+b+c;
}

function execFn(fnName, context){
    args = Array.prototype.slice.call(arguments,2);
    return context[fnName].apply(context, args);
}

function ReadtheFuncNum(){
    // Read the JavaScript file
    const file = fs.readFileSync('./script.js', 'utf-8');

    // Use Esprima to parse the file
    const ast = esprima.parseScript(file);

    // Iterate over the functions in the file
    ast.body.forEach(element => {
        if (element.type === 'FunctionDeclaration'){
            if(element.id.name !== "ReadtheFuncNum" && element.id.name !== "execFn"){
                console.log(`Function name: ${element.id.name}`);
                console.log(`Parameters: ${element.params.map(p => p.name).join(', ')}`);    
            } 
           }
    });
}

context={};
context["add"]=add;
context["mul"]=mul;
context["sub"]=sub;
context["div"]=div;
context["mod"]=mod;
context["addthree"]=addthree;
context["RTFN"]=ReadtheFuncNum;

fnName = process.argv[2];
inputArgs = process.argv[3];
if(inputArgs !== undefined ){
    args = inputArgs.split(',')
}else{
    args = []
}

result = execFn(fnName,context, ...args);
if(fnName !== "RTFN") console.log(result)
