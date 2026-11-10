<!-- web/src/components/guild/progression/EncounterProgress.svelte -->
<!-- One boss on the Progression tab (guild dense spec, h-guild-tabs): its figures (kills, best
     kill and its night, how much faster than the first kill, deaths a pull, first kill), then
     every night it was pulled on one dense table, newest first: the kill time on a Meter against
     the slowest, that time against the night before, pulls, wipes and deaths. The guild's best
     figure per role follows it. -->
<script lang="ts">
  import type { GuildParseRef, GuildProgressionEncounter } from '../../../lib/guild/api';
  import { progressionNightDay } from '../../../lib/guild/overview-view';
  import { outputFigure, outputUnit } from '../../../lib/guild/output-figure';
  import { fasterSinceFirst, progressionNights, slowestKillMs } from '../../../lib/guild/progression-nights';
  import { deltaWords, killClock } from '../../../lib/guild/raid-nights-table';
  import { specClassName } from '../../../lib/guild/spec-name';
  import { classColorVar } from '../../../lib/report/format';
  import { counted } from '../../../lib/plural';
  import type { StatFigure } from '../../../lib/ui/stat-table';
  import ClassCrestRing from '../../character/ClassCrestRing.svelte';
  import Figure from '../../ui/Figure.svelte';
  import Meter from '../../ui/Meter.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let { encounter, now }: { encounter: GuildProgressionEncounter; now: Date } = $props();

  const nights = $derived(progressionNights(encounter));
  const slowest = $derived(slowestKillMs(nights));
  const faster = $derived(fasterSinceFirst(encounter));
  const bestNight = $derived(nights.find((n) => n.best) ?? null);
  const firstKill = $derived(encounter.first_kill_at);
  const ROLE_WORD = { tank: 'Tank', healer: 'Healer', dps: 'Damage' } as const;
  const roles = $derived(
    (['tank', 'healer', 'dps'] as const).flatMap((role) => {
      const parse: GuildParseRef | undefined = encounter.best_by_role[role];
      return parse === undefined ? [] : [{ role, parse }];
    }),
  );
  const DASH: StatFigure = { value: '–', muted: true };
  const columns = [
    { id: 'kill', label: 'Kill' },
    { id: 'delta', label: 'vs last' },
    { id: 'pulls', label: 'Pulls' },
    { id: 'wipes', label: 'Wipes' },
    { id: 'deaths', label: 'Deaths' },
  ];
  const meta = $derived(
    `${encounter.zone} · ${encounter.pulls} pulls · ${encounter.pulls_by_night.length} nights`,
  );
</script>

<Section title={encounter.name} {meta} panel frame="framed" flush testid="guild-progression-encounter">
  <div class="flex flex-wrap gap-x-10 gap-y-3 px-5 pb-3" data-testid="guild-progression-figures">
    <Figure value={String(encounter.kills)} label="Kills" />
    {#if typeof encounter.best_kill_ms === 'number'}
      <Figure
        value={killClock(encounter.best_kill_ms)}
        label="Best kill"
        tone="gold"
        caption={bestNight === null ? undefined : progressionNightDay(bestNight.date, now)}
      />
    {/if}
    {#if faster !== null && faster > 0}<Figure
        value={`${faster} s`}
        label="Faster than the first"
        tone="kill"
      />{/if}
    <Figure value={encounter.deaths_per_pull.toFixed(1)} label="Deaths a pull" />
    {#if typeof firstKill === 'string'}
      <Figure value={progressionNightDay(firstKill, now)} label="First kill" />
    {/if}
  </div>
  {#if nights.length > 0}
    <StatTable {columns} nameHead="Night" barHead="Kill time" dense label={`${encounter.name} by night`}>
      {#each nights as night (night.reportId)}
        <StatRow
          figures={[
            night.killMs === null
              ? DASH
              : { value: killClock(night.killMs), tone: night.best ? 'gold' : undefined },
            night.deltaS === null ? DASH : deltaWords(night.deltaS),
            { value: String(night.pulls) },
            { value: String(night.wipes) },
            { value: String(night.deaths) },
          ]}
          phone={`${counted(night.pulls, 'pull')} · ${counted(night.deaths, 'death')}`}
          testid="guild-progression-night"
        >
          {#snippet name()}{progressionNightDay(night.date, now)}{/snippet}
          {#snippet sub()}{night.killMs === null
              ? 'no kill'
              : night.first
                ? 'first kill'
                : night.best
                  ? 'best kill'
                  : 'kill'}{/snippet}
          {#snippet meter()}{#if night.killMs !== null && slowest !== null}<Meter
                value={night.killMs / slowest}
                tone="gold"
              />{/if}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  {/if}
  {#if roles.length > 0}
    <div class="border-line-soft border-t" data-testid="guild-progression-best-parse">
      <StatTable
        columns={[{ id: 'figure', label: 'Best kill figure' }]}
        nameHead={`Best on ${encounter.name}`}
        bar={false}
        dense
      >
        {#each roles as { role, parse } (role)}
          <StatRow
            bar={false}
            figures={[{ value: `${outputFigure(parse.value)} ${outputUnit(parse.metric)}` }]}
          >
            {#snippet icon()}<ClassCrestRing characterClass={parse.class} size={28} />{/snippet}
            {#snippet name()}<span style:color={classColorVar(parse.class)}>{parse.name}</span>{/snippet}
            {#snippet sub()}{ROLE_WORD[role]}{#if parse.spec}
                · {specClassName(parse.spec, parse.class)}{/if}{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </div>
  {/if}
</Section>
