// web/src/lib/app-data/changelog-file.ts
// /data/changelog.json (Wave 3 contract W3-C1): the site's changelog, newest first, for the
// apps' What changed list. Each entry's body is split into plain-text blocks (a paragraph, a
// heading, a list item) the guide reader draws without a Markdown renderer.
import { readFileSync, readdirSync } from 'node:fs';
import path from 'node:path';
import { changelogSchema } from '../../content.config';
import { repoRoot } from '../data/build-files';
import { plainText } from '../guides/overview';
import { isoDay } from './guide-file';
import { splitFrontmatter } from './guide-content';

export type ChangelogBlockType = 'p' | 'h' | 'li';

export interface ChangelogBlock {
  type: ChangelogBlockType;
  text: string;
}

export interface AppChangelogEntry {
  /** The file's name without `.md`, unique and stable. */
  id: string;
  date: string;
  title: string;
  /** `site` for a change to this site; `blizzard`, `datamined` or `community` for the game. */
  kind: string;
  blocks: ChangelogBlock[];
  sources: { label: string; url: string; kind: string }[];
}

export interface AppChangelog {
  entries: AppChangelogEntry[];
}

const CHANGELOG_DIR = 'web/src/content/changelog';

/** A Markdown body as blocks: `#` lines are headings, `- ` lines items, the rest paragraphs. */
export function markdownBlocks(body: string): ChangelogBlock[] {
  const blocks: ChangelogBlock[] = [];
  let paragraph: string[] = [];
  const flush = (): void => {
    const text = plainText(paragraph.join(' '));
    if (text !== '') blocks.push({ type: 'p', text });
    paragraph = [];
  };
  for (const line of body.split(/\r?\n/)) {
    const heading = /^#{1,6}\s+(.+)$/.exec(line);
    const item = /^\s*[-*]\s+(.+)$/.exec(line);
    if (heading !== null || item !== null || line.trim() === '') flush();
    if (heading !== null) blocks.push({ type: 'h', text: plainText(heading[1]!) });
    else if (item !== null) blocks.push({ type: 'li', text: plainText(item[1]!) });
    else if (line.trim() !== '') paragraph.push(line);
  }
  flush();
  return blocks;
}

function readEntry(dir: string, name: string): AppChangelogEntry {
  const id = name.replace(/\.md$/, '');
  const { frontmatter, body } = splitFrontmatter(readFileSync(path.join(dir, name), 'utf8'));
  const parsed = changelogSchema.safeParse(frontmatter);
  if (!parsed.success) throw new Error(`changelog: ${id} fails the changelog schema: ${parsed.error.message}`);
  const data = parsed.data;
  return {
    id,
    date: isoDay(data.date),
    title: data.title,
    kind: data.kind,
    blocks: markdownBlocks(body),
    sources: data.sources.map((link) => ({ label: link.label, url: link.url, kind: link.kind })),
  };
}

/** Every changelog entry, newest first (a tie keeps the later file first). */
export function buildChangelog(root: string = repoRoot()): AppChangelog {
  const dir = path.join(root, CHANGELOG_DIR);
  const entries = readdirSync(dir)
    .filter((name) => name.endsWith('.md'))
    .map((name) => readEntry(dir, name))
    .sort((a, b) => b.date.localeCompare(a.date) || b.id.localeCompare(a.id));
  return { entries };
}
