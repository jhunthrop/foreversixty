<!-- web/src/components/ui/TalentEditor.svelte -->
<!-- The editable talent trees on the page's one Plate (planner spec 2026-11-08 section 3A, 4.4; the
     mock's TalentEditor). From 768 one column per tree split by hairlines; below 768 the tree tabs
     and the open tree (`active`). Each tree: a head (spec icon, name, points, the per-tree Reset
     when it has points), the client's art at 50% behind the grid, the prerequisite arrows (gold
     when met) on a --bg halo, then tiers × columns of gridcells holding TalentButtons. The grid
     owns the keyboard: one tab stop per tree on its last-focused cell, arrows move, Backspace or
     Delete removes (Enter is the button's own click). `foot` renders in `.tedit-f` (the status
     line `.st2`, `.st2.wipe`, and the right-hand `.rs` group are styled here). -->
<!-- Its styles travel with its code (css="injected"): it renders only after the island hydrates, and a
     page inlines only the styles of what its server render drew. -->
<svelte:options css="injected" />

<script lang="ts">
  import type { Snippet } from 'svelte';
  import { radioKeyTarget } from '../../lib/ui/control-bar';
  import {
    gridKeyAction,
    gridRows,
    moveCell,
    tabStopId,
    type EditorTree,
    type TalentEditorCopy,
  } from '../../lib/ui/talent-editor';
  import Plate from './Plate.svelte';
  import TalentButton from './TalentButton.svelte';

  let {
    trees,
    active = $bindable(0),
    label = 'Talent trees',
    copy,
    readOnly = false,
    onadd,
    onremove,
    onresettree,
    foot,
    testid,
  }: {
    trees: readonly EditorTree[];
    /** The open tree below 768. */
    active?: number;
    /** The Plate's accessible name. */
    label?: string;
    copy: TalentEditorCopy;
    readOnly?: boolean;
    onadd: (id: number) => void;
    onremove: (id: number) => void;
    /** The per-tree Reset in each tree head (shown when the tree has points and not read-only). */
    onresettree?: (treeIndex: number) => void;
    foot?: Snippet;
    testid?: string;
  } = $props();

  // The last-focused cell per tree id: where Tab re-enters that tree's grid.
  let lastFocused = $state<Readonly<Record<number, number>>>({});
  const uid = $props.id();
  // Tree ids whose art failed to load: the Plate's own ground shows instead of a broken image.
  let brokenArt = $state<ReadonlySet<number>>(new Set());

  function rememberFocus(treeId: number, cellId: number): void {
    lastFocused = { ...lastFocused, [treeId]: cellId };
  }

  function onGridKey(event: KeyboardEvent, tree: EditorTree): void {
    const action = gridKeyAction(event.key);
    if (action === null) return;
    const current = tabStopId(tree, lastFocused[tree.id]);
    if (current === undefined) return;
    if (action.kind === 'remove') {
      if (readOnly) return;
      event.preventDefault();
      onremove(current);
      return;
    }
    event.preventDefault();
    const target = moveCell(tree.cells, current, action.dx, action.dy);
    rememberFocus(tree.id, target);
    (event.currentTarget as HTMLElement)
      .querySelector<HTMLElement>(`[data-testid="talent-${target}"]`)
      ?.focus();
  }

  function onTabKey(event: KeyboardEvent, index: number): void {
    const target = radioKeyTarget(event.key, index, trees.length);
    if (target === null) return;
    event.preventDefault();
    active = target;
    const tablist = (event.currentTarget as HTMLElement).parentElement;
    tablist?.querySelector<HTMLElement>(`[data-testid="tree-tab-${target}"]`)?.focus();
  }
</script>

<Plate {label} {testid}>
  <div class="tedit">
    <div class="tedit-tabs" role="tablist" aria-label={label}>
      {#each trees as tree, index (tree.id)}
        <button
          type="button"
          role="tab"
          id={`${uid}-tab-${index}`}
          aria-selected={index === active}
          aria-controls={`${uid}-panel-${index}`}
          tabindex={index === active ? 0 : -1}
          data-testid={`tree-tab-${index}`}
          onclick={() => (active = index)}
          onkeydown={(event) => onTabKey(event, index)}
        >
          <span><img src={tree.iconSrc} alt="" width="28" height="28" /><b>{tree.points}</b></span>{tree.name}
        </button>
      {/each}
    </div>
    <div class="tedit-cols" style:--n={trees.length}>
      {#each trees as tree, index (tree.id)}
        {@const stop = tabStopId(tree, lastFocused[tree.id])}
        <div
          class="tedit-tree"
          class:on={index === active}
          role="tabpanel"
          id={`${uid}-panel-${index}`}
          aria-label={tree.name}
        >
          <div class="tedit-h">
            <img src={tree.iconSrc} alt="" width="28" height="28" />
            <h2>{tree.name}</h2>
            <b class:zero={tree.points === 0} data-testid={`tree-points-${tree.id}`}>{tree.points}</b>
            {#if tree.points > 0 && !readOnly && onresettree}
              <button
                type="button"
                class="trs"
                aria-label={copy.resetTreeLabel(tree.name)}
                data-testid={`tree-reset-${tree.id}`}
                onclick={() => onresettree(index)}>{copy.resetTree}</button
              >
            {/if}
          </div>
          <div
            class="tedit-g"
            role="grid"
            tabindex={-1}
            aria-label={`${tree.name} talents`}
            data-testid={`tree-${tree.id}`}
            style:--cols={tree.columns}
            onkeydown={(event) => onGridKey(event, tree)}
          >
            {#if tree.artSrc && !brokenArt.has(tree.id)}
              <img
                class="art"
                src={tree.artSrc}
                alt=""
                aria-hidden="true"
                loading="eager"
                decoding="async"
                data-testid={`tree-art-${tree.id}`}
                onerror={() => (brokenArt = new Set([...brokenArt, tree.id]))}
              />
            {/if}
            <svg viewBox={tree.viewBox} preserveAspectRatio="none" aria-hidden="true">
              {#each tree.connectors as connector, c (c)}
                <path class="halo" d={connector.d} />
                <path class={connector.met ? 'on' : 'off'} d={connector.d} data-met={connector.met} />
              {/each}
            </svg>
            {#each gridRows(tree) as row, tier (tier)}
              <div role="row">
                {#each row as cell, column (column)}
                  <span role="gridcell"
                    >{#if cell}<TalentButton
                        {cell}
                        {copy}
                        {readOnly}
                        focused={cell.id === stop}
                        {onadd}
                        {onremove}
                        onfocuscell={(id) => rememberFocus(tree.id, id)}
                      />{/if}</span
                  >
                {/each}
              </div>
            {/each}
          </div>
        </div>
      {/each}
    </div>
    {#if foot}<div class="tedit-f">{@render foot()}</div>{/if}
  </div>
</Plate>

<style>
  .tedit {
    --cell: 44px;
    --cgap: 8px;
  }
  .tedit-tabs {
    display: flex;
    border: 1px solid var(--color-line);
    border-radius: var(--radius-field);
    overflow: hidden;
    margin: 12px 12px 0;
  }
  .tedit-tabs button {
    all: unset;
    box-sizing: border-box;
    flex: 1;
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 2px;
    min-height: 60px;
    padding: 8px 4px;
    border-left: 1px solid var(--color-line-soft);
    cursor: pointer;
    color: var(--color-nav);
    font: 600 12px/16px var(--font-body);
    text-align: center;
  }
  .tedit-tabs button:first-child {
    border-left: 0;
  }
  .tedit-tabs button:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: -2px;
  }
  .tedit-tabs button > span {
    display: flex;
    align-items: center;
    gap: 6px;
  }
  .tedit-tabs b {
    font: 700 15px/20px var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--color-strong);
  }
  .tedit-tabs button[aria-selected='true'] {
    background: var(--color-selected);
    color: var(--color-gold);
    box-shadow: inset 0 -2px 0 var(--color-gold);
  }
  .tedit-tabs button[aria-selected='true'] b {
    color: var(--color-gold);
  }
  .tedit-tabs img,
  .tedit-h img {
    display: block;
    width: 28px;
    height: 28px;
    border-radius: 4px;
    flex: none;
  }
  .tedit-cols {
    display: grid;
    grid-template-columns: minmax(0, 1fr);
  }
  .tedit-tree {
    display: flex;
    flex-direction: column;
    align-items: center;
    gap: 14px;
    padding: 14px 12px 18px;
    min-width: 0;
  }
  .tedit-h {
    align-self: stretch;
    display: flex;
    align-items: center;
    gap: 10px;
    min-width: 0;
  }
  .tedit-h h2 {
    font: 700 15px/20px var(--font-body);
    color: var(--color-strong);
    margin: 0;
    min-width: 0;
  }
  .tedit-h b {
    margin-left: auto;
    font: 700 18px/24px var(--font-body);
    font-variant-numeric: tabular-nums;
    color: var(--color-gold);
  }
  .tedit-h b.zero {
    color: var(--color-muted);
  }
  .trs {
    all: unset;
    cursor: pointer;
    font: 600 13px/18px var(--font-body);
    color: var(--color-gold);
    border-radius: 2px;
  }
  .trs:hover {
    color: var(--color-gold-hover);
  }
  .trs:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  .tedit-g {
    position: relative;
    display: grid;
    grid-template-columns: repeat(var(--cols), var(--cell));
    gap: var(--cgap);
    padding: 6px;
  }
  .tedit-g:focus {
    outline: none;
  }
  .art {
    position: absolute;
    inset: 0;
    width: 100%;
    height: 100%;
    object-fit: cover;
    border-radius: 4px;
    opacity: 0.5;
    pointer-events: none;
  }
  svg {
    position: absolute;
    inset: 6px;
    width: calc(100% - 12px);
    height: calc(100% - 12px);
    pointer-events: none;
    overflow: visible;
  }
  path {
    fill: none;
    stroke-linecap: round;
    vector-effect: non-scaling-stroke;
  }
  path.halo {
    stroke: var(--color-bg);
    stroke-width: 5;
  }
  path.on {
    stroke: var(--color-gold);
    stroke-width: 2;
  }
  path.off {
    stroke: #3a4256;
    stroke-width: 2;
  }
  [role='row'] {
    display: contents;
  }
  [role='gridcell'] {
    position: relative;
    width: var(--cell);
    height: var(--cell);
  }
  .tedit-f {
    display: flex;
    flex-wrap: wrap;
    align-items: center;
    gap: 6px 16px;
    padding: 12px 16px;
    border-top: 1px solid var(--color-line-soft);
  }
  .tedit-f :global(.st2) {
    flex-basis: 100%;
    min-height: 20px;
    margin: 0;
    font-size: 14px;
    line-height: 20px;
    color: var(--color-strong);
  }
  .tedit-f :global(.st2.wipe) {
    color: var(--color-wipe);
  }
  .tedit-f :global(.rs) {
    margin-left: auto;
    display: flex;
    gap: 10px;
    align-items: center;
    flex-wrap: wrap;
  }
  @media (max-width: 767px) {
    .tedit-tree:not(.on) {
      display: none;
    }
    .tedit-h {
      display: none;
    }
  }
  @media (min-width: 768px) {
    .tedit {
      --cell: 48px;
    }
    .tedit-tabs {
      display: none;
    }
    .tedit-cols {
      grid-template-columns: repeat(var(--n), minmax(0, 1fr));
    }
    .tedit-tree + .tedit-tree {
      border-left: 1px solid var(--color-line-soft);
    }
  }
  @media (min-width: 1280px) {
    .tedit-tree {
      padding: 16px 20px 20px;
    }
  }
  @media (min-width: 1760px) {
    .tedit {
      --cell: 56px;
      --cgap: 10px;
    }
  }
</style>
