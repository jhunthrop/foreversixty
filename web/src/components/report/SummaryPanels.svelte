<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The Summary tab's ability panels under the chart (logs spec 2026-11-07 6.2; owner review
     2026-11-09): `metric="done"` Damage done by ability, every ability the scope's players hit
     with; `metric="taken"` Damage taken by ability, every ability that hit them. Each a StatTable
     (spell icon, name, a Meter, the share of the scope's total, Per sec, Amount), the top eight. -->
<script lang="ts">
  import EmptyLine from './EmptyLine.svelte';
  import { formatAmount } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { Actor, Summary } from '../../lib/report/types';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import SpellIcon from './SpellIcon.svelte';

  let {
    metric,
    summary,
    approximate = false,
    durationMs,
    players,
    spellIcon,
  }: {
    metric: 'done' | 'taken';
    summary: Summary;
    durationMs: number;
    players: ReadonlySet<string>;
    /** A brushed window: the by-ability split is prorated from the whole fight's, and says so. */
    approximate?: boolean;
    spellIcon: SpellIconOf;
  } = $props();

  const ROWS = 8;

  interface AbilityTotal {
    name: string;
    spellId: number;
    school?: number;
    total: number;
  }

  /** Every ability that hit a player, summed over the players it hit, largest first. */
  function abilityTotals(table: readonly Actor[]): AbilityTotal[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, AbilityTotal>();
    for (const actor of table) {
      if (players.size > 0 && !players.has(actor.guid)) continue;
      for (const ability of actor.abilities) {
        const key = ability.name === '' ? 'Melee' : ability.name;
        const found = totals.get(key);
        if (found === undefined)
          totals.set(key, {
            name: key,
            spellId: ability.spell_id,
            school: ability.school,
            total: ability.effective,
          });
        else found.total += ability.effective;
      }
    }
    return [...totals.values()].filter((row) => row.total > 0).sort((a, b) => b.total - a.total);
  }
  const TITLE = { done: 'Damage done by ability', taken: 'Damage taken by ability' } as const;
  const EMPTY = {
    done: 'Nobody did damage in this window.',
    taken: 'Nobody took damage in this window.',
  } as const;
  const byAbility = $derived.by(() => {
    const rows = abilityTotals(metric === 'done' ? summary.damage_done : summary.damage_taken);
    const total = rows.reduce((sum, row) => sum + row.total, 0);
    return rows.slice(0, ROWS).map((row) => ({ ...row, share: total === 0 ? 0 : row.total / total }));
  });
  const seconds = $derived(Math.max(durationMs, 1) / 1000);
  const COLUMNS = [
    { id: 'share', label: 'Share' },
    { id: 'rate', label: 'Per sec' },
    { id: 'amount', label: 'Amount' },
  ] as const;
</script>

<Section title={TITLE[metric]} panel flush testid={metric === 'done' ? 'panel-done' : 'panel-taken'}>
  {#if byAbility.length === 0}
    <EmptyLine text={EMPTY[metric]} />
  {:else}
    <StatTable columns={COLUMNS} nameHead="Ability" label={TITLE[metric]}>
      {#each byAbility as row (row.name)}
        <StatRow
          figures={[
            { value: `${(row.share * 100).toFixed(1)}%` },
            { value: formatRate(row.total / seconds) },
            { value: `${approximate ? '~' : ''}${formatAmount(row.total)}` },
          ]}
          phone={`${formatAmount(row.total)} damage`}
        >
          {#snippet icon()}<SpellIcon url={spellIcon(row.spellId, row.name, row.school)} />{/snippet}
          {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
          {#snippet meter()}<Meter
              value={row.share}
              tone={metric === 'done' ? 'accent' : 'muted'}
            />{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  {/if}
</Section>
