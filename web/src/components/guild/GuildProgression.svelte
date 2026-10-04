<!-- web/src/components/guild/GuildProgression.svelte -->
<!-- Guild control-centre spec §4.D: tier bar with an honest caveat, Onyxia's own
     pulls-to-kill trend, and the Barrow Deeps/Hyjal Summit pulls-only panel. Shown to
     every role. -->
<script lang="ts">
  import { characterHref, splitUnitName, type Region, type Ruleset } from '../../lib/characters';
  import { classColorVar, formatAmount, rowLink } from '../../lib/report/format';
  import type { GuildProgressionPage } from '../../lib/guild/api';
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

  function clock(ms: number | null): string {
    if (ms === null) return '—';
    const totalSeconds = Math.round(ms / 1000);
    return `${Math.floor(totalSeconds / 60)}:${String(totalSeconds % 60).padStart(2, '0')}`;
  }
</script>

<section class="flex flex-col gap-6" data-testid="guild-progression-tab">
  <h2 class="section-title text-[18px]">Progression</h2>
  {#if progression === null}
    <EmptyState message="No pulls recorded yet." testid="guild-progression-empty" />
  {:else}
    {@const tier = progression.tier}
    <div class="flex flex-col gap-2" data-testid="guild-progression-tier">
      <div class="text-muted flex justify-between text-[13px]">
        <span>Named encounters down this tier</span>
        <span class="tabular font-mono text-[14px]"
          >{tier.down} · {progression.encounters.map((e) => e.name).join(', ') || '—'}</span
        >
      </div>
      <div class="border-line-soft h-1.5 overflow-hidden rounded-full">
        <div
          class="bg-gold h-full"
          style={`width:${tier.named_encounters === 0 ? 0 : Math.min(100, (tier.down / tier.named_encounters) * 100)}%`}
        ></div>
      </div>
      <p class="text-muted text-[12px]">
        Barrow Deeps and Hyjal Summit have no published encounter list yet
        (data/curated/loot/forever-raid-phases.json) — their pulls count toward the tier total above but
        cannot be named or ranked per boss until Forever publishes them.
      </p>
    </div>

    {#each progression.encounters as encounter (encounter.encounter_id)}
      <div
        class="border-line-soft bg-raised flex flex-col gap-3 rounded-md border p-5"
        data-testid="guild-progression-encounter"
      >
        <h3 class="section-title text-[16px]">{encounter.name} · pulls to kill</h3>
        {#if encounter.first_kill_at !== null}
          <p class="text-[13px]">
            First kill {encounter.first_kill_at.slice(0, 10)} · best kill time {clock(encounter.best_kill_ms)}
          </p>
        {/if}
        <div class="flex flex-col">
          {#each encounter.pulls_by_night as night (night.report_id)}
            <div class="border-line-soft flex items-center gap-4 border-b py-1.5">
              <span class="text-muted tabular min-w-[90px] font-mono text-[13px]">{night.date}</span>
              <span class="text-[13px]">{night.pulls} pull{night.pulls === 1 ? '' : 's'}</span>
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
          Deaths per pull: {encounter.deaths_per_pull.toFixed(1)} average. Top death causes need the report's own
          event data (which spell killed whom) — death <em>count</em> is stored today, the <em>cause</em> is
          not (<em>NEEDS NEW DATA</em>).
        </p>
        {#if encounter.best_by_role.dps !== undefined || encounter.best_by_role.healer !== undefined || encounter.best_by_role.tank !== undefined}
          <p class="text-[13px]">
            Guild best parse:
            {#each Object.entries(encounter.best_by_role) as [role, parse] (role)}
              {#if parse !== undefined}
                <span style={`color:${classColorVar(parse.class)};font-weight:600`}>{parse.name}</span>
                {parse.value.toFixed(1)}
                {parse.metric.toUpperCase()}
              {/if}
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
        <h3 class="section-title text-[16px]">Barrow Deeps / Hyjal Summit</h3>
        {#each progression.unnamed as zone (zone.zone)}
          <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
            <span class="flex-1 text-[14px] font-semibold">{zone.zone}</span>
            <span class="text-muted tabular font-mono text-[13px]">{zone.pulls} pulls logged</span>
            <span class="text-muted text-[12px]">no named encounters yet</span>
          </div>
        {/each}
      </div>
    {/if}
  {/if}

  {#if rosterBest.length > 0}
    <div class="flex flex-col gap-2" data-testid="guild-roster-best">
      <h3 class="section-title text-[16px]">Roster bests</h3>
      {#each rosterBest as row (`${row.player.key}-${row.encounter_id}-${row.metric}`)}
        <div class="border-line-soft flex items-center gap-3 border-b py-1.5">
          <a
            class="{rowLink} truncate font-semibold"
            style={`color:${classColorVar(row.player.class)}`}
            href={characterHref(region, ruleset, row.player.name)}
          >
            {splitUnitName(row.player.name).name}
          </a>
          <span class="text-muted hidden truncate text-[13px] md:inline"
            >{row.encounter} · {row.player.spec}</span
          >
          <span class="tabular ml-auto font-mono text-[14px]">{formatAmount(Math.round(row.value))}</span>
          <span
            class="text-muted tabular hidden text-right font-mono text-[13px] md:inline"
            title={executionTitle(row.execution_score)}
          >
            {executionLabel(row.execution_score)}
          </span>
        </div>
      {/each}
    </div>
  {/if}
</section>
