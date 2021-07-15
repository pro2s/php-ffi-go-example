require('./wasm_exec');
const fs = require("fs");
const data = require("./data");
const wasm = './go/web.wasm';
const go = new Go();

const start = async (productId) => {
    result = await WebAssembly.instantiate(fs.readFileSync(wasm), go.importObject);
    go.run(result.instance);

    const start = process.hrtime();
    const combinations = wasmCombine(productId, data);
    const [endSec, end] = process.hrtime(start);

    console.log(JSON.stringify({
        combinations,
        time: endSec + end / 1e9,
    }));

    go.exit(0);
}

const hash = process.argv.slice(2)[0] || 1;
start(1);
