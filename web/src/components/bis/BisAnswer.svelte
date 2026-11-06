<!-- web/src/components/bis/BisAnswer.svelte -->
<!-- The page head's answer (web layout spec 2026-11-05 §5.3 item 2, §5.10 item 6, §5.11): the
     set's figure for the chosen sources, the trust line under it, and the two links. A tank band
     shows its four figures instead of one; a healer band its HPS with the mana and overheal
     cells and the profile disclosure. -->
<script lang="ts">
  import Figure from '../ui/Figure.svelte';
  import type { BisAnswerModel, GainUnit } from '../../lib/bis/page-model';
  import type { SetFigure } from '../../lib/bis/source-choice';

  let {
    answer,
    figure,
    unit,
    raceNote,
  }: {
    answer: BisAnswerModel;
    /** The DPS/HPS figure for the chosen sources (source-choice.ts), absent on a tank band. */
    figure: SetFigure | undefined;
    unit: GainUnit;
    /** "Simmed as a Night Elf" when the race control names another race. */
    raceNote?: string;
  } = $props();
</script>

<div class="bis-answer" data-testid="bis-this-set">
  {#if answer.tank}
    <dl class="bis-tank-figures">
      {#each answer.tank.figures as tank (tank.id)}
        <div class="bis-tank-figure" title={tank.title} data-testid={`bis-tank-figure-${tank.id}`}>
          <dt class="tank-figure-label">{tank.label}</dt>
          <dd class="bis-tank-value">{tank.value}</dd>
        </div>
      {/each}
    </dl>
    <p class="bis-answer-line" title={answer.tank.secondaryTitle} data-testid="bis-tank-secondary">
      {answer.tank.secondaryLine}
    </p>
  {:else if figure}
    <div data-testid={answer.healer ? 'bis-healer-figure' : 'bis-set-figure'}>
      <Figure
        value={figure.value}
        {unit}
        label={figure.label}
        caption={figure.caption}
        size="big"
        align="end"
        testid="bis-answer-figure"
      />
    </div>
    {#if answer.healer}
      <p class="bis-answer-line" data-testid="bis-healer-caption">{answer.healer.caption}</p>
      {#if answer.healer.metrics.length > 0}
        <dl class="bis-healer-metrics" data-testid="bis-healer-metrics">
          {#each answer.healer.metrics as metric (metric.key)}
            <div title={metric.line} data-testid={`bis-healer-${metric.key}`}>
              <dt>{metric.label}</dt>
              <dd class="tabular-nums">{metric.value}</dd>
              <dd class="sr-only">{metric.line}</dd>
            </div>
          {/each}
        </dl>
      {/if}
    {/if}
  {/if}
  <p class="bis-answer-line" data-testid="bis-trust-line">
    {answer.trust.join(' · ')}{#if raceNote} · {raceNote}{/if}{#if answer.rankingSimHref} · <a
        href={answer.rankingSimHref}
        data-testid="bis-ranking-sim-link">View ranking sim ›</a
      >{/if}
  </p>
  {#if answer.healer?.disclosure}
    <details class="bis-answer-profile" data-testid="bis-healer-profile">
      <summary>{answer.healer.disclosure.summary}</summary>
      <p>{answer.healer.disclosure.profileSummary}</p>
      <dl>
        {#each answer.healer.disclosure.rows as row (row.heading)}
          <dt>{row.heading}</dt>
          <dd>{row.figures}</dd>
          {#if row.reason !== ''}<dd class="text-muted">{row.reason}</dd>{/if}
        {/each}
      </dl>
      <p class="text-muted">{answer.healer.disclosure.notes}</p>
    </details>
  {/if}
  {#if answer.presetGroups && answer.presetGroups.length > 0}
    <details class="bis-answer-profile" data-testid="bis-preset-effects">
      <summary>What this profile assumes</summary>
      <dl>
        {#each answer.presetGroups as group (group.heading)}
          <dt>{group.heading}</dt>
          <dd>{group.labels.join(', ')}</dd>
        {/each}
      </dl>
      {#if answer.presetNotes}<p class="text-muted">{answer.presetNotes}</p>{/if}
    </details>
  {/if}
  <p class="bis-answer-links">
    {#if answer.simHref}<a href={answer.simHref} data-testid="bis-sim-link">Open in simulator</a> · {/if}<a
      href={answer.plannerHref}
      data-testid="bis-planner-link">Talents in planner</a
    >
  </p>
</div>

<style>
  .bis-answer {
    display: flex;
    flex-direction: column;
    align-items: flex-end;
    gap: 4px;
    text-align: right;
    max-width: 640px;
  }
  .bis-answer-line {
    color: var(--color-muted);
    font-size: 13px;
    line-height: 18px;
  }
  .bis-answer-links {
    font-size: 14px;
    font-weight: 600;
  }
  .bis-answer a {
    color: var(--color-gold);
  }
  .bis-tank-figures,
  .bis-healer-metrics {
    display: flex;
    flex-wrap: wrap;
    justify-content: flex-end;
    gap: 4px 20px;
  }
  .bis-tank-figure dt,
  .bis-healer-metrics dt {
    color: var(--color-muted);
    font-size: 12px;
    font-weight: 700;
    letter-spacing: 0.06em;
    text-transform: uppercase;
  }
  .bis-tank-value {
    font-family: var(--font-figure, inherit);
    font-size: 24px;
    font-weight: 700;
    font-variant-numeric: tabular-nums;
  }
  .bis-answer-profile {
    font-size: 13px;
    color: var(--color-muted);
  }
  .bis-answer-profile summary {
    cursor: pointer;
    color: var(--color-gold);
  }
  @media (max-width: 767px) {
    .bis-answer {
      align-items: flex-start;
      text-align: left;
    }
    .bis-tank-figures,
    .bis-healer-metrics {
      justify-content: flex-start;
    }
  }
</style>
