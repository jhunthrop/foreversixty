// web/src/lib/guides/overview.ts
// A guide's Overview is a summary sentence followed by three `###` blocks of bullets
// (Strengths, Weaknesses, In a raid). The page lays them out as cards and uses the sentence
// as the header's summary; this splits the section's Markdown into those parts.

export interface OverviewParts {
  /** The one-sentence summary, plain text. */
  summary: string;
  /** The bullet lists as Markdown, one per block, `''` when the guide has no such block. */
  strengths: string;
  weaknesses: string;
  raid: string;
}

const BLOCKS: Readonly<Record<string, keyof Omit<OverviewParts, 'summary'>>> = {
  strengths: 'strengths',
  weaknesses: 'weaknesses',
  'in a raid': 'raid',
};

/** Markdown inline marks off a sentence, for the header's plain-text summary. */
export function plainText(markdown: string): string {
  return markdown
    .replace(/\[([^\]]+)\]\([^)]*\)/g, '$1')
    .replace(/[*_`]/g, '')
    .replace(/\s+/g, ' ')
    .trim();
}

/** `sectionMarkdown` is the raw `## Overview` section, heading line included. */
export function splitOverview(sectionMarkdown: string): OverviewParts {
  const lines = sectionMarkdown.split(/\r?\n/);
  const parts: OverviewParts = { summary: '', strengths: '', weaknesses: '', raid: '' };
  const summary: string[] = [];
  let current: keyof Omit<OverviewParts, 'summary'> | undefined;
  const bodies: Record<keyof Omit<OverviewParts, 'summary'>, string[]> = {
    strengths: [],
    weaknesses: [],
    raid: [],
  };
  for (const line of lines.slice(1)) {
    const heading = /^###\s+(.+?)\s*$/.exec(line);
    if (heading !== null) {
      current = BLOCKS[heading[1]!.toLowerCase()];
      continue;
    }
    if (current === undefined) summary.push(line);
    else bodies[current].push(line);
  }
  parts.summary = plainText(summary.join(' '));
  for (const key of Object.keys(bodies) as (keyof typeof bodies)[]) {
    parts[key] = bodies[key].join('\n').trim();
  }
  return parts;
}
