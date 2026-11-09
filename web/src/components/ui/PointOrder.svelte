<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  import { tooltipLeftPx } from '../../lib/items/tooltip-position';
  import {
    drawsTalent,
    holdsPoint,
    ORDER_MARKERS,
    orderIndices,
    type PointOrderRow,
  } from '../../lib/ui/point-order';
  import { isPlaceholderIcon, talentTipSide, talentTipSideLeft } from '../../lib/ui/talent-editor';
  import TalentTip from './TalentTip.svelte';

  let {
    rows,
    capAfter = -1,
    capText = '',
    testid,
  }: {
    rows: readonly PointOrderRow[];
    /** The row index the character's level line follows; -1 for none. */
    capAfter?: number;
    capText?: string;
    testid?: string;
  } = $props();

  const indices = $derived(orderIndices(rows));
  // Order indices whose icon failed or is the pipeline's 1×1 placeholder: the dark socket.
  let broken = $state<ReadonlySet<number>>(new Set());

  function markBroken(index: number): void {
    broken = new Set([...broken, index]);
  }

  // The open talent tooltip: its socket (`r-s`) and where it sits, beside the socket when there is
  // room (as the client's), else under it inside the viewport.
  const TIP_WIDTH = 300;
  let tip = $state<{ key: string; left: number; beside: boolean } | null>(null);
  function openTip(key: string, event: Event): void {
    const box = (event.currentTarget as HTMLElement).getBoundingClientRect();
    const width = document.documentElement.clientWidth;
    const side = talentTipSide(box.left, box.right, TIP_WIDTH, width);
    tip =
      side === null
        ? { key, left: tooltipLeftPx('start', box.left, box.width, width, TIP_WIDTH), beside: false }
        : { key, left: talentTipSideLeft(side, box.width, TIP_WIDTH), beside: true };
  }
  function closeTip(key: string): void {
    if (tip?.key === key) tip = null;
  }
</script>

<!-- web/src/components/ui/PointOrder.svelte -->
<!-- The build's spend order (planner spec 2026-11-08 section 3A, 4.5; the mock's PointOrder): a
     faint 0 to 9 marker over the first row (the last digit of each socket's level), then rows of
     ten levels, a 120 px label column and ten 36 px sockets (30 below 768). A spent point is its
     icon with the rank it reached, a tab stop with its name; with a `tip`, the game's talent
     tooltip opens beside it on hover and focus. A planned one (past the character's level) dims;
     an `ahead` one (the guide's point past the build) dims further on a dashed ring; an empty
     level is dashed --line, the next point dashed gold. The gold cap line follows row `capAfter`
     when `capText` is given. Nothing scrolls sideways. -->

<div class="porder" data-testid={testid}>
  <div class="porder-m" aria-hidden="true">
    <span></span>
    <div class="porder-s">
      {#each ORDER_MARKERS as digit (digit)}<span>{digit}</span>{/each}
    </div>
  </div>
  {#each rows as row, r (r)}
    <div class="porder-r">
      <span class="lbl">{row.label}</span>
      <div class="porder-s" role="list">
        {#each row.slots as slot, s (s)}
          {@const index = indices[r][s]}
          {#if drawsTalent(slot)}
            {@const key = `${r}-${s}`}
            {@const iconKey = index ?? 1000 + r * 10 + s}
            <!-- Each point is a tab stop so its name (the talent and the level) and its tooltip reach the keyboard (spec 5.4). -->
            <!-- svelte-ignore a11y_no_noninteractive_tabindex -->
            <span
              role="listitem"
              class={slot.kind === 'planned' ? 'pl' : slot.kind === 'ahead' ? 'ah' : undefined}
              tabindex="0"
              aria-label={slot.name}
              title={slot.tip === undefined ? slot.name : undefined}
              data-kind={slot.kind}
              data-testid={holdsPoint(slot) ? `order-point-${index}` : `order-ahead-${r * 10 + s}`}
              onmouseenter={(event) => slot.tip && openTip(key, event)}
              onmouseleave={() => closeTip(key)}
              onfocus={(event) => slot.tip && openTip(key, event)}
              onblur={() => closeTip(key)}
            >
              {#if slot.iconSrc && !broken.has(iconKey)}
                <img
                  src={slot.iconSrc}
                  alt=""
                  width="36"
                  height="36"
                  decoding="async"
                  onerror={() => markBroken(iconKey)}
                  onload={(event) => {
                    if (isPlaceholderIcon(event.currentTarget)) markBroken(iconKey);
                  }}
                />
              {:else}
                <span class="ic" aria-hidden="true"></span>
              {/if}
              {#if slot.rank !== undefined}<b>{slot.rank}</b>{/if}
              {#if slot.tip && tip?.key === key}
                <span class="po-tip" class:beside={tip.beside} style:left={`${tip.left}px`}>
                  <TalentTip name={slot.tip.name} rankLine={slot.tip.rankLine} now={slot.tip.now} />
                </span>
              {/if}
            </span>
          {:else}
            <span role="listitem" aria-label={slot.name} title={slot.name} data-kind={slot.kind}
              ><span class={['em', slot.kind === 'next' && 'nx']}></span></span
            >
          {/if}
        {/each}
      </div>
    </div>
    {#if r === capAfter && capText}<p class="porder-cap">{capText}</p>{/if}
  {/each}
</div>

<style>
  .porder {
    display: flex;
    flex-direction: column;
  }
  .porder-m {
    display: grid;
    grid-template-columns: 120px minmax(0, 1fr);
    column-gap: 16px;
    padding-top: 4px;
  }
  .porder-m .porder-s > span {
    height: auto;
    text-align: center;
    font-size: 11px;
    line-height: 14px;
    color: var(--color-rule);
    font-variant-numeric: tabular-nums;
  }
  .porder-r {
    display: grid;
    grid-template-columns: 120px minmax(0, 1fr);
    align-items: center;
    column-gap: 16px;
    padding: 8px 0;
    border-bottom: 1px solid var(--color-line-soft);
  }
  .porder-r:last-child {
    border-bottom: 0;
  }
  .lbl {
    font-size: 13px;
    line-height: 20px;
    color: var(--color-muted);
    font-variant-numeric: tabular-nums;
  }
  .porder-s {
    display: grid;
    grid-template-columns: repeat(10, 36px);
    gap: 6px;
  }
  .porder-s > span {
    position: relative;
    width: 36px;
    height: 36px;
    border-radius: 5px;
  }
  .porder-s > span:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .porder-s img,
  .porder-s .ic {
    display: block;
    width: 36px;
    height: 36px;
    border-radius: 5px;
  }
  .porder-s .ic {
    background: var(--color-bg);
  }
  .em {
    display: block;
    width: 36px;
    height: 36px;
    border-radius: 5px;
    border: 1px dashed var(--color-line);
  }
  .em.nx {
    border-color: var(--color-gold);
  }
  .pl img,
  .pl .ic {
    opacity: 0.55;
  }
  .ah img,
  .ah .ic {
    opacity: 0.4;
  }
  .ah::after {
    content: '';
    position: absolute;
    inset: 0;
    border-radius: 5px;
    border: 1px dashed var(--color-line);
  }
  .ah b {
    color: var(--color-muted);
    opacity: 0.7;
  }
  .po-tip {
    position: absolute;
    top: calc(100% + 8px);
    z-index: 30;
    width: 300px;
    text-align: left;
  }
  .po-tip.beside {
    top: 0;
  }
  .pl::after {
    content: '';
    position: absolute;
    inset: 0;
    border-radius: 5px;
    box-shadow: inset 0 0 0 1px var(--color-line);
  }
  b {
    position: absolute;
    right: -4px;
    bottom: -5px;
    padding: 0 3px;
    border-radius: 999px;
    background: var(--color-bg);
    font: 700 12px/14px var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--color-muted);
  }
  .porder-cap {
    display: flex;
    align-items: center;
    gap: 10px;
    margin: 0;
    padding: 10px 0;
    font-size: 13px;
    line-height: 18px;
    color: var(--color-gold);
    border-bottom: 1px solid var(--color-line-soft);
  }
  .porder-cap::before {
    content: '';
    flex: none;
    width: 10px;
    height: 2px;
    background: var(--color-gold);
  }
  @media (max-width: 767px) {
    .porder-m {
      grid-template-columns: minmax(0, 1fr);
    }
    .porder-m > span:first-child {
      display: none;
    }
    .porder-r {
      grid-template-columns: minmax(0, 1fr);
      row-gap: 6px;
    }
    .porder-s {
      grid-template-columns: repeat(10, minmax(0, 30px));
      gap: 4px;
    }
    /* 30 px sockets, shrinking with the column below 336 px so a row never runs past it. */
    .porder-s > span,
    .porder-s img,
    .porder-s .ic,
    .em {
      width: 100%;
      max-width: 30px;
      height: auto;
      aspect-ratio: 1;
    }
    .porder-m .porder-s > span {
      height: auto;
      aspect-ratio: auto;
    }
  }
</style>
