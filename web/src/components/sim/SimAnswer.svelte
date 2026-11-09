<!-- web/src/components/sim/SimAnswer.svelte -->
<!-- The page head's answer on /sim and /sim/<id> (spec 2026-11-06-sim §4.3 to §4.5, §5.1): a caps
     label, the big Figure, the band captioned once as 95% (ruling S2), a fix line after a failure,
     the run Meter, then the actions. The slot keeps one size from the first state to the last
     (152 px, 169 px from 900 px wide, min 360 px wide), so a run never moves the bar under it.
     aria-live sits on the captions, never on the figure: a number that changes ten times a
     second would make a screen reader unusable. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import type { AnswerModel } from '../../lib/sim/answer';
  import Figure from '../ui/Figure.svelte';
  import Meter from '../ui/Meter.svelte';

  let {
    model,
    unit = 'DPS',
    extra,
    actions,
  }: {
    model: AnswerModel;
    unit?: string;
    /** A line after the captions (the fidelity note, the "Saved" stamp). */
    extra?: Snippet;
    actions?: Snippet;
  } = $props();
</script>

<div class="sim-answer" data-testid="sim-answer" data-state={model.state}>
  <span class="text-label text-muted font-bold" data-testid="sim-answer-label">{model.label}</span>
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
    {#if model.caption}
      <span
        class={['text-[13px] leading-5', model.caption.tone === 'wipe' ? 'text-wipe' : 'text-muted']}
        data-testid="sim-error"
        >{#if model.caption.bold}<b class="text-text font-semibold">{model.caption.bold}</b>{/if}{model
          .caption.rest ?? ''}</span
      >
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
  {#if actions}<span class="sim-answer-actions">{@render actions()}</span>{/if}
</div>

<style>
  .sim-answer {
    display: flex;
    flex-direction: column;
    gap: 4px;
    min-width: 0;
    min-height: 152px;
  }
  .sim-answer-captions {
    display: flex;
    flex-direction: column;
  }
  .sim-answer-meter {
    display: block;
    width: 100%;
    max-width: 360px;
    margin-top: 6px;
  }
  .sim-answer-actions {
    display: flex;
    gap: 16px;
    align-items: center;
    flex-wrap: wrap;
    margin-top: 6px;
  }
  @media (min-width: 900px) {
    .sim-answer {
      min-width: 360px;
      min-height: 169px;
      align-items: flex-end;
      text-align: right;
    }
    .sim-answer-meter {
      align-self: flex-end;
    }
    .sim-answer-actions {
      justify-content: flex-end;
    }
  }
</style>
