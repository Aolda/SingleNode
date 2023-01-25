function add(a,b){
    a = Number(a)
    b = Number(b)
    return a+b
}

function sub(a,b){
    return a-b;
}

function mul(a,b){
    return a*b;
}

function div(a,b){
    return a/b;
}

function mod(a,b){
    return a%b;
}

function addthree(a,b,c){
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


args = inputArgs.split(',')


result = execFn(fnName,context, ...args);