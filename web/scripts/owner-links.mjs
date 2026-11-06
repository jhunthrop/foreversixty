// web/scripts/owner-links.mjs
// The repository goes private after launch, and nothing the public reaches may link to it or
// lead back to its owner (memory: forever-legal-entity). This is the check, as a pure function
// of a list of directories so src/lib/site/owner-links.test.ts can drive it (a plain .mjs: the
// postbuild script runs it under node before vitest ever starts, as removed-routes.mjs does).
import { readdirSync, readFileSync } from 'node:fs';
import { extname, join, relative } from 'node:path';

/** The owner's code host account. Nothing the site serves may contain it. */
export const OWNER_LINK = 'github.com/jhunthrop';

/** Files read as text. Compressed and binary assets (wasm, fonts, images) cannot carry a link a reader follows. */
export const TEXT_EXTENSIONS = new Set([
  '.html',
  '.json',
  '.js',
  '.mjs',
  '.css',
  '.txt',
  '.xml',
  '.webmanifest',
  '.md',
  '.mdx',
  '.yml',
  '.yaml',
  '.svg',
  '.lua',
  '.toc',
]);

function* textFiles(dir) {
  for (const entry of readdirSync(dir, { withFileTypes: true })) {
    const path = join(dir, entry.name);
    if (entry.isDirectory()) yield* textFiles(path);
    else if (entry.isFile() && TEXT_EXTENSIONS.has(extname(entry.name).toLowerCase())) yield path;
  }
}

/** Every text file under `roots` that contains OWNER_LINK, relative to `base`, with its first line number. */
export function ownerLinkFailures(roots, base) {
  const failures = [];
  for (const root of roots) {
    for (const file of textFiles(root)) {
      const lines = readFileSync(file, 'utf8').split('\n');
      const index = lines.findIndex((line) => line.toLowerCase().includes(OWNER_LINK));
      if (index >= 0) failures.push(`${relative(base, file)}:${index + 1}`);
    }
  }
  return failures;
}
