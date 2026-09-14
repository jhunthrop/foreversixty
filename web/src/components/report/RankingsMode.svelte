<!-- web/src/components/report/RankingsMode.svelte -->
<!-- Where this kill sits, from the same endpoint the standalone rankings pages read. This
     report's own row is marked, because the question is always "where am I on this
     list" -- and a moderation state (contract Amendments: ok | at_risk | removed) is
     printed inline rather than silently filtered, the same way a live fight says "Live"
     instead of hiding until it closes. -->
<script lang="ts">
  import { characterHref, guildHref, rulesetLabel, splitUnitName } from '../../lib/characters';
  import { classColorVar, formatAmount, formatDuration, percentileToken } from '../../lib/report/format';
  import { fetchRankings, type RankingMetric, type RankingsPage } from '../../lib/rankings/api';
  import type { FightEntry } from '../../lib/report/types';

  let {
    fight,
    reportId,
    encounterSlug,
  }: { fight: FightEntry; reportId: string; encounterSlug: string } = $props();

  let metric = $state<RankingMetric>('dps');
  let page = $state<RankingsPage | null>(null);
  let status = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let error = $state('');

  /**
   * The metric picker is not the only thing that can change while a request is in
   * flight: the visitor can pick a different fight in the selector, which stays visible
   * in every mode. Two rankings requests can therefore be in flight at once and settle
   * in either order, the same hazard ReportView.svelte's own Effect 3 guards against on
   * both its success and its rejection path. `wantedFight` and `wantedMetric`, captured
   * before the request and re-checked against the live props/state after it, are that
   * guard here: an answer for a fight or a metric nobody has selected any more is
   * dropped rather than painted under the current selection's name, on both paths.
   */
  $effect(() => {
    const wantedMetric = metric;
    const wantedFight = fight.index;
    const encounterId = fight.encounter_id;
    const difficulty = fight.difficulty;
    if (encounterId === undefined) {
      status = 'idle';
      return;
    }
    status = 'loading';
    void fetchRankings({ encounter: encounterSlug, difficulty, metric: wantedMetric, page: 1 })
      .then((result) => {
        if (fight.index !== wantedFight || metric !== wantedMetric) return;
        page = result;
        status = 'ready';
      })
      .catch((thrown: unknown) => {
        if (fight.index !== wantedFight || metric !== wantedMetric) return;
        status = 'failed';
        error = thrown instanceof Error ? thrown.message : 'Rankings did not load.';
      });
  });
</script>

<div class="flex flex-col gap-3" data-testid="rankings-mode">
  {#if fight.encounter_id === undefined}
    <p class="text-muted text-[14px]">Trash is not ranked. Pick an encounter to see its rankings.</p>
  {:else}
    <div class="flex flex-wrap items-center gap-x-4 gap-y-2">
      <label class="label text-muted flex items-center gap-2" for="rankings-metric">
        Metric
        <select
          id="rankings-metric"
          class="border-line-warm bg-raised rounded-control text-text h-11 px-2 text-[13px] md:h-9"
          bind:value={metric}
          data-testid="rankings-metric"
        >
          <option value="dps">Damage</option>
          <option value="hps">Healing</option>
          <option value="damage_taken">Damage taken</option>
        </select>
      </label>
      <a class="text-[13px]" href={`/rankings/${encounterSlug}`}>Full rankings for {fight.name}</a>
    </div>

    {#if status === 'loading'}
      <p class="text-muted text-[14px]">Loading rankings.</p>
    {:else if status === 'failed'}
      <p class="text-[14px]" role="alert">{error}</p>
    {:else if page !== null}
      <p class="text-muted text-[13px]">
        <span class="font-mono tabular">{page.total}</span> ranked kills · updated
        <span class="font-mono tabular">{page.updated_at.slice(0, 10)}</span>
      </p>
      <ul class="flex flex-col" data-testid="rankings-rows">
        {#each page.rows as row (`${row.report_id}-${row.fight_index}-${row.player.key}`)}
          <li
            class="border-line-soft grid min-h-11 grid-cols-[40px_minmax(0,1fr)_auto] items-center gap-x-3 gap-y-1 border-b px-2 py-2 text-[14px] md:grid-cols-[40px_minmax(120px,1.4fr)_minmax(100px,1fr)_88px_96px_72px]"
            class:bg-card-top={row.report_id === reportId}
            data-testid={row.report_id === reportId ? 'rankings-mine' : 'rankings-row'}
          >
            <span
              class="font-mono tabular text-[12px]"
              style={`color: ${percentileToken(Math.max(0, 100 - ((row.rank - 1) / Math.max(page.total, 1)) * 100))}`}
            >
              <!-- Decorative and approximate: a rank's position on this page, not the
                   API's own percentile (GET /v1/rankings/percentile). Good enough to tint
                   the digit; not a substitute for that endpoint if a precise percentile is
                   ever needed here. -->
              {row.rank}
            </span>
            <span class="flex min-w-0 items-center gap-2">
              <a
                class="truncate font-semibold"
                style={`color: ${classColorVar(row.player.class)}`}
                href={characterHref(row.guild?.region ?? 'us', row.guild?.ruleset ?? 'normal', row.player.name)}
              >
                {splitUnitName(row.player.name).name}
              </a>
              {#if row.report_id === reportId}
                <!-- The row's own background also marks this, but a colour alone is not a
                     safe way to say "this is the report you are reading": this text does
                     the same job in words. -->
                <span class="pill pill-site shrink-0">This report</span>
              {/if}
            </span>
            <span class="text-muted truncate text-[13px]">
              {#if row.guild}
                <a href={guildHref(row.guild.region, row.guild.ruleset, row.guild.name)}>{row.guild.name}</a>
                · {rulesetLabel(row.guild.ruleset)}
              {/if}
            </span>
            <span class="text-muted font-mono tabular text-right text-[13px]">{row.talent_split}</span>
            <span class="font-mono tabular text-right">{formatAmount(Math.round(row.value))}</span>
            <span class="text-muted font-mono tabular text-right text-[13px]">{formatDuration(row.duration_ms)}</span>
            {#if row.state !== 'ok'}
              <span class="pill pill-sample col-span-full md:col-auto">{row.state.replace('_', ' ')}</span>
            {/if}
          </li>
        {/each}
      </ul>
    {/if}
  {/if}
</div>
