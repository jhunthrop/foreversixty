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
  import { viewerZone, weekdayDayMonth } from '../../../lib/dates';
  import Crest from '../../ui/Crest.svelte';
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
  /** Rows shown before All: the most overdue. */
  const FIRST = 12;
  let all = $state(false);
  const allRows = $derived(fairnessRows(roster, loot, grid));
  const rows = $derived.by(() => {
    if (all) return allRows;
    let left = FIRST;
    return allRows.flatMap((r) => {
      if (left <= 0) return [];
      const shared = r.items === 0 && r.raiders.length >= 3;
      const raiders = shared ? r.raiders : r.raiders.slice(0, left);
      left -= shared ? 1 : raiders.length;
      return [{ ...r, raiders }];
    });
  });
  const total = $derived(allRows.reduce((n, r) => n + r.raiders.length, 0));
  const perNight = $derived(guildItemsPerNight(roster, loot, grid));
  const meta = $derived(`${perNight.toFixed(2)} items per raider night`);

  function day(iso: string): string {
    return weekdayDayMonth(new Date(iso), timeZone ?? viewerZone());
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
        {#if row.items === 0 && row.raiders.length >= 3 && row.raiders.every((r) => r.present === row.raiders[0].present && r.nights === row.raiders[0].nights)}
          <StatRow
            bar={false}
            figures={[
              { value: '–', muted: true },
              { value: `${row.raiders[0].present}/${row.raiders[0].nights}` },
            ]}
          >
            {#snippet name()}{row.raiders.length} raiders, no item yet{/snippet}
            {#snippet sub()}{#each row.raiders as r, i (r.key)}<b
                  class="whitespace-nowrap"
                  style:color={classColorVar(r.cls)}>{r.name}</b
                >{i < row.raiders.length - 1 ? ', ' : ''}{/each}{/snippet}
          </StatRow>
        {:else}
          {#each row.raiders as one (one.key)}
            <StatRow
              bar={false}
              figures={figures(one)}
              phone={counted(one.items, 'item')}
              testid={`guild-fair-${one.name.toLowerCase()}`}
            >
              {#snippet icon()}<Crest cls={one.cls} size={28} />{/snippet}
              {#snippet name()}<span style:color={classColorVar(one.cls)}>{one.name}</span>{/snippet}
              {#snippet sub()}{one.spec}{one.last === null
                  ? ''
                  : ` · last ${one.last.item}, ${day(one.last.at)}`}{/snippet}
            </StatRow>
          {/each}
        {/if}
      {/each}
    </StatTable>
    {#if total > FIRST}
      <div class="pt-2 pr-5 pb-4 pl-[56px]">
        <Button tone="ghost" size="sm" onclick={() => (all = !all)} pressed={all}
          >{all ? `First ${FIRST}` : `All ${total}`}</Button
        >
      </div>
    {/if}
  </div>
</Section>
