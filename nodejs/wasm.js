require('./wasm_exec');
const fs = require("fs");
const data = require("./data");
const wasm = './go/util.wasm';
const go = new Go();

const start = async () => {
    result = await WebAssembly.instantiate(fs.readFileSync(wasm), go.importObject);
    go.run(result.instance);

    const combinations = wasmCombine(1, data);
    console.log(combinations.total);
    console.log(combinations['65e36ba5']);

    go.exit(0);
}

start();
