<!-- web/src/components/guild/GuildProgression.svelte -->
<!-- Guild control-centre spec §4.D: tier bar with an honest caveat, Onyxia's own
     pulls-to-kill trend, and the Barrow Deeps/Hyjal Summit pulls-only panel. Shown to
     every role. -->
<script lang="ts">
  import { characterHref, splitUnitName, type Region, type Ruleset } from '../../lib/characters';
  import { classColorVar, rowLink } from '../../lib/report/format';
  import { dayWord } from '../../lib/dates';
  import type { GuildParseRef, GuildProgressionPage } from '../../lib/guild/api';
  import { progressionNightDay } from '../../lib/guild/overview-view';
  import { rosterBestLabels } from '../../lib/guild/roster-best';
  import { counted } from '../../lib/plural';
  import type { GuildRosterBest } from '../../lib/rankings/api';
  import { executionLabel, executionTitle } from '../../lib/sim/execution';
  import EmptyState from '../ui/EmptyState.svelte';

  /** `rosterBest` is v1's own "Roster bests" feature (`GET .../<guild>`'s existing,
   *  working `roster_best` rows) -- carried forward unchanged here (tenet 10) rather than
   *  dropped while the new control-centre `/progression` endpoint is still deploying. */
  let {
    progression,
    rosterBest = [],
    region,
    ruleset,
  }: {
    progression: GuildProgressionPage | null;
    rosterBest?: GuildRosterBest[];
    region: Region;
    ruleset: Ruleset;
  } = $props();

  /** One row per raider per encounter: the API also sends each raider's best with no encounter
   *  behind it (id 0, no name), which reads as a second row of dashes and is not a boss. */
  const encounterBests = $derived(rosterBest.filter((row) => row.encounter_id > 0 && row.encounter !== ''));

  /** The tank, healer and damage bests that exist, in that order. */
  function bestParses(best: {
    tank?: GuildParseRef;
    healer?: GuildParseRef;
    dps?: GuildParseRef;
  }): GuildParseRef[] {
    return [best.tank, best.healer, best.dps].filter((parse): parse is GuildParseRef => parse !== undefined);
  }

  function clock(ms: number | null): string {
    if (ms === null) return '—';
    const totalSeconds = Math.round(ms / 1000);
    return `${Math.floor(totalSeconds / 60)}:${String(totalSeconds % 60).padStart(2, '0')}`;
  }
</script>

<section class="flex flex-col gap-6" data-testid="guild-progression-tab">
  <h2 class="section-title">Progression</h2>
  {#if progression === null}
    <EmptyState message="No pulls recorded yet." testid="guild-progression-empty" />
  {:else}
    {@const tier = progression.tier}
    <div class="flex flex-col gap-2" data-testid="guild-progression-tier">
      <div class="text-muted flex justify-between text-[13px]">
        <span>Named encounters down this tier</span>
        <span class="numeric text-[14px]"
          >{tier.down} · {progression.encounters.map((e) => e.name).join(', ') || '—'}</span
        >
      </div>
      <div class="border-line-soft h-1.5 overflow-hidden rounded-full">
        <div
          class="bg-gold h-full"
          style={`width:${tier.named_encounters === 0 ? 0 : Math.min(100, (tier.down / tier.named_encounters) * 100)}%`}
        ></div>
      </div>
      {#if progression.unnamed.length > 0}
        <!-- Our to-do, so only when this guild has pulls it applies to (P.1, player review S2-5). -->
        <p class="text-muted text-[12px]" data-testid="guild-progression-unnamed-note">
          Barrow Deeps and Hyjal Summit have no encounter list yet. Their pulls count toward the tier total
          above but cannot be named or ranked per boss.
        </p>
      {/if}
    </div>

    {#each progression.encounters as encounter (encounter.encounter_id)}
      <div
        class="border-line-soft bg-raised flex flex-col gap-3 rounded-md border p-5"
        data-testid="guild-progression-encounter"
      >
        <h3 class="section-title">{encounter.name} · pulls to kill</h3>
        {#if encounter.first_kill_at !== null}
          <p class="text-[13px]">
            First kill {dayWord(new Date(encounter.first_kill_at))} · best kill time {clock(
              encounter.best_kill_ms,
            )}
          </p>
        {/if}
        <div class="flex flex-col">
          {#each encounter.pulls_by_night as night (night.report_id)}
            <div class="border-line-soft flex items-center gap-4 border-b py-1.5">
              <span class="text-muted tabular min-w-[90px] text-[13px]"
                >{progressionNightDay(night.date)}</span
              >
              <span class="text-[13px]">{counted(night.pulls, 'pull')}</span>
              {#if night.killed}
                <span
                  class="pill"
                  style="color:#7bff5c;background:rgba(30,255,0,.10);border-color:rgba(30,255,0,.25)"
                  >Kill</span
                >
              {/if}
            </div>
          {/each}
        </div>
        <p class="text-[13px]">
          Deaths per pull: {encounter.deaths_per_pull.toFixed(1)} on average. We do not track what killed people
          yet.
        </p>
        {#if bestParses(encounter.best_by_role).length > 0}
          <p class="text-[13px]" data-testid="guild-progression-best-parse">
            Guild best parse:
            {#each bestParses(encounter.best_by_role) as parse, index (index)}
              {#if index > 0}<span class="hidden md:inline" aria-hidden="true">&nbsp;·&nbsp;</span>{/if}
              <span class="block md:inline">
                <span style={`color:${classColorVar(parse.class)};font-weight:600`}>{parse.name}</span>
                {parse.value.toFixed(1)}
                {parse.metric.toUpperCase()}
              </span>
            {/each}
          </p>
        {/if}
      </div>
    {/each}

    {#if progression.unnamed.length > 0}
      <div
        class="border-line-soft bg-raised flex flex-col gap-2 rounded-md border p-5"
        data-testid="guild-progression-unnamed"
      >
        <h3 class="section-title">Barrow Deeps / Hyjal Summit</h3>
        {#each progression.unnamed as zone (zone.zone)}
          <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
            <span class="flex-1 text-[14px] font-semibold">{zone.zone}</span>
            <span class="text-muted numeric text-[13px]">{counted(zone.pulls, 'pull')} logged</span>
            <span class="text-muted text-[12px]">no named encounters yet</span>
          </div>
        {/each}
      </div>
    {/if}
  {/if}

  {#if encounterBests.length > 0}
    <div class="flex flex-col gap-2" data-testid="guild-roster-best">
      <h3 class="section-title">Roster bests</h3>
      {#each encounterBests as row (`${row.player.key}-${row.encounter_id}-${row.metric}`)}
        {@const labels = rosterBestLabels(row)}
        <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
          <a
            class="{rowLink} truncate font-semibold"
            style={`color:${classColorVar(row.player.class)}`}
            href={characterHref(region, ruleset, row.player.name)}
          >
            {splitUnitName(row.player.name).name}
          </a>
          <span class="text-muted min-w-0 truncate text-[13px]">{labels.what}</span>
          <span class="numeric ml-auto text-[14px]" title={labels.hover}>{labels.value}</span>
          <span
            class="text-muted numeric hidden text-right text-[13px] md:inline"
            title={executionTitle(row.execution_score)}
          >
            {executionLabel(row.execution_score)}
          </span>
        </div>
      {/each}
    </div>
  {/if}
</section>
