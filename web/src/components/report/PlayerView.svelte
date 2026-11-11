<!-- web/src/components/report/PlayerView.svelte -->
<!-- One player's view (logs spec 2026-11-07 6.2, decision R5: the bar's Source is one player on
     Damage done, Healing or Damage taken). The Plate: Section "{name}'s damage" over the player's
     abilities (spell icon, name, "{hits} hits · {crits} crits · {min}–{max}", a Meter, Amount, the
     rate, Share, Hits, Crit). Under a window or a filter the split is prorated from the summary
     until it is measured from the fight's events, which it does by itself. `part="gear"` is the
     Section Gear at the pull under the chart: the shared paperdoll frame in report mode, or on a pull
     with no gear in the log (every trash pull) the one line that says why. -->
<script lang="ts">
  import EmptyState from '../ui/EmptyState.svelte';
  import CopyCsv from './CopyCsv.svelte';
  import { unitLabel } from '../../lib/report/names.svelte';
  import { schoolName } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { ExactSplit } from '../../lib/report/exact';
  import type { SpellIconOf } from '../../lib/guild/report-spell-icons';
  import {
    abilityKey,
    type Ability,
    type Actor,
    type CombatantRow,
    type RosterRow,
  } from '../../lib/report/types';
  import type { StatColumn, StatFigure, StatSort } from '../../lib/ui/stat-table';
  import { compact, percent } from '../../lib/format/number';
  import Meter from '../ui/Meter.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import GearList from './GearList.svelte';
  import SpellIcon from '../ui/SpellIcon.svelte';

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
    plannerHref,
    onChart,
    charted = '',
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
    /** The worn set at the pull as a planner build; null when the class has no planner data. */
    plannerHref?: string | null;
    /** Puts this ability's line on the chart under the table, or takes it off. */
    onChart?: (actor: Actor, ability: Ability) => void;
    /** `${guid}|${abilityKey(ability)}` of the ability on the chart; '' for none. */
    charted?: string;
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
    (exact?.abilities ?? actor?.abilities ?? []).filter(
      (ability) => ability.effective > 0 || landed(ability) > 0,
    ),
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
      { value: `${mark}${compact(ability.effective)}` },
      { value: formatRate(ability.effective / seconds) },
      { value: total === 0 ? '0%' : percent((ability.effective / total) * 100) },
      { value: String(hits) },
      { value: hits === 0 ? '–' : `${Math.round((ability.crits / hits) * 100)}%`, muted: hits === 0 },
    ];
  }
  /** What did not land, or how much of a heal was over: the old row expansion's notes column. */
  function notesOf(ability: Ability): string[] {
    const notes: string[] = [];
    if (ability.absorbed) notes.push(`${compact(ability.absorbed)} absorbed`);
    if (ability.overheal !== undefined && ability.total > 0) {
      if (ability.overheal > 0)
        notes.push(
          `${percent((ability.overheal / ability.total) * 100)} over · ${compact(ability.overheal)}`,
        );
      return notes;
    }
    if (ability.blocked) notes.push(`${compact(ability.blocked)} blocked`);
    for (const [type, count] of Object.entries(ability.misses ?? {}).sort(([a], [b]) => a.localeCompare(b)))
      notes.push(`${count} ${type.toLowerCase()}`);
    return notes;
  }
  /** Names two spell ids share: each such row carries its id to tell them apart. */
  const sameName = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const seen = new Map<string, Set<number>>();
    for (const ability of abilities)
      // eslint-disable-next-line svelte/prefer-svelte-reactivity
      seen.set(ability.name, (seen.get(ability.name) ?? new Set<number>()).add(ability.spell_id));
    return new Set([...seen].filter(([, ids]) => ids.size > 1).map(([name]) => name));
  });
  /** The player's amount by spell school ("Physical 80% · Fire 20%"), when it has more than one. */
  const schools = $derived.by(() => {
    // eslint-disable-next-line svelte/prefer-svelte-reactivity
    const sums = new Map<string, number>();
    for (const ability of abilities) {
      if (ability.effective <= 0) continue;
      const name = schoolName(ability.school) || 'Physical';
      sums.set(name, (sums.get(name) ?? 0) + ability.effective);
    }
    if (sums.size < 2 || total === 0) return '';
    return [...sums]
      .sort((a, b) => b[1] - a[1])
      .map(([name, amount]) => `${name} ${Math.round((amount / total) * 100)}%`)
      .join(' · ');
  });
  function subOf(ability: Ability): string {
    const hits = landed(ability);
    return [
      `${hits} ${hits === 1 ? 'hit' : 'hits'}`,
      `${ability.crits} ${ability.crits === 1 ? 'crit' : 'crits'}`,
      ability.max > 0 ? `${compact(ability.min)}–${compact(ability.max)}` : '',
      hits > 0 && ability.effective > 0 ? `${compact(ability.effective / hits)} each` : '',
      ability.via ? `via ${ability.via}` : '',
      ...notesOf(ability),
    ]
      .filter((part) => part !== '')
      .join(' · ');
  }
  /** The abilities table as lines, the notes as words. */
  function abilityCsv(): string[][] {
    return [
      ['Ability', 'Spell id', 'Via', 'Amount', 'Share %', 'Hits', 'Crit %', 'Avg', 'Max', 'Notes'],
      ...rows.map((ability) => {
        const hits = landed(ability);
        return [
          ability.name,
          String(ability.spell_id),
          ability.via ?? '',
          String(ability.effective),
          total === 0 ? '0' : ((ability.effective / total) * 100).toFixed(1),
          String(hits),
          hits === 0 ? '0' : ((ability.crits / hits) * 100).toFixed(1),
          hits === 0 ? '0' : Math.round(ability.effective / hits).toString(),
          String(ability.max),
          notesOf(ability).join('; '),
        ];
      }),
    ];
  }
  const chartKey = (ability: Ability): string => `${actor?.guid ?? ''}|${abilityKey(ability)}`;

  const gearMeta = $derived(
    [row?.spec, row?.item_level ? `ilvl ${row.item_level}` : ''].filter(Boolean).join(' · '),
  );
</script>

{#if part === 'abilities'}
  <Section
    title={`${name}’s ${NOUN[tab]}`}
    meta={[`${mark}${compact(total)} ${NOUN[tab]} · ${formatRate(total / seconds)} ${RATE[tab]}`, schools]
      .filter((part) => part !== '')
      .join(' · ')}
    testid="report-table-card"
    panel
    flush
  >
    {#snippet head()}{#if rows.length > 0}<CopyCsv lines={abilityCsv} inHead />{/if}{/snippet}
    <div>
      {#if rows.length === 0}
        <EmptyState inset message="Nothing in this window." testid="table-empty" />
      {:else}
        <StatTable
          columns={COLUMNS}
          nameHead="Ability"
          {sort}
          {onsort}
          label="Abilities"
          testid="row-abilities"
        >
          {#each rows as ability (`${ability.spell_id}|${ability.via ?? ''}|${ability.name}`)}
            <StatRow
              figures={figures(ability)}
              phone={`${formatRate(ability.effective / seconds)} ${RATE[tab]}`}
            >
              {#snippet icon()}<SpellIcon
                  socket
                  tip={ability.spell_id}
                  url={spellIcon(ability.spell_id, ability.name, ability.school)}
                  size={28}
                />{/snippet}
              {#snippet name()}<b class="text-strong">{ability.name === '' ? 'Melee' : ability.name}</b
                >{#if sameName.has(ability.name)}<span class="text-muted ml-1 text-[12px]"
                    >{ability.spell_id === 0 ? 'swing' : `#${ability.spell_id}`}</span
                  >{/if}{#if onChart !== undefined && actor !== undefined}<button
                    type="button"
                    class="text-gold ml-2 text-[13px] font-semibold [@media(pointer:coarse)]:-my-3 [@media(pointer:coarse)]:py-3"
                    title="Draw this ability's amount per second on the chart below"
                    aria-pressed={charted === chartKey(ability)}
                    data-testid="ability-chart"
                    onclick={() => {
                      if (actor !== undefined) onChart(actor, ability);
                    }}>{charted === chartKey(ability) ? 'Off the chart' : 'On the chart'}</button
                  >{/if}{/snippet}
              {#snippet sub()}{subOf(ability)}{/snippet}
              {#snippet meter()}<Meter value={top === 0 ? 0 : ability.effective / top} />{/snippet}
            </StatRow>
          {/each}
        </StatTable>
      {/if}
    </div>
    {#if measureError !== ''}
      <p class="text-small text-wipe px-5 pt-2" role="alert">{measureError}</p>
    {/if}
  </Section>
{:else}
  <Section
    title="Gear at the pull"
    meta={gearMeta || undefined}
    action={plannerHref ? { label: 'Open in the planner', href: plannerHref } : undefined}
    testid="report-gear"
    panel
  >
    {#if combatant === undefined || combatant.gear.every((item) => item.ID <= 0)}
      <EmptyState inset message={NO_GEAR} testid="report-no-gear" />
    {:else}
      <GearList
        gear={combatant.gear}
        className={row?.class}
        {dataBuild}
        {name}
        spec={row?.class}
        figure={total > 0 ? { value: formatRate(total / seconds), unit: RATE[tab] } : undefined}
      />
    {/if}
  </Section>
{/if}
