// web/src/lib/home/floor-load.ts
// Build-time inputs for the feed's floor (lib/home/floor.ts): each spec's published BiS file
// date, the guides' `updated`, the site changelog and the raid opening. Node only (reads the BiS
// files from disk); index.astro passes the content collections in.
import activeBuild from '../../data/active-build.json';
import dates from '../../data/dates.json';
import { loadBisFile } from '../bis/load';
import { SPECS } from '../sim/specs';
import type { FloorInputs } from './floor';

export interface GuideEntry {
  classSlug: string;
  spec?: string;
  updated: Date;
}

export interface ChangelogEntry {
  id: string;
  title: string;
  date: Date;
}

/** The raid opening while it is still ahead of `now`. */
export function raidsOpenIso(now: Date): string | null {
  const raids = dates.find((entry) => 'role' in entry && entry.role === 'raids');
  if (raids === undefined) return null;
  const at = new Date(`${raids.iso}T00:00:00Z`);
  return at.getTime() > now.getTime() ? at.toISOString() : null;
}

export function floorInputs(
  guides: readonly GuideEntry[],
  changelog: readonly ChangelogEntry[],
  now: Date,
  build: string = activeBuild.build,
): FloorInputs {
  const iconOf = new Map(SPECS.map((spec) => [spec.spec, spec.icon]));
  const bis = SPECS.flatMap((spec) => {
    const file = loadBisFile(spec.spec, build);
    if (file === null) return [];
    return [
      {
        spec: spec.spec,
        generatedAt: file.generated_at,
        href: `/bis/${spec.class_slug}/${spec.spec_slug}`,
        icon: spec.icon,
      },
    ];
  });
  const specGuides = guides.flatMap((guide) => {
    if (guide.spec === undefined) return [];
    const key = `${guide.classSlug}-${guide.spec}`;
    return [
      {
        spec: key,
        updated: guide.updated.toISOString(),
        href: `/guides/${guide.classSlug}/${guide.spec}`,
        icon: iconOf.get(key) ?? '',
      },
    ];
  });
  return {
    raidsOpen: raidsOpenIso(now),
    bis,
    guides: specGuides,
    changelog: changelog.map((entry) => ({
      id: entry.id,
      title: entry.title,
      date: entry.date.toISOString(),
    })),
  };
}
