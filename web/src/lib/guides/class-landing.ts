// web/src/lib/guides/class-landing.ts
// The class landing's content contract (spec design/specs/2026-11-02-class-landing.md §6,
// §10): a class's index.md is `##` sections drawn from a fixed heading vocabulary and nothing
// before them. There is no intro paragraph (owner 2026-11-09: a "<Class> is the <noun>: a, b and c"
// blurb tells a veteran what they know). The page lays each part out in its own region, so a file
// that strays from the shape fails the build rather than rendering a section nobody designed.
import { guidesCopy } from './copy';

/** The section headings a class landing may carry, in the order they must appear. */
export const CLASS_LANDING_HEADINGS = [
  'What changed from 1.12',
  'Races',
  'In a raid',
  'Raids and leveling',
] as const;

export type ClassLandingHeading = (typeof CLASS_LANDING_HEADINGS)[number];

/** The headings that render as a side panel; the rest render as the ruled fact list. */
export const CLASS_LANDING_PANELS: readonly ClassLandingHeading[] = [
  'Races',
  'In a raid',
  'Raids and leveling',
];

export interface ClassLandingSection {
  heading: ClassLandingHeading;
  /** The section's Markdown, heading line excluded. */
  markdown: string;
}

export interface ClassLanding {
  sections: ClassLandingSection[];
}

/** A spec card's line under its name: the role for a healer or a tank only (every class has a
 *  damage spec, so "DPS" under each one tells a veteran nothing), then the guide build's point
 *  split, "5/13/33", when its code decodes (h-audit 2026-11-09). */
export function specCardLine(
  role: 'dps' | 'healer' | 'tank',
  trees: readonly { points: number }[] = [],
): string | undefined {
  const parts = [
    role === 'dps' ? undefined : guidesCopy.roleLabel(role),
    trees.length === 0 ? undefined : trees.map((tree) => tree.points).join('/'),
  ].filter((part): part is string => part !== undefined);
  return parts.length === 0 ? undefined : parts.join(' · ');
}

const HEADING = /^##[ \t]+(.+?)[ \t]*$/gm;

function isHeading(text: string): text is ClassLandingHeading {
  return (CLASS_LANDING_HEADINGS as readonly string[]).includes(text);
}

function parseSections(markdown: string, file: string): ClassLandingSection[] {
  const matches = [...markdown.matchAll(HEADING)];
  const sections = matches.map((match, index): ClassLandingSection => {
    const heading = match[1]!;
    if (!isHeading(heading)) throw new Error(`${file}: "${heading}" is not a class landing heading`);
    const start = (match.index ?? 0) + match[0].length;
    const end = matches[index + 1]?.index ?? markdown.length;
    return { heading, markdown: markdown.slice(start, end).trim() };
  });
  const order = sections.map((s) => CLASS_LANDING_HEADINGS.indexOf(s.heading));
  if (order.some((rank, index) => index > 0 && rank <= order[index - 1]!)) {
    throw new Error(
      `${file}: headings out of order or repeated (${sections.map((s) => s.heading).join(', ')})`,
    );
  }
  // Races may be the heading alone: the page draws the class's race list under it (GuideRaceList).
  if (sections.some((s) => s.markdown === '' && s.heading !== 'Races'))
    throw new Error(`${file}: a section has no content`);
  return sections;
}

/** `body` is the index.md source minus frontmatter; `file` names it in every error. */
export function splitClassLanding(body: string, file: string): ClassLanding {
  const main = body.trim();
  const firstHeading = main.search(/^##[ \t]/m);
  const lead = firstHeading < 0 ? main : main.slice(0, firstHeading);
  if (lead.trim() !== '')
    throw new Error(`${file}: a class landing has no intro paragraph before its first heading`);
  return { sections: firstHeading < 0 ? [] : parseSections(main.slice(firstHeading), file) };
}

/** A paragraph that is only bold text ("**Race picks.**") labels the list under it: the Prose
 *  primitive's `glabel`. */
export function proseLabels(html: string): string {
  return html.replace(/<p><strong>([^<]+)<\/strong><\/p>/g, '<p class="glabel">$1</p>');
}
