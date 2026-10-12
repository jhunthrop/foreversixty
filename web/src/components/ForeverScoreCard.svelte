<!-- web/src/components/ForeverScoreCard.svelte -->
<!-- The Forever Score (spec 2026-10-11 §4.6) as a design-system card: raised ground, 1px border,
     the gold top highlight, 6px radius. The profile mounts it under its own heading and caption;
     /ratings mounts it bare (`heading={false}`) as the sample that opens the page.

     With a score: the number (48px from md), "of 2100" and the tier, "170 to Legendary (2000)"
     under them (fix round 4 item 1); the tier strip under it,
     current tier marked; the score over time as a gold sparkline when score_history has two or
     more points (Task L8.4 ruling 3, fix round 2 items 1 and 2: the content column's width at a
     fixed height, the newest point marked in its tier's colour, "Up 330 since 10 Nov" and the
     newest date under it, an image named for a screen reader with the table beside it, and no
     box at all with fewer than two, "First score recorded 12 Oct." with one: fix round 3 item 12);
     the five parts as bars against their ceilings. From 1280 the card is two columns (fix round 3
     item 13): the number, tier word and parts on the left, the sparkline over the tier strip on
     the right. One part is marked as
     the one to work on, and it is always the part "Most to gain" names (lib/score/score.ts's
     mostToGain: the largest point gap, the earlier part on a tie). It takes one neutral colour
     that is never a tier's, and a "+88 pts" at the row's end so colour is not the only signal. With
     every part at its top nothing is marked and there is no line. The footer carries "Most to
     gain" left and "What the score counts" right.

     With no score (a 404, or a 200 whose score is null): the same shape -- "No logs yet" where
     the number would be, "of 2100" greyed, the five parts as empty tracks with a dash -- then
     one button: the owner's "Upload logs" (the companion uploads fights automatically), anyone
     else's "Look up another character". The addon's no-logs sentence lives in the tooltip mock
     below, so the card does not repeat it. -->
<script lang="ts">
  import { scoreCopy } from '../lib/score/copy';
  import {
    COMPONENT_MAX,
    NEEDS_WORK_COLOR,
    SCORE_MAX,
    hasScore,
    mostToGain,
    nextTierLine,
    shortDate,
    tierLabel,
    tierTextColor,
    type ForeverScore,
    type ScorePoint,
  } from '../lib/score/score';
  import {
    SPARKLINE_HEIGHT,
    SPARKLINE_WIDTH,
    scoreHistoryFirst,
    scoreSparkline,
  } from '../lib/score/sparkline';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import ScoreBadge from './ScoreBadge.svelte';
  import TierStrip from './TierStrip.svelte';

  let {
    forever,
    history = null,
    bisHref = null,
    owned = false,
    name = '',
    region = '',
    ruleset = '',
    heading = true,
    sample = false,
  }: {
    forever: ForeverScore | null;
    /** The character's score_history, oldest first; the sparkline needs two points or more. */
    history?: readonly ScorePoint[] | null;
    /** The spec's best-in-slot page, when the spec is known: where "Most to gain: Gear" goes. */
    bisHref?: string | null;
    /** The viewer's own character: "Your whole record" and the owner's no-logs action. */
    owned?: boolean;
    /** The character's name, for "Look up another character" (the lookup starts filled). */
    name?: string;
    /** The character's region and ruleset, carried to the lookup with the name (round 3 item 3). */
    region?: string;
    ruleset?: string;
    /** False on /ratings, where the page's own heading introduces the sample. */
    heading?: boolean;
    /** Marks the card as an example, with the design system's Sample pill. */
    sample?: boolean;
  } = $props();

  /** A bar's fill, to one decimal place: 412 of 500 is 82.4%. */
  function share(points: number, max: number): number {
    return Math.round(Math.min(1, Math.max(0, points / max)) * 1000) / 10;
  }

  const CARD =
    'bg-raised border-line rounded-panel flex flex-col gap-4 border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:p-5';
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). */
  const LINK = 'text-gold inline-flex min-h-11 items-center underline underline-offset-2 md:min-h-0';
  /** The last column holds the "+88 pts" cue on one line (fix round 4 item 2). */
  const ROW = 'grid grid-cols-[88px_minmax(0,1fr)_80px_56px] items-center gap-x-3 gap-y-1 text-[13px]';

  const date = $derived(shortDate(forever?.as_of));
  const gain = $derived(mostToGain(forever?.components));
  const gainLabel = $derived(scoreCopy.components.find((c) => c.key === gain?.key)?.label ?? '');
  const sparkline = $derived(scoreSparkline(history));
  const firstScore = $derived(scoreHistoryFirst(history));
  const gainHref = $derived(gain?.key === 'gear' && bisHref !== null ? bisHref : scoreCopy.explainHref);
</script>

{#snippet card()}
  <div class={CARD} data-testid="forever-score-card">
    {#if hasScore(forever)}
      {@const tier = tierLabel(forever.tier)}
      <!-- Fix round 3 item 13: from 1280 two columns -- left the number, the tier word and the
           parts; right the sparkline (the full column) over the tier strip. Below 1280 one
           column: the two wrappers are `display: contents`, so their children are this flex
           column's own and `order` draws them as before (number, tier strip, sparkline, parts). -->
      <div class="flex flex-col gap-4 xl:grid xl:grid-cols-2 xl:gap-x-8" data-testid="forever-score-columns">
        <div class="contents xl:flex xl:flex-col xl:gap-4" data-testid="forever-score-main">
          <div
            class="order-1 flex flex-wrap items-baseline gap-x-2 gap-y-1 xl:order-none"
            data-testid="forever-score-headline"
          >
            <ScoreBadge score={forever.score} tier={forever.tier} size="xl" />
            <span class="text-muted tabular font-mono text-[14px]">{scoreCopy.outOf(SCORE_MAX)}</span>
            {#if tier !== null}<span
                class="text-[15px] font-semibold"
                style={`color: ${tierTextColor(forever.tier)}`}
                data-testid="forever-score-tier">{tier}</span
              >{/if}
            {#if sample}<span class="pill pill-sample ml-auto">{scoreCopy.samplePill}</span>{/if}
            <!-- Fix round 4 item 1: under the score, on a line of its own, the points to the next
                 tier ("170 to Legendary (2000)"), or "Top tier". -->
            <p class="text-muted basis-full text-[13px]" data-testid="forever-score-next-tier">
              {nextTierLine(forever.score)}
            </p>
          </div>
          {#if forever.components !== null}
            <ul
              class="order-4 flex max-w-[600px] flex-col gap-2 xl:order-none"
              data-testid="forever-score-parts"
            >
              {#each scoreCopy.components as component (component.key)}
                {@const points = forever.components[component.key]}
                {@const max = COMPONENT_MAX[component.key]}
                {@const marked = gain !== null && component.key === gain.key}
                {#if typeof points === 'number'}
                  <li
                    class={ROW}
                    data-testid={`forever-score-${component.key}`}
                    data-needs-work={marked ? 'true' : undefined}
                  >
                    <span>{component.label}</span>
                    <span class="bg-line-soft h-2" aria-hidden="true"
                      ><span
                        class={`block h-full ${marked ? '' : 'bg-gold/60'}`}
                        style={`width: ${share(points, max)}%${marked ? `; background-color: ${NEEDS_WORK_COLOR}` : ''}`}
                      ></span></span
                    >
                    <span class="text-muted tabular text-right font-mono whitespace-nowrap"
                      >{scoreCopy.ofMax(points, max)}</span
                    >
                    <span
                      class="tabular text-right font-mono text-[12px] font-semibold whitespace-nowrap"
                      style={marked && gain !== null ? `color: ${NEEDS_WORK_COLOR}` : undefined}
                      aria-label={marked && gain !== null ? scoreCopy.gainCueLabel(gain.gap) : undefined}
                      data-testid={marked ? 'forever-score-gain-cue' : undefined}
                      >{marked && gain !== null ? scoreCopy.gainCue(gain.gap) : ''}</span
                    >
                    {#if component.key === 'experience' && forever.experience !== null}
                      <span
                        class="text-muted col-span-3 col-start-2 text-[12px]"
                        data-testid="forever-score-experience-detail"
                        >{scoreCopy.experienceBreakdown(
                          forever.experience.dungeons,
                          forever.experience.raid_kills,
                          forever.experience.alts,
                        )}</span
                      >
                    {/if}
                  </li>
                {/if}
              {/each}
            </ul>
          {/if}
        </div>
        <div class="contents xl:flex xl:flex-col xl:gap-3" data-testid="forever-score-side">
          {#if sparkline !== null}
            <!-- Fix round 2 items 1 and 2: the tier strip's width from md, the whole content width
                 on a phone, the whole right column from 1280 (fix round 3 item 13); 40px high,
                 48px from 1024px. The 8px end dot sits on the right edge and may overhang it. The
                 caption ("Up 330 since 10 Nov", which carries the first date; ", 2 readings" under
                 five points) leads the date row in the text colour and the newest date ends it,
                 mono muted (fix round 3 item 12). -->
            <figure
              class="order-3 m-0 flex w-full max-w-[460px] flex-col gap-1 xl:order-none xl:max-w-none"
              data-testid="forever-score-history"
            >
              <div class="relative h-10 lg:h-12" data-testid="forever-score-plot">
                <svg
                  class="absolute inset-0 block h-full w-full overflow-visible"
                  viewBox={`0 0 ${SPARKLINE_WIDTH} ${SPARKLINE_HEIGHT}`}
                  preserveAspectRatio="none"
                  role="img"
                  aria-label={sparkline.ariaLabel}
                  data-testid="forever-score-sparkline"
                >
                  <path
                    d={sparkline.path}
                    fill="none"
                    stroke="var(--color-gold)"
                    stroke-width="2"
                    stroke-linejoin="round"
                    stroke-linecap="round"
                    vector-effect="non-scaling-stroke"
                  />
                </svg>
                <span
                  class="absolute size-2 -translate-x-1/2 -translate-y-1/2 rounded-full"
                  style={`left: ${sparkline.last.x}%; top: ${sparkline.last.y}%; background-color: ${sparkline.last.color ?? 'var(--color-gold)'}`}
                  aria-hidden="true"
                  data-testid="forever-score-sparkline-last"
                ></span>
              </div>
              <div
                class="flex items-baseline justify-between gap-3"
                aria-hidden="true"
                data-testid="forever-score-history-dates"
              >
                <span class="text-text text-[12px]" data-testid="forever-score-history-caption"
                  >{sparkline.caption}</span
                >
                <span
                  class="text-muted tabular font-mono text-[12px]"
                  data-testid="forever-score-history-last">{sparkline.lastLabel}</span
                >
              </div>
              <table class="sr-only" data-testid="forever-score-history-table">
                <caption>{scoreCopy.historyCaption}</caption>
                <thead>
                  <tr
                    ><th scope="col">{scoreCopy.historyDate}</th><th scope="col">{scoreCopy.historyScore}</th
                    ></tr
                  >
                </thead>
                <tbody>
                  {#each sparkline.rows as row, index (index)}
                    <tr><td>{row.date}</td><td>{row.score}</td></tr>
                  {/each}
                </tbody>
              </table>
            </figure>
          {:else if firstScore !== null}
            <!-- One point is not a line (fix round 3 item 12): the date the record starts, where
                 the line would be. -->
            <p class="text-text order-3 text-[12px] xl:order-none" data-testid="forever-score-first">
              {firstScore}
            </p>
          {/if}
          <!-- -mt-1 below 1280 keeps the strip 12px under the number, as before the two columns. -->
          <div
            class="order-2 max-w-[460px] max-xl:-mt-1 xl:order-none xl:max-w-none"
            data-testid="forever-score-tiers"
          >
            <TierStrip current={forever.tier} />
          </div>
        </div>
      </div>
      <div class="flex flex-wrap items-center justify-between gap-x-4 gap-y-1">
        {#if gain !== null}
          <p class="text-[14px] leading-[22px]" data-testid="forever-score-most-to-gain">
            {scoreCopy.mostToGainLead}
            <a class={LINK} href={gainHref}>{gainLabel}</a>{scoreCopy.mostToGainRest(gain.gap)}
          </p>
        {/if}
        {#if !sample}
          <a class={`${LINK} ml-auto text-[14px]`} href={scoreCopy.explainHref}>{scoreCopy.explainLink}</a>
        {/if}
      </div>
    {:else}
      <div class="flex flex-wrap items-baseline gap-x-2 gap-y-1" data-testid="forever-score-empty">
        <span class="text-strong text-[24px] leading-none font-semibold md:text-[28px]"
          >{scoreCopy.noLogsTitle}</span
        >
        <span class="text-muted tabular font-mono text-[14px]" data-testid="forever-score-empty-max"
          >{scoreCopy.outOf(SCORE_MAX)}</span
        >
      </div>
      <ul class="flex max-w-[600px] flex-col gap-2">
        {#each scoreCopy.components as component (component.key)}
          <li class={ROW} data-testid={`forever-score-empty-${component.key}`}>
            <span class="text-muted">{component.label}</span>
            <span class="bg-line-soft h-2" aria-hidden="true"></span>
            <span class="text-muted tabular text-right font-mono">{scoreCopy.emptyPart}</span>
            <span></span>
          </li>
        {/each}
      </ul>
      <div class="flex flex-wrap items-center justify-between gap-x-4 gap-y-3">
        <div class="flex flex-col items-start gap-3">
          <p class="text-[14px] leading-[22px]" data-testid="forever-score-empty-hint">
            {owned ? scoreCopy.noLogsOwner : scoreCopy.noLogsHint}
          </p>
          {#if !owned}
            <!-- Fix round 4 item 5: a visitor's next step, one line under the message: the addon
                 and the logs upload. The owner's "Upload logs" button below stands alone. -->
            <p class="text-muted text-[14px] leading-[22px]" data-testid="forever-score-empty-next">
              {scoreCopy.noLogsNext.lead}
              <a class={LINK} href={scoreCopy.noLogsNext.addonHref}>{scoreCopy.noLogsNext.addon}</a
              >{scoreCopy.separator}<a class={LINK} href={scoreCopy.noLogsNext.uploadHref}
                >{scoreCopy.noLogsNext.upload}</a
              >
            </p>
          {/if}
          {#if owned}
            <a
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4`}
              href={scoreCopy.noLogsUploadHref}
              data-testid="forever-score-empty-action">{scoreCopy.noLogsUpload}</a
            >
          {:else}
            <a
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4`}
              href={scoreCopy.lookupHref(name, region, ruleset)}
              data-testid="forever-score-empty-action">{scoreCopy.noLogsLookup}</a
            >
          {/if}
        </div>
        {#if !sample}
          <a class={`${LINK} ml-auto self-end text-[14px]`} href={scoreCopy.explainHref}
            >{scoreCopy.explainLink}</a
          >
        {/if}
      </div>
    {/if}
  </div>
{/snippet}

{#if heading}
  <!-- -order-2: on the profile (a flex column) the score card comes right after the header,
       above the visitor's sign-in band, at every width (Character.svelte, fix round 3 item 4,
       L8.4 fix round 2 item 7). -->
  <section class="-order-2 flex flex-col gap-2" data-testid="forever-score">
    <h2 class="section-title text-[18px]">{scoreCopy.heading}</h2>
    <p class="text-muted text-[13px]" data-testid="forever-score-caption">{scoreCopy.caption(date, owned)}</p>
    {@render card()}
  </section>
{:else}
  {@render card()}
{/if}
