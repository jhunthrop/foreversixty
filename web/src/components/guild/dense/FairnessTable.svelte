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
  import Button from '../../ui/Button.svelte';
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

  /** Folded on a phone until asked for: the council above it is the decision. */
  let open = $state(false);
  const columns = [
    { id: 'items', label: 'Items' },
    { id: 'nights', label: 'Nights' },
  ];
  const rows = $derived(fairnessRows(roster, loot, grid));
  const perNight = $derived(guildItemsPerNight(roster, loot, grid));
  const meta = $derived(`${perNight.toFixed(2)} items per raider night`);

  function day(iso: string): string {
    return new Date(iso).toLocaleDateString('en-GB', {
      weekday: 'short',
      day: 'numeric',
      month: 'short',
      timeZone,
    });
  }

  const blank = columns.map(() => ({ value: '' }));

  function figures(one: FairnessRaider) {
    return [
      one.items === 0 ? { value: '–', muted: true } : { value: String(one.items) },
      { value: `${one.present}/${one.nights}` },
    ];
  }
</script>

<Section title="Fairness, this tier" {meta} panel flush testid="guild-loot-fairness">
  {#snippet head()}
    <span class="lg:hidden"
      ><Button tone="ghost" size="sm" onclick={() => (open = !open)} pressed={open}
        >{open ? 'Hide' : 'Show'}</Button
      ></span
    >
  {/snippet}
  <div class={open ? '' : 'max-lg:hidden'}>
    <StatTable {columns} nameHead="Raider" bar={false} dense label="Fairness this tier">
      {#each rows as row, index (`${row.since}:${row.items}:${row.raiders[0].key}`)}
        {#if index === 0 || rows[index - 1].since !== row.since}
          <StatRow total bar={false} figures={blank}>
            {#snippet name()}{counted(row.since, 'night')} since an item{/snippet}
          </StatRow>
        {/if}
        {#each row.raiders as one (one.key)}
          <StatRow
            bar={false}
            figures={figures(one)}
            phone={counted(one.items, 'item')}
            testid={`guild-fair-${one.name.toLowerCase()}`}
          >
            {#snippet icon()}<ClassCrestRing characterClass={one.cls} size={28} />{/snippet}
            {#snippet name()}<span style:color={classColorVar(one.cls)}>{one.name}</span>{/snippet}
            {#snippet sub()}{one.spec}{one.last === null
                ? ''
                : ` · last ${one.last.item}, ${day(one.last.at)}`}{/snippet}
          </StatRow>
        {/each}
      {/each}
    </StatTable>
  </div>
</Section>
