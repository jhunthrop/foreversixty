import type { z } from 'astro/zod';
import type { sourceKinds, sourceSchema } from '../content.config';

export type SourceKind = (typeof sourceKinds)[number];
export type Source = z.infer<typeof sourceSchema>;

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
