<!-- web/src/components/guild/dense/LastNightMovers.svelte -->
<!-- Last night (dense spec 5.1): the kill's head line, then the raiders whose rating moved most
     against the night before, Up then Down, each with the weakest part of their rating and their
     figure. -->
<script lang="ts">
  import { classColorVar } from '../../../lib/report/format';
  import type { Mover } from '../../../lib/guild/dense/week-board';
  import { grouped } from '../../../lib/format/number';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';
  import Crest from '../../ui/Crest.svelte';

  let {
    meta,
    vsLabel,
    up,
    down,
    reportHref,
    title,
  }: {
    /** `Onyxia 3:12, best yet · 0 wipes · 2 deaths · 36 raiders`. */
    meta: string;
    /** The change column's head: `vs 30 Dec`. */
    vsLabel: string;
    up: readonly Mover[];
    down: readonly Mover[];
    reportHref: string;
    /** `Wed 6 Jan`: the raid the movers are from. */
    title: string;
  } = $props();

  const columns = $derived([
    { id: 'rating', label: 'Rating' },
    { id: 'change', label: vsLabel },
    { id: 'value', label: 'DPS or HPS' },
  ]);
  const blank = [{ value: '' }, { value: '' }, { value: '' }];

  function figures(m: Mover) {
    return [
      { value: String(m.rating) },
      {
        value: m.change > 0 ? `+${m.change}` : `−${-m.change}`,
        tone: m.change > 0 ? ('kill' as const) : ('wipe' as const),
      },
      { value: grouped(m.value) },
    ];
  }
</script>

{#snippet rows(list: readonly Mover[], label: string)}
  {#if list.length > 0}
    <StatRow total bar={false} figures={blank}>{#snippet name()}{label}{/snippet}</StatRow>
    {#each list as m (m.key)}
      <StatRow
        bar={false}
        figures={figures(m)}
        phone={`${figures(m)[1].value} ${vsLabel}`}
        testid={`guild-mover-${m.name.toLowerCase()}`}
      >
        {#snippet icon()}<Crest cls={m.cls} size={28} />{/snippet}
        {#snippet name()}<span style:color={classColorVar(m.cls)}>{m.name}</span>{/snippet}
        {#snippet sub()}{m.spec}{#if m.weakest}{` · weakest ${m.weakest.part} ${m.weakest.value}`}{/if}{/snippet}
      </StatRow>
    {/each}
  {/if}
{/snippet}

{#if up.length + down.length > 0}
  <Section
    {title}
    {meta}
    panel
    flush
    action={{ label: 'Report', href: reportHref }}
    testid="guild-week-last-night"
  >
    <StatTable {columns} nameHead="Raider" bar={false} dense label="Rating changes last night">
      {@render rows(up, 'Up')}
      {@render rows(down, 'Down')}
    </StatTable>
  </Section>
{/if}
