<!-- web/src/components/ForeverScoreCard.svelte -->
<!-- The character page's Forever Score (spec 2026-10-11 §4.6), as a design-system card: raised
     ground, 1px border, the gold top highlight, 6px radius.

     With a score: the badge, "of 2100", the tier and its date; the five parts as bars against
     their ceilings, the weakest share drawn in the tier colour and the rest in 60% gold so the
     part to work on reads at a glance; under Experience what it counts; then "Most to gain"
     (lib/score/score.ts's mostToGain). From lg the tier strip sits to the right, current tier
     marked; below lg it stacks.

     With no score (a 404, or a 200 whose score is null): the same card, "No logs yet", the
     addon's no-logs line where the number would be (never a zero), and what to do about it.
     From md the card keeps one minimum height in both states. -->
<script lang="ts">
  import { scoreCopy } from '../lib/score/copy';
  import {
    COMPONENT_MAX,
    SCORE_MAX,
    hasScore,
    mostToGain,
    shortDate,
    tierColor,
    tierLabel,
    tierTextColor,
    weakestPart,
    type ForeverScore,
  } from '../lib/score/score';
  import ScoreBadge from './ScoreBadge.svelte';
  import TierStrip from './TierStrip.svelte';

  let {
    forever,
    bisHref = null,
  }: {
    forever: ForeverScore | null;
    /** The spec's best-in-slot page, when the spec is known: where "Most to gain: Gear" goes. */
    bisHref?: string | null;
  } = $props();

  /** A bar's fill, to one decimal place: 412 of 500 is 82.4%. */
  function share(points: number, max: number): number {
    return Math.round(Math.min(1, Math.max(0, points / max)) * 1000) / 10;
  }

  const CARD =
    'bg-raised border-line rounded-panel flex flex-col border p-4 shadow-[inset_0_1px_0_rgba(229,185,85,.35)] md:min-h-[244px] md:p-5';
  /** In-text links are underlined: colour alone does not mark a link (WCAG 1.4.1). */
  const LINK = 'text-gold inline-flex min-h-11 items-center underline underline-offset-2 md:min-h-0';

  const date = $derived(shortDate(forever?.as_of));
  const gain = $derived(mostToGain(forever?.components));
  const weakest = $derived(weakestPart(forever?.components));
  const gainLabel = $derived(scoreCopy.components.find((c) => c.key === gain?.key)?.label ?? '');
  const gainHref = $derived(gain?.key === 'gear' && bisHref !== null ? bisHref : scoreCopy.explainHref);
</script>

<section class="flex flex-col gap-2" data-testid="forever-score">
  <h2 class="section-title text-[18px]">{scoreCopy.heading}</h2>
  <p class="text-muted text-[13px]" data-testid="forever-score-caption">{scoreCopy.caption(date)}</p>
  <div class={CARD} data-testid="forever-score-card">
    {#if hasScore(forever)}
      {@const tier = tierLabel(forever.tier)}
      {@const weakColor = tierColor(forever.tier)}
      <div class="grid gap-6 lg:grid-cols-[minmax(0,1fr)_minmax(0,380px)] lg:items-start">
        <div class="flex flex-col gap-4">
          <div class="flex flex-wrap items-baseline gap-x-2 gap-y-1">
            <ScoreBadge score={forever.score} tier={forever.tier} size="lg" />
            <span class="text-muted tabular font-mono text-[14px]">{scoreCopy.outOf(SCORE_MAX)}</span>
            <span class="text-[14px]" data-testid="forever-score-tier">
              {#if tier !== null}<span class="font-semibold" style={`color: ${tierTextColor(forever.tier)}`}
                  >{tier}</span
                >{/if}{#if tier !== null && date !== null}{scoreCopy.separator}{/if}{#if date !== null}<span
                  class="text-muted">{scoreCopy.asOf(date)}</span
                >{/if}
            </span>
          </div>
          {#if forever.components !== null}
            <ul class="flex max-w-[560px] flex-col gap-2">
              {#each scoreCopy.components as component (component.key)}
                {@const points = forever.components[component.key]}
                {@const max = COMPONENT_MAX[component.key]}
                {@const weak = component.key === weakest}
                {#if typeof points === 'number'}
                  <li
                    class="grid grid-cols-[88px_minmax(0,1fr)_88px] items-center gap-x-3 gap-y-1 text-[13px]"
                    data-testid={`forever-score-${component.key}`}
                    data-weakest={weak ? 'true' : undefined}
                  >
                    <span>{component.label}</span>
                    <span class="bg-line-soft h-2" aria-hidden="true"
                      ><span
                        class={`block h-full ${weak && weakColor !== null ? '' : weak ? 'bg-gold' : 'bg-gold/60'}`}
                        style={`width: ${share(points, max)}%${weak && weakColor !== null ? `; background-color: ${weakColor}` : ''}`}
                      ></span></span
                    >
                    <span class="text-muted tabular text-right font-mono whitespace-nowrap"
                      >{scoreCopy.ofMax(points, max)}</span
                    >
                    {#if component.key === 'experience' && forever.experience !== null}
                      <span
                        class="text-muted col-span-2 col-start-2 text-[12px]"
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
          {#if gain !== null}
            <p class="text-[14px]" data-testid="forever-score-most-to-gain">
              {scoreCopy.mostToGainLead}
              <a class={LINK} href={gainHref}>{gainLabel}</a>{scoreCopy.mostToGainRest(gain.gap)}
            </p>
          {/if}
        </div>
        <TierStrip current={forever.tier} />
      </div>
    {:else}
      <div class="flex flex-col gap-2" data-testid="forever-score-empty">
        <p class="text-strong text-[16px] font-semibold">{scoreCopy.noLogsTitle}</p>
        <p class="text-muted text-[14px]" data-testid="forever-score-none">{scoreCopy.noLogs}</p>
        <p class="text-[14px]">
          {scoreCopy.noLogsCheckLead}
          <a class={LINK} href={scoreCopy.noLogsCheckHref}>{scoreCopy.noLogsCheckLink}</a>.
        </p>
      </div>
    {/if}
  </div>
  <a class={`${LINK} w-fit text-[14px]`} href={scoreCopy.explainHref}>{scoreCopy.explainLink}</a>
</section>
