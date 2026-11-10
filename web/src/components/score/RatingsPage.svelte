<!-- web/src/components/score/RatingsPage.svelte -->
<!-- /ratings (logs spec 2026-11-07 section 8, decision T1), rendered once at build time with no
     client script: the PageHeader with the sample Forever Score as its answer (in its tier's
     colour) and the two-anchor jump row, then a split Page of panels (round 5,
     design/mocks/directions/report_tail.py b1-ratings). Main: the Forever Score (the lead panel:
     the five parts as a table), the Performance Rating (framed: three Figures), the six components
     with their weights by role and what each measures, and how a fight is judged (ranked,
     absolute, wipes, a part left out, the cap) as Rows. Rail: the tiers in the game's item quality
     colours, the sample's tinted in its own colour (MUST 6), and how logs get here. Words and
     numbers: lib/score/copy.ts, lib/rating/copy.ts and lib/score/ratings-page.ts. -->
<script lang="ts">
  import { ratingCopy } from '../../lib/rating/copy';
  import { scoreCopy } from '../../lib/score/copy';
  import {
    SAMPLE_SCORE,
    SAMPLE_TREND,
    SAMPLE_WEAKEST,
    WEIGHTS,
    ratingsPageCopy,
  } from '../../lib/score/ratings-page';
  import { COMPONENT_MAX, TIER_FROM, tierRanges, type Tier } from '../../lib/score/score';
  import Button from '../ui/Button.svelte';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import StatRow from '../ui/StatRow.svelte';
  import StatTable from '../ui/StatTable.svelte';
  import TabStrip from '../ui/TabStrip.svelte';

  /** The game's item quality colours, exactly (tokens.css --color-rarity-*). */
  const TIER_COLOR: Readonly<Record<Tier, string>> = {
    common: 'var(--color-rarity-common)',
    uncommon: 'var(--color-rarity-uncommon)',
    rare: 'var(--color-rarity-rare)',
    epic: 'var(--color-rarity-epic)',
    legendary: 'var(--color-rarity-legendary)',
  };
  const RATING_TOP = 100;
  const count = (n: number): string => n.toLocaleString('en-US');

  const parts = scoreCopy.components;
  const maxTotal = Object.values(COMPONENT_MAX).reduce((sum, max) => sum + max, 0);
  const tier = SAMPLE_SCORE.tier as Tier;
  const tierName = tier.charAt(0).toUpperCase() + tier.slice(1);
  const gap = parts
    .map((part) => ({ part, gap: COMPONENT_MAX[part.key] - SAMPLE_SCORE.components[part.key] }))
    .sort((a, b) => b.gap - a.gap)[0];
  const latest = SAMPLE_TREND[SAMPLE_TREND.length - 1] ?? 0;
  const rise = latest - (SAMPLE_TREND[0] ?? 0);
  /** The next tier up and the points to it, for the sample's row in the Tiers rail. */
  const nextTier = (Object.keys(TIER_FROM) as Tier[]).find((each) => TIER_FROM[each] > SAMPLE_SCORE.score);
  const toNext =
    nextTier === undefined
      ? ''
      : `${count(TIER_FROM[nextTier] - SAMPLE_SCORE.score)} to ${nextTier.charAt(0).toUpperCase()}${nextTier.slice(1)}`;
  const tabsList = [
    { id: 'forever-score', label: scoreCopy.explainerHeading, href: '#forever-score' },
    { id: 'performance-rating', label: scoreCopy.explainerRatingHeading, href: '#performance-rating' },
  ];
</script>

<Page layout="split" panels testid="ratings-page">
  {#snippet header()}
    <PageHeader title={scoreCopy.explainerPageTitle}>
      {#snippet meta()}
        <span>{scoreCopy.explainerScoreCaption}</span><span>{scoreCopy.explainerRatingCaption}</span>
      {/snippet}
      {#snippet answer()}
        <Figure
          label="{scoreCopy.samplePill} · {scoreCopy.explainerHeading}"
          value={count(SAMPLE_SCORE.score)}
          unit="of {count(maxTotal)} · {tierName}"
          color={TIER_COLOR[tier]}
          size="big"
          align="end"
          testid="ratings-sample-score"
        />
        {#if gap}<span class="text-small text-muted"
            >{scoreCopy.mostToGainLead} {gap.part.label}{scoreCopy.mostToGainRest(gap.gap)}</span
          >{/if}
      {/snippet}
      {#snippet tabs()}
        <TabStrip
          tabs={tabsList}
          active="forever-score"
          label={scoreCopy.explainerJumpLabel}
          testid="ratings-jump"
        />
      {/snippet}
    </PageHeader>
  {/snippet}

  <Section
    title={scoreCopy.explainerHeading}
    anchor="forever-score"
    testid="ratings-forever-score"
    panel
    frame="framed"
    meta={ratingsPageCopy.sampleNote}
  >
    <div>
      <StatTable
        nameHead="Part"
        barHead="Of its top"
        icons={false}
        open
        columns={[
          { id: 'points', label: 'Points' },
          { id: 'top', label: 'Top' },
        ]}
        label={scoreCopy.explainerPartsHeading}
      >
        {#each parts as part (part.key)}
          <StatRow
            figures={[
              { value: count(SAMPLE_SCORE.components[part.key]) },
              { value: count(COMPONENT_MAX[part.key]) },
            ]}
            phone={`of ${COMPONENT_MAX[part.key]}`}
            testid={`ratings-forever-score-${part.key}`}
          >
            {#snippet name()}<b>{part.label}</b>{/snippet}
            {#snippet sub()}{ratingsPageCopy.parts[part.key]}{/snippet}
            {#snippet meter()}<Meter
                value={SAMPLE_SCORE.components[part.key] / COMPONENT_MAX[part.key]}
              />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </div>
    <p class="text-small text-muted pt-3" data-testid="ratings-launch">{scoreCopy.explainerLaunch}</p>
  </Section>

  <Section
    title={scoreCopy.explainerRatingHeading}
    anchor="performance-rating"
    testid="ratings-performance-rating"
    panel
  >
    <div class="grid grid-cols-3 gap-x-6 gap-y-5" data-testid="ratings-rating-sample">
      <Figure value={String(latest)} label="Latest fight" unit={ratingCopy.outOf100} tone="gold" size="big" />
      <div class="flex min-w-0 flex-col gap-[2px]">
        <span class="text-label text-muted font-bold">Last {SAMPLE_TREND.length} fights</span>
        <span
          class="text-strong text-[15px] leading-5 font-semibold tabular-nums"
          title={`out of ${RATING_TOP}`}>{SAMPLE_TREND.join(' · ')}</span
        >
        <span class="text-small text-muted">Up {rise} since the first</span>
      </div>
      <div class="flex min-w-0 flex-col gap-[2px]">
        <span class="text-label text-muted font-bold">Weakest part</span>
        <span class="text-strong text-[15px] leading-5 font-semibold">{SAMPLE_WEAKEST}</span>
      </div>
    </div>
    <p class="text-small text-muted pt-4" data-testid="ratings-visibility">
      {ratingsPageCopy.visibility}
      {ratingCopy.cappedRule}
    </p>
  </Section>

  <Section title="Weights, by role" meta={ratingsPageCopy.adjustable} testid="ratings-components" panel flush>
    <StatTable
      nameHead="Component"
      bar={false}
      icons={false}
      columns={[
        { id: 'dps', label: 'DPS' },
        { id: 'healer', label: 'Healer' },
        { id: 'tank', label: 'Tank' },
      ]}
      label="Weights, by role"
      testid="ratings-weights"
    >
      {#each WEIGHTS as row (row.component)}
        {@const measures = ratingsPageCopy.components.find((item) => item.part === row.component)}
        <StatRow
          bar={false}
          figures={[{ value: String(row.dps) }, { value: String(row.healer) }, { value: String(row.tank) }]}
          phone="DPS"
        >
          {#snippet name()}<b>{row.component}</b>{/snippet}
          {#snippet sub()}{measures?.body ?? ''}<span class="ui-st-phone"
              ><br />Healer {row.healer} · Tank {row.tank}</span
            >{/snippet}
        </StatRow>
      {/each}
    </StatTable>
  </Section>

  <Section title="How a fight is judged" testid="ratings-basis" panel>
    {#each ratingsPageCopy.rules as rule (rule.title)}
      <Row title={rule.title} sub={rule.body} open />
    {/each}
  </Section>

  {#snippet rail()}
    <Section title={`${scoreCopy.tierColumn}s`} testid="ratings-forever-tiers" panel>
      {#each tierRanges() as { tier: each, range } (each)}
        <div
          class="ratings-tier"
          data-current={each === tier ? 'true' : undefined}
          style:--tier={TIER_COLOR[each]}
        >
          <Row title={each} sub={range}>
            {#snippet titleContent()}<span class="text-[19px] leading-6" style:color={TIER_COLOR[each]}
                >{each.charAt(0).toUpperCase() + each.slice(1)}</span
              >{/snippet}
            {#snippet trail()}{#if each === tier}<span class="text-small text-muted text-right"
                  >The sample{#if toNext !== ''}<br /><span
                      class="text-text tabular-nums"
                      data-testid="ratings-to-next">{toNext}</span
                    >{/if}</span
                >{/if}{/snippet}
          </Row>
        </div>
      {/each}
    </Section>
    <Section title={scoreCopy.explainerLogsHeading} panel>
      <p class="text-small text-muted" data-testid="ratings-forever-logs">
        {scoreCopy.explainerLogsBody}
      </p>
      <p class="mt-3">
        <Button href={scoreCopy.explainerLogsHref} testid="ratings-forever-logs-button"
          >{scoreCopy.explainerLogsButton}</Button
        >
      </p>
    </Section>
  {/snippet}
</Page>

<style>
  .ratings-tier[data-current='true'] {
    background: color-mix(in srgb, var(--tier) 12%, transparent);
    box-shadow: inset 3px 0 0 var(--tier);
  }
</style>
