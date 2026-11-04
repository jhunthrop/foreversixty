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
  import { homeSpecKeyFor, loadBisContextFor, type BisContext } from '../../lib/home/upgrades-loader';
  import { levelingHomeCopy } from '../../lib/home-panel-copy';
  import { levelingPlanFor } from '../../lib/home/header-answer';
  import { classGuideHref, talentStepFor, type LevelingPlan } from '../../lib/leveling/leveling-plan';
  import { FIRST_BAND_LEVEL } from '../../lib/bis/landing-bands';
  import { FIRST_POINT_LEVEL } from '../../lib/planner/types';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import activeBuild from '../../data/active-build.json';
  import GuideTooltips from '../guides/GuideTooltips.svelte';
  import LevelingTalentRow from './LevelingTalentRow.svelte';
  import { talentDeltaFor } from '../../lib/home/talent-delta';
  import ActionLink from '../ui/ActionLink.svelte';
  import Card from '../ui/Card.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import { presetLabelFor } from '../../lib/bis/presets';

  let { hero }: { hero: MeCharacter } = $props();

  /** Below the first gear range the cards teach leveling (spec 2026-11-04 §4b, §4c). */
  const leveling = $derived((hero.level ?? 1) < FIRST_BAND_LEVEL);
  /** How many points of the order the card shows: five fit one row at 390. */
  const ORDER_SHOWN = 5;

  let plan = $state<LevelingPlan | null>(null);
  let planStatus = $state<'loading' | 'ready' | 'failed'>('loading');

  $effect(() => {
    const current = hero;
    if (!leveling) return;
    planStatus = 'loading';
    void levelingPlanFor(current).then((loaded) => {
      if (current !== hero) return;
      plan = loaded ?? null;
      planStatus = loaded === undefined ? 'failed' : 'ready';
    });
  });

  const step = $derived(plan === null ? null : talentStepFor(plan, hero.level ?? 1, hero.build?.talents));
  const shownPoints = $derived(
    plan === null || step === null || step.nextIndex < 0
      ? []
      : plan.points.slice(step.nextIndex, step.nextIndex + ORDER_SHOWN),
  );
  const levelingCaption = $derived.by(() => {
    if (plan === null || step === null) return '';
    const level = hero.level ?? 1;
    const lead = levelingHomeCopy.levelingAs(plan.specName);
    if (step.nextIndex < 0) return `${lead} · ${levelingHomeCopy.orderDone}`;
    if (level < FIRST_POINT_LEVEL) return `${lead} · ${levelingHomeCopy.firstPointAt(FIRST_POINT_LEVEL)}`;
    if (step.toSpend === 0) return `${lead} · ${levelingHomeCopy.nextPointAt(step.nextPointLevel)}`;
    return `${lead} · ${levelingHomeCopy.toSpend(step.toSpend)}`;
  });
  const otherSpecsAction = $derived(
    hero.class === undefined
      ? undefined
      : { label: levelingHomeCopy.otherSpecs, href: classGuideHref(classSlugFromName(hero.class)) },
  );
  const tipsUrl = $derived(plan === null ? null : plan.tipsUrl);

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
  let bandStatus = $state<BandStatus>(
    untrack(() => (homeSpecKeyFor(hero) === undefined ? 'no-spec' : 'loading')),
  );
  let ctx = $state<BisContext | null>(null);

  $effect(() => {
    const current = hero;
    ctx = null;
    if (homeSpecKeyFor(current) === undefined) {
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
    latestSim === null ? null : simCardFigures(latestSim, { specKey: homeSpecKeyFor(hero), band: ctx?.band }),
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
    const specKey = homeSpecKeyFor(hero);
    return specKey === undefined ? '/planner' : `/planner?spec=${specKey}`;
  });

  /** Fix round 1 item B.4 (ux-designer): the card's own answer is the one primary element and
   *  reads at the design system's headline-figure treatment (`design/DESIGN-SYSTEM.md`'s type
   *  table; polish spec 2026-10-28 §4.4 moved figures to Barlow Condensed 800) -- 22px here,
   *  not the `text-[15px] font-semibold` body-text size the mock's own three cards never use
   *  for their headline number. The line underneath it stays Barlow 12px muted, unchanged. */
  /** The character's own sim is showing, so the link opens the simulator rather than asking for a first run. */
  const hasSimFigure = $derived(simStatus === 'ready' && simFigures !== null && simFigures.nowDps !== null);

  const CARD_FIGURE_CLASS = 'figure-stat text-[22px] text-strong';
  /** Each card's answer box, one height in every state (alignment stage 2, tenet 13): the ready
   *  figure (22 px) over its one-line caption (22), 12 apart, so the loading skeleton, a one-line
   *  state and the ready figure all draw 56 px and the row the slot reserves never moves. As
   *  Cards the old 72 px talents skeleton swapped for a 22 px line and shifted the hero 50 px. */
  const CARD_ANSWER_CLASS = 'flex min-h-14 flex-col gap-3';
  const SKELETON_ROW = 'h-[22px]';
</script>

<div class="reveal grid grid-cols-1 items-start gap-3 sm:grid-cols-2" data-testid="home-hero-cards">
  {#if leveling}
    <Card
      glyph="tree"
      title={homeHeroCardsCopy.talentsLabel}
      compact
      action={otherSpecsAction}
      testid="home-hero-card-talents"
    >
      <div class={CARD_ANSWER_CLASS} data-testid="home-talents-leveling">
        {#if planStatus === 'loading'}
          <Skeleton lines={2} rowHeight={SKELETON_ROW} testid="home-hero-card-talents-skeleton" />
        {:else if plan === null}
          <span class="text-muted text-[13px]">{levelingHomeCopy.talentOrderFailed}</span>
        {:else}
          {#if shownPoints.length > 0}
            <LevelingTalentRow
              points={shownPoints}
              lit={(hero.level ?? 1) >= FIRST_POINT_LEVEL && (step?.toSpend ?? 0) > 0}
            />
          {/if}
          <span class="text-muted text-[13px]" data-testid="home-talents-caption">{levelingCaption}</span>
        {/if}
      </div>
      <ActionLink href={plan?.plannerHref ?? talentsHref} data-testid="home-hero-card-talents-open">
        {levelingHomeCopy.talentOrder}
      </ActionLink>
    </Card>
    <Card glyph="bolt" title={homeHeroCardsCopy.simulatorLabel} compact quiet testid="home-hero-card-sim">
      <div class={CARD_ANSWER_CLASS} data-testid="home-sim-locked">
        <span class="text-muted text-[13px]">{levelingHomeCopy.simLocked(FIRST_BAND_LEVEL)}</span>
      </div>
    </Card>
    {#if tipsUrl !== null}
      <GuideTooltips build={activeBuild.build} bundleUrl={tipsUrl} />
    {/if}
  {:else}
    <Card glyph="tree" title={homeHeroCardsCopy.talentsLabel} compact testid="home-hero-card-talents">
      <div class={CARD_ANSWER_CLASS}>
        {#if bandStatus === 'no-spec'}
          <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNoListYet}</span>
        {:else if hero.build?.talents === undefined}
          <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNotAvailable}</span>
        {:else if bandStatus === 'loading'}
          <!-- Fix round 1 item C.9 (tenet 13, no layout shift): the ready figure and its caption,
           22 px each (`CARD_ANSWER_CLASS`). -->
          <Skeleton lines={2} rowHeight={SKELETON_ROW} testid="home-hero-card-talents-skeleton" />
        {:else if bandStatus === 'no-list-yet'}
          <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNoListYet}</span>
        {:else if talentDelta !== null && ctx !== null}
          <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-talents-value">
            {talentDelta === 0
              ? homeHeroCardsCopy.talentsOptimizedFigure
              : homeHeroCardsCopy.talentsUnoptimizedFigure}
          </span>
          <span class="text-muted text-[13px]">
            {homeHeroCardsCopy.talentsLine(talentDelta, ctx.band.talent_points, bandLabel)}
          </span>
        {/if}
      </div>
      <ActionLink href={talentsHref} data-testid="home-hero-card-talents-open">
        {homeHeroCardsCopy.talentsAction}
      </ActionLink>
    </Card>
    <Card glyph="bolt" title={homeHeroCardsCopy.simulatorLabel} compact testid="home-hero-card-sim">
      <div class={CARD_ANSWER_CLASS}>
        {#if simStatus === 'loading'}
          <Skeleton lines={2} rowHeight={SKELETON_ROW} testid="home-hero-card-sim-skeleton" />
        {:else if simStatus === 'failed'}
          <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
        {:else if simFigures !== null && simFigures.nowDps !== null}
          <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-sim-value">
            {homeHeroCardsCopy.simulatorNowFigure(formatAmount(Math.round(simFigures.nowDps)))}
          </span>
          {#if simFigures.bandDps !== null && ctx !== null}
            <span class="text-muted numeric text-[13px]" title={presetLabelFor(ctx.bisFile, ctx.band)}
              >{homeHeroCardsCopy.simulatorBandLine(simFigures.bandDps)}</span
            >
          {/if}
        {:else if simFigures !== null && simFigures.bandDps !== null && ctx !== null}
          <span class={CARD_FIGURE_CLASS} data-testid="home-hero-card-sim-value">
            {homeHeroCardsCopy.simulatorBandOnlyFigure(simFigures.bandDps)}
          </span>
          <span class="text-muted numeric text-[13px]" title={presetLabelFor(ctx.bisFile, ctx.band)}
            >{homeHeroCardsCopy.simulatorBandOnlyLine}</span
          >
        {:else}
          <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
        {/if}
      </div>
      <ActionLink
        href={armorySimHref(hero.key)}
        data-testid={hasSimFigure ? 'home-hero-card-sim-open' : 'home-hero-card-sim-run'}
      >
        {hasSimFigure ? homeHeroCardsCopy.simulatorAction : homePanelCopy.runAction}
      </ActionLink>
    </Card>
  {/if}
</div>
