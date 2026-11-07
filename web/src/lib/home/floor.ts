// web/src/lib/home/floor.ts
// The feed's floor (design/specs/2026-11-07-home-b.md F1): real dated site events, built at build
// time from what the site itself publishes, so a quiet filter never looks dead and never shows
// invented player activity. Pure: lib/home/floor-load.ts gathers the inputs.
import type { FeedRow } from './feed';
import { FIRST_RAIDS, specName } from './home-b-view';

export interface FloorInputs {
  /** The first raid opening, while it is still ahead (dates.json role `raids`). */
  raidsOpen: string | null;
  /** Each spec's published best-in-slot file. */
  bis: readonly { spec: string; generatedAt: string; href: string; icon: string }[];
  /** Each spec guide's `updated` date. */
  guides: readonly { spec: string; updated: string; href: string; icon: string }[];
  /** The site changelog. */
  changelog: readonly { id: string; title: string; date: string }[];
}

/** Rows that are about no one spec, kept newest first. */
const GENERAL_KEPT = 10;
/** Rows kept per spec (its BiS refresh and its guide update). */
const PER_SPEC_KEPT = 2;
const CHANGELOG_ICON = 'inv_misc_book_09';

function row(partial: Omit<FeedRow, 'crest' | 'quality' | 'date_only'> & Partial<FeedRow>): FeedRow {
  return { crest: null, quality: null, date_only: false, ...partial };
}

function raidRow(at: string): FeedRow {
  return row({
    id: 'floor:raids',
    kind: 'site',
    at,
    tag: 'patch',
    icon: FIRST_RAIDS[2].icon,
    href: '/guides',
    date_only: true,
    segments: [
      { text: 'First raids open', style: 'strong' },
      { text: `: ${FIRST_RAIDS.map((raid) => raid.name).join(', ')}` },
    ],
  });
}

function specRows(inputs: FloorInputs): FeedRow[] {
  const bis = inputs.bis.map((file) =>
    row({
      id: `floor:bis:${file.spec}`,
      kind: 'site',
      at: file.generatedAt,
      tag: 'best_in_slot',
      icon: file.icon,
      href: file.href,
      spec: file.spec,
      segments: [
        { text: 'Best in slot refreshed', style: 'strong' },
        { text: ` for ${specName(file.spec)}` },
      ],
    }),
  );
  const guides = inputs.guides.map((guide) =>
    row({
      id: `floor:guide:${guide.spec}`,
      kind: 'site',
      at: guide.updated,
      tag: 'guides',
      icon: guide.icon,
      href: guide.href,
      spec: guide.spec,
      segments: [{ text: `${specName(guide.spec)} guide`, style: 'strong' }, { text: ' updated' }],
    }),
  );
  return [...bis, ...guides];
}

const newestFirst = (a: FeedRow, b: FeedRow) => b.at.localeCompare(a.at);

/** The floor, in the order it fills: the raid opening, then the newest events. Every spec keeps
 *  its own newest rows so Your spec can floor; the rest keeps the newest few. */
export function buildFloor(inputs: FloorInputs): FeedRow[] {
  const changelog = inputs.changelog.map((entry) =>
    row({
      id: `floor:changelog:${entry.id}`,
      kind: 'site',
      at: entry.date,
      tag: 'patch',
      icon: CHANGELOG_ICON,
      href: `/changelog#${entry.id}`,
      segments: [{ text: entry.title }],
    }),
  );
  const perSpec = new Map<string, FeedRow[]>();
  for (const specRow of specRows(inputs).sort(newestFirst)) {
    const list = perSpec.get(specRow.spec ?? '') ?? [];
    if (list.length < PER_SPEC_KEPT) perSpec.set(specRow.spec ?? '', [...list, specRow]);
  }
  const events = [...changelog, ...[...perSpec.values()].flat()].sort(newestFirst);
  const general = events.slice(0, GENERAL_KEPT);
  const specOnly = events.slice(GENERAL_KEPT).filter((event) => event.spec !== undefined);
  const lead = inputs.raidsOpen === null ? [] : [raidRow(inputs.raidsOpen)];
  return [...lead, ...general, ...specOnly];
}
