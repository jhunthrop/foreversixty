// web/src/lib/guides/overview.ts
// A guide's Overview is three `###` blocks of bullets (Strengths, Weaknesses, In a raid), and
// nothing before them: no opening "<Spec> is the <noun>: a, b and c" sentence (owner 2026-11-09:
// AI slop; the class landings' guard is lib/guides/class-landing.ts). This splits the section's
// Markdown into the blocks and fails loudly on text before the first one.

/** An opening "<Name> is the ..." line: the blurb the owner struck from every guide. */
export const IS_THE_BLURB = /^[A-Z][\w'’ ()-]* is the /;

export interface OverviewParts {
  /** The bullet lists as Markdown, one per block, `''` when the guide has no such block. */
  strengths: string;
  weaknesses: string;
  raid: string;
}

const BLOCKS: Readonly<Record<string, keyof OverviewParts>> = {
  strengths: 'strengths',
  weaknesses: 'weaknesses',
  'in a raid': 'raid',
};

/** Markdown inline marks off a sentence, as plain text. */
export function plainText(markdown: string): string {
  return markdown
    .replace(/\[([^\]]+)\]\([^)]*\)/g, '$1')
    .replace(/[*_`]/g, '')
    .replace(/\s+/g, ' ')
    .trim();
}

/** `sectionMarkdown` is the raw `## Overview` section, heading line included; `file` names the
 *  guide in the error. Throws on an opening "<Name> is the" line, and on any text before the
 *  first block when the section has blocks. */
export function splitOverview(sectionMarkdown: string, file = 'a guide'): OverviewParts {
  const lines = sectionMarkdown.split(/\r?\n/);
  const parts: OverviewParts = { strengths: '', weaknesses: '', raid: '' };
  const before: string[] = [];
  let current: keyof OverviewParts | undefined;
  let seenBlock = false;
  const bodies: Record<keyof OverviewParts, string[]> = {
    strengths: [],
    weaknesses: [],
    raid: [],
  };
  for (const line of lines.slice(1)) {
    const heading = /^###\s+(.+?)\s*$/.exec(line);
    if (heading !== null) {
      seenBlock = true;
      current = BLOCKS[heading[1]!.toLowerCase()];
      continue;
    }
    if (!seenBlock) before.push(line);
    else if (current !== undefined) bodies[current].push(line);
  }
  const opening = before.join(' ').trim();
  if (IS_THE_BLURB.test(opening)) throw new Error(`${file}: the Overview opens on a "<Name> is the" blurb`);
  if (seenBlock && opening !== '') throw new Error(`${file}: the Overview has text before its first block`);
  for (const key of Object.keys(bodies) as (keyof typeof bodies)[]) {
    parts[key] = bodies[key].join('\n').trim();
  }
  return parts;
}
