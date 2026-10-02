<!-- web/src/components/character/HomeHeroCards.svelte -->
<!-- Home rebuild spec §3.B.2: the signed-in hero's three next-action cards (Best in slot /
     Talents / Simulator). Extracted out of HomeAccountPanel.svelte (review round 3,
     "islands" item): the Simulator card's own fetch (lib/sim/api.ts's `listMySims`) and its
     dependencies (lib/home/next-steps.ts's `simCardLine`, ui/Skeleton.svelte) were
     previously imported at HomeAccountPanel's own top level, so every signed-out visitor's
     browser downloaded that code too, even though the component never reaches the branch
     that uses it -- Vite/Rollup bundles a file's static imports into its own chunk
     regardless of which conditional branch actually runs. Mounted by HomeAccountPanel.svelte
     into its own `home-hero-cards-slot` only once `ready && me !== null && hero !== null`,
     the same dynamic-import-once-signed-in trick it already uses for
     HomeSwitchCharacterPanel, so a signed-out page never fetches this module or its
     dependencies either.

     Best in slot and Talents now read a real comparison (`lib/home/upgrades.ts`/
     `talent-delta.ts`, over a band fetched once by `lib/home/upgrades-loader.ts` and shared
     by both cards -- the identical real gap when it is missing, the identical band once it
     is not). Three honest non-figure states remain, each its own real reason (never a
     generic spinner forever, tenet 8): no spec at all, a spec but no addon export yet, or a
     spec and an export but no published BiS file/band for it. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { homePanelCopy, homeHeroCardsCopy } from '../../lib/home-panel-copy';
  import { armorySimHref } from '../../lib/sim/url';
  import { listMySims } from '../../lib/sim/api';
  import { simCardLine } from '../../lib/home/next-steps';
  import type { SimListRow } from '../../lib/sim/types';
  import Skeleton from '../ui/Skeleton.svelte';
  import { loadBisContextFor, specKeyForCharacter, type BisContext } from '../../lib/home/upgrades-loader';
  import { upgradesFor } from '../../lib/home/upgrades';
  import { talentDeltaFor } from '../../lib/home/talent-delta';
  import { bisCopy } from '../../lib/bis/copy';

  let { hero }: { hero: MeCharacter } = $props();

  // The Simulator card's own data: the visitor's latest saved sim, independent per-card
  // loading, the same rule HomeNextSteps.svelte's own cards used to follow before that
  // component was removed.
  let latestSim = $state<SimListRow | null>(null);
  let simStatus = $state<'loading' | 'ready' | 'failed'>('loading');

  $effect(() => {
    const current = hero;
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
  });

  type BandStatus = 'no-spec' | 'loading' | 'no-list-yet' | 'ready';

  // Determined synchronously from `hero.spec` alone, not inside the `$effect` below: SSR
  // runs no effects at all, and "no spec" is knowable from the prop the moment this
  // component renders, so there is no reason to show an infinite loading skeleton for a
  // fact that will never change without a different `hero`. `untrack` says explicitly that
  // only this INITIAL value of `hero` matters here -- the `$effect` below is what keeps
  // `bandStatus` correct for every later `hero` (a character switch), the same deliberate
  // read-once pattern `ItemHover.svelte`'s own `hoverId` already uses.
  let bandStatus = $state<BandStatus>(untrack(() => (hero.spec === undefined ? 'no-spec' : 'loading')));
  let ctx = $state<BisContext | null>(null);

  $effect(() => {
    const current = hero;
    ctx = null;
    if (current.spec === undefined) {
      bandStatus = 'no-spec';
      return;
    }
    bandStatus = 'loading';
    void loadBisContextFor(current).then((loaded) => {
      if (current !== hero) return;
      if (loaded === null) {
        bandStatus = 'no-list-yet';
        return;
      }
      ctx = loaded;
      bandStatus = 'ready';
    });
  });

  const bandLabel = $derived(ctx === null ? '' : bisCopy.bandRangeLabel(ctx.band.band));
  const upgrades = $derived(
    ctx === null || hero.build?.gear === undefined ? null : upgradesFor(hero, ctx.band, ctx.items),
  );
  const talentDelta = $derived(
    ctx === null || hero.build?.talents === undefined
      ? null
      : talentDeltaFor(hero.build.talents.trees, ctx.band.talents),
  );

  /** Planner rebuild spec §4.I: "Compare in the planner" carries `?spec=` so the planner
   *  opens already mid-comparison against the visitor's own band -- `?band=` is dropped
   *  (owner ruling §12; `BandCompare` always reads the build's own level band). The same
   *  class+spec-name -> canonical spec key lookup `loadBisContextFor` already uses for
   *  this exact character (`specKeyForCharacter`), not `hero.spec` itself -- that field is
   *  a display name ("Marksmanship"), not the `hunter-marksmanship` key `?spec=` needs. No
   *  spec learned yet, or an unrecognised class/spec pair, links to a bare planner. */
  const talentsHref = $derived.by(() => {
    const specKey = specKeyForCharacter(hero);
    return specKey === undefined ? '/planner' : `/planner?spec=${specKey}`;
  });

  const CARD_CLASS =
    'flex flex-col gap-[6px] p-[14px_16px] rounded-panel border border-line bg-gradient-to-b from-card-top to-raised shadow-[inset_0_-1px_0_rgba(229,185,85,.35)] text-strong hover:border-gold-deep transition-colors duration-150';

  /** Fix round 1 item B.4 (ux-designer): the card's own answer is the one primary element and
   *  reads at the design system's "stat figure" treatment (`design/DESIGN-SYSTEM.md`'s type
   *  table: "Display | Cinzel 600-800 | ... stat figures 22px") -- Cinzel, uppercase, 22px,
   *  not the `text-[15px] font-semibold` body-text size the mock's own three cards never use
   *  for their headline number. The line underneath it stays Barlow 12px muted, unchanged. */
  const CARD_FIGURE_CLASS = 'font-display text-[22px] font-bold uppercase tracking-[0.01em] text-strong';
</script>

<div class="reveal grid grid-cols-1 gap-3 sm:grid-cols-3" data-testid="home-hero-cards">
  <a href="#upgrades" class={CARD_CLASS} data-testid="home-hero-card-bis">
    <span class="label text-gold">{homeHeroCardsCopy.bestInSlotLabel}</span>
    {#if bandStatus === 'no-spec'}
      <span class="text-strong text-[15px] font-semibold">{homeHeroCardsCopy.pickASpec}</span>
      <span class="text-muted text-[12px]">{homeHeroCardsCopy.pickASpecLine}</span>
    {:else if hero.build?.gear === undefined}
      <span class="text-muted text-[13px]">{homeHeroCardsCopy.bestInSlotNotAvailable}</span>
    {:else if bandStatus === 'loading'}
      <!-- Fix round 1 item C.9 (tenet 13, no layout shift): sized for the ready figure's
           own 22px Cinzel line (`CARD_FIGURE_CLASS`), not the 15px body-text line this
           skeleton was built for before item B.4's display-stat treatment landed -- `h-3`
           left the ready state ~30px taller than its own loading skeleton. -->
      <Skeleton lines={3} rowHeight="h-4" testid="home-hero-card-bis-skeleton" />
    {:else if bandStatus === 'no-list-yet'}
      <span class="text-muted text-[13px]">{homeHeroCardsCopy.bestInSlotNoListYet}</span>
    {:else if upgrades !== null}
      {#if upgrades.upgrades.length === 0}
        <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-bis-value">
          {homeHeroCardsCopy.bestInSlotAllMatchFigure}
        </span>
        <span class="text-muted text-[12px]">{homeHeroCardsCopy.bestInSlotAllMatchLine(bandLabel)}</span>
      {:else}
        <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-bis-value">
          {homeHeroCardsCopy.bestInSlotFigure(upgrades.upgrades.length)}
        </span>
        <span class="text-muted text-[12px]">
          {homeHeroCardsCopy.bestInSlotUpgradesLine(
            bandLabel,
            upgrades.totalGainDps,
            upgrades.notSimCheckedCount,
          )}
        </span>
      {/if}
    {/if}
  </a>
  <a href={talentsHref} class={CARD_CLASS} data-testid="home-hero-card-talents">
    <span class="label text-gold">{homeHeroCardsCopy.talentsLabel}</span>
    {#if bandStatus === 'no-spec'}
      <span class="text-strong text-[15px] font-semibold">{homeHeroCardsCopy.pickASpec}</span>
      <span class="text-muted text-[12px]">{homeHeroCardsCopy.pickASpecLine}</span>
    {:else if hero.build?.talents === undefined}
      <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNotAvailable}</span>
    {:else if bandStatus === 'loading'}
      <!-- Fix round 1 item C.9: same height fix as the Best in slot card's own skeleton. -->
      <Skeleton lines={3} rowHeight="h-4" testid="home-hero-card-talents-skeleton" />
    {:else if bandStatus === 'no-list-yet'}
      <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNoListYet}</span>
    {:else if talentDelta !== null && ctx !== null}
      <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-talents-value">
        {talentDelta === 0
          ? homeHeroCardsCopy.talentsOptimizedFigure
          : homeHeroCardsCopy.talentsUnoptimizedFigure}
      </span>
      <span class="text-muted text-[12px]">
        {homeHeroCardsCopy.talentsLine(talentDelta, ctx.band.talent_points, bandLabel)}
      </span>
    {/if}
  </a>
  <a href={armorySimHref(hero.key)} class={CARD_CLASS} data-testid="home-hero-card-sim">
    <span class="label text-gold">{homeHeroCardsCopy.simulatorLabel}</span>
    {#if simStatus === 'loading'}
      <Skeleton lines={2} rowHeight="h-3" testid="home-hero-card-sim-skeleton" />
    {:else if simStatus === 'failed'}
      <span class="text-muted text-[12px]">{homePanelCopy.noSimYet}</span>
    {:else if latestSim !== null}
      <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-sim-value">
        {simCardLine(latestSim)}
      </span>
      {#if ctx !== null}
        <span class="text-muted font-mono text-[12px]"
          >{homeHeroCardsCopy.simulatorBandSuffix(ctx.band.set_dps)}</span
        >
      {/if}
    {:else}
      <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
      <span class="text-nav text-[12px] font-semibold">{homePanelCopy.runAction}</span>
    {/if}
  </a>
</div>
