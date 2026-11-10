// web/src/lib/home/raid-facts-load.ts
// Build-time inputs for the raid plaque's facts (lib/home/raid-facts.ts): the build's loot table
// and each class's item ids. Node only; index.astro calls it.
import activeBuild from '../../data/active-build.json';
import classes from '../../data/classes.json';
import { itemDetails, loadLootFile } from '../bis/load';
import { raidFacts, type RaidFact } from './raid-facts';

export function loadRaidFacts(build: string = activeBuild.build): RaidFact[] {
  const loot = loadLootFile(build);
  const classItems = new Map(
    classes.map((entry) => [entry.slug, new Set(itemDetails(build, entry.slug).keys())] as const),
  );
  return raidFacts(loot.sources, classItems);
}
