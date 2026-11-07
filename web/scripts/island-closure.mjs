// web/scripts/island-closure.mjs
// An island's startup closure (owner-delegated budget, 2026-11-07): its entry file plus every
// chunk it imports statically, followed transitively, i.e. what a browser must fetch before the
// island can run. Dynamic `import()` chunks load later and are not counted. Each file is gzipped
// on its own, as each is fetched on its own; a shared chunk counts once per closure.
import { readFile } from 'node:fs/promises';
import path from 'node:path';
import { gzipSync } from 'node:zlib';

/** A static `import`/`export ... from` of a relative .js chunk (minified or not); never `import(`. */
const STATIC_IMPORT =
  /\b(?:import|export)\s*(?:\{[^}]*\}\s*from\s*|\*\s*(?:as\s+[\w$]+\s*)?from\s*|[\w$]+\s*(?:,\s*\{[^}]*\})?\s*from\s*)?["'](\.{1,2}\/[^"']+\.js)["']/g;

/** The relative chunks `source` imports statically. */
export function staticImports(source) {
  return [...source.matchAll(STATIC_IMPORT)].map((match) => match[1]);
}

/** The entry and its static closure: `{ files, gzipped }`, files relative to `root`. */
export async function startupClosure(entry, root) {
  const seen = new Set();
  const pending = [path.resolve(entry)];
  let gzipped = 0;
  while (pending.length > 0) {
    const file = pending.pop();
    if (seen.has(file)) continue;
    seen.add(file);
    const source = await readFile(file);
    gzipped += gzipSync(source, { level: 9 }).length;
    for (const specifier of staticImports(source.toString('utf8'))) {
      pending.push(path.resolve(path.dirname(file), specifier));
    }
  }
  return { files: [...seen].map((file) => path.relative(root, file)), gzipped };
}
