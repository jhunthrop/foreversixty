<!-- web/src/components/guild/GuildRaidNights.svelte -->
<!-- The Raids tab's logged nights as one dense table (guild dense spec, h-guild-tabs): the night
     (a button that opens its pulls in place), the boss time on a Meter against the slowest night,
     that time against the night before, wipes, deaths, raiders, the top HPS and DPS in class colour
     and the items the council handed out. A night matched to a guild raid links its raid page from
     the open pulls. A night in a zone with no published encounters names its pulls plainly. -->
<script lang="ts">
  import CharacterName from '../ui/CharacterName.svelte';
  import ZoneIcon from '../ui/ZoneIcon.svelte';
  import type { GuildRaidRow } from '../../lib/guild/api';
  import { guildRaidHref } from '../../lib/characters';
  import { attendanceCopy } from '../../lib/guild/attendance-copy';
  import { raidNightDay } from '../../lib/guild/overview-view';
  import { outputFigure } from '../../lib/guild/output-figure';
  import { TRASH_LABEL, nightFights } from '../../lib/guild/raid-view';
  import {
    bossPulls,
    deltaWords,
    raidNightFigures,
    slowestMs,
    type NightMisses,
    type RaidNightFigures,
  } from '../../lib/guild/raid-nights-table';
  import { counted } from '../../lib/plural';
  import { SEP } from '../../lib/guild/raid-copy';
  import type { StatFigure } from '../../lib/ui/stat-table';
  import { LINK, RULE_BOTTOM } from '../../lib/ui/primitives';
  import Meter from '../ui/Meter.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import Tag from '../ui/Tag.svelte';
  import activeBuild from '../../data/active-build.json';
  import { clockDuration } from '../../lib/format/duration';

  let {
    rows,
    bossIds,
    guildPath,
    lootByReport = null,
    absentByRaid = null,
    expanded,
    onToggle,
  }: {
    rows: readonly GuildRaidRow[];
    bossIds: ReadonlySet<number> | null;
    guildPath: string;
    /** Items awarded per report id (a member's loot nights read); null leaves the column out. */
    lootByReport?: ReadonlyMap<string, number> | null;
    /** Officers: raiders absent per guild raid id; null leaves the column out. */
    absentByRaid?: ReadonlyMap<number, NightMisses> | null;
    expanded: string | null;
    onToggle: (id: string) => void;
  } = $props();

  const figures = $derived(raidNightFigures(rows, bossIds));
  const slowest = $derived(slowestMs(figures));
  const columns = $derived([
    { id: 'kill', label: 'Kill time' },
    { id: 'delta', label: 'vs last' },
    { id: 'wipes', label: 'Wipes' },
    { id: 'deaths', label: 'Deaths' },
    { id: 'raiders', label: 'Raiders' },
    ...(absentByRaid === null ? [] : [{ id: 'absent', label: 'Absent' }]),
    ...(lootByReport === null ? [] : [{ id: 'loot', label: 'Loot' }]),
  ]);
  const DASH: StatFigure = { value: '–', muted: true };

  /** The night's top healer and damage dealer, healing first, for the sub line. */
  const tops = (row: GuildRaidRow) =>
    [
      { word: 'Top HPS', parse: row.top_hps, kind: 'hps' },
      { word: 'Top DPS', parse: row.top_dps, kind: 'dps' },
    ].filter((top) => top.parse != null);

  function missesOf(row: GuildRaidRow): NightMisses | undefined {
    return row.raid_id == null ? undefined : absentByRaid?.get(row.raid_id);
  }

  function absentCell(row: GuildRaidRow): StatFigure {
    const misses = missesOf(row);
    if (misses === undefined) return DASH;
    const absent = misses.absent.length;
    return { value: String(absent), tone: absent > 0 ? 'wipe' : undefined };
  }

  function cells(f: RaidNightFigures): StatFigure[] {
    const delta = f.deltaS === null ? DASH : { ...deltaWords(f.deltaS) };
    return [
      f.bossMs === null ? DASH : { value: clockDuration(f.bossMs), tone: f.best ? 'gold' : undefined },
      delta,
      { value: String(f.wipes) },
      { value: String(f.row.deaths) },
      { value: String(f.row.raiders) },
      ...(absentByRaid === null ? [] : [absentCell(f.row)]),
      ...(lootByReport === null ? [] : [{ value: String(lootByReport.get(f.row.id) ?? 0) }]),
    ];
  }
</script>

<StatTable {columns} nameHead="Night" barHead="Boss time" dense label="Raid nights" testid="guild-raids-rows">
  {#each figures as f (f.row.id)}
    {@const isOpen = expanded === f.row.id}
    {@const misses = missesOf(f.row)}
    <StatRow
      figures={cells(f)}
      phone={`${counted(f.wipes, 'wipe')} · ${counted(f.row.deaths, 'death')}`}
      testid={`guild-raid-night-row-${f.row.id}`}
    >
      {#snippet icon()}<ZoneIcon zone={f.row.zone} build={activeBuild.build} size={28} />{/snippet}
      {#snippet name()}
        <button
          type="button"
          class="text-strong relative min-h-6 text-left underline-offset-4 after:absolute after:inset-x-0 after:top-1/2 after:h-11 after:-translate-y-1/2 after:content-[''] hover:underline"
          aria-expanded={isOpen}
          onclick={() => onToggle(f.row.id)}
          data-testid={`guild-raid-row-${f.row.id}`}>{raidNightDay(f.row.created_at)}</button
        >
      {/snippet}
      {#snippet sub()}{counted(bossPulls(f.row, bossIds), 'pull')}{#if f.best}{SEP}<span class="text-gold"
            >best kill</span
          >{/if}{#each tops(f.row) as top (top.kind)}{SEP}<span
            data-testid={`guild-raid-top-${top.kind}-${f.row.id}`}
            >{top.word}&#32;<CharacterName
              name={top.parse?.name ?? ''}
              cls={top.parse?.class}
              weight="inherit"
              class="font-bold"
            />&#32;{outputFigure(top.parse?.value ?? 0)}</span
          >{/each}{/snippet}
      {#snippet meter()}{#if f.bossMs !== null && slowest !== null}<Meter
            value={f.bossMs / slowest}
            tone="gold"
          />{/if}{/snippet}
    </StatRow>
    {#if isOpen}
      <div
        id={`guild-raid-night-${f.row.id}`}
        class="{RULE_BOTTOM} px-4 pt-1 pb-3 md:pl-14"
        data-testid={`guild-raid-expanded-${f.row.id}`}
      >
        {#if f.row.raid_id !== undefined && f.row.raid_id !== null}
          <a
            href={guildRaidHref(guildPath, f.row.raid_id)}
            class="{LINK} inline-flex min-h-9 items-center text-[13px]"
            data-testid={`guild-raid-link-${f.row.id}`}>{attendanceCopy.match.openRaid}</a
          >
        {/if}
        {#if misses !== undefined && (misses.absent.length > 0 || misses.bench.length > 0)}
          <p class="text-[13px] leading-5" data-testid={`guild-raid-misses-${f.row.id}`}>
            {#if misses.absent.length > 0}<span class="text-muted">Absent</span>
              <span class="text-wipe">{misses.absent.join(', ')}</span
              >{/if}{#if misses.absent.length > 0 && misses.bench.length > 0}{SEP}{/if}{#if misses.bench.length > 0}<span
                class="text-muted">Bench</span
              >
              {misses.bench.join(', ')}{/if}
          </p>
        {/if}
        <ul class="max-w-[560px]" data-testid={`guild-raid-fights-${f.row.id}`}>
          {#each nightFights(f.row, bossIds) as fight (fight.index)}
            <li
              class="grid min-h-8 grid-cols-[minmax(0,1fr)_48px_56px] items-center gap-3 text-[13px]"
              data-testid={fight.trash ? 'guild-raid-fight-trash' : 'guild-raid-fight'}
            >
              <span class={fight.trash || fight.unnamed ? 'text-muted' : 'text-strong font-semibold'}
                >{fight.label}</span
              >
              {#if fight.unnamed}
                <span class="text-muted">no named encounter published</span>
              {:else}
                <span class="text-muted numeric text-right">{clockDuration(fight.duration_ms)}</span>
                <span class="justify-self-end">
                  {#if fight.trash}<span class="text-muted">{TRASH_LABEL}</span>
                  {:else}<Tag tone={fight.kill ? 'success' : 'danger'}>{fight.kill ? 'Kill' : 'Wipe'}</Tag
                    >{/if}
                </span>
              {/if}
            </li>
          {/each}
        </ul>
      </div>
    {/if}
  {/each}
</StatTable>
