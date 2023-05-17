function sleep(sec) {
  return new Promise(resolve => setTimeout(resolve, sec * 1000));
} 

async function add(a,b) {
    a = Number(a);
    b = Number(b);
        await sleep(2); // 2초대기
      return a+b;
}
module.exports = { add };