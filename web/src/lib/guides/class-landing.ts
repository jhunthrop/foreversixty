// web/src/lib/guides/class-landing.ts
// The class landing's content contract (spec design/specs/2026-11-02-class-landing.md §6,
// §10): a class's index.md is one plain identity paragraph, then `##` sections drawn from a
// fixed heading vocabulary. The page lays each part out in its own region, so a file that strays from the shape fails the build
// rather than rendering a section nobody designed.
import { specCardLine } from './overview';
import { splitSpecSections } from './sections';

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

/** The identity line is the header's summary: one or two plain sentences. */
export const IDENTITY_MAX_CHARS = 300;

export interface ClassLandingSection {
  heading: ClassLandingHeading;
  /** The section's Markdown, heading line excluded. */
  markdown: string;
}

export interface ClassLanding {
  identity: string;
  sections: ClassLandingSection[];
}

const HEADING = /^##[ \t]+(.+?)[ \t]*$/gm;
const INLINE_MARKUP = /[*_`[\]]/;

function isHeading(text: string): text is ClassLandingHeading {
  return (CLASS_LANDING_HEADINGS as readonly string[]).includes(text);
}

function parseIdentity(lead: string, file: string): string {
  const identity = lead.replace(/\s+/g, ' ').trim();
  if (identity === '') throw new Error(`${file}: no identity paragraph before the first heading`);
  if (/\n[ \t]*\n/.test(lead.trim())) throw new Error(`${file}: the identity is more than one paragraph`);
  if (INLINE_MARKUP.test(identity))
    throw new Error(`${file}: the identity carries Markdown; it renders as plain text`);
  if (identity.length > IDENTITY_MAX_CHARS) {
    throw new Error(`${file}: the identity is ${identity.length} characters, over ${IDENTITY_MAX_CHARS}`);
  }
  return identity;
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
  if (sections.some((s) => s.markdown === '')) throw new Error(`${file}: a section has no content`);
  return sections;
}

/** A spec card's line fits a 354px phone card in at most six lines. */
export const CARD_LINE_MAX_CHARS = 200;

/** The spec card's line from the spec guide's own body (frontmatter excluded): its Overview
 *  summary minus link-only sentences (`specCardLine`). */
export function cardLineFor(specBody: string, file: string): string {
  const overview = splitSpecSections(specBody).get('Overview');
  if (overview === undefined) throw new Error(`${file}: no ## Overview for the class landing card`);
  const line = specCardLine(overview);
  if (line === '') throw new Error(`${file}: the Overview has no summary for the class landing card`);
  if (line.length > CARD_LINE_MAX_CHARS) {
    throw new Error(`${file}: the card line is ${line.length} characters, over ${CARD_LINE_MAX_CHARS}`);
  }
  return line;
}

/** `body` is the index.md source minus frontmatter; `file` names it in every error. */
export function splitClassLanding(body: string, file: string): ClassLanding {
  const main = body.trim();
  const firstHeading = main.search(/^##[ \t]/m);
  const lead = firstHeading < 0 ? main : main.slice(0, firstHeading);
  return {
    identity: parseIdentity(lead, file),
    sections: firstHeading < 0 ? [] : parseSections(main.slice(firstHeading), file),
  };
}

/** A paragraph that is only bold text ("**Race picks.**") labels the list under it: the Prose
 *  primitive's `glabel`. */
export function proseLabels(html: string): string {
  return html.replace(/<p><strong>([^<]+)<\/strong><\/p>/g, '<p class="glabel">$1</p>');
}
