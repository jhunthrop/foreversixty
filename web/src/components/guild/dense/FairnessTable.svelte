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
  import { fairnessRows, guildItemsPerNight, type FairnessRow } from '../../../lib/guild/dense/fairness';
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

  /** A shared row's nights: the median raider's. */
  function medianNights(row: FairnessRow): string {
    const sorted = [...row.raiders].sort((a, b) => a.present - b.present);
    const mid = sorted[Math.floor(sorted.length / 2)];
    return `${mid.present}/${mid.nights}`;
  }

  function figures(row: FairnessRow) {
    const one = row.raiders.length === 1 ? row.raiders[0] : null;
    return [
      { value: String(row.since) },
      { value: String(row.items) },
      one === null ? { value: medianNights(row) } : { value: `${one.present}/${one.nights}` },
    ];
  }
</script>

<Section title="Fairness, this tier" {meta} panel flush testid="guild-loot-fairness">
  <StatTable {columns} nameHead="Raider" barHead="" dense label="Fairness this tier">
    {#each rows as row (row.raiders.map((r) => r.key).join(','))}
      {@const one = row.raiders.length === 1 ? row.raiders[0] : null}
      <StatRow figures={figures(row)} phone={counted(row.items, 'item')}>
        {#snippet icon()}{#if one !== null}<ClassCrestRing
              characterClass={one.cls}
              size={28}
            />{/if}{/snippet}
        {#snippet name()}
          {#if one !== null}<span style:color={classColorVar(one.cls)}>{one.name}</span>{:else}{row.raiders
              .length} raiders{/if}
        {/snippet}
        {#snippet sub()}
          {#if one !== null}{one.spec}{one.last === null
              ? ' · no item yet'
              : ` · last ${one.last.item}, ${day(one.last.at)}`}{:else}{#each row.raiders as r, i (r.key)}<b
                class="whitespace-nowrap"
                style:color={classColorVar(r.cls)}>{r.name}</b
              >{i < row.raiders.length - 1 ? ', ' : ''}{/each}{/if}
        {/snippet}
        {#snippet meter()}<Meter
            value={row.since / top}
            tone="gold"
            label={`${row.since} nights since an item`}
          />{/snippet}
      </StatRow>
    {/each}
  </StatTable>
</Section>
