import { spawn } from 'node:child_process';
import { existsSync, statSync, watch } from 'node:fs';
import { readFile } from 'node:fs/promises';
import { createServer } from 'node:http';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const webRoot = path.dirname(fileURLToPath(import.meta.url));
const distRoot = path.join(webRoot, 'dist');
const languageRoot = path.resolve(webRoot, '../Better-Language');
const host = process.env.HOST || '127.0.0.1';
const port = Number(process.env.PORT || 3000);
const reloadClients = new Set();
const childProcesses = new Set();

const mimeTypes = {
	'.css': 'text/css; charset=utf-8',
	'.html': 'text/html; charset=utf-8',
	'.js': 'text/javascript; charset=utf-8',
	'.json': 'application/json; charset=utf-8',
	'.map': 'application/json; charset=utf-8',
	'.wasm': 'application/wasm',
};

function run(name, command, args, options = {}) {
	const child = spawn(command, args, {
		cwd: options.cwd || webRoot,
		env: { ...process.env, ...options.env },
		stdio: 'inherit',
	});

	childProcesses.add(child);
	child.once('exit', (code, signal) => {
		childProcesses.delete(child);
		if (!options.persistent && code !== 0) {
			console.error(`[dev] ${name} failed with exit code ${code}`);
		}
		if (options.persistent && !signal) {
			console.error(`[dev] ${name} watcher stopped with exit code ${code}`);
		}
	});

	return child;
}

function startAssetWatchers() {
	run(
		'TypeScript',
		process.execPath,
		['./node_modules/typescript/bin/tsc', '--watch', '--preserveWatchOutput'],
		{ persistent: true },
	);
	run(
		'Tailwind',
		process.execPath,
		[
			'./node_modules/tailwindcss/lib/cli.js',
			'-i',
			'./main.css',
			'-o',
			'./dist/output.css',
			'--watch',
		],
		{ persistent: true },
	);
}

let goBuildTimer;
let goBuildRunning = false;
let goBuildQueued = false;

function buildWasm() {
	if (goBuildRunning) {
		goBuildQueued = true;
		return;
	}

	goBuildRunning = true;
	console.log('[dev] Building WebAssembly...');
	const child = run(
		'WebAssembly build',
		'go',
		['build', '-o', path.join(distRoot, 'main.wasm'), '.'],
		{
			cwd: languageRoot,
			env: { GOOS: 'js', GOARCH: 'wasm' },
		},
	);

	child.once('exit', (code) => {
		goBuildRunning = false;
		if (code === 0) console.log('[dev] WebAssembly build complete.');
		if (goBuildQueued) {
			goBuildQueued = false;
			buildWasm();
		}
	});
}

function scheduleWasmBuild() {
	clearTimeout(goBuildTimer);
	goBuildTimer = setTimeout(buildWasm, 100);
}

function watchGoSources() {
	watch(languageRoot, { recursive: true }, (_event, filename) => {
		if (filename?.endsWith('.go') || filename === 'go.mod' || filename === 'go.sum') {
			scheduleWasmBuild();
		}
	});
}

let reloadTimer;

function scheduleReload() {
	clearTimeout(reloadTimer);
	reloadTimer = setTimeout(() => {
		for (const response of reloadClients) response.write('data: reload\n\n');
	}, 75);
}

function injectReloadClient(html) {
	const script = `<script>
		new EventSource('/__live_reload').onmessage = () => location.reload();
	</script>`;
	return html.includes('</body>') ? html.replace('</body>', `${script}\n</body>`) : html + script;
}

const server = createServer(async (request, response) => {
	const requestUrl = new URL(request.url || '/', 'http://localhost');

	if (requestUrl.pathname === '/__live_reload') {
		response.writeHead(200, {
			'Cache-Control': 'no-cache',
			Connection: 'keep-alive',
			'Content-Type': 'text/event-stream',
		});
		response.write(': connected\n\n');
		reloadClients.add(response);
		request.on('close', () => reloadClients.delete(response));
		return;
	}

	const pathname = requestUrl.pathname === '/' ? '/index.html' : requestUrl.pathname;
	const filePath = path.resolve(distRoot, `.${decodeURIComponent(pathname)}`);
	const relativePath = path.relative(distRoot, filePath);

	if (relativePath.startsWith('..') || path.isAbsolute(relativePath)) {
		response.writeHead(403).end('Forbidden');
		return;
	}

	try {
		const resolvedPath = statSync(filePath).isDirectory()
			? path.join(filePath, 'index.html')
			: filePath;
		let contents = await readFile(resolvedPath);
		const extension = path.extname(resolvedPath);

		if (extension === '.html') {
			contents = Buffer.from(injectReloadClient(contents.toString()));
		}

		response.writeHead(200, {
			'Cache-Control': 'no-cache',
			'Content-Type': mimeTypes[extension] || 'application/octet-stream',
		});
		response.end(contents);
	} catch (error) {
		if (error.code === 'ENOENT') {
			response.writeHead(404).end('Not found');
			return;
		}
		console.error(error);
		response.writeHead(500).end('Internal server error');
	}
});

if (!existsSync(distRoot)) {
	console.error('[dev] Missing web/dist directory.');
	process.exit(1);
}

startAssetWatchers();
buildWasm();
watchGoSources();
watch(distRoot, { recursive: true }, scheduleReload);

server.listen(port, host, () => {
	console.log(`[dev] Playground running at http://${host}:${port}`);
	console.log('[dev] Changes to TypeScript, CSS, HTML, or Go files reload the browser.');
});

function shutDown() {
	for (const child of childProcesses) child.kill('SIGTERM');
	for (const response of reloadClients) response.end();
	server.close(() => process.exit(0));
}

process.on('SIGINT', shutDown);
process.on('SIGTERM', shutDown);
