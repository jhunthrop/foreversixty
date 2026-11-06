<!-- web/src/components/bis/BisWeights.svelte -->
<!-- "What to stack" (web layout spec 2026-11-05 §5.10 item 11): one sentence naming the stats in
     order, and one column per stat (name, scale figure, bar, per point; hit cap and sim error on
     the stat's title) that stacks as rows on a phone. The addon link
     stays a link: the addon reads its weights from its own data, it has no paste format. -->
<script lang="ts">
  import Section from '../ui/Section.svelte';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';
  import { bisCopy } from '../../lib/bis/copy';
  import type { BisWeightsModel } from '../../lib/bis/page-model';
  import type { ScaleRow } from '../../lib/bis/scale-view';

  let { weights, testid }: { weights: BisWeightsModel; testid: string } = $props();

  const lead = $derived(weights.rows.find((row) => row.significant));

  /** A stat's hover: its rating factor and sim error, or why it carries no figure; the hit stat
   *  also names the cap. */
  function weightTitle(row: ScaleRow): string {
    if (!row.significant) {
      return `${row.label} ${row.reading === 'no-gain' ? 'adds nothing' : 'not significant'}`;
    }
    const parts = [row.ratingFactorTitle, `Sim error ±${row.scaleError.toFixed(2)}`];
    if (row.stat.startsWith('hit') && weights.hitCap) parts.push(weights.hitCap.title);
    return parts.filter((part) => part !== undefined && part !== '').join('. ');
  }
</script>

<Section title="What to stack" {testid}>
  {#if weights.sentence}<p class="bis-weights-sentence" data-testid="bis-weights-sentence">
      {weights.sentence}
    </p>{/if}
  {#if weights.rows.length > 0}
    <p class="bis-weights-caption text-muted">
      {lead ? `Normalized to ${lead.label} = 1.00. ` : ''}{weights.note}{weights.haste
        ? ` ${weights.haste}`
        : ''}
    </p>
    <ul class="bis-weights-columns" aria-label="Stat weights">
      {#each weights.rows as row (row.stat)}
        <li
          class={['bis-weight', !row.significant && 'faint']}
          title={weightTitle(row)}
          data-testid={`bis-weight-row-${row.stat}`}
        >
          <span class="bis-weight-name">{row.label}</span>
          {#if row.significant}
            <Figure value={row.scaleFactor.toFixed(2)} size="sm" />
            <span class="bar" aria-hidden="true"><Meter value={row.barPercent / 100} /></span>
            <span class="muted"
              >{row.dpsPerPoint !== undefined
                ? `${row.dpsPerPoint.toFixed(3)} ${row.perPointUnit} per point`
                : ''}</span
            >
          {:else}
            <span class="muted"
              >{row.reading === 'no-gain' ? bisCopy.weightsNoGain : bisCopy.weightsNotSignificant}</span
            >
          {/if}
        </li>
      {/each}
    </ul>
    {#if weights.expertiseCap}
      <p
        class="text-muted bis-weights-line"
        title={weights.expertiseCap.title}
        data-testid={`${testid}-expertise-cap`}
      >
        {weights.expertiseCap.text}
      </p>
    {/if}
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
  .bis-weights-columns {
    display: grid;
    grid-template-columns: repeat(auto-fit, minmax(150px, 1fr));
    gap: 16px;
    margin-top: 12px;
    padding: 0;
    list-style: none;
  }
  .bis-weight {
    display: flex;
    flex-direction: column;
    gap: 6px;
    min-width: 0;
    font-size: 14px;
  }
  .bis-weight-name {
    font-size: 15px;
    font-weight: 600;
  }
  .bis-weight.faint {
    opacity: 0.6;
  }
  .muted {
    color: var(--color-muted);
  }
  @media (max-width: 767px) {
    .bis-weights-columns {
      grid-template-columns: 1fr;
      gap: 8px;
    }
    .bis-weight {
      display: grid;
      grid-template-columns: minmax(96px, 1fr) 52px minmax(60px, 1fr);
      align-items: center;
      gap: 12px;
    }
    .bis-weight .bar {
      display: none;
    }
  }
</style>
