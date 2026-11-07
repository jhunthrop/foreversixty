<!-- web/src/components/ui/TalentButton.svelte -->
<!-- One talent (planner spec 2026-11-08 section 3A, the mock's TalentCell2): the icon, a 2 px inset
     ring and a rank pill in the state's colour (avail, part, max, lock), the gold dot when it
     differs from the band, a dashed dimmed pill when planned past the character's level. Click
     adds; right-click or a 500 ms primary press removes, exactly once per gesture. A locked cell
     is aria-disabled but its click still reaches `onadd`: the caller refuses and says why. The
     tooltip opens 150 ms after hover or focus and closes on leave, blur or Escape, placed with the
     site's one edge-aware rule (lib/items/tooltip-position). Arrow keys and Backspace belong to
     the grid (TalentEditor), which owns the one tab stop per tree. Read-only: no gesture, the
     tooltip still opens. -->
<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  import { onDestroy } from 'svelte';
  import { tooltipLeftPx, tooltipVerticalPlacementFor } from '../../lib/items/tooltip-position';
  import {
    cellLabel,
    isPlaceholderIcon,
    LONG_PRESS_MS,
    TIP_DELAY_MS,
    type EditorCell,
    type TalentEditorCopy,
  } from '../../lib/ui/talent-editor';
  import TalentTip from './TalentTip.svelte';

  let {
    cell,
    copy,
    focused = false,
    readOnly = false,
    onadd,
    onremove,
    onfocuscell,
  }: {
    cell: EditorCell;
    copy: Pick<TalentEditorCopy, 'rankOf' | 'nextRank'>;
    /** This cell holds its tree's one tab stop. */
    focused?: boolean;
    readOnly?: boolean;
    onadd: (id: number) => void;
    onremove: (id: number) => void;
    onfocuscell?: (id: number) => void;
  } = $props();

  const tipId = $derived(`talent-tip-${cell.id}`);

  let open = $state(false);
  let iconBroken = $state(false);
  let button: HTMLButtonElement | undefined = $state();
  let tipBox: HTMLDivElement | undefined = $state();

  interface TipPlacement {
    left: number;
    below: boolean;
    maxHeightPx: number | null;
  }
  // Unshifted and unclamped: reset on every close so the next opening measures the natural box.
  const NATURAL: TipPlacement = { left: 0, below: true, maxHeightPx: null };
  let placement = $state<TipPlacement>(NATURAL);

  $effect(() => {
    if (!open || !tipBox || !button) {
      placement = NATURAL;
      return;
    }
    const anchor = button.getBoundingClientRect();
    const root = document.documentElement;
    const { side, maxHeightPx } = tooltipVerticalPlacementFor(
      anchor.top,
      anchor.bottom,
      tipBox.offsetHeight,
      root.clientHeight,
    );
    placement = {
      left: tooltipLeftPx('start', anchor.left, anchor.width, root.clientWidth, tipBox.offsetWidth),
      below: side === 'below',
      maxHeightPx,
    };
  });

  let tipTimer: ReturnType<typeof setTimeout> | undefined;

  function showTip(): void {
    clearTimeout(tipTimer);
    tipTimer = setTimeout(() => (open = true), TIP_DELAY_MS);
  }

  function hideTip(): void {
    clearTimeout(tipTimer);
    tipTimer = undefined;
    open = false;
  }

  let pressTimer: ReturnType<typeof setTimeout> | undefined;
  // True once the gesture in progress has removed its point. The click that ends a long press
  // and the platform's touch-and-hold contextmenu would otherwise each remove another; whichever
  // fires first sets this so one gesture is always exactly one removal.
  let gestureRemoved = false;

  function startPress(event: PointerEvent): void {
    // Only a primary press can become a long press: a secondary button raises contextmenu itself.
    if (readOnly || event.button !== 0) return;
    gestureRemoved = false;
    pressTimer = setTimeout(() => {
      disarmPress();
      gestureRemoved = true;
      onremove(cell.id);
    }, LONG_PRESS_MS);
  }

  /** Cancels a pending long press; the click still follows, so `gestureRemoved` stays. */
  function disarmPress(): void {
    clearTimeout(pressTimer);
    pressTimer = undefined;
  }

  /** The press ended elsewhere: no click follows, nothing is left to suppress. */
  function abandonPress(): void {
    disarmPress();
    gestureRemoved = false;
  }

  function removeOnContextMenu(event: MouseEvent): void {
    event.preventDefault();
    if (readOnly || gestureRemoved) return;
    // A contextmenu during a pending press is that press's own touch gesture and a click still
    // follows it; a right-click arms no timer and no click follows.
    const duringPress = pressTimer !== undefined;
    disarmPress();
    gestureRemoved = duringPress;
    onremove(cell.id);
  }

  function add(): void {
    if (readOnly) return;
    if (gestureRemoved) {
      gestureRemoved = false;
      return;
    }
    onadd(cell.id);
  }

  function closeOnEscape(event: KeyboardEvent): void {
    if (event.key !== 'Escape' || (!open && tipTimer === undefined)) return;
    event.stopPropagation();
    hideTip();
  }

  onDestroy(() => {
    clearTimeout(tipTimer);
    clearTimeout(pressTimer);
  });
</script>

<button
  bind:this={button}
  type="button"
  class={['tc2', cell.state, cell.differs && 'diff', cell.planned && 'plan']}
  aria-label={cellLabel(cell)}
  aria-disabled={cell.state === 'lock' && cell.plannable !== true ? 'true' : undefined}
  aria-describedby={open ? tipId : undefined}
  tabindex={focused ? 0 : -1}
  data-testid={`talent-${cell.id}`}
  data-rank={cell.rank}
  data-state={cell.state}
  onclick={add}
  oncontextmenu={removeOnContextMenu}
  onfocus={() => {
    showTip();
    onfocuscell?.(cell.id);
  }}
  onblur={hideTip}
  onmouseenter={showTip}
  onmouseleave={() => {
    hideTip();
    abandonPress();
  }}
  onpointerdown={startPress}
  onpointerup={disarmPress}
  onpointercancel={abandonPress}
  onkeydown={closeOnEscape}
>
  {#if iconBroken || cell.iconSrc === ''}
    <span class="ic" aria-hidden="true"></span>
  {:else}
    <img
      src={cell.iconSrc}
      alt=""
      width="48"
      height="48"
      loading="eager"
      decoding="async"
      onerror={() => (iconBroken = true)}
      onload={(event) => {
        if (isPlaceholderIcon(event.currentTarget)) iconBroken = true;
      }}
    />
  {/if}
  <span class="rk" aria-hidden="true">{cell.rank}/{cell.max}</span>
</button>
{#if open}
  <div
    bind:this={tipBox}
    class="tc2-tip"
    class:above={!placement.below}
    class:scroll={placement.maxHeightPx !== null}
    style:left={`${placement.left}px`}
    style:max-height={placement.maxHeightPx === null ? undefined : `${placement.maxHeightPx}px`}
  >
    <TalentTip
      id={tipId}
      talentId={cell.id}
      name={cell.name}
      rankLine={copy.rankOf(cell.rank, cell.max)}
      now={cell.now}
      next={cell.next}
      nextLabel={copy.nextRank}
      req={cell.req}
      band={cell.band}
    />
  </div>
{/if}

<style>
  .tc2 {
    all: unset;
    box-sizing: border-box;
    position: relative;
    display: block;
    width: var(--cell, 48px);
    height: var(--cell, 48px);
    border-radius: 5px;
    cursor: pointer;
    background: #10141d;
  }
  .tc2 img,
  .tc2 .ic {
    display: block;
    width: 100%;
    height: 100%;
    border-radius: 5px;
    border: 0;
  }
  .tc2::after {
    content: '';
    position: absolute;
    inset: 0;
    border-radius: 5px;
    box-shadow: inset 0 0 0 2px var(--ring, var(--color-line));
  }
  .rk {
    position: absolute;
    right: -5px;
    bottom: -5px;
    z-index: 1;
    padding: 0 4px;
    border-radius: 999px;
    border: 1px solid var(--ring, var(--color-line));
    background: var(--color-bg);
    font: 700 12px/14px var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--rk, var(--color-muted));
    white-space: nowrap;
  }
  .tc2.avail {
    --ring: #4fae5c;
    --rk: var(--color-kill);
  }
  .tc2.part {
    --ring: var(--color-kill);
    --rk: var(--color-kill);
  }
  .tc2.max {
    --ring: var(--color-gold);
    --rk: var(--color-gold);
  }
  .tc2.lock {
    cursor: default;
  }
  .tc2.lock img,
  .tc2.lock .ic {
    filter: grayscale(1) brightness(0.55);
  }
  .tc2.diff::before {
    content: '';
    position: absolute;
    left: -4px;
    top: -4px;
    z-index: 1;
    width: 9px;
    height: 9px;
    border-radius: 50%;
    background: var(--color-gold);
    box-shadow: 0 0 0 2px var(--color-bg);
  }
  .tc2:hover:not(.lock) {
    filter: brightness(1.15);
  }
  .tc2:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .tc2:active {
    transform: translateY(1px);
  }
  .tc2.plan .rk {
    opacity: 0.6;
    border-style: dashed;
  }
  .tc2-tip {
    position: absolute;
    top: 100%;
    z-index: 30;
    margin-top: 8px;
  }
  .tc2-tip.above {
    top: auto;
    bottom: 100%;
    margin: 0 0 8px;
  }
  .tc2-tip.scroll {
    overflow-y: auto;
  }
</style>
