# Better Language Playground

A browser-based editor and interpreter for [Better Programming Language](https://github.com/Chanadu/better-language).

The playground compiles the Go interpreter to WebAssembly, so BPL programs run locally in the browser without a backend or remote execution service.

## Features

- Write and run BPL code directly in the browser
- Go interpreter compiled to WebAssembly
- Built-in Hello World, Fibonacci, and Factorial examples
- Editor line numbers and four-space Tab indentation
- `Ctrl + Enter` execution shortcut
- Clear output and runtime status controls
- Persistent light and dark themes
- Responsive playground and language documentation

## Getting started

### Requirements

- Node.js and npm
- Go 1.23 or newer

### Run locally

```sh
git clone https://github.com/Chanadu/better-language-playground.git
cd better-language-playground/web
npm install
npm run dev
```

Open [http://localhost:3000](http://localhost:3000).

Use a different port by setting `PORT`:

```sh
PORT=4000 npm run dev
```

The development server watches the HTML, TypeScript, CSS, static assets, and Go source. Changes are rebuilt and reloaded automatically.

## Production build

```sh
cd web
npm ci
npm run build
```

The deployable static site is generated in `web/dist/`. It contains the HTML pages, compiled JavaScript and CSS, Go WebAssembly runtime, and interpreter binary.

To remove generated output:

```sh
npm run clean
```

## Project structure

```text
.
├── Better-Language/     # Scanner, parser, interpreter, and Go WASM entry point
└── web/
    ├── public/          # Static browser assets
    ├── scripts/         # Build and asset-copying utilities
    ├── src/
    │   ├── docs.html    # Language documentation
    │   ├── index.html   # Playground page
    │   ├── index.ts     # Editor and interpreter integration
    │   ├── theme.ts     # Light and dark theme behavior
    │   └── styles.css   # Application styles
    ├── dev-server.mjs   # Local server and file watchers
    ├── package.json
    ├── tailwind.config.js
    └── tsconfig.json
```

`web/dist/` is generated output and is not committed.

## Available commands

Run commands from `web/`.

| Command | Description |
| --- | --- |
| `npm run dev` | Start the local server with rebuilds and live reload |
| `npm run build` | Create a clean production build |
| `npm test` | Verify that the complete project builds |
| `npm run clean` | Remove generated files from `dist/` |
| `npm run build:js` | Compile TypeScript |
| `npm run build:css` | Compile Tailwind CSS |
| `npm run build:wasm` | Compile the Go interpreter to WebAssembly |
| `npm run build:static` | Copy HTML and public assets into `dist/` |

## BPL example

```text
function fib(n) {
    if (n <= 1) return n
    return fib(n - 2) + fib(n - 1)
}

for (var i = 0; i < 10; i = i + 1) {
    print fib(i)
}
```

See the in-app documentation for variables, functions, control flow, operators, data types, comments, and built-in functions.

## Deployment

The production output is a static site and can be hosted on GitHub Pages, Netlify, Vercel, Cloudflare Pages, or any static web server.

Configure the host to:

1. Run `cd web && npm ci && npm run build`.
2. Publish the `web/dist` directory.
3. Serve `.wasm` files with the `application/wasm` MIME type.

No application server, database, or environment variables are required.

## Related project

The standalone language repository and its version history are available at [Chanadu/better-language](https://github.com/Chanadu/better-language).
