<!-- web/src/components/bis/BisWeights.svelte -->
<!-- "What to stack" (web layout spec 2026-11-05 §5.10 item 11): one sentence naming the stats in
     order, and a disclosure with the weights table (bar, scale, per point, error). The addon link
     stays a link: the addon reads its weights from its own data, it has no paste format. -->
<script lang="ts">
  import Section from '../ui/Section.svelte';
  import Meter from '../ui/Meter.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import type { BisWeightsModel } from '../../lib/bis/page-model';

  let { weights, testid }: { weights: BisWeightsModel; testid: string } = $props();

  const lead = $derived(weights.rows.find((row) => row.significant));
</script>

<Section title="What to stack" {testid}>
  {#if weights.sentence}<p class="bis-weights-sentence" data-testid="bis-weights-sentence">
      {weights.sentence}
    </p>{/if}
  {#if weights.rows.length > 0}
    <details class="bis-weights-fold" data-testid="bis-weights-fold">
      <summary>Show the weights and their error</summary>
      <p class="bis-weights-caption text-muted">
        {lead ? `Normalized to ${lead.label} = 1.00. ` : ''}{weights.note}{weights.haste
          ? ` ${weights.haste}`
          : ''}
      </p>
      <div class="bis-weights-table" role="table" aria-label="Stat weights">
        <div class="bis-weights-row bis-weights-header" role="row">
          <span role="columnheader">Stat</span>
          <span aria-hidden="true"></span>
          <span role="columnheader" class="num">Scale</span>
          <span role="columnheader" class="num">Per point</span>
          <span role="columnheader" class="num">Error</span>
        </div>
        {#each weights.rows as row (row.stat)}
          <div
            class={['bis-weights-row', !row.significant && 'faint']}
            role="row"
            title={row.significant
              ? row.ratingFactorTitle
              : `${row.label} ${row.reading === 'no-gain' ? 'adds nothing' : 'not significant'}`}
            data-testid={`bis-weight-row-${row.stat}`}
          >
            <span role="cell">{row.label}</span>
            <span class="bar" aria-hidden="true"
              >{#if row.significant}<Meter value={row.barPercent / 100} />{/if}</span
            >
            {#if row.significant}
              <span role="cell" class="num">{row.scaleFactor.toFixed(2)}</span>
              <span role="cell" class="num muted"
                >{row.dpsPerPoint !== undefined
                  ? `${row.dpsPerPoint.toFixed(3)} ${row.perPointUnit}`
                  : ''}</span
              >
              <span role="cell" class="num muted">{`±${row.scaleError.toFixed(2)}`}</span>
            {:else}
              <span role="cell" class="num muted span3"
                >{row.reading === 'no-gain' ? bisCopy.weightsNoGain : bisCopy.weightsNotSignificant}</span
              >
            {/if}
          </div>
        {/each}
      </div>
      {#if weights.expertiseCap}
        <p
          class="text-muted bis-weights-line"
          title={weights.expertiseCap.title}
          data-testid={`${testid}-expertise-cap`}
        >
          {weights.expertiseCap.text}
        </p>
      {/if}
    </details>
  {/if}
  <p class="bis-weights-line text-muted">
    {bisCopy.weightsAddonBlurb}
    <a href="/addon" data-testid="bis-weights-addon-button">{bisCopy.weightsAddonButton} ›</a>
  </p>
</Section>

<style>
  .bis-weights-sentence {
    font-size: 15px;
    line-height: 24px;
  }
  .bis-weights-fold summary {
    cursor: pointer;
    color: var(--color-gold);
    font-weight: 600;
    margin-top: 8px;
  }
  .bis-weights-fold summary:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .bis-weights-caption,
  .bis-weights-line {
    font-size: 13px;
    line-height: 20px;
    margin-top: 8px;
  }
  .bis-weights-line a {
    color: var(--color-gold);
    font-weight: 600;
  }
  .bis-weights-table {
    margin-top: 8px;
    max-width: 760px;
  }
  .bis-weights-row {
    display: grid;
    grid-template-columns: minmax(120px, 1.4fr) minmax(60px, 2fr) 64px minmax(96px, 1.2fr) 64px;
    gap: 12px;
    align-items: center;
    min-height: 32px;
    font-size: 14px;
  }
  .bis-weights-header {
    color: var(--color-muted);
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  .bis-weights-row.faint {
    opacity: 0.6;
  }
  .num {
    text-align: right;
    font-variant-numeric: tabular-nums;
  }
  .muted {
    color: var(--color-muted);
  }
  .span3 {
    grid-column: span 3;
  }
  .bar {
    min-width: 0;
  }
  @media (max-width: 767px) {
    .bis-weights-row {
      grid-template-columns: minmax(96px, 1fr) 52px minmax(80px, 1fr) 52px;
    }
    .bar {
      display: none;
    }
  }
</style>
