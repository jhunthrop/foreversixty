<!-- web/src/components/sim/SimAurasCard.svelte -->
<!-- Buffs and debuffs (spec 2026-11-04-sim-results 3.10, build ruling R2): uptime on you and on the
     target, with the same full buff list and the same "No effect on your damage" rule as Settings. -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { loadWeights } from '../../lib/planner/load';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { auraCard } from '../../lib/sim/aura-card';
  import { damageSchools, type SpecWeights } from '../../lib/sim/buff-effect';
  import { CATALOGUE } from '../../lib/sim/buffs';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { aurasCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { loadSimBuffs, type SimBuffRow } from '../../lib/sim/sim-buffs';
  import type { SimResult } from '../../lib/sim/types';
  import AuraTable from '../report/AuraTable.svelte';
  import Card from '../ui/Card.svelte';

  let { result, names }: { result: SimResult; names: ActionNames | null } = $props();

  const build = activeBuild.build;
  const FALLBACK_TARGET_IDS: ReadonlySet<string> = new Set(
    CATALOGUE.filter((row) => row.group === 'debuffs').map((row) => row.id),
  );

  const tables = createCardTables(
    () => result,
    () => names,
    () => result.summary.auras.map((track) => track.name),
  );

  let buffs = $state.raw<Readonly<Record<string, SimBuffRow>>>({});
  let weights = $state.raw<SpecWeights | null>(null);
  $effect(() => {
    const { spec } = result.request;
    loadSimBuffs(build).then(
      (file) => (buffs = file.entries),
      (error: unknown) => console.error('sim auras: the buff table did not load', error),
    );
    loadWeights(build).then(
      (rows) => (weights = rows.find((row) => row.spec === spec)?.weights ?? null),
      (error: unknown) => console.error('sim auras: the stat weights did not load', error),
    );
  });

  const card = $derived(
    auraCard({
      result,
      buffs,
      weights,
      schools: damageSchools(result.summary.damage_done),
      names: tables.names,
      fallbackTargetIds: FALLBACK_TARGET_IDS,
    }),
  );
</script>

<Card
  glyph="spark"
  label={copy.label}
  title={copy.title}
  meta={copy.meta}
  body="none"
  class="mx-[18px] md:mx-0"
  testid="sim-auras-card"
>
  <div class="border-line-soft divide-line-soft grid grid-cols-1 divide-y border-t lg:grid-cols-2 lg:divide-x lg:divide-y-0">
    <AuraTable
      variant="bars"
      kind="BUFF"
      tracks={result.summary.auras}
      durationMs={result.summary.duration_ms}
      bars={card.player}
      heading={copy.onYou}
      {tables}
    />
    <AuraTable
      variant="bars"
      kind="DEBUFF"
      tracks={result.summary.auras}
      durationMs={result.summary.duration_ms}
      bars={card.target}
      heading={copy.onTarget}
      {tables}
      note={card.targetFromRequest ? copy.targetNeedsNewerRun : undefined}
    />
  </div>
</Card>
