<!-- web/src/components/guild/dense/FairnessTable.svelte -->
<!-- Fairness this tier (dense spec 5.5), the loot council's who-is-due read: every raider by the
     nights they came since their last item, then fewest items; raiders level on both share one row
     with their names. One grid for both row kinds (since, items, nights); the last item rides the
     raider's sub line. -->
<script lang="ts">
  import type { GuildRosterRow } from '../../../lib/guild/api';
  import type { LootNights } from '../../../lib/guild/loot-council-api';
  import type { GuildRecentAttendance } from '../../../lib/guild/overview-api';
  import { classColorVar } from '../../../lib/report/format';
  import { counted } from '../../../lib/plural';
  import { fairnessRows, guildItemsPerNight, type FairnessRaider } from '../../../lib/guild/dense/fairness';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Meter from '../../ui/Meter.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let {
    roster,
    loot,
    grid,
    timeZone,
  }: {
    roster: readonly GuildRosterRow[];
    loot: LootNights;
    grid: GuildRecentAttendance;
    timeZone?: string;
  } = $props();

  const columns = [
    { id: 'since', label: 'Nights since an item' },
    { id: 'items', label: 'Items' },
    { id: 'nights', label: 'Nights' },
  ];
  const rows = $derived(fairnessRows(roster, loot, grid));
  const top = $derived(Math.max(1, ...rows.map((r) => r.since)));
  const items = $derived(new Set(loot.nights.flatMap((n) => n.decisions.map((d) => d.award_id))).size);
  const perNight = $derived(guildItemsPerNight(roster, loot, grid));
  const meta = $derived(
    `${counted(items, 'item')} over ${counted(grid.raids.length, 'night')} · ${perNight.toFixed(2)} per raider night`,
  );

  function day(iso: string): string {
    return new Date(iso).toLocaleDateString('en-GB', {
      weekday: 'short',
      day: 'numeric',
      month: 'short',
      timeZone,
    });
  }

  const blank = columns.map(() => ({ value: '' }));

  function figures(since: number, one: FairnessRaider) {
    return [
      { value: String(since) },
      { value: String(one.items) },
      { value: `${one.present}/${one.nights}` },
    ];
  }
</script>

<Section title="Fairness, this tier" {meta} panel flush testid="guild-loot-fairness">
  <StatTable {columns} nameHead="Raider" barHead="" dense label="Fairness this tier">
    {#each rows as row, index (`${row.since}:${row.items}:${row.raiders[0].key}`)}
      {#if index === 0 || rows[index - 1].since !== row.since}
        <StatRow total figures={blank}>
          {#snippet name()}{counted(row.since, 'night')} since an item{/snippet}
        </StatRow>
      {/if}
      {#each row.raiders as one (one.key)}
        <StatRow
          figures={figures(row.since, one)}
          phone={counted(one.items, 'item')}
          testid={`guild-fair-${one.name.toLowerCase()}`}
        >
          {#snippet icon()}<ClassCrestRing characterClass={one.cls} size={28} />{/snippet}
          {#snippet name()}<span style:color={classColorVar(one.cls)}>{one.name}</span>{/snippet}
          {#snippet sub()}{one.spec}{one.last === null
              ? ' · no item yet'
              : ` · last ${one.last.item}, ${day(one.last.at)}`}{/snippet}
          {#snippet meter()}<Meter
              value={row.since / top}
              tone="gold"
              label={`${row.since} nights since an item`}
            />{/snippet}
        </StatRow>
      {/each}
    {/each}
  </StatTable>
</Section>
