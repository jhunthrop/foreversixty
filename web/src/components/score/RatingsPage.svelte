<!-- web/src/components/score/RatingsPage.svelte -->
<!-- /ratings (logs spec 2026-11-07 section 8, decision T1), rendered once at build time with no
     client script: the PageHeader with the sample Forever Score as its answer (in its tier's
     colour) and the two-anchor jump row, then a split Page of panels (round 5, the logs lane's
     panel form). Main: the Forever Score (the page's lead panel: the five parts as a table), the
     Performance Rating (framed: three Figures), and the reference panels (each component, weights
     by role, ranked or absolute, the cap). Rail: the tiers in the game's item quality colours on
     their panel, the sample's tinted in its own colour (MUST 6), and how logs get here. Words and numbers: lib/score/copy.ts, lib/rating/copy.ts and
     lib/score/ratings-page.ts. -->
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
  import { COMPONENT_MAX, tierRanges, type Tier } from '../../lib/score/score';
  import Button from '../ui/Button.svelte';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Spark from '../ui/Spark.svelte';
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

  const parts = scoreCopy.components;
  const maxTotal = Object.values(COMPONENT_MAX).reduce((sum, max) => sum + max, 0);
  const tier = SAMPLE_SCORE.tier as Tier;
  const tierName = tier.charAt(0).toUpperCase() + tier.slice(1);
  const gap = parts
    .map((part) => ({ part, gap: COMPONENT_MAX[part.key] - SAMPLE_SCORE.components[part.key] }))
    .sort((a, b) => b.gap - a.gap)[0];
  const latest = SAMPLE_TREND[SAMPLE_TREND.length - 1] ?? 0;
  const rise = latest - (SAMPLE_TREND[0] ?? 0);
  const tabsList = [
    { id: 'forever-score', label: scoreCopy.explainerHeading, href: '#forever-score' },
    { id: 'performance-rating', label: scoreCopy.explainerRatingHeading, href: '#performance-rating' },
  ];
  const count = (n: number): string => n.toLocaleString('en-US');
</script>

{#snippet prose(
  paragraphs: readonly string[],
  muted: readonly boolean[] = [],
  testids: readonly string[] = [],
)}
  <div class="flex flex-col gap-3 [&:not(:first-child)]:pt-3">
    {#each paragraphs as paragraph, index (index)}
      <p
        class={['max-w-[72ch] text-[15px] leading-6', muted[index] && 'text-muted']}
        data-testid={testids[index]}
      >
        {paragraph}
      </p>
    {/each}
  </div>
{/snippet}

<Page layout="split" panels testid="ratings-page">
  {#snippet header()}
    <PageHeader title={scoreCopy.explainerPageTitle}>
      {#snippet meta()}
        <span>{scoreCopy.explainerScoreCaption}</span><span>{scoreCopy.explainerRatingCaption}</span>
      {/snippet}
      {#snippet answer()}
        <!-- the answer column holds 360 px from 900, as the mock's -->
        <span class="hidden min-[900px]:block min-[900px]:w-[360px]" aria-hidden="true"></span>
        <span class="text-label text-muted font-bold"
          >{scoreCopy.samplePill} · {scoreCopy.explainerHeading}</span
        >
        <span
          class="text-[36px] leading-10 font-bold tabular-nums"
          style:color={TIER_COLOR[tier]}
          data-testid="ratings-sample-score"
          >{count(SAMPLE_SCORE.score)}<small class="text-muted ml-1 text-[13px] font-semibold"
            >of {count(maxTotal)} · {tierName}</small
          ></span
        >
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
    meta={scoreCopy.explainerScoreCaption}
    anchor="forever-score"
    testid="ratings-forever-score"
    panel
    frame="lead"
  >
    <p class="max-w-[62ch] text-[17px] leading-[26px]" data-testid="ratings-forever-lead">
      {scoreCopy.explainerLead}
    </p>
    <div class="mt-4">
      <StatTable
        nameHead="Part"
        barHead="Of its top"
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
              { value: scoreCopy.upTo(COMPONENT_MAX[part.key]) },
            ]}
            phone={`${SAMPLE_SCORE.components[part.key]} ${scoreCopy.upTo(COMPONENT_MAX[part.key])}`}
            testid={`ratings-forever-score-${part.key}`}
          >
            {#snippet name()}<b>{part.label}</b>{/snippet}
            {#snippet sub()}{part.moves}{/snippet}
            {#snippet meter()}<Meter
                value={SAMPLE_SCORE.components[part.key] / COMPONENT_MAX[part.key]}
              />{/snippet}
          </StatRow>
        {/each}
      </StatTable>
    </div>
    {@render prose([scoreCopy.explainerWhat, scoreCopy.explainerLaunch], [false, true])}
  </Section>

  <Section
    title={scoreCopy.explainerRatingHeading}
    meta={scoreCopy.explainerRatingCaption}
    anchor="performance-rating"
    testid="ratings-performance-rating"
    panel
    frame="framed"
  >
    <p class="max-w-[62ch] text-[17px] leading-[26px]" data-testid="ratings-rating-bridge">
      {`${scoreCopy.explainerRatingBridge} ${scoreCopy.explainerPreparationShared}`}
    </p>
    <div class="mt-4 grid grid-cols-3 gap-x-6 gap-y-5" data-testid="ratings-rating-sample">
      <Figure value={String(latest)} label="Latest fight" unit={ratingCopy.outOf100} tone="gold" size="big" />
      <div class="flex min-w-0 flex-col gap-[2px]">
        <span class="text-label text-muted font-bold">Last {SAMPLE_TREND.length} fights</span>
        <Spark values={SAMPLE_TREND} top={RATING_TOP} />
        <span class="text-small text-muted">Up {rise} since the first</span>
      </div>
      <div class="flex min-w-0 flex-col gap-[2px]">
        <span class="text-label text-muted font-bold">Weakest part</span>
        <span class="text-strong text-[15px] leading-5 font-semibold">{SAMPLE_WEAKEST}</span>
      </div>
    </div>
    {@render prose(
      [ratingsPageCopy.visibility, ...ratingsPageCopy.oneNumberFirst],
      [true],
      ['ratings-visibility'],
    )}
  </Section>

  <Section title="What each component measures" testid="ratings-components" panel>
    {#each ratingsPageCopy.components as item (item.part)}
      <Row title={item.part} sub={item.body} open />
    {/each}
  </Section>

  <Section title="Weights, by role" panel>
    <StatTable
      nameHead="Component"
      bar={false}
      open
      columns={[
        { id: 'dps', label: 'DPS' },
        { id: 'healer', label: 'Healer' },
        { id: 'tank', label: 'Tank' },
      ]}
      label="Weights, by role"
      testid="ratings-weights"
    >
      {#each WEIGHTS as row (row.component)}
        <StatRow
          bar={false}
          figures={[{ value: String(row.dps) }, { value: String(row.healer) }, { value: String(row.tank) }]}
          phone={`${row.dps} / ${row.healer} / ${row.tank}`}
        >
          {#snippet name()}<b>{row.component}</b>{/snippet}
        </StatRow>
      {/each}
    </StatTable>
    {@render prose(
      [ratingsPageCopy.roles, ratingsPageCopy.adjustable],
      [true, true],
      ['ratings-roles', 'ratings-adjustable'],
    )}
  </Section>

  <Section title="Ranked against your spec, or absolute" testid="ratings-basis" panel>
    {@render prose(
      ratingsPageCopy.rankedOrAbsolute,
      [],
      ['ratings-basis-rule', 'ratings-small-samples', 'ratings-wipes'],
    )}
  </Section>

  <Section title="The cap" panel>
    {@render prose([ratingCopy.cappedRule])}
  </Section>

  {#snippet rail()}
    <Section title={`${scoreCopy.tierColumn}s`} testid="ratings-forever-tiers" panel frame="framed">
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
            {#snippet trail()}{#if each === tier}<span class="text-small text-muted">The sample</span
                >{/if}{/snippet}
          </Row>
        </div>
      {/each}
      <p class="text-small text-muted mt-[10px]" data-testid="ratings-forever-tiers-note">
        {scoreCopy.tiersNote}
      </p>
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
