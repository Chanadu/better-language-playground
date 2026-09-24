# better-language-playground
Better Language Playground is a fully functional, interactive, online playground for my custom built programming language, [Better Language](https://www.github.com/Chanadu/better-language). 
The interactive playground was written using web assembly and Astro JS.

The pure code is embedded within this repo, and the version history for the programming language is in the [Better Language Github Repo](https://www.github.com/Chanadu/better-language).

## Local development

The development server rebuilds TypeScript, Tailwind CSS, and the Go WebAssembly binary as their source files change. It automatically reloads the browser whenever a build finishes or an HTML file changes.

```sh
cd web
npm install
npm run dev
```

Then open [http://localhost:3000](http://localhost:3000). To use a different port, run `PORT=4000 npm run dev`.
