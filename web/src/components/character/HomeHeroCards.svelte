<!-- web/src/components/character/HomeHeroCards.svelte -->
<!-- The signed-in home panel's two next-action cards, Talents and Simulator, as one row under the
     shared character header (character header spec 2026-10-11 section 4: the Best in slot card
     is the header's answer and button now). Extracted out of HomeAccountPanel.svelte (review
     round 3, "islands" item): the Simulator card's own fetch (lib/sim/api.ts's `listMySims`)
     and its dependencies (lib/report/format.ts, ui/Skeleton.svelte) would otherwise be bundled
     into every signed-out visitor's page. Mounted by HomeAccountPanel.svelte into its own
     `home-hero-cards-slot` only once signed in, the same dynamic-import-once-signed-in trick it
     uses for the upgrades table.

     Talents reads a real comparison (`lib/home/talent-delta.ts`, over a band fetched once by
     `lib/home/upgrades-loader.ts`). Three honest non-figure states remain, each its own real
     reason (never a generic spinner forever, tenet 8): no spec at all, a spec but no addon
     export yet, or a spec and an export but no published file/band for it. -->
<script lang="ts">
  import { untrack } from 'svelte';
  import type { MeCharacter } from '../../lib/account/api';
  import { homePanelCopy, homeHeroCardsCopy } from '../../lib/home-panel-copy';
  import { armorySimHref } from '../../lib/sim/url';
  import { listMySims } from '../../lib/sim/api';
  import { formatAmount } from '../../lib/report/format';
  import { simCardFigures } from '../../lib/home/sim-card';
  import type { SimListRow } from '../../lib/sim/types';
  import Skeleton from '../ui/Skeleton.svelte';
  import { loadBisContextFor, specKeyForCharacter, type BisContext } from '../../lib/home/upgrades-loader';
  import { talentDeltaFor } from '../../lib/home/talent-delta';
  import { ACTION_LINK, ACTION_LINK_CHEVRON } from '../../lib/ui/link';
  import Glyph from '../ui/Glyph.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import { presetLabelFor } from '../../lib/bis/presets';

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

  const simFigures = $derived(
    latestSim === null
      ? null
      : simCardFigures(latestSim, { specKey: specKeyForCharacter(hero), band: ctx?.band }),
  );
  const bandLabel = $derived(ctx === null ? '' : bisCopy.bandRangeLabel(ctx.band.band));
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
   *  reads at the design system's headline-figure treatment (`design/DESIGN-SYSTEM.md`'s type
   *  table; polish spec 2026-10-28 §4.4 moved figures to Barlow Condensed 800) -- 22px here,
   *  not the `text-[15px] font-semibold` body-text size the mock's own three cards never use
   *  for their headline number. The line underneath it stays Barlow 12px muted, unchanged. */
  const CARD_FIGURE_CLASS = 'figure-stat text-[22px] text-strong';
</script>

<!-- The card is itself the link to the simulator, so its action is an action-link look-alike (gold,
     chevron, 44 px) rather than a nested anchor. -->
{#snippet runAction()}
  <span class={ACTION_LINK} data-testid="home-hero-card-sim-run">
    {homePanelCopy.runAction}<Glyph name="chevron" size={16} class={ACTION_LINK_CHEVRON} />
  </span>
{/snippet}

<div class="reveal grid grid-cols-1 gap-3 sm:grid-cols-2" data-testid="home-hero-cards">
  <a href={talentsHref} class={CARD_CLASS} data-testid="home-hero-card-talents">
    <span class="label text-gold">{homeHeroCardsCopy.talentsLabel}</span>
    {#if bandStatus === 'no-spec'}
      <span class="text-strong text-[15px] font-semibold">{homeHeroCardsCopy.pickASpec}</span>
      <span class="text-muted text-[12px]">{homeHeroCardsCopy.pickASpecLine}</span>
    {:else if hero.build?.talents === undefined}
      <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNotAvailable}</span>
    {:else if bandStatus === 'loading'}
      <!-- Fix round 1 item C.9 (tenet 13, no layout shift): sized for the ready figure's own
           22px figure line (`CARD_FIGURE_CLASS`). -->
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
      <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
      {@render runAction()}
    {:else if simFigures !== null && simFigures.nowDps !== null}
      <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-sim-value">
        {homeHeroCardsCopy.simulatorNowFigure(formatAmount(Math.round(simFigures.nowDps)))}
      </span>
      {#if simFigures.bandDps !== null && ctx !== null}
        <span class="text-muted numeric text-[12px]" title={presetLabelFor(ctx.bisFile, ctx.band)}
          >{homeHeroCardsCopy.simulatorBandLine(simFigures.bandDps)}</span
        >
      {/if}
    {:else if simFigures !== null && simFigures.bandDps !== null && ctx !== null}
      <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-sim-value">
        {homeHeroCardsCopy.simulatorBandOnlyFigure(simFigures.bandDps)}
      </span>
      <span class="text-muted numeric text-[12px]" title={presetLabelFor(ctx.bisFile, ctx.band)}
        >{homeHeroCardsCopy.simulatorBandOnlyLine}</span
      >
      {@render runAction()}
    {:else}
      <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
      {@render runAction()}
    {/if}
  </a>
</div>
