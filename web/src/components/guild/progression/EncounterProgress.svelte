<!-- web/src/components/guild/progression/EncounterProgress.svelte -->
<!-- One boss on the Progression tab (guild dense spec, h-guild-tabs): its figures (kills, best
     kill and its night, how much faster than the first kill, deaths a pull, first kill), then
     every night it was pulled on one dense table, newest first: the kill time on a Meter against
     the slowest, that time against the night before, pulls, wipes and deaths. The guild's best
     figure per role is the rail's (BestByRole). -->
<script lang="ts">
  import type { GuildProgressionEncounter } from '../../../lib/guild/api';
  import { progressionNightDay } from '../../../lib/guild/overview-view';
  import { fasterSinceFirst, progressionNights, slowestKillMs } from '../../../lib/guild/progression-nights';
  import { deltaWords } from '../../../lib/guild/raid-nights-table';
  import { counted } from '../../../lib/plural';
  import { SEP } from '../../../lib/guild/raid-copy';
  import type { StatFigure } from '../../../lib/ui/stat-table';
  import { clockDuration } from '../../../lib/format/duration';
  import Meter from '../../ui/Meter.svelte';
  import Section from '../../ui/Section.svelte';
  import StatRow from '../../ui/StatRow.svelte';
  import StatTable from '../../ui/StatTable.svelte';

  let { encounter, now }: { encounter: GuildProgressionEncounter; now: Date } = $props();

  const nights = $derived(progressionNights(encounter));
  const slowest = $derived(slowestKillMs(nights));
  const faster = $derived(fasterSinceFirst(encounter));
  const bestNight = $derived(nights.find((n) => n.best) ?? null);
  const DASH: StatFigure = { value: '–', muted: true };
  const columns = [
    { id: 'kill', label: 'Kill' },
    { id: 'delta', label: 'vs last' },
    { id: 'pulls', label: 'Pulls' },
    { id: 'wipes', label: 'Wipes' },
    { id: 'deaths', label: 'Deaths' },
  ];
  /** The boss's head line: kills, the best kill and its night, how much faster than the first,
   *  deaths a pull. */
  const meta = $derived(
    [
      counted(encounter.kills, 'kill'),
      typeof encounter.best_kill_ms === 'number'
        ? `best ${clockDuration(encounter.best_kill_ms)}${bestNight === null ? '' : `, ${progressionNightDay(bestNight.date, now)}`}`
        : '',
      faster !== null && faster > 0 ? `${faster} s faster than the first` : '',
      `${encounter.deaths_per_pull.toFixed(1)} deaths a pull`,
    ]
      .filter((part) => part !== '')
      .join(SEP),
  );
</script>

<Section title={encounter.name} {meta} panel flush testid="guild-progression-encounter">
  {#if nights.length > 0}
    <StatTable {columns} nameHead="Night" barHead="Kill time" dense label={`${encounter.name} by night`}>
      {#each nights as night (night.reportId)}
        <StatRow
          figures={[
            night.killMs === null
              ? DASH
              : { value: clockDuration(night.killMs), tone: night.best ? 'gold' : undefined },
            night.deltaS === null ? DASH : deltaWords(night.deltaS),
            { value: String(night.pulls) },
            { value: String(night.wipes) },
            { value: String(night.deaths) },
          ]}
          phone={`${counted(night.pulls, 'pull')} · ${counted(night.deaths, 'death')}`}
          testid="guild-progression-night"
        >
          {#snippet name()}{progressionNightDay(night.date, now)}{/snippet}
          {#snippet sub()}{!night.killed
              ? 'no kill'
              : night.first
                ? 'first kill'
                : night.best
                  ? 'best kill'
                  : ''}{/snippet}
          {#snippet meter()}{#if night.killMs !== null && slowest !== null}<Meter
                value={night.killMs / slowest}
                tone="gold"
              />{/if}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  {/if}
</Section>
