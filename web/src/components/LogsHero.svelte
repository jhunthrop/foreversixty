<!-- web/src/components/LogsHero.svelte -->
<!-- Logs landing spec (2026-10-04) §4.A: the header band's own content -- a real report,
     never a blank title. Signed in with a report of your own, that report (newest first,
     `listMyReports(1).rows[0]`); every other case, the canonical sample
     (`data/sample-report.json`), fetched through the public `GET /v1/reports/{id}`.

     The signed-in hero also carries §4.A.1's per-player hook: a *third*, deferred fetch
     (the hero report's own last kill, then that fight's summary.json) that must never
     block the hero's own first paint -- it starts only after the hero's primary content
     (title, facts, button) is already known, and fades in on its own with `.reveal` or is
     omitted entirely on failure (never an error banner for a bonus fact). The sample hero
     never runs this fetch at all (§4.A.1's own reasoning: a third round trip on every
     anonymous pageview is a real LCP/TBT cost for a fact that is not about the visitor's
     own character); it shows the static sentence in the hook's place instead. -->
<script lang="ts">
  import activeBuild from '../data/active-build.json';
  import sampleReport from '../data/sample-report.json';
  import { fetchMeOnce, listMyReports, type Me } from '../lib/account/api';
  import { logsHeroCopy } from '../lib/reports/copy';
  import {
    heroFactsFromMyReport,
    heroFactsFromReportMeta,
    heroTitle,
    reportDay,
    selectOwnHero,
    type HeroFacts,
  } from '../lib/reports/hero';
  import { buildHookLinks, pickLastKillFight, selectHookRow, type HookLinks } from '../lib/reports/hero-hook';
  import { loadTalents } from '../lib/planner/load';
  import { fetchReportMeta, fetchSummary } from '../lib/report/load';
  import { resolveTreeSizes } from '../lib/report/tree-sizes';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let status = $state<'loading' | 'ready' | 'failed'>('loading');
  let facts = $state<HeroFacts | null>(null);
  let isSample = $state(false);

  let hookStatus = $state<'idle' | 'loading' | 'ready' | 'failed'>('idle');
  let hook = $state<HookLinks | null>(null);
  let hookFightName = $state('');

  async function loadHook(reportId: string, me: Me): Promise<void> {
    hookStatus = 'loading';
    try {
      const meta = await fetchReportMeta(reportId);
      const fight = pickLastKillFight(meta.fights);
      if (fight === null) {
        hookStatus = 'failed';
        return;
      }
      const summary = await fetchSummary(meta.data_base_url, fight.index, meta.engine_version);
      const row = selectHookRow(
        summary.roster,
        me.characters.map((character) => character.name),
      );
      if (row === null) {
        hookStatus = 'failed';
        return;
      }
      // Only the one class this row needs -- the report page's own effect (ReportView.svelte)
      // resolves every class a fight's roster contains because its tables render every row;
      // this hook renders exactly one.
      const wanted = row.class === undefined ? [] : [row.class];
      const resolved = await resolveTreeSizes(wanted, new Set(), loadTalents, activeBuild.build);
      const sizes = new Map(resolved);
      hook = buildHookLinks({
        reportId,
        fightIndex: fight.index,
        dataBuild: activeBuild.build,
        row,
        combatants: summary.combatants,
        treeSizesFor: (className) => (className === undefined ? [] : (sizes.get(className) ?? [])),
      });
      hookFightName = fight.name;
      hookStatus = 'ready';
    } catch {
      // A bonus fact, not the region's job (§4.A.1): the hero's own primary content never
      // waited on this, and a failure here is silently omitted, not an error banner.
      hookStatus = 'failed';
    }
  }

  async function load(): Promise<void> {
    status = 'loading';
    hookStatus = 'idle';
    hook = null;
    try {
      const me = await fetchMeOnce();
      const own = me === null ? null : selectOwnHero((await listMyReports(1)).rows);
      if (own !== null && me !== null) {
        isSample = false;
        facts = heroFactsFromMyReport(own);
        status = 'ready';
        void loadHook(own.id, me);
      } else {
        isSample = true;
        const meta = await fetchReportMeta(sampleReport.id);
        facts = heroFactsFromReportMeta(meta);
        status = 'ready';
      }
    } catch {
      status = 'failed';
    }
  }

  $effect(() => {
    void load();
  });

  const title = $derived(facts === null ? '' : heroTitle(facts.title, facts.zone));
</script>

<!-- §8: no layout shift on hydration. The outer wrapper itself carries the reserved
     height (not just the Skeleton) so loading -> ready AND loading -> failed both land
     on the same floor -- measured against the real ready hero at 360px (mobile, the
     Lighthouse/lighthouserc.json emulation width) and 1024px+ (desktop, where the
     description wraps to fewer lines): 292px below 1024px, 232px from it. -->
<div class="flex min-h-[292px] flex-col gap-3 lg:min-h-[232px]" data-testid="logs-hero">
  {#if status === 'loading'}
    <Skeleton lines={4} rowHeight="h-4" testid="logs-hero-skeleton" />
  {:else if status === 'failed'}
    <LoadError message={logsHeroCopy.failed} onRetry={() => void load()} testid="logs-hero-error" />
  {:else if facts !== null}
    <p class="label flex items-center gap-[10px]" style="color: var(--color-gold)">
      <i class="inline-block h-px w-[28px]" style="background: var(--color-gold)" aria-hidden="true"></i>
      {logsHeroCopy.eyebrow}
    </p>
    <h1 class="font-display text-strong text-[22px] leading-[1.1] font-bold" data-testid="logs-hero-title">
      {title}{#if isSample}<span
          class="pill pill-sample ml-[10px] align-middle"
          data-testid="logs-hero-sample-pill">Sample</span
        >{/if}
    </h1>
    <p class="tabular text-muted font-mono text-[14px]" data-testid="logs-hero-facts">
      {reportDay(facts.createdAt)} · {facts.fightCount} fights · {facts.killCount} kills
    </p>
    <p class="max-w-[62ch] text-[14px]" style="color: #c9c2b2" data-testid="logs-hero-description">
      {logsHeroCopy.description}
      <a class="text-nav underline" href="#companion">Log live with the companion</a>
      or
      <a class="text-nav underline" href="#upload">upload a log file</a>.
    </p>
    <div class="min-h-[20px]" data-testid="logs-hero-hook">
      {#if isSample}
        <p class="text-muted reveal text-[13px]">{logsHeroCopy.sampleHookLine}</p>
      {:else if hookStatus === 'loading'}
        <Skeleton lines={1} rowHeight="h-5" testid="logs-hero-hook-skeleton" />
      {:else if hookStatus === 'ready' && hook !== null}
        <p
          class="reveal text-muted flex flex-wrap items-center gap-2 text-[13px]"
          data-testid="logs-hero-hook-line"
        >
          <span>Last kill, {hookFightName}:</span>
          <span class="font-display font-bold" style={`color:${hook.classColor}`}>{hook.name}</span>
          <span aria-hidden="true">·</span>
          <span class="tabular text-text font-mono">{hook.dpsText} DPS</span>
          {#if hook.plannerLink !== null}
            <span aria-hidden="true">·</span>
            <a
              class="text-gold text-[12px] font-bold tracking-[0.06em] uppercase"
              href={hook.plannerLink.href}
            >
              {hook.plannerLink.label}
            </a>
          {/if}
          <span aria-hidden="true">|</span>
          <a class="text-gold text-[12px] font-bold tracking-[0.06em] uppercase" href={hook.simLink.href}>
            {hook.simLink.label}
          </a>
        </p>
      {/if}
    </div>
    <a
      class="{SECONDARY_BUTTON_FIXED} border-line-warm-strong text-gold w-fit px-4"
      href={`/reports/${facts.id}`}
      data-testid="logs-hero-open"
    >
      {logsHeroCopy.openButton}
    </a>
  {/if}
</div>
