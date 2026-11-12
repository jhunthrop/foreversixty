<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The Summary tab's ability panels under the chart (logs spec 2026-11-07 6.2; owner review
     2026-11-09): `metric="done"` Damage done by ability, every ability the scope's players hit
     with; `metric="taken"` Damage taken by ability, every ability that hit them; `metric="healing"`
     Healing by ability, every heal they cast (effective healing, overheal left out). Each a
     StatTable (spell icon, name, a Meter, the share of the scope's total, Per sec, Amount), the
     top eight, with the full tab's link in the head. The amount tabs show their own under the chart,
     every ability and the log site's columns beside the amount: Hits, Crit, Avg and Max (density
     pass 2026-11-09), Overheal on Healing, no Crit on Damage taken. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import { compact, grouped } from '../../lib/format/number';
  import { formatRate } from '../../lib/report/frame';
  import type { Actor, Summary } from '../../lib/report/types';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';

  type Metric = 'done' | 'taken' | 'healing';

  let {
    metric,
    tabHref,
    summary,
    approximate = false,
    durationMs,
    players,
    spellIcon,
  }: {
    metric: Metric;
    /** The full tab's url (Damage done, Damage taken, Healing); none on that tab itself. */
    tabHref?: string;
    summary: Summary;
    durationMs: number;
    players: ReadonlySet<string>;
    /** A brushed window: the by-ability split is prorated from the whole fight's, and says so. */
    approximate?: boolean;
    spellIcon: SpellIconOf;
  } = $props();

  const ROWS = 8;
  const TITLE: Readonly<Record<Metric, string>> = {
    done: 'Damage done by ability',
    taken: 'Damage taken by ability',
    healing: 'Healing by ability',
  };
  const TAB: Readonly<Record<Metric, string>> = {
    done: 'Damage done',
    taken: 'Damage taken',
    healing: 'Healing',
  };
  const EMPTY: Readonly<Record<Metric, string>> = {
    done: 'Nobody did damage in this window.',
    taken: 'Nobody took damage in this window.',
    healing: 'Nobody healed in this window.',
  };
  const TONE: Readonly<Record<Metric, 'accent' | 'muted' | 'kill'>> = {
    done: 'accent',
    taken: 'muted',
    healing: 'kill',
  };

  interface AbilityTotal {
    name: string;
    spellId: number;
    school?: number;
    total: number;
    /** Hits and ticks: what landed. */
    landed: number;
    crits: number;
    max: number;
    overheal: number;
  }

  /** The tab's own panel (no link to a tab): every ability, with the detail columns. */
  const detail = $derived(tabHref === undefined);

  /** Every ability in `table` for the scope's players, summed by name, largest first. */
  function abilityTotals(table: readonly Actor[]): AbilityTotal[] {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const totals = new Map<string, AbilityTotal>();
    for (const actor of table) {
      if (players.size > 0 && !players.has(actor.guid)) continue;
      for (const ability of actor.abilities) {
        const key = ability.name === '' ? 'Melee' : ability.name;
        const found = totals.get(key);
        const landed = ability.hits + ability.ticks;
        if (found === undefined)
          totals.set(key, {
            name: key,
            spellId: ability.spell_id,
            school: ability.school,
            total: ability.effective,
            landed,
            crits: ability.crits,
            max: ability.max,
            overheal: ability.overheal ?? 0,
          });
        else
          totals.set(key, {
            ...found,
            total: found.total + ability.effective,
            landed: found.landed + landed,
            crits: found.crits + ability.crits,
            max: Math.max(found.max, ability.max),
            overheal: found.overheal + (ability.overheal ?? 0),
          });
      }
    }
    return [...totals.values()].filter((row) => row.total > 0).sort((a, b) => b.total - a.total);
  }

  const source = (which: Metric): readonly Actor[] =>
    which === 'done' ? summary.damage_done : which === 'taken' ? summary.damage_taken : summary.healing;
  const byAbility = $derived.by(() => {
    const rows = abilityTotals(source(metric));
    const total = rows.reduce((sum, row) => sum + row.total, 0);
    return (detail ? rows : rows.slice(0, ROWS)).map((row) => ({
      ...row,
      share: total === 0 ? 0 : row.total / total,
    }));
  });
  const seconds = $derived(Math.max(durationMs, 1) / 1000);
  const BASE_COLUMNS = [
    { id: 'share', label: 'Share' },
    { id: 'rate', label: 'Per sec' },
    { id: 'amount', label: 'Amount' },
  ] as const;
  /** The detail columns each metric reads, after the base three. */
  const DETAIL: Readonly<Record<Metric, readonly ('overheal' | 'hits' | 'crit' | 'avg' | 'max')[]>> = {
    done: ['hits', 'crit', 'avg', 'max'],
    taken: ['hits', 'avg', 'max'],
    healing: ['overheal', 'hits', 'crit', 'max'],
  };
  const DETAIL_LABEL = { overheal: 'Overheal', hits: 'Hits', crit: 'Crit', avg: 'Avg', max: 'Max' } as const;
  /** The metric's detail columns, less Overheal when every row reads 0%: a column of zeros says nothing. */
  const details = $derived(
    DETAIL[metric].filter(
      (id) => id !== 'overheal' || byAbility.some((row) => detailValue(row, id) !== '0%'),
    ),
  );
  const columns = $derived([
    ...BASE_COLUMNS,
    ...(detail ? details.map((id) => ({ id, label: DETAIL_LABEL[id] })) : []),
  ]);
  /** A detail's word on the phone's sub line, by its count: "1 hit", "12 hits", "48% crit". */
  function phoneLabel(row: AbilityTotal, id: (typeof DETAIL)[Metric][number]): string {
    return id === 'hits' && row.landed === 1 ? 'hit' : DETAIL_LABEL[id].toLowerCase();
  }
  /** One detail figure; a prorated window marks the amounts it scales, as the Amount column does. */
  function detailValue(row: AbilityTotal, id: (typeof DETAIL)[Metric][number]): string {
    const mark = approximate ? '~' : '';
    if (id === 'hits') return `${mark}${grouped(row.landed)}`;
    if (id === 'crit') return row.landed === 0 ? '0%' : `${Math.round((row.crits / row.landed) * 100)}%`;
    if (id === 'avg') return row.landed === 0 ? '0' : compact(Math.round(row.total / row.landed));
    if (id === 'max') return compact(row.max);
    const raw = row.total + row.overheal;
    return raw === 0 ? '0%' : `${Math.round((row.overheal / raw) * 100)}%`;
  }
</script>

<Section
  title={TITLE[metric]}
  action={tabHref === undefined ? undefined : { label: TAB[metric], href: tabHref, arrow: true }}
  panel
  flush
  testid={`panel-${metric}`}
>
  {#if byAbility.length === 0}
    <EmptyState inset message={EMPTY[metric]} />
  {:else}
    <StatTable {columns} nameHead="Ability" label={TITLE[metric]}>
      {#each byAbility as row (row.name)}
        {@const figures = [
          { value: `${(row.share * 100).toFixed(1)}%` },
          { value: formatRate(row.total / seconds) },
          { value: `${approximate ? '~' : ''}${compact(row.total)}` },
          ...(detail ? details.map((id) => ({ value: detailValue(row, id) })) : []),
        ]}
        {@const phone = compact(row.total)}
        {#if detail}
          <!-- On a phone the detail columns fold into the sub line under the name. -->
          <StatRow {figures} {phone}>
            {#snippet icon()}<SpellIcon
                socket
                tip={row.spellId}
                size={24}
                url={spellIcon(row.spellId, row.name, row.school)}
              />{/snippet}
            {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
            {#snippet sub()}{#each details as id, i (id)}<span class="ui-st-phone"
                  >{i > 0 ? ' · ' : ''}{detailValue(row, id)} {phoneLabel(row, id)}</span
                >{/each}{/snippet}
            {#snippet meter()}<Meter value={row.share} tone={TONE[metric]} />{/snippet}
          </StatRow>
        {:else}
          <StatRow {figures} {phone}>
            {#snippet icon()}<SpellIcon
                socket
                tip={row.spellId}
                size={24}
                url={spellIcon(row.spellId, row.name, row.school)}
              />{/snippet}
            {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
            {#snippet meter()}<Meter value={row.share} tone={TONE[metric]} />{/snippet}
          </StatRow>
        {/if}
      {/each}
    </StatTable>
  {/if}
</Section>
