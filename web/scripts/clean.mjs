import { rm } from 'node:fs/promises';
import path from 'node:path';
import { fileURLToPath } from 'node:url';

const scriptsRoot = path.dirname(fileURLToPath(import.meta.url));
const distRoot = path.resolve(scriptsRoot, '../dist');

await rm(distRoot, { recursive: true, force: true });
