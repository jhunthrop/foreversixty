// web/src/content/guides/_prose.test.ts
// Underscore-prefixed so Astro's content loader (glob: '**/*.md') never treats it as a guide;
// it reads the files straight off disk for the reason _sections.test.ts gives.
//
// The guide prose check: a guide reads like the best player of the class wrote it, so the
// text a guide page renders carries no dashes, no talk about the site or its method and no
// AI tells. It scans, per class:
//   - every guide markdown file in web/src/content/guides/<class>/ (the class index and each
//     spec guide): the body, minus HTML comments and link targets, and the frontmatter text
//     the page shows (title, description, source labels);
//   - every note under `rotation` in data/curated/apl/<class>-<spec>.json for each spec guide
//     (the priority notes the Rotation block and the Play it panel render).
//
// OPT-IN PER CLASS. Only the classes in ENFORCED_CLASSES fail the suite, so main stays green
// while the class lanes rewrite their guides. A lane that lands its rewrite adds its class
// slug (the guide folder name, e.g. 'warrior') to ENFORCED_CLASSES in the same commit; once
// every class is in, the list can go and the check applies to all.
//
// To extend the rules, add to BANNED_PROSE: a string matches case-insensitively as a whole
// phrase, a RegExp matches as written. Every failure names the file, the line and the phrase.
//
// Report mode: `GUIDE_PROSE_REPORT=<path> npx vitest run src/content/guides/_prose.test.ts`
// writes the per-class violation counts for every class, enforced or not, to <path>.
import { existsSync, readdirSync, readFileSync, writeFileSync } from 'node:fs';
import { dirname, join, relative } from 'node:path';
import { fileURLToPath } from 'node:url';
import { describe, expect, it } from 'vitest';

/** Class slugs whose guides must pass. Each class lane adds its own when it lands. */
export const ENFORCED_CLASSES: readonly string[] = ['mage', 'paladin'];

/** Everything a guide's prose must not contain, by rule. */
export const BANNED_PROSE = {
  dashes: ['—', '–', /\s--\s/],
  siteMeta: [
    'this site',
    'our simulator',
    'the simulator',
    'the sim found',
    'search run',
    'the search',
    'our search',
    'the gate',
    'the APL',
    'pipeline',
    'preset',
    'presets',
    // "band" as a level word ("the 40 band", "band 30"); item names ("Band of Accuria") pass.
    /\b[Bb]ands?\b(?! of\b)/,
  ],
  aiTells: [
    "it's worth noting",
    'delve',
    'delves',
    'crucial',
    'robust',
    'seamless',
    'seamlessly',
    'leverage',
    'in summary',
    "it's important to note",
  ],
} as const satisfies Record<string, readonly (string | RegExp)[]>;

export interface ProseViolation {
  file: string;
  line: number;
  rule: keyof typeof BANNED_PROSE;
  phrase: string;
}

const guidesRoot = dirname(fileURLToPath(import.meta.url));
const repoRoot = join(guidesRoot, '../../../..');
const aplRoot = join(repoRoot, 'data/curated/apl');

function escapeRegExp(text: string): string {
  return text.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}

function matcherFor(entry: string | RegExp): RegExp {
  if (typeof entry !== 'string') return new RegExp(entry.source, entry.flags.replace('g', '') + 'g');
  // Whole-phrase match: a word phrase needs word edges, a bare glyph does not.
  const edged = /\w/.test(entry) ? `(?<![\\w-])${escapeRegExp(entry)}(?![\\w-])` : escapeRegExp(entry);
  return new RegExp(edged, 'gi');
}

const MATCHERS = (
  Object.entries(BANNED_PROSE) as [keyof typeof BANNED_PROSE, readonly (string | RegExp)[]][]
).flatMap(([rule, entries]) => entries.map((entry) => ({ rule, pattern: matcherFor(entry) })));

/** Every banned phrase in one rendered line of text. */
export function violationsIn(text: string, file: string, line: number): ProseViolation[] {
  const normalized = text.replace(/[‘’]/g, "'");
  return MATCHERS.flatMap(({ rule, pattern }) =>
    [...normalized.matchAll(pattern)].map((match) => ({ file, line, rule, phrase: match[0].trim() })),
  );
}

const FRONTMATTER_TEXT = /^\s*(?:-\s+)?(?:title|description|label):\s*(.*)$/;

/** The text a reader sees on a body line: link targets and inline HTML comments dropped. */
function visibleText(line: string): string {
  return line.replace(/<!--.*?-->/g, '').replace(/\]\([^)]*\)/g, ']');
}

/** The prose violations of one guide markdown file, by its raw source lines. */
export function markdownViolations(raw: string, file: string): ProseViolation[] {
  const lines = raw.split(/\r?\n/);
  const found: ProseViolation[] = [];
  let fences = 0;
  let inComment = false;
  lines.forEach((source, index) => {
    if (source.trim() === '---' && fences < 2) {
      fences += 1;
      return;
    }
    const lineNo = index + 1;
    if (fences === 1) {
      const text = FRONTMATTER_TEXT.exec(source)?.[1];
      if (text !== undefined) found.push(...violationsIn(text, file, lineNo));
      return;
    }
    let line = source;
    if (inComment) {
      const end = line.indexOf('-->');
      if (end < 0) return;
      line = line.slice(end + 3);
      inComment = false;
    }
    const open = line.indexOf('<!--');
    if (open >= 0 && line.indexOf('-->', open) < 0) {
      line = line.slice(0, open);
      inComment = true;
    }
    found.push(...violationsIn(visibleText(line), file, lineNo));
  });
  return found;
}

function rotationNotes(node: unknown, into: Set<string>): Set<string> {
  if (Array.isArray(node)) {
    for (const item of node) rotationNotes(item, into);
  } else if (typeof node === 'object' && node !== null) {
    for (const [key, value] of Object.entries(node)) {
      if (key === 'notes' && typeof value === 'string') into.add(value);
      else rotationNotes(value, into);
    }
  }
  return into;
}

const NOTES_LINE = /^\s*"notes":\s*("(?:[^"\\]|\\.)*")\s*,?\s*$/;

/** The prose violations of the rendered notes (everything under `rotation`) in one APL file. */
export function aplViolations(raw: string, file: string): ProseViolation[] {
  const notes = rotationNotes((JSON.parse(raw) as { rotation?: unknown }).rotation, new Set());
  const found: ProseViolation[] = [];
  raw.split(/\r?\n/).forEach((source, index) => {
    const literal = NOTES_LINE.exec(source)?.[1];
    if (literal === undefined) return;
    const note = JSON.parse(literal) as string;
    if (notes.has(note)) found.push(...violationsIn(note, file, index + 1));
  });
  return found;
}

function classSlugs(): string[] {
  return readdirSync(guidesRoot, { withFileTypes: true })
    .filter((entry) => entry.isDirectory())
    .map((entry) => entry.name)
    .sort();
}

/** Every prose violation a class's guide pages render. */
export function classViolations(classSlug: string): ProseViolation[] {
  const dir = join(guidesRoot, classSlug);
  const found: ProseViolation[] = [];
  for (const name of readdirSync(dir)
    .filter((file) => file.endsWith('.md'))
    .sort()) {
    const path = join(dir, name);
    const raw = readFileSync(path, 'utf8');
    found.push(...markdownViolations(raw, relative(repoRoot, path)));
    if (name === 'index.md') continue;
    const apl = join(aplRoot, `${classSlug}-${name.slice(0, -'.md'.length)}.json`);
    if (existsSync(apl)) found.push(...aplViolations(readFileSync(apl, 'utf8'), relative(repoRoot, apl)));
  }
  return found;
}

function describeViolation(v: ProseViolation): string {
  return `${v.file}:${v.line}: "${v.phrase}" (${v.rule})`;
}

function reportMarkdown(slugs: readonly string[]): string {
  const rules = Object.keys(BANNED_PROSE) as (keyof typeof BANNED_PROSE)[];
  const rows = slugs.map((slug) => {
    const found = classViolations(slug);
    const counts = rules.map((rule) => found.filter((v) => v.rule === rule).length);
    return `| ${slug} | ${ENFORCED_CLASSES.includes(slug) ? 'yes' : 'no'} | ${counts.join(' | ')} | ${found.length} |`;
  });
  return [
    '# Guide prose baseline',
    '',
    'Violations of web/src/content/guides/_prose.test.ts per class (guide markdown plus rendered APL notes).',
    'Regenerate: `cd web && GUIDE_PROSE_REPORT=<path> npx vitest run src/content/guides/_prose.test.ts`.',
    '',
    `| Class | Enforced | ${rules.join(' | ')} | Total |`,
    `| --- | --- | ${rules.map(() => '---').join(' | ')} | --- |`,
    ...rows,
    '',
  ].join('\n');
}

describe('guide prose rules', () => {
  it('flags dashes, site-meta phrases and AI tells with the phrase that matched', () => {
    const found = violationsIn("This site's simulator — it's worth noting — is crucial.", 'x.md', 3);
    expect(found.map((v) => [v.rule, v.phrase])).toEqual([
      ['dashes', '—'],
      ['dashes', '—'],
      ['siteMeta', 'This site'],
      ['aiTells', "it's worth noting"],
      ['aiTells', 'crucial'],
    ]);
  });

  it('reads "band" as a level word but passes item names', () => {
    expect(
      violationsIn('At the 40 band, take it. Bands 20 to 30 too.', 'x.md', 1).map((v) => v.phrase),
    ).toEqual(['band', 'Bands']);
    expect(violationsIn('Wear Band of Accuria and the Band of Forced Concentration.', 'x.md', 1)).toEqual([]);
  });

  it('matches whole phrases only, and curly apostrophes too', () => {
    expect(violationsIn('Delvers and robustness and presetting pass.', 'x.md', 1)).toEqual([]);
    expect(violationsIn('It’s important to note this.', 'x.md', 1).map((v) => v.phrase)).toEqual([
      "It's important to note",
    ]);
    expect(
      violationsIn('Ranks 1-4 and Fire-Resist pass; a -- dash does not.', 'x.md', 1).map((v) => v.phrase),
    ).toEqual(['--']);
  });

  it('scans frontmatter text, the body, and skips comments, link targets and other keys', () => {
    const raw = [
      '---',
      "title: 'Arms — Warrior'",
      "build: 'FS1:pipeline'",
      'sources:',
      '  - label: The search results',
      '    url: https://example.com/preset',
      '---',
      '<!-- this site -->',
      'See [the planner](/planner?preset=raid).',
      '<!--',
      'the gate',
      '-->',
      'Robust play.',
    ].join('\n');
    expect(markdownViolations(raw, 'g.md').map(describeViolation)).toEqual([
      'g.md:2: "—" (dashes)',
      'g.md:5: "The search" (siteMeta)',
      'g.md:13: "Robust" (aiTells)',
    ]);
  });

  it('scans only the rotation notes of an APL file, on their own lines', () => {
    const raw = JSON.stringify(
      {
        notes: 'the search found this',
        rotation: { priorityList: [{ notes: 'Fire it on cooldown -- the gate holds.', action: {} }] },
      },
      null,
      1,
    );
    expect(aplViolations(raw, 'a.json').map(describeViolation)).toEqual([
      'a.json:6: "--" (dashes)',
      'a.json:6: "the gate" (siteMeta)',
    ]);
  });

  it('names only real guide classes in ENFORCED_CLASSES', () => {
    const slugs = classSlugs();
    for (const slug of ENFORCED_CLASSES) expect(slugs, `ENFORCED_CLASSES names "${slug}"`).toContain(slug);
  });

  for (const slug of ENFORCED_CLASSES) {
    it(`${slug} guides carry no dashes, site-meta phrases or AI tells`, () => {
      const found = classViolations(slug);
      expect(found.map(describeViolation), `${found.length} prose violations in ${slug}`).toEqual([]);
    });
  }
});

const reportPath = process.env.GUIDE_PROSE_REPORT;
if (reportPath !== undefined && reportPath !== '') {
  describe('guide prose report', () => {
    it('writes per-class violation counts', () => {
      writeFileSync(reportPath, reportMarkdown(classSlugs()));
    });
  });
}
