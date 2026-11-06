// web/scripts/parity-mocks-stale.mjs
// Whether the committed parity mocks (design/mocks/web, the human reference and the side-by-side
// source) still match what the generators render from the build the parity specs served
// (tests/e2e/support/parity-mocks.ts writes those to test-results/parity/mocks/<page>/). The gate
// itself compares against the test-time render, so a stale reference never fails a landing; this
// is the report that keeps the reference honest (design/parity.md, "The reference mocks").
// Tested by src/lib/site/parity-mocks-stale.test.ts; run by scripts/check-parity-mocks.mjs.
import { existsSync, readdirSync, readFileSync } from 'node:fs';
import path from 'node:path';

/** Each page's test-time folder and the exact commands that refresh its committed mocks. */
export const PARITY_PAGES = {
  bis: {
    prefix: 'web-bis-',
    regen: [
      '(cd web && npm run build && npm run preview -- --port 4500)   # in its own terminal',
      'node design/mocks/bis_facts.mjs http://localhost:4500 && python3 design/mocks/ws_bis.py',
    ],
  },
  sim: {
    prefix: 'web-sim',
    regen: ['python3 design/mocks/ws_sim.py'],
  },
};

// A game icon is drawn when the checkout carries its file and as an empty socket when it does not
// (design/mocks/wide_lib.py, ws_sim.py gicon): same box, no role, so not a difference of design.
const ICON_IMG = /<img class="(ic[^"]*)" src="data:[^"]*"[^>]*>/g;
const ICON_SOCKET = /<span class="(ic[^"]*)" style="[^"]*"><\/span>/g;

/** A mock's HTML with what depends on the checkout (which icon files it holds) taken out. */
export function normalizeMock(html) {
  return html.replace(ICON_IMG, '<icon class="$1">').replace(ICON_SOCKET, '<icon class="$1">');
}

/** About 60 characters of text either side of the first difference, tags and data URIs dropped. */
function firstDifference(committed, rendered) {
  let i = 0;
  while (i < committed.length && committed[i] === rendered[i]) i += 1;
  const around = (s) =>
    s
      .slice(Math.max(0, i - 400), i + 400)
      .replace(/data:[^"]*/g, '')
      .replace(/<[^>]*>/g, ' ')
      .replace(/\s+/g, ' ')
      .trim()
      .slice(0, 160);
  return { committed: around(committed), rendered: around(rendered) };
}

/**
 * The stale mocks of one page: each test-time mock whose committed copy is missing or differs.
 * Throws when the test-time folder is missing: the parity specs have not run.
 */
export function staleMocks(renderedDir, committedDir, prefix) {
  if (!existsSync(renderedDir)) {
    throw new Error(`no test-time mocks in ${renderedDir}: run the parity specs first`);
  }
  return readdirSync(renderedDir)
    .filter((file) => file.startsWith(prefix) && file.endsWith('.html'))
    .sort()
    .flatMap((file) => {
      const committedFile = path.join(committedDir, file);
      const rendered = normalizeMock(readFileSync(path.join(renderedDir, file), 'utf8'));
      if (!existsSync(committedFile)) return [{ file, missing: true }];
      const committed = normalizeMock(readFileSync(committedFile, 'utf8'));
      return committed === rendered
        ? []
        : [{ file, missing: false, ...firstDifference(committed, rendered) }];
    });
}

/** The report: which mocks are stale and the commands that refresh them. */
export function formatStaleReport(stale) {
  return Object.entries(stale)
    .filter(([, files]) => files.length > 0)
    .map(([page, files]) => {
      const lines = files.map((f) =>
        f.missing
          ? `  ${f.file}: not committed`
          : `  ${f.file}\n    committed: …${f.committed}…\n    build:     …${f.rendered}…`,
      );
      const regen = PARITY_PAGES[page].regen.map((c) => `    ${c}`).join('\n');
      return `${page}: the committed mocks are behind the build\n${lines.join('\n')}\n  refresh them from the repo root, then commit design/mocks/web:\n${regen}`;
    })
    .join('\n\n');
}
