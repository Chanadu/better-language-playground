import { copyFile, mkdir, readdir } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const scriptsRoot = path.dirname(fileURLToPath(import.meta.url));
const webRoot = path.resolve(scriptsRoot, '..');
const sourceRoot = path.join(webRoot, 'src');
const publicRoot = path.join(webRoot, 'public');
const distRoot = path.join(webRoot, 'dist');

async function copyDirectory(source, destination) {
	await mkdir(destination, { recursive: true });
	const entries = await readdir(source, { withFileTypes: true });

	await Promise.all(entries.map(async (entry) => {
		const sourcePath = path.join(source, entry.name);
		const destinationPath = path.join(destination, entry.name);

		if (entry.isDirectory()) {
			await copyDirectory(sourcePath, destinationPath);
		} else {
			await copyFile(sourcePath, destinationPath);
		}
	}));
}

export async function copyStaticAssets() {
	await mkdir(distRoot, { recursive: true });
	await Promise.all([
		copyFile(path.join(sourceRoot, 'index.html'), path.join(distRoot, 'index.html')),
		copyFile(path.join(sourceRoot, 'docs.html'), path.join(distRoot, 'docs.html')),
		copyDirectory(publicRoot, distRoot),
	]);
}

if (process.argv[1] === fileURLToPath(import.meta.url)) {
	await copyStaticAssets();
}
