const input = document.getElementById('bplInput');
const button = document.getElementById('runButton');
const output = document.getElementById('bplOutput');
window.goPrint = function (msg, isError) {
    if (isError) {
        console.log('Go Error: ' + msg);
    }
    else {
        console.log('Go Print: ' + msg);
    }
    if (isError) {
        output.textContent += '<span style="color: red;">' + msg + '</span>';
    }
    else {
        output.textContent += msg;
    }
    // output.textContent += '<span style="color: red;">' + msg + '</span>';
    // output.textContent += '<span style="color: red;">' + msg + '</span>';
};
function runGo(code) {
    const go = new Go();
    go.argv = ['better-language', code];
    WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject).then((result) => {
        console.log('WASM');
        go.run(result.instance);
    });
}
button.addEventListener('click', () => {
    console.log('Run button clicked');
    output.textContent = ''; // Clear previous output
    const code = input.value.trim();
    runGo(code);
});
export {};
//# sourceMappingURL=index.js.map