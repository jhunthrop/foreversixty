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
