<!-- web/src/components/HomeNextSteps.svelte -->
<!-- Spec 2026-09-24 §2.3: the four signed-in cards, each keyed to the main character.
     Occludes the same-shaped signed-out "four things" grid index.astro renders statically,
     the identical grid-overlay trick HomeAccountPanel already uses for the hero -- kept as
     its own small effect here rather than an extraction, since the two occluded elements
     (#home-signed-out, #home-next-steps-signed-out) are different ids on different DOM
     nodes and the eight-line body has nothing left to share once that's the only variable. -->
<script lang="ts">
  import { createHomeHero } from '../lib/account/home-hero.svelte';
  import { armorySimHref } from '../lib/sim/url';
  import { listMySims, fetchSimInput } from '../lib/sim/api';
  import { listMyReports, type MyReport } from '../lib/account/api';
  import { classSlugFromName } from '../lib/report/tree-sizes';
  import { logsCardLine, plannerPointsLabel, ratingCardValue, simCardLine } from '../lib/home/next-steps';
  import { homePanelCopy } from '../lib/home-panel-copy';
  import { homeProductPanels } from '../lib/home-landing-copy';
  import type { SimInput, SimListRow } from '../lib/sim/types';

  export const HOME_NEXT_STEPS_SIGNED_OUT_ID = 'home-next-steps-signed-out';

  const homeHero = createHomeHero();
  const me = $derived(homeHero.me);
  const ready = $derived(homeHero.ready);
  const hero = $derived(homeHero.hero);
  const heroPath = $derived(homeHero.heroPath);
  const rating = $derived(homeHero.rating);

  $effect(() => {
    const el = document.getElementById(HOME_NEXT_STEPS_SIGNED_OUT_ID);
    if (el === null) return;
    if (ready && me !== null && hero !== null) {
      el.setAttribute('inert', '');
      el.setAttribute('aria-hidden', 'true');
    } else {
      el.removeAttribute('inert');
      el.removeAttribute('aria-hidden');
      if (ready) el.removeAttribute('data-session-hide');
    }
  });

  // Per-character reads, one effect keyed on the hero object's own reference: Svelte's
  // $effect only re-runs when a tracked value's identity changes, so this never re-fires
  // for the same hero (Global Constraints: "must not re-run for the same object").
  let latestSim = $state<SimListRow | null>(null);
  let simStatus = $state<'loading' | 'ready' | 'failed'>('loading');
  let simInput = $state<SimInput | null>(null);
  let simInputStatus = $state<'loading' | 'ready' | 'failed' | 'none'>('none');
  let latestReport = $state<MyReport | null>(null);
  let reportsStatus = $state<'loading' | 'ready' | 'failed'>('loading');

  $effect(() => {
    const current = hero;
    const path = heroPath;
    if (current === null) return;

    simStatus = 'loading';
    void listMySims(1)
      .then((page) => {
        if (current !== hero) return;
        latestSim = page.rows[0] ?? null;
        simStatus = 'ready';
      })
      .catch(() => {
        if (current !== hero) return;
        simStatus = 'failed';
      });

    reportsStatus = 'loading';
    void listMyReports(1)
      .then((page) => {
        if (current !== hero) return;
        latestReport = page.rows[0] ?? null;
        reportsStatus = 'ready';
      })
      .catch(() => {
        if (current !== hero) return;
        reportsStatus = 'failed';
      });

    if (current.build === undefined || path === null) {
      simInput = null;
      simInputStatus = 'none';
      return;
    }
    simInputStatus = 'loading';
    void fetchSimInput(path)
      .then((input) => {
        if (current !== hero) return;
        simInput = input;
        simInputStatus = 'ready';
      })
      .catch(() => {
        if (current !== hero) return;
        simInputStatus = 'failed';
      });
  });

  const [plannerPanel, simPanel, logsPanel, rankingsPanel] = homeProductPanels;
  const plannerPoints = $derived(simInput === null ? '' : plannerPointsLabel(simInput.talents));
  const ratingValue = $derived(ratingCardValue(rating));
  const classSlug = $derived(hero?.class === undefined ? '' : classSlugFromName(hero.class));
</script>

{#if ready && me !== null && hero !== null}
  <div
    class="reveal grid grid-cols-1 items-stretch gap-4 [grid-area:1/1] sm:grid-cols-2 lg:grid-cols-4"
    data-testid="home-next-steps"
  >
    <div
      class="bg-raised border-line rounded-panel flex flex-col gap-2 border px-[18px] py-[16px]"
      data-testid="home-next-simulator"
    >
      <div class="flex items-center gap-2">
        <span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"
        ></span>
        <span class="label text-gold">{simPanel.label}</span>
      </div>
      <p class="text-muted text-[13px]">{simPanel.sentence}</p>
      {#if simStatus === 'loading'}
        <span class="skeleton-block h-4 w-2/3" aria-hidden="true"></span>
      {:else if latestSim !== null}
        <span class="text-strong text-[13px] font-semibold" data-testid="home-next-simulator-value"
          >{simCardLine(latestSim)}</span
        >
      {:else}
        <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
      {/if}
      <a class="text-nav mt-auto w-fit text-[13px] font-semibold" href={armorySimHref(hero.key)}>
        {latestSim === null ? homePanelCopy.runAction : simPanel.linkLabel}
      </a>
    </div>

    <div
      class="bg-raised border-line rounded-panel flex flex-col gap-2 border px-[18px] py-[16px]"
      data-testid="home-next-planner"
    >
      <div class="flex items-center gap-2">
        <span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"
        ></span>
        <span class="label text-gold">{plannerPanel.label}</span>
      </div>
      <p class="text-muted text-[13px]">{plannerPanel.sentence}</p>
      {#if simInputStatus === 'loading'}
        <span class="skeleton-block h-4 w-1/2" aria-hidden="true"></span>
      {:else if plannerPoints !== ''}
        <span class="text-strong text-[13px] font-semibold" data-testid="home-next-planner-value"
          >{plannerPoints}</span
        >
      {:else}
        <span class="text-muted text-[13px]">{homePanelCopy.noBuildYet}</span>
      {/if}
      <a
        class="text-nav mt-auto w-fit text-[13px] font-semibold"
        href={hero.build === undefined ? '/account#characters' : '/planner'}
      >
        {hero.build === undefined ? homePanelCopy.getTheBuild : homePanelCopy.continuePlanning}
      </a>
    </div>

    <div
      class="bg-raised border-line rounded-panel flex flex-col gap-2 border px-[18px] py-[16px]"
      data-testid="home-next-logs"
    >
      <div class="flex items-center gap-2">
        <span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"
        ></span>
        <span class="label text-gold">{logsPanel.label}</span>
      </div>
      <p class="text-muted text-[13px]">{logsPanel.sentence}</p>
      {#if reportsStatus === 'loading'}
        <span class="skeleton-block h-4 w-3/4" aria-hidden="true"></span>
      {:else if latestReport !== null}
        <span class="text-strong text-[13px] font-semibold" data-testid="home-next-logs-value"
          >{logsCardLine(latestReport)}</span
        >
      {:else}
        <span class="text-muted text-[13px]">{homePanelCopy.noLogsYet}</span>
      {/if}
      <a class="text-nav mt-auto w-fit text-[13px] font-semibold" href="/logs">
        {latestReport === null ? homePanelCopy.uploadALog : homePanelCopy.openAction}
      </a>
    </div>

    <div
      class="bg-raised border-line rounded-panel flex flex-col gap-2 border px-[18px] py-[16px]"
      data-testid="home-next-rankings"
    >
      <div class="flex items-center gap-2">
        <span class="bg-gold h-2 w-2 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]" aria-hidden="true"
        ></span>
        <span class="label text-gold">{rankingsPanel.label}</span>
      </div>
      <p class="text-muted text-[13px]">{rankingsPanel.sentence}</p>
      {#if ratingValue !== ''}
        <span class="text-strong text-[13px] font-semibold" data-testid="home-next-rankings-value"
          >{ratingValue}</span
        >
      {:else}
        <span class="text-muted text-[13px]">{homePanelCopy.notRatedYet}</span>
      {/if}
      <a class="text-nav mt-auto w-fit text-[13px] font-semibold" href={`/rankings?class=${classSlug}`}>
        {homePanelCopy.rankingsForClass(hero.class ?? '')}
      </a>
    </div>
  </div>
{:else}
  <div class="pointer-events-none min-h-[220px] [grid-area:1/1]" aria-hidden="true"></div>
{/if}
