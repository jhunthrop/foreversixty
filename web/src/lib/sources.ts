import type { z } from 'astro/zod';
import type { sourceKinds, sourceSchema } from '../content.config';

export type SourceKind = (typeof sourceKinds)[number];
export type Source = z.infer<typeof sourceSchema>;

/**
 * What a tool card's status pill claims about its data: either this site maintains it or
 * the community does. A subset of SourceKind so both reuse the same `pill-*` classes.
 */
export type ToolStatusKind = Extract<SourceKind, 'site' | 'community'>;

const labels: Record<SourceKind, string> = {
  blizzard: 'Blizzard',
  datamined: 'Datamined',
  community: 'Community',
  site: 'This site',
};

export function pillClassFor(kind: SourceKind): string {
  return `pill-${kind}`;
}

export function pillLabelFor(kind: SourceKind): string {
  return labels[kind];
}

/** The publishers a source row's kind pill already names: "Blizzard" for a Blizzard source, and
 *  this site for its own datamined tables and figures. */
const PILL_PUBLISHERS: readonly string[] = ['Blizzard, ', 'Forever Sixty, '];

/** A source row's link text without the publisher its pill already says: "Blizzard, Class Deep
 *  Dives: Hunter and Druid" beside a Blizzard pill is "Class Deep Dives: Hunter and Druid" (h-audit
 *  2026-11-09). A community source keeps its site's name: the pill says only "Community". */
export function sourceTitle(label: string): string {
  const prefix = PILL_PUBLISHERS.find((publisher) => label.startsWith(publisher));
  if (prefix === undefined) return label;
  const rest = label.slice(prefix.length);
  return rest.charAt(0).toUpperCase() + rest.slice(1);
}
