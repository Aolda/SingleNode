functionName = process.argv[2];
parameter = [process.argv[3], process.argv[4]];

eval(functionName + "(" + parameter[0] + "," + parameter[1] + ")");

function add(a, b) {
    a = parseInt(a);
    b = parseInt(b);
    console.log(a + b);
}