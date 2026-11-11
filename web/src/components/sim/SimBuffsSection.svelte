<!-- web/src/components/sim/SimBuffsSection.svelte -->
<!-- Buffs and debuffs (sim spec 2026-11-06 §4.5, §4.9; spec 2026-11-04-sim-results 3.10, build
     ruling R2): two columns, On you and On the target. What came and went is an open StatTable of
     uptime and times applied; what stayed up is an IconList, the ones with no effect on this spec's
     damage dimmed and last (lib/sim/buff-effect.ts, the same rule as This run). -->
<script lang="ts">
  import activeBuild from '../../data/active-build.json';
  import { loadWeights } from '../../lib/planner/load';
  import type { ActionNames } from '../../lib/sim/action-names';
  import { actionTip, resolveIconUrl } from '../../lib/sim/action-mark';
  import { auraCard, type AuraColumn, type AuraEntry } from '../../lib/sim/aura-card';
  import { damageSchools, type SpecWeights } from '../../lib/sim/buff-effect';
  import { CATALOGUE } from '../../lib/sim/buffs';
  import { createCardTables } from '../../lib/sim/card-tables.svelte';
  import { aurasCardCopy as copy } from '../../lib/sim/input-cards-copy';
  import { oneDecimal } from '../../lib/format/number';
  import { loadSimBuffs, type SimBuffRow } from '../../lib/sim/sim-buffs';
  import type { SimResult } from '../../lib/sim/types';
  import type { IconListItem } from '../../lib/ui/facts';
  import IconList from '../ui/IconList.svelte';
  import Section from '../ui/Section.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import ActionStatRow from './ActionStatRow.svelte';

  let { result, names }: { result: SimResult; names: ActionNames | null } = $props();

  const build = activeBuild.build;
  const FALLBACK_TARGET_IDS: ReadonlySet<string> = new Set(
    CATALOGUE.filter((row) => row.group === 'debuffs').map((row) => row.id),
  );
  const COLUMNS = [
    { id: 'uptime', label: '' },
    { id: 'applied', label: '' },
  ];
  const UP_ALL_FIGHT_ICON = 20;

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

  function iconItem(entry: AuraEntry, dim: boolean): IconListItem {
    return {
      id: entry.key,
      name: entry.stacks > 1 ? `${entry.name} ${copy.stacks(entry.stacks)}` : entry.name,
      icon: resolveIconUrl(entry.actionKey, tables.visuals, entry.buffIcon),
      dim,
      tip: actionTip(entry.actionKey, (id) => tables.spells.fact(id) !== undefined).tip,
    };
  }

  /** What stayed up, the ones that help first and the no-effect ones dimmed after them. */
  const upAllFight = (column: AuraColumn): IconListItem[] => [
    ...column.upAllFight.map((entry) => iconItem(entry, false)),
    ...column.noEffect.map((entry) => iconItem(entry, true)),
  ];
</script>

{#snippet side(column: AuraColumn, heading: string, key: 'player' | 'target', note?: string)}
  {@const stayed = upAllFight(column)}
  <div class="min-w-0" data-testid={`aura-column-${key}`}>
    <h3 class="text-label text-muted pt-4 pb-1 font-bold">
      {heading} · {copy.comeAndGo(column.bars.length)}
    </h3>
    {#if note !== undefined}<p class="text-muted text-[13px]" data-testid="aura-column-note">{note}</p>{/if}
    {#if column.bars.length > 0}
      <StatTable columns={COLUMNS} open label={heading}>
        {#each column.bars as entry (entry.key)}
          <ActionStatRow
            action={entry.actionKey}
            label={entry.name}
            {tables}
            fallbackIcon={entry.buffIcon}
            figures={[
              { value: `${oneDecimal(entry.uptimePct)}%` },
              { value: copy.applied(entry.applications) },
            ]}
            fraction={entry.uptimePct / 100}
            testid="aura-bar-row"
          />
        {/each}
      </StatTable>
    {/if}
    {#if stayed.length > 0}
      <IconList
        label={copy.upAllFight}
        items={stayed}
        iconSize={UP_ALL_FIGHT_ICON}
        testid={`aura-up-all-fight-${key}`}
      />
    {/if}
  </div>
{/snippet}

<Section panel title={copy.title} testid="sim-auras-card">
  <div class="grid grid-cols-1 gap-x-[var(--gap)] lg:grid-cols-2">
    {@render side(card.player, copy.onYou, 'player')}
    {@render side(
      card.target,
      copy.onTarget,
      'target',
      card.targetFromRequest ? copy.targetNeedsNewerRun : undefined,
    )}
  </div>
</Section>
