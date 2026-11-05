// web/src/lib/sim/card-tables.svelte.ts
// What each of the saved run's lazy cards reads to draw an action: the names with any shard-only
// ids filled in, its icon and rank, and whether the build carries its game tooltip. The page's own
// loads are shared through the planner's fetch cache, so every card asking costs one request.
import activeBuild from '../../data/active-build.json';
import { withShardNames, type ActionNames } from './action-names';
import { actionVisuals, loadVisualTables, pendingVisual, type ActionVisuals, type VisualTables } from './action-visuals';
import { NO_SPELL_INFO, loadSpellInfo, type SpellInfo } from './spell-info';
import type { SimResult } from './types';

export interface CardTables {
  readonly names: ActionNames | null;
  readonly visuals: ActionVisuals;
  readonly spells: SpellInfo;
}

export function createCardTables(
  result: () => SimResult,
  names: () => ActionNames | null,
  keys: () => readonly string[],
): CardTables {
  let tables = $state.raw<VisualTables | null>(null);
  let spells = $state.raw<SpellInfo>(NO_SPELL_INFO);
  let shardNames = $state.raw<ActionNames | null>(null);

  const quietly = (what: string) => (error: unknown) => console.error(`sim result: ${what} did not load`, error);
  $effect(() => {
    const current = result();
    loadVisualTables(current).then((loaded) => (tables = loaded), quietly('the icons'));
    loadSpellInfo(activeBuild.build, current.request.character.class).then(
      (loaded) => (spells = loaded),
      quietly('the spell data'),
    );
  });
  $effect(() => {
    const parent = names();
    const wanted = keys();
    if (parent === null) return;
    shardNames = parent;
    withShardNames(parent, wanted).then((loaded) => (shardNames = loaded), quietly('the name shards'));
  });

  const liveNames = $derived(shardNames ?? names());
  const visuals = $derived(tables === null ? pendingVisual : actionVisuals(tables, liveNames));
  return {
    get names() {
      return liveNames;
    },
    get visuals() {
      return visuals;
    },
    get spells() {
      return spells;
    },
  };
}
