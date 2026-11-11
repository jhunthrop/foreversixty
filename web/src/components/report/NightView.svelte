<!-- web/src/components/report/NightView.svelte -->
<!-- The whole night, the Summary over "All boss pulls" (logs spec 2026-11-07 6.3, round 5
     design/mocks/directions/report_tail.py b1-report-night): the framed panel Players across every
     boss pull (one StatTable: the crest, the name, the spec, a damage bar in the class colour,
     Pulls, Damage and DPS, Healing and HPS, Taken, Active, Deaths), then the panel Bosses (Pulls,
     Kills, Wipes, Time, Deaths, the best kill or the lowest wipe, each boss's pulls one click away).
     Trash is left out for the reason the rankings leave it out; a zone's trash entry gets the same
     view of its own pulls (`trash`), with no boss table. Not reachable on the real log (no boss
     pull, gap G14). -->
<script lang="ts">
  import { unitLabel } from '../../lib/report/names.svelte';
  import { specCell, type UnitClass } from '../../lib/report/roster-class';
  import { classColorVar } from '../../lib/report/format';
  import { formatRate } from '../../lib/report/frame';
  import type { Night, NightBoss, NightPlayer } from '../../lib/report/night';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import { clockDuration } from '../../lib/format/duration';
  import { compact, percent } from '../../lib/format/number';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import Crest from '../ui/Crest.svelte';
  import CopyCsv from './CopyCsv.svelte';

  let {
    night,
    trash = false,
    loading,
    onSelect,
    onSelectPlayer,
    unitClasses = new Map<string, UnitClass>(),
    inScope = () => true,
  }: {
    night: Night;
    /** One zone's trash folded together rather than the night's boss pulls. */
    trash?: boolean;
    loading: boolean;
    onSelect: (index: number) => void;
    onSelectPlayer: (guid: string) => void;
    unitClasses?: ReadonlyMap<string, UnitClass>;
    /** The page's Source scope: the players it keeps (the group, by default). */
    inScope?: (guid: string) => boolean;
  } = $props();

  const players = $derived(night.players.filter((player) => inScope(player.guid)));

  const PLAYER_COLUMNS = [
    { id: 'pulls', label: 'Pulls' },
    { id: 'damage', label: 'Damage' },
    { id: 'dps', label: 'DPS' },
    { id: 'healing', label: 'Healing' },
    { id: 'hps', label: 'HPS' },
    { id: 'taken', label: 'Taken' },
    { id: 'active', label: 'Active' },
    { id: 'deaths', label: 'Deaths' },
  ];
  const BOSS_COLUMNS = [
    { id: 'pulls', label: 'Pulls' },
    { id: 'kills', label: 'Kills' },
    { id: 'wipes', label: 'Wipes' },
    { id: 'time', label: 'Time' },
    { id: 'deaths', label: 'Deaths' },
  ];

  const meta = $derived(
    [
      trash ? 'across every trash pull' : 'across every boss pull',
      ...(loading ? [`loading ${night.fights} of ${night.expected}`] : []),
    ].join(' · '),
  );

  const csvLines = (): string[][] => [
    ['Player', 'Class', ...PLAYER_COLUMNS.map((column) => column.label)],
    ...players.map((player) => [
      unitLabel(player.name).text,
      specCell(player, unitClasses).text,
      ...playerFigures(player).map((figure) => figure.value),
    ]),
  ];

  function playerFigures(player: NightPlayer): StatFigure[] {
    return [
      { value: String(player.fights) },
      { value: compact(player.damage_done) },
      { value: formatRate(player.dps) },
      { value: compact(player.healing_done) },
      { value: formatRate(player.hps) },
      { value: compact(player.damage_taken) },
      { value: percent(player.activity_pct) },
      { value: String(player.deaths), muted: player.deaths === 0 },
    ];
  }
  function bossFigures(boss: NightBoss): StatFigure[] {
    return [
      { value: String(boss.pulls) },
      { value: String(boss.kills), muted: boss.kills === 0 },
      { value: String(boss.wipes), muted: boss.wipes === 0 },
      { value: clockDuration(boss.time_ms) },
      { value: String(boss.deaths), muted: boss.deaths === 0 },
    ];
  }
  function bossSub(boss: NightBoss): string {
    if (boss.best) return `Best kill ${clockDuration(boss.best.duration_ms)}`;
    if (boss.lowest_wipe_pct !== undefined) return `Best wipe ${Math.round(boss.lowest_wipe_pct)}%`;
    return 'No kill';
  }
</script>

<Section title="Players" {meta} metaTestid="night-summary" testid="night-view" panel flush>
  {#snippet head()}<CopyCsv lines={csvLines} inHead />{/snippet}
  <div>
    <StatTable columns={PLAYER_COLUMNS} nameHead="Player" bar={false} label="Players" testid="night-players">
      {#each players as player (player.guid)}
        {@const display = unitLabel(player.name)}
        {@const cell = specCell(player, unitClasses)}
        <StatRow
          bar={false}
          figures={playerFigures(player)}
          phone={`${formatRate(player.dps)} DPS · ${player.fights} pulls`}
          testid={`night-player-${player.guid}`}
        >
          {#snippet icon()}<Crest cls={cell.class} size={28} />{/snippet}
          {#snippet name()}<button
              type="button"
              class="text-left font-bold hover:underline"
              style:color={classColorVar(cell.class)}
              title={display.title}
              data-testid="night-player-name"
              onclick={() => onSelectPlayer(player.guid)}>{display.text}</button
            >{/snippet}
          {#snippet sub()}{cell.text}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </div>
</Section>

{#if !trash}
  <Section title="Bosses" testid="night-bosses" panel flush>
    <StatTable columns={BOSS_COLUMNS} nameHead="Boss" bar={false} icons={false} label="Bosses">
      {#each night.bosses as boss (boss.name)}
        <StatRow
          bar={false}
          figures={bossFigures(boss)}
          phone={bossSub(boss)}
          testid={`night-boss-${boss.name}`}
        >
          {#snippet name()}<b class="text-strong">{boss.name}</b>{/snippet}
          {#snippet sub()}{bossSub(boss)} · {#each boss.fights as index (index)}<button
                type="button"
                class="text-gold mr-2 hover:underline"
                title="Open this pull"
                onclick={() => onSelect(index)}>Pull {index}</button
              >{/each}{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>
{/if}
