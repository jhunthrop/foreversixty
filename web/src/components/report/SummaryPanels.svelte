<!-- web/src/components/report/SummaryPanels.svelte -->
<!-- The Summary tab's two open Sections around its Plate (logs spec 2026-11-07 6.2):
     `part="taken"` (main, under the chart) Damage taken by ability: an open StatTable of each
     ability that hit a player in the scope (spell icon, name, a Meter, Share of everything every
     player took whatever the scope shows, Amount, Per sec) with the Damage taken › link;
     `part="deaths"` (the rail) Deaths: one Row per death in time order (the time, the crest, the
     name, "{ability} · {source} · {amount}", a melee swing reads Melee) with the Deaths › link. -->
<script lang="ts">
  import EmptyLine from './EmptyLine.svelte';
  import { splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { Actor, Summary } from '../../lib/report/types';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import Meter from '../ui/Meter.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import ClassIcon from './ClassIcon.svelte';
  import SpellIcon from './SpellIcon.svelte';

  /** The log's null unit: a fall, a hazard, or a source the log did not name. */
  const NULL_GUID = /^0+$/;

  let {
    part,
    summary,
    everyone,
    approximate = false,
    durationMs,
    players,
    tabHref,
    spellIcon,
  }: {
    part: 'taken' | 'deaths';
    summary: Summary;
    durationMs: number;
    players: ReadonlySet<string>;
    /** The unscoped window: a source scope narrows the rows, never the total they share. */
    everyone?: Summary;
    /** A brushed window: the by-ability split is prorated from the whole fight's, and says so. */
    approximate?: boolean;
    /** The url of a deeper tab, kept with the fight, source and window. */
    tabHref: (tab: 'damage-taken' | 'deaths') => string;
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
        if (found === undefined) totals.set(key, { name: key, spellId: ability.spell_id, total: ability.effective });
        else found.total += ability.effective;
      }
    }
    return [...totals.values()].filter((row) => row.total > 0).sort((a, b) => b.total - a.total);
  }
  const takenByAbility = $derived.by(() => {
    const rows = abilityTotals(summary.damage_taken);
    const total = abilityTotals((everyone ?? summary).damage_taken).reduce((sum, row) => sum + row.total, 0);
    return rows.slice(0, ROWS).map((row) => ({ ...row, share: total === 0 ? 0 : row.total / total }));
  });
  const seconds = $derived(Math.max(durationMs, 1) / 1000);
  const deaths = $derived([...summary.deaths].sort((a, b) => a.at_ms - b.at_ms));
  const COLUMNS = [
    { id: 'share', label: 'Share' },
    { id: 'amount', label: 'Amount' },
    { id: 'rate', label: 'Per sec' },
  ] as const;

  function cause(death: Summary['deaths'][number]): string {
    const blow = death.killing_blow;
    if (blow === undefined) return '';
    const spell =
      NULL_GUID.test(blow.source_guid) && blow.spell_name === ''
        ? 'a fall, a hazard or an untracked source'
        : blow.spell_name === ''
          ? 'Melee'
          : blow.spell_name;
    return `${spell} · ${splitUnitName(blow.source_name).name} · ${formatAmount(blow.amount)}`;
  }
</script>

{#if part === 'taken'}
  <Section
    title="Damage taken by ability"
    action={{ label: 'Damage taken', href: tabHref('damage-taken') }}
    testid="panel-taken"
  >
    {#if takenByAbility.length === 0}
      <EmptyLine text="Nobody took damage in this window." section />
    {:else}
      <StatTable columns={COLUMNS} nameHead="Ability" open label="Damage taken by ability">
        {#each takenByAbility as row (row.name)}
          <StatRow
            figures={[
              { value: `${(row.share * 100).toFixed(1)}%` },
              { value: `${approximate ? '~' : ''}${formatAmount(row.total)}` },
              { value: formatRate(row.total / seconds) },
            ]}
            phone={`${(row.share * 100).toFixed(1)}%`}
          >
            {#snippet icon()}<SpellIcon url={spellIcon(row.spellId)} />{/snippet}
            {#snippet name()}<b class="text-strong">{row.name}</b>{/snippet}
            {#snippet meter()}<Meter value={row.share} />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
      <p class="text-small text-muted pt-2" data-testid="panel-share-note">
        Share is of everything every player took in this window, whatever the source scope shows.
      </p>
    {/if}
  </Section>
{:else}
  <Section
    title="Deaths"
    action={deaths.length === 0 ? undefined : { label: 'Deaths', href: tabHref('deaths') }}
    testid="panel-deaths"
  >
    {#if deaths.length === 0}
      <EmptyLine text="Nobody died." section />
    {:else}
      {#each deaths as death, i (`${death.guid}-${death.at_ms}-${i}`)}
        <Row title={splitUnitName(death.name).name} sub={cause(death)} open>
          {#snippet lead()}<span class="text-muted tabular-nums text-[13px]">{formatDuration(death.at_ms)}</span
            ><ClassIcon className={death.class} size={28} />{/snippet}
          {#snippet titleContent()}<span style:color={classColorVar(death.class)}
              >{splitUnitName(death.name).name}</span
            >{/snippet}
        </Row>
      {/each}
    {/if}
  </Section>
{/if}
