export {};

declare global {
	interface Window {
		goPrint: (msg: string, isError: boolean) => void;
	}
}

const input = document.getElementById('bplInput')! as HTMLTextAreaElement;
const button = document.getElementById('runButton')!;
const output = document.getElementById('bplOutput')!;

window.goPrint = function (msg: string, isError: boolean) {
	if (isError) {
		console.log('Go Error: ' + msg);
	} else {
		console.log('Go Print: ' + msg);
	}

	const newChild = document.createElement('span');
	newChild.textContent = msg;

	if (isError) {
		// output.textContent += '<span style="color: red;">' + msg + '</span>';
		newChild.style = 'color: red;';
	} else {


		// output.textContent += msg;
	}

	output.appendChild(newChild);

	// output.textContent += '<span style="color: red;">' + msg + '</span>';

	// output.textContent += '<span style="color: red;">' + msg + '</span>';
};

function runGo(code: string) {
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
	// document.body.document.createElement('div');

	runGo(code);
});
		               
input.addEventListener('keydown', function(e) {
  if (e.key == 'Tab') {
    e.preventDefault();
    var start = this.selectionStart;
    var end = this.selectionEnd;

    // set textarea value to: text before caret + tab + text after caret
    this.value = this.value.substring(0, start) +
      "\t" + this.value.substring(end);

    // put caret at right position again
    this.selectionStart =
      this.selectionEnd = start + 1;
  }
});		               
