
const start = performance.now();
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

context={};
context["add"]=add;
context["mul"]=mul;
context["sub"]=sub;
context["div"]=div;
context["mod"]=mod;
context["addthree"]=addthree;


fnName = process.argv[2];
inputArgs = process.argv[3];
args = inputArgs.split(',')

result = execFn(fnName,context, ...args);
console.log(result)
const end = performance.now();
const duration = (end - start).toFixed(2);
console.log(`실행 속도: ${duration}ms`);