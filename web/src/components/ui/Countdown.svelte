<!-- web/src/components/ui/Countdown.svelte -->
<!-- A countdown panel (home B spec 2026-11-07 §1): a caps eyebrow (with a live dot), a clock of
     days, hours and minutes to a known instant (seconds only in the final hour, lib/ui/countdown.ts)
     and the panel's own content at the right. It ticks in the browser from the instant it is given;
     `now` freezes it (tests). With no instant, or once it has passed, only the content shows. The
     live dot pulses only when the reader allows motion. -->
<script lang="ts">
  import { onMount, type Snippet } from 'svelte';
  import { clockLabel, clockParts, nextTickMs } from '../../lib/ui/countdown';

  let {
    eyebrow,
    target,
    live = false,
    tone = 'neutral',
    now: fixedNow,
    testid,
    children,
  }: {
    eyebrow: string;
    /** An ISO instant; null: no clock. */
    target: string | null;
    live?: boolean;
    tone?: 'raid' | 'guild' | 'neutral';
    now?: Date;
    testid?: string;
    children?: Snippet;
  } = $props();

  let ticking = $state(new Date());
  const now = $derived(fixedNow ?? ticking);
  const parts = $derived(target === null ? null : clockParts(new Date(target), now));

  onMount(() => {
    if (fixedNow !== undefined || target === null) return;
    let timer: ReturnType<typeof setTimeout>;
    const tick = () => {
      ticking = new Date();
      timer = setTimeout(tick, nextTickMs(new Date(target), ticking));
    };
    tick();
    return () => clearTimeout(timer);
  });
</script>

<div class="ui-countdown" data-tone={tone} data-testid={testid}>
  <p class="ui-countdown-eyebrow">
    {#if live}<span class="ui-countdown-live" aria-hidden="true"></span>{/if}{eyebrow}
  </p>
  <div class="ui-countdown-row">
    {#if parts !== null}
      <div class="ui-countdown-clock" role="timer" aria-label={clockLabel(parts)}>
        {#each parts as part (part.unit)}
          <span class="ui-countdown-part" aria-hidden="true"><b>{part.value}</b><span>{part.unit}</span></span
          >
        {/each}
      </div>
    {/if}
    {#if children}<div class="ui-countdown-aside">{@render children()}</div>{/if}
  </div>
</div>

<style>
  .ui-countdown {
    position: relative;
    min-width: 0;
    padding: 20px 24px;
    border: 1px solid var(--color-line);
    border-radius: 8px;
    background: var(--color-raised);
  }
  /* Round 5's raid plaque (the guild's next raid): the cave mouth and its lava glow in the
     plaque's right edge, under the cave veil (text 11.9:1). The first-raids countdown stays plain. */
  .ui-countdown[data-tone='guild'] {
    background:
      linear-gradient(var(--color-veil-cave), var(--color-veil-cave)),
      var(--art-cave, none) right center / 160% auto no-repeat,
      var(--color-raised);
  }
  .ui-countdown-eyebrow {
    margin: 0;
    font-weight: 700;
    font-size: 12px;
    line-height: 16px;
    letter-spacing: 0;
    color: var(--color-muted);
  }
  .ui-countdown-live {
    display: inline-block;
    width: 8px;
    height: 8px;
    margin-right: 8px;
    border-radius: 50%;
    background: var(--color-kill);
    vertical-align: 1px;
  }
  @media (prefers-reduced-motion: no-preference) {
    .ui-countdown-live {
      animation: ui-countdown-pulse 2s infinite;
    }
  }
  @keyframes ui-countdown-pulse {
    0% {
      box-shadow: 0 0 0 0 #7fd48a88;
    }
    70% {
      box-shadow: 0 0 0 8px #7fd48a00;
    }
    100% {
      box-shadow: 0 0 0 0 #7fd48a00;
    }
  }
  .ui-countdown-row {
    display: flex;
    flex-wrap: wrap;
    align-items: flex-end;
    gap: 16px 28px;
    margin-top: 10px;
  }
  .ui-countdown-clock {
    display: flex;
    gap: 14px;
    font-variant-numeric: tabular-nums;
  }
  .ui-countdown-part {
    display: flex;
    flex-direction: column;
    align-items: center;
  }
  .ui-countdown-part b {
    font-weight: 700;
    font-size: 40px;
    line-height: 42px;
    color: var(--color-strong);
  }
  .ui-countdown-part span {
    font-size: 12px;
    line-height: 16px;
    letter-spacing: 0;
    color: var(--color-muted);
  }
  /* Grows beside the clock and wraps its own lines (a long "your time" line on a viewer outside
     the guild's zone); it drops below the clock only when the row cannot hold 200 px of it. */
  .ui-countdown-aside {
    flex: 1 1 200px;
    min-width: 0;
    margin-left: auto;
  }
</style>
