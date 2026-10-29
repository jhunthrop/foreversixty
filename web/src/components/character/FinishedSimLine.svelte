<!-- web/src/components/character/FinishedSimLine.svelte -->
<!-- The band's footer row (experience spec 2026-10-24 section 3.4): "Your Frost Mage sim
     finished: 1,204 DPS  Open result →", `and N more` when several wait, and a dismiss. It
     fills the slot CharacterHeader gives it (full width, flush inside the panel border); it
     owns its tint, its copy and the dismiss. Following the link marks that run seen on the
     result page (3.6), so nothing here posts on a link click. Dismiss posts the seen call for
     every run the line names. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import type { FinishedSim } from '../../lib/account/api';
  import { DISMISS_FAILURE_MS, markAllSeen } from '../../lib/character-header/finished-sim-line';
  import { markSimSeen } from '../../lib/sim/api';
  import { serverRunsCopy } from '../../lib/sim/server-runs-copy';
  import { specLabel } from '../../lib/sim/spec-label';

  let {
    runs,
    onDismissed,
    markSeen = markSimSeen,
  }: {
    /** Newest first; the first is the one the sentence names. */
    runs: readonly [FinishedSim, ...FinishedSim[]];
    onDismissed: () => void;
    markSeen?: (simId: string) => Promise<void>;
  } = $props();

  type Phase = 'idle' | 'busy' | 'failed';
  let phase = $state<Phase>('idle');
  let restoreTimer: ReturnType<typeof setTimeout> | undefined;

  const newest = $derived(runs[0]);
  const label = $derived(specLabel(newest.spec));
  const olderCount = $derived(runs.length - 1);
  const sentence = $derived(serverRunsCopy.bandLineParts(label, newest.dps));

  async function dismiss(): Promise<void> {
    if (phase === 'busy') return;
    clearTimeout(restoreTimer);
    phase = 'busy';
    if (await markAllSeen(runs, markSeen)) {
      onDismissed();
      return;
    }
    phase = 'failed';
    clearTimeout(restoreTimer);
    restoreTimer = setTimeout(() => {
      phase = 'idle';
    }, DISMISS_FAILURE_MS);
  }

  onDestroy(() => clearTimeout(restoreTimer));
</script>

<div class="fsl" role="status" data-testid="finished-sim-line">
  {#if phase === 'failed'}
    <p class="fsl-failed" data-testid="finished-sim-failed">{serverRunsCopy.dismissFailed}</p>
  {:else}
    <div class="fsl-text">
      <a
        class="fsl-link"
        href={`/sim/${newest.sim_id}`}
        title={newest.title === '' ? undefined : newest.title}
        data-testid="finished-sim-link"
      >
        <span class="fsl-sentence">{sentence.prefix}<span class="fsl-figure">{sentence.figure}</span></span>
        <span class="fsl-open fsl-open-wide" data-testid="finished-sim-open">{serverRunsCopy.bandOpen}</span>
        <span class="fsl-open fsl-open-phone" data-testid="finished-sim-open-phone"
          >{serverRunsCopy.bandOpenPhone}</span
        >
      </a>
      {#if olderCount > 0}
        <a class="fsl-more" href="/sim#server-runs" data-testid="finished-sim-more"
          >{serverRunsCopy.moreLabel(olderCount)}</a
        >
      {/if}
    </div>
  {/if}
  <button
    type="button"
    class="fsl-dismiss"
    aria-label={serverRunsCopy.dismissLabel(label)}
    aria-busy={phase === 'busy'}
    onclick={dismiss}
    data-testid="finished-sim-dismiss"
    ><span aria-hidden="true"
      ><svg
        width="16"
        height="16"
        viewBox="0 0 16 16"
        fill="none"
        stroke="currentColor"
        stroke-width="1.75"
        stroke-linecap="round"><path d="M3.5 3.5l9 9M12.5 3.5l-9 9" /></svg
      ></span
    ></button
  >
</div>

<style>
  .fsl {
    display: grid;
    grid-template-columns: minmax(0, 1fr) auto;
    align-items: center;
    box-sizing: border-box;
    min-height: var(--ch-foot-phone-h);
    padding: 6px 14px;
    border-top: 1px solid var(--color-line-soft);
    background: color-mix(in srgb, var(--color-gold) 8%, var(--color-raised));
    font-size: 13px;
    color: var(--color-text);
  }
  @media (min-width: 600px) {
    .fsl {
      min-height: var(--ch-foot-h);
      padding: 0 22px;
    }
  }
  .fsl-text {
    display: flex;
    flex-direction: column;
    align-items: flex-start;
    min-width: 0;
  }
  @media (min-width: 600px) {
    .fsl-text {
      flex-direction: row;
      align-items: center;
      gap: 8px;
    }
  }
  .fsl-link,
  .fsl-more {
    display: inline-flex;
    align-items: center;
    flex-wrap: wrap;
    gap: 0 6px;
    min-height: 44px;
    color: var(--color-text);
  }
  .fsl-link:hover .fsl-sentence {
    text-decoration: underline;
    text-underline-offset: 4px;
  }
  .fsl-figure {
    font-family: var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--color-strong);
  }
  .fsl-open {
    font-weight: 700;
    color: var(--color-gold);
    white-space: nowrap;
  }
  .fsl-link:hover .fsl-open,
  .fsl-link:active .fsl-open {
    color: var(--color-gold-hover);
  }
  .fsl-open-phone {
    display: none;
  }
  @media (max-width: 599px) {
    .fsl-open-wide {
      display: none;
    }
    .fsl-open-phone {
      display: inline;
    }
  }
  .fsl-more {
    font-size: 12px;
    color: var(--color-muted);
    text-decoration: underline;
  }
  @media (min-width: 600px) {
    .fsl-more {
      font-size: 13px;
    }
  }
  .fsl-more:hover,
  .fsl-more:active {
    color: var(--color-strong);
  }
  .fsl-link:focus-visible,
  .fsl-more:focus-visible,
  .fsl-dismiss:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
    border-radius: 4px;
  }
  .fsl-failed {
    margin: 0;
    color: var(--color-muted);
  }
  .fsl-dismiss {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 44px;
    height: 44px;
    margin-right: -14px;
    padding: 0;
    border: 0;
    border-radius: 999px;
    background: transparent;
    font-size: 16px;
    line-height: 1;
    color: var(--color-muted);
    cursor: pointer;
  }
  @media (min-width: 600px) {
    .fsl-dismiss {
      margin-right: -22px;
    }
  }
  .fsl-dismiss span {
    display: inline-flex;
    align-items: center;
    justify-content: center;
    width: 32px;
    height: 32px;
    border-radius: 999px;
  }
  .fsl-dismiss:hover {
    color: var(--color-strong);
  }
  .fsl-dismiss:hover span {
    background: rgb(255 255 255 / 6%);
  }
  .fsl-dismiss:active span {
    background: rgb(255 255 255 / 12%);
  }
</style>
