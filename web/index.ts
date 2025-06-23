export {};

declare global {
  interface Window {
    goPrint: (msg: string) => void;
  }
}

const input = document.getElementById("bplInput")! as HTMLTextAreaElement;
const button = document.getElementById("runButton")!;
const output = document.getElementById("bplOutput")!;

window.goPrint = function (msg: string) {
  console.log("Go Print: " + msg);
  output.textContent += msg;
};

function runGo(code: string) {
  const go = new Go();
  go.argv = ["better-language", code];
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then(
    (result) => {
      console.log("Starting Wasm");
      go.run(result.instance);
    },
  );
}

button.addEventListener("click", () => {
  console.log("Run button clicked");
  output.textContent = ""; // Clear previous output
  const code = input.value.trim();
  runGo(code);
});
