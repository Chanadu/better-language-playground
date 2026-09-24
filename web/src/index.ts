export {};

declare global {
	interface Window {
		goPrint: (msg: string, isError: boolean) => void;
	}
}

const examples: Record<string, string> = {
	hello: `// Your first Better Language program
var greeting = "Hello, world!"
var language = "BPL"

print greeting
print "Running with " + language`,
	fibonacci: `var startTime = clock()

function fib(n) {
  if (n <= 1) return n
  return fib(n - 2) + fib(n - 1)
}

for (var i = 0; i < 10; i = i + 1) {
  print fib(i)
}

print "Time: " + (clock() - startTime) + "ms"`,
	factorial: `function factorial(n) {
  if (n == 0) {
    return 1
  }
  return n * factorial(n - 1)
}

print factorial(8)`,
};

const input = document.getElementById('bplInput') as HTMLTextAreaElement;
const runButton = document.getElementById('runButton') as HTMLButtonElement;
const clearButton = document.getElementById('clearButton') as HTMLButtonElement;
const output = document.getElementById('bplOutput') as HTMLDivElement;
const placeholder = document.getElementById('outputPlaceholder') as HTMLDivElement;
const lineNumbers = document.getElementById('lineNumbers') as HTMLDivElement;
const editorStats = document.getElementById('editorStats') as HTMLSpanElement;
const runStatus = document.getElementById('runStatus') as HTMLSpanElement;
const exampleSelect = document.getElementById('exampleSelect') as HTMLSelectElement;

function updateEditorMeta() {
	const lines = input.value.split('\n').length;
	lineNumbers.textContent = Array.from({ length: lines }, (_, index) => index + 1).join('\n');
	editorStats.textContent = `${lines} ${lines === 1 ? 'line' : 'lines'} · BPL`;
}

function clearOutput() {
	output.textContent = '';
	placeholder.hidden = false;
	runStatus.classList.remove('running');
	runStatus.innerHTML = '<i></i> Runtime ready';
}

window.goPrint = (msg: string, isError: boolean) => {
	placeholder.hidden = true;
	const line = document.createElement('span');
	line.className = isError ? 'output-line error' : 'output-line';
	line.textContent = msg;
	output.appendChild(line);
};

async function runGo(code: string) {
	const go = new Go();
	go.argv = ['better-language', code];
	try {
		const result = await WebAssembly.instantiateStreaming(fetch('main.wasm'), go.importObject);
		await go.run(result.instance);
		runStatus.classList.remove('running');
		runStatus.innerHTML = '<i></i> Finished successfully';
	} catch (error) {
		window.goPrint(`Runtime error: ${String(error)}`, true);
		runStatus.classList.remove('running');
		runStatus.innerHTML = '<i></i> Run failed';
	}
}

async function runProgram() {
	output.textContent = '';
	placeholder.hidden = true;
	const code = input.value.trim();
	runStatus.classList.add('running');
	runStatus.innerHTML = '<i></i> Running…';
	runButton.disabled = true;
	try {
		await runGo(code);
	} finally {
		runButton.disabled = false;
	}
}

input.value = examples.hello;
updateEditorMeta();

runButton.addEventListener('click', runProgram);
clearButton.addEventListener('click', clearOutput);
exampleSelect.addEventListener('change', () => {
	input.value = examples[exampleSelect.value];
	updateEditorMeta();
	input.focus();
});

input.addEventListener('input', updateEditorMeta);
input.addEventListener('scroll', () => { lineNumbers.scrollTop = input.scrollTop; });
input.addEventListener('keydown', function (event) {
	if (event.key === 'Tab') {
		event.preventDefault();
		const start = this.selectionStart;
		const end = this.selectionEnd;
		this.value = `${this.value.substring(0, start)}    ${this.value.substring(end)}`;
		this.selectionStart = this.selectionEnd = start + 4;
		updateEditorMeta();
	}
	if (event.key === 'Enter' && event.ctrlKey) {
		event.preventDefault();
		runProgram();
	}
});
