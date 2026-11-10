<!-- web/src/components/sim/SimAnswer.svelte -->
<!-- The page head's answer on /sim and /sim/<id> (spec 2026-11-06-sim §4.3 to §4.5, §5.1), on the
     compact answer contract (ui/PageHeader): a caps label, the big Figure, one caption line (the
     band captioned once as 95%, ruling S2, with the spec's fidelity Tag beside it), a fix line
     after a failure, the run Meter. The actions are the head's own (PageHeader `actions`, beside
     the answer). A browser result has no label (the default lane says nothing); its row is held so
     every state keeps one height, and the Meter hangs in the band's foot under the caption, so a
     run never moves the bar under it. aria-live sits on the captions, never on the figure: a
     number that changes ten times a second would make a screen reader unusable. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { AnswerModel } from '../../lib/sim/answer';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';

  let {
    model,
    unit = 'DPS',
    badge,
    extra,
  }: {
    model: AnswerModel;
    unit?: string;
    /** A Tag on the first caption's line (the spec's fidelity). */
    badge?: Snippet;
    /** A short line after the captions (a message the run left). */
    extra?: Snippet;
  } = $props();
</script>

<div class="sim-answer" data-testid="sim-answer" data-state={model.state}>
  <span class="sim-answer-label text-label text-muted font-bold" data-testid="sim-answer-label"
    >{model.label}</span
  >
  {#if model.figure}
    <span data-testid="sim-dps">
      <Figure
        value={model.figure.value}
        unit={model.figure.unit ? unit : undefined}
        tone={model.figure.tone}
        size="big"
        align="end"
      />
    </span>
  {/if}
  <span class="sim-answer-captions" aria-live="polite">
    {#if model.caption || badge}
      <span class="sim-answer-line">
        {#if model.caption}<span
            class={['text-[13px] leading-5', model.caption.tone === 'wipe' ? 'text-wipe' : 'text-muted']}
            data-testid="sim-error"
            >{#if model.caption.bold}<b class="text-text font-semibold">{model.caption.bold}</b>{/if}{model
              .caption.rest ?? ''}</span
          >{/if}
        {#if badge}{@render badge()}{/if}
      </span>
    {/if}
    {#if model.fix}<span class="text-muted text-[13px] leading-5" data-testid="sim-answer-fix"
        >{model.fix}</span
      >{/if}
  </span>
  {#if extra}{@render extra()}{/if}
  {#if model.meter}
    <span class="sim-answer-meter" data-testid="sim-progress-bar">
      <Meter value={model.meter.value} tone={model.meter.tone} label="Runs done" />
    </span>
  {/if}
</div>

<style>
  .sim-answer {
    position: relative;
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
  }
  .sim-answer-label {
    min-height: 16px;
  }
  .sim-answer-captions {
    display: flex;
    flex-direction: column;
  }
  /* The caption's line keeps one 24 px row with or without its Tag, so the head's actions centre on it. */
  .sim-answer-line {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 4px 8px;
    min-height: 24px;
  }
  /* The run Meter hangs under the caption in the band's foot: it never adds a row. */
  .sim-answer-meter {
    position: absolute;
    top: calc(100% + 6px);
    left: 0;
    display: block;
    width: 100%;
    min-width: 200px;
    max-width: 360px;
  }
  @media (min-width: 900px) {
    .sim-answer {
      align-items: flex-end;
      text-align: right;
    }
    .sim-answer-meter {
      left: auto;
      right: 0;
    }
    .sim-answer-line {
      justify-content: flex-end;
    }
  }
</style>
