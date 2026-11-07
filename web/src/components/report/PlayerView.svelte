<!-- web/src/components/report/PlayerView.svelte -->
<!-- One player's view (logs spec 2026-11-07 6.2, decision R5: the bar's Source is one player on
     Damage done, Healing or Damage taken). The Plate: Section "{name}'s damage" over the player's
     abilities (spell icon, name, "{hits} hits · {crits} crits · {min}–{max}", a Meter, Amount, the
     rate, Share, Hits, Crit). Under a window or a filter the split is prorated from the summary
     until it is measured from the fight's events, which it does by itself. `part="gear"` is the
     Section Gear at the pull under the chart: the shared Paperdoll in report mode, or on a pull
     with no gear in the log (every trash pull) the one line that says why. -->
<script lang="ts">
  import EmptyLine from './EmptyLine.svelte';
  import { unitLabel } from '../../lib/report/names.svelte';
  import { formatAmount, formatPercent } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { ExactSplit } from '../../lib/report/exact';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import type { Ability, Actor, CombatantRow, RosterRow } from '../../lib/report/types';
  import type { StatColumn, StatFigure, StatSort } from '../../lib/ui/stat-table';
  import Meter from '../ui/Meter.svelte';
  import Plate from '../ui/Plate.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import GearList from './GearList.svelte';
  import SpellIcon from './SpellIcon.svelte';

  type ActorTab = 'damage-done' | 'healing' | 'damage-taken';

  let {
    part = 'abilities',
    tab,
    actor,
    row,
    combatant,
    durationMs,
    approximate = false,
    measure,
    spellIcon,
    dataBuild,
  }: {
    part?: 'abilities' | 'gear';
    tab: ActorTab;
    actor: Actor | undefined;
    row: RosterRow | undefined;
    combatant: CombatantRow | undefined;
    durationMs: number;
    approximate?: boolean;
    measure?: (actor: Actor) => Promise<ExactSplit>;
    spellIcon: SpellIconOf;
    dataBuild: string;
  } = $props();

  const NOUN: Readonly<Record<ActorTab, string>> = {
    'damage-done': 'damage',
    healing: 'healing',
    'damage-taken': 'damage taken',
  };
  const RATE: Readonly<Record<ActorTab, string>> = {
    'damage-done': 'DPS',
    healing: 'HPS',
    'damage-taken': 'Taken / s',
  };
  const NO_GEAR =
    'No gear in this pull. The game writes each player’s gear into the log when a boss pull starts.';

  const name = $derived(unitLabel(actor?.name ?? row?.name ?? '').name);
  const COLUMNS = $derived<StatColumn[]>([
    { id: 'amount', label: 'Amount', sortable: true },
    { id: 'rate', label: RATE[tab], sortable: true },
    { id: 'share', label: 'Share', sortable: true },
    { id: 'hits', label: 'Hits', sortable: true },
    { id: 'crit', label: 'Crit', sortable: true },
  ]);

  let exact = $state<ExactSplit | null>(null);
  let measureError = $state('');
  $effect(() => {
    const wanted = actor;
    exact = null;
    measureError = '';
    if (!approximate || wanted === undefined || measure === undefined) return;
    let current = true;
    measure(wanted).then(
      (split) => {
        if (current) exact = split;
      },
      (thrown: unknown) => {
        if (current) measureError = thrown instanceof Error ? thrown.message : 'The measure did not run.';
      },
    );
    return () => {
      current = false;
    };
  });

  const landed = (ability: Ability): number => ability.hits + ability.ticks;
  const abilities = $derived(
    (exact?.abilities ?? actor?.abilities ?? []).filter((ability) => ability.effective > 0 || landed(ability) > 0),
  );
  const total = $derived(abilities.reduce((sum, ability) => sum + ability.effective, 0));
  const top = $derived(abilities.reduce((most, ability) => Math.max(most, ability.effective), 0));
  const seconds = $derived(Math.max(actor?.time_ms ?? durationMs, 1) / 1000);
  const mark = $derived(approximate && exact === null ? '~' : '');

  let sort = $state<StatSort>({ column: 'amount', direction: 'desc' });
  function onsort(column: string): void {
    sort =
      sort.column === column
        ? { column, direction: sort.direction === 'desc' ? 'asc' : 'desc' }
        : { column, direction: 'desc' };
  }
  function value(ability: Ability, column: string): number {
    if (column === 'hits') return landed(ability);
    if (column === 'crit') return landed(ability) === 0 ? 0 : ability.crits / landed(ability);
    return ability.effective;
  }
  const rows = $derived.by(() => {
    const sign = sort.direction === 'desc' ? -1 : 1;
    return [...abilities].sort((a, b) => sign * (value(a, sort.column) - value(b, sort.column)));
  });

  function figures(ability: Ability): StatFigure[] {
    const hits = landed(ability);
    return [
      { value: `${mark}${formatAmount(ability.effective)}` },
      { value: formatRate(ability.effective / seconds) },
      { value: total === 0 ? '0%' : formatPercent((ability.effective / total) * 100) },
      { value: String(hits) },
      { value: hits === 0 ? '–' : `${Math.round((ability.crits / hits) * 100)}%`, muted: hits === 0 },
    ];
  }
  function subOf(ability: Ability): string {
    const hits = landed(ability);
    return [
      `${hits} ${hits === 1 ? 'hit' : 'hits'}`,
      `${ability.crits} ${ability.crits === 1 ? 'crit' : 'crits'}`,
      ability.max > 0 ? `${formatAmount(ability.min)}–${formatAmount(ability.max)}` : '',
      ability.via ?? '',
    ]
      .filter((part) => part !== '')
      .join(' · ');
  }
  const gearMeta = $derived(
    [row?.spec, row?.item_level ? `item level ${row.item_level}` : ''].filter(Boolean).join(' · '),
  );
</script>

{#if part === 'abilities'}
  <Section
    title={`${name}’s ${NOUN[tab]}`}
    meta={`${mark}${formatAmount(total)} ${NOUN[tab]} · ${formatRate(total / seconds)} ${RATE[tab]}`}
    testid="report-table-card"
  >
    <Plate label="Abilities">
      {#if rows.length === 0}
        <EmptyLine text="Nothing in this window." testid="table-empty" />
      {:else}
        <StatTable columns={COLUMNS} nameHead="Ability" {sort} {onsort} label="Abilities" testid="row-abilities">
          {#each rows as ability (`${ability.spell_id}|${ability.via ?? ''}|${ability.name}`)}
            <StatRow
              figures={figures(ability)}
              phone={`${formatRate(ability.effective / seconds)} ${RATE[tab]}`}
            >
              {#snippet icon()}<SpellIcon url={spellIcon(ability.spell_id)} size={28} />{/snippet}
              {#snippet name()}<b class="text-strong">{ability.name === '' ? 'Melee' : ability.name}</b>{/snippet}
              {#snippet sub()}{subOf(ability)}{/snippet}
              {#snippet meter()}<Meter value={top === 0 ? 0 : ability.effective / top} />{/snippet}
            </StatRow>
          {/each}
        </StatTable>
      {/if}
    </Plate>
    {#if measureError !== ''}
      <p class="text-small text-wipe pt-2" role="alert">{measureError}</p>
    {/if}
  </Section>
{:else}
  <Section title="Gear at the pull" meta={gearMeta || undefined} testid="report-gear">
    {#if combatant === undefined || combatant.gear.every((item) => item.ID <= 0)}
      <EmptyLine text={NO_GEAR} section testid="report-no-gear" />
    {:else}
      <GearList gear={combatant.gear} className={row?.class} {dataBuild} {name} framed={false} />
    {/if}
  </Section>
{/if}
