<script lang="ts">
  import { onDestroy, tick } from 'svelte';
  import { positionTooltipPanel } from '../../lib/items/tooltip-position';
  import {
    cellLabel,
    cellTipKey,
    cellTipModel,
    isPlaceholderIcon,
    LONG_PRESS_MS,
    TIP_DELAY_MS,
    type TalentCellView,
  } from '../../lib/ui/talent-tree';
  import GameTooltip from '../GameTooltip.svelte';

  let {
    cell,
    build,
    editable = false,
    focusable = true,
    onadd,
    onremove,
    onfocuscell,
  }: {
    cell: TalentCellView;
    /** The data build the tooltip's icon comes from. */
    build: string;
    /** The planner: a click adds, right-click or a long press removes, the tooltip is its own. */
    editable?: boolean;
    /** This cell holds its tree's one tab stop. */
    focusable?: boolean;
    onadd?: (id: number) => void;
    onremove?: (id: number) => void;
    onfocuscell?: (id: number) => void;
  } = $props();

  const tipId = $derived(`talent-tip-${cell.id}`);
  const tip = $derived(editable ? cellTipModel(cell) : null);

  let open = $state(false);
  let iconBroken = $state(false);
  let wrap: HTMLSpanElement | undefined = $state();

  // Placed by the site's one rule (beside the cell when there is room, else below or above), as
  // every read-only cell's tooltip is placed by the page host.
  $effect(() => {
    if (!open || !wrap || tip === null) return;
    const host = wrap;
    void tick().then(() => positionTooltipPanel(host));
  });

  let tipTimer: ReturnType<typeof setTimeout> | undefined;

  function showTip(delay = TIP_DELAY_MS): void {
    clearTimeout(tipTimer);
    tipTimer = setTimeout(() => (open = true), delay);
  }

  function hideTip(): void {
    clearTimeout(tipTimer);
    tipTimer = undefined;
    open = false;
  }

  // A tap on a touch screen adds the point and shows what it bought; a tap anywhere else closes it.
  $effect(() => {
    if (!open) return;
    const close = (event: PointerEvent): void => {
      if (wrap && !wrap.contains(event.target as Node)) hideTip();
    };
    document.addEventListener('pointerdown', close);
    return () => document.removeEventListener('pointerdown', close);
  });

  let pressTimer: ReturnType<typeof setTimeout> | undefined;
  // True once the gesture in progress has removed its point. The click that ends a long press
  // and the platform's touch-and-hold contextmenu would otherwise each remove another; whichever
  // fires first sets this so one gesture is always exactly one removal.
  let gestureRemoved = false;

  function startPress(event: PointerEvent): void {
    // Only a primary press can become a long press: a secondary button raises contextmenu itself.
    if (event.button !== 0) return;
    gestureRemoved = false;
    pressTimer = setTimeout(() => {
      disarmPress();
      gestureRemoved = true;
      onremove?.(cell.id);
    }, LONG_PRESS_MS);
  }

  /** Cancels a pending long press; the click still follows, so `gestureRemoved` stays. */
  function disarmPress(): void {
    clearTimeout(pressTimer);
    pressTimer = undefined;
  }

  function endPress(event: PointerEvent): void {
    disarmPress();
    if (event.pointerType === 'touch') showTip(0);
  }

  /** The press ended elsewhere: no click follows, nothing is left to suppress. */
  function abandonPress(): void {
    disarmPress();
    gestureRemoved = false;
  }

  function removeOnContextMenu(event: MouseEvent): void {
    event.preventDefault();
    if (gestureRemoved) return;
    // A contextmenu during a pending press is that press's own touch gesture and a click still
    // follows it; a right-click arms no timer and no click follows.
    const duringPress = pressTimer !== undefined;
    disarmPress();
    gestureRemoved = duringPress;
    onremove?.(cell.id);
  }

  function add(): void {
    if (gestureRemoved) {
      gestureRemoved = false;
      return;
    }
    onadd?.(cell.id);
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

<!-- web/src/components/ui/TalentButton.svelte -->
<!-- One talent, the same in every tree on the site (talent-tree lane 2026-11-10; planner spec
     2026-11-08 section 3A, the mock's TalentCell2): the icon, a 2 px inset ring and a rank pill in
     the state's colour (avail, part, max, lock), the gold dot when it differs from the band, a
     dashed dimmed pill when planned past the character's level. Its tooltip is the site's one game
     tooltip (GameTooltip), placed beside the cell by the one placement rule (`data-tip-beside`).
     Read-only (`editable` false): the cell carries its `data-game-tip` key and a page host opens the
     tooltip on hover, keyboard focus and a tap, as for every spell and item. Editable: the cell
     opens its own tooltip 150 ms after hover or keyboard focus, and after a tap on a touch screen
     (which also adds); a click adds; right-click or a 500 ms primary press removes, exactly once
     per gesture. A locked cell is aria-disabled but its click still reaches `onadd`: the caller
     refuses and says why. Arrow keys and Backspace belong to the grid (TalentTree), which owns the
     one tab stop per tree. -->

{#snippet face()}
  {#if iconBroken || cell.iconSrc === ''}
    <span class="tt-ic" aria-hidden="true"></span>
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
  <span class="tt-rk" aria-hidden="true">{cell.rank}/{cell.max}</span>
{/snippet}

{#if editable}
  <span class="tt-wrap" bind:this={wrap} data-tip-beside>
    <button
      type="button"
      class={['tt-cell', cell.state, cell.differs && 'diff', cell.planned && 'plan']}
      aria-label={cellLabel(cell)}
      aria-disabled={cell.state === 'lock' && cell.plannable !== true ? 'true' : undefined}
      aria-describedby={open ? tipId : undefined}
      tabindex={focusable ? 0 : -1}
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
      onmouseenter={() => showTip()}
      onmouseleave={() => {
        hideTip();
        abandonPress();
      }}
      onpointerdown={startPress}
      onpointerup={endPress}
      onpointercancel={abandonPress}
      onkeydown={closeOnEscape}
    >
      {@render face()}
    </button>
    {#if open && tip !== null}
      <GameTooltip model={tip} {build} id={tipId} />
    {/if}
  </span>
{:else}
  <button
    type="button"
    class={['tt-cell', cell.state, cell.differs && 'diff']}
    aria-label={cellLabel(cell)}
    tabindex={focusable ? 0 : -1}
    data-testid={`talent-${cell.id}`}
    data-rank={cell.rank}
    data-state={cell.state}
    data-game-tip={cellTipKey(cell)}
    data-tip-beside
    onfocus={() => onfocuscell?.(cell.id)}
  >
    {@render face()}
  </button>
{/if}
