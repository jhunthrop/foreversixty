// web/src/lib/sim/bis-set-rail.ts
// The /sim landing's "Sim a best-in-slot set" rail (spec 2026-11-06-sim §4.2, ruling S6): one row
// per class with a damage spec the sim can rerun, in the client's class order, each spec linking
// to its level 60 band's "View ranking sim" (the tier list's own link, lib/tiers/tier-load.ts).
// No figures: the tier list owns the ranking (one job per page). Read at build time (node fs).
import { bandEntry, loadBisFile, readSpecCatalog } from '../bis/load';
import type { BisBand, BisFile, SpecCatalogEntry } from '../bis/types';
import { classRows } from '../planner/reference';
import { tierRankingSimHref } from '../tiers/tier-load';

/** The band the published ranking sims run: level 60, raid-ready (the tier list's band). */
const RAIL_BAND = 60;
/** The faction whose set the link carries; the race only changes the racial, not the set. */
const RAIL_FACTION = 'alliance';

export interface BisSetSpec {
  name: string;
  href: string;
}

export interface BisSetClass {
  slug: string;
  name: string;
  specs: BisSetSpec[];
}

export interface BisSetRailDeps {
  catalog: readonly SpecCatalogEntry[];
  classes: readonly { slug: string; name: string }[];
  loadFile: (spec: string) => BisFile | null;
  simHref: (entry: SpecCatalogEntry, band: BisBand) => string | undefined;
}

export function buildBisSetRail(deps: BisSetRailDeps): BisSetClass[] {
  return deps.classes.flatMap((klass) => {
    const specs = deps.catalog
      .filter((entry) => entry.class_slug === klass.slug && entry.role === 'dps')
      .flatMap((entry) => {
        const file = deps.loadFile(entry.spec);
        const band = file === null ? undefined : bandEntry(file, RAIL_BAND, RAIL_FACTION);
        const href = band === undefined ? undefined : deps.simHref(entry, band);
        return href === undefined ? [] : [{ name: entry.name, href }];
      });
    return specs.length === 0 ? [] : [{ slug: klass.slug, name: klass.name, specs }];
  });
}

export function loadBisSetRail(build: string): BisSetClass[] {
  return buildBisSetRail({
    catalog: readSpecCatalog(),
    classes: classRows,
    loadFile: (spec) => loadBisFile(spec, build),
    simHref: tierRankingSimHref(build),
  });
}
