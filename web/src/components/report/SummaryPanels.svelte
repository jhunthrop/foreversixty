<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The Summary tab's panel under the chart (logs spec 2026-11-07 6.2; directions b1-report-summary):
     Damage taken by ability, a StatTable of each ability that hit a player in the scope (spell
     icon, name, a Meter, the share of the scope's damage taken, Per sec, Amount). -->
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
    summary,
    approximate = false,
    durationMs,
    players,
    spellIcon,
  }: {
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
          totals.set(key, { name: key, spellId: ability.spell_id, total: ability.effective });
        else found.total += ability.effective;
      }
    }
    return [...totals.values()].filter((row) => row.total > 0).sort((a, b) => b.total - a.total);
  }
  const takenByAbility = $derived.by(() => {
    const rows = abilityTotals(summary.damage_taken);
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

<Section title="Damage taken by ability" panel flush testid="panel-taken">
  {#if takenByAbility.length === 0}
    <EmptyLine text="Nobody took damage in this window." />
  {:else}
    <StatTable columns={COLUMNS} nameHead="Ability" label="Damage taken by ability">
      {#each takenByAbility as row (row.name)}
        <StatRow
          figures={[
            { value: `${(row.share * 100).toFixed(1)}%` },
            { value: formatRate(row.total / seconds) },
            { value: `${approximate ? '~' : ''}${formatAmount(row.total)}` },
          ]}
          phone={`${(row.share * 100).toFixed(1)}%`}
        >
          {#snippet icon()}<SpellIcon url={spellIcon(row.spellId)} />{/snippet}
          {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
          {#snippet meter()}<Meter value={row.share} tone="muted" />{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  {/if}
</Section>
