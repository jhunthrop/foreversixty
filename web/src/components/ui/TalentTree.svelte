<script lang="ts">
  import { onMount } from 'svelte';
  import {
    gridKeyAction,
    gridRows,
    moveCell,
    tabStopId,
    type TalentTreeView,
    type TalentTreesCopy,
  } from '../../lib/ui/talent-tree';
  import TalentButton from './TalentButton.svelte';

  let {
    tree,
    build,
    editable = false,
    copy,
    onadd,
    onremove,
    onreset,
  }: {
    tree: TalentTreeView;
    build: string;
    editable?: boolean;
    /** `editable` only: the per-tree Reset's words. */
    copy?: TalentTreesCopy;
    onadd?: (id: number) => void;
    onremove?: (id: number) => void;
    /** The per-tree Reset in the head (shown when the tree has points). */
    onreset?: () => void;
  } = $props();

  // The cell that holds the tree's one tab stop. Until the island runs (a static page: a guide's At
  // a glance) no script moves focus, so every talent is its own tab stop.
  let lastFocused = $state<number | undefined>(undefined);
  let roving = $state(false);
  onMount(() => {
    roving = true;
  });
  const stop = $derived(tabStopId(tree, lastFocused));
  // Art that failed to load: the panel's own ground shows instead of a broken image.
  let artBroken = $state(false);

  function onGridKey(event: KeyboardEvent): void {
    const action = gridKeyAction(event.key);
    if (action === null || stop === undefined) return;
    if (action.kind === 'remove') {
      if (!editable) return;
      event.preventDefault();
      onremove?.(stop);
      return;
    }
    event.preventDefault();
    const target = moveCell(tree.cells, stop, action.dx, action.dy);
    lastFocused = target;
    (event.currentTarget as HTMLElement)
      .querySelector<HTMLElement>(`[data-testid="talent-${target}"]`)
      ?.focus();
  }
</script>

<!-- web/src/components/ui/TalentTree.svelte -->
<!-- One talent tree, the same everywhere (talent-tree lane 2026-11-10; planner spec 2026-11-08
     section 3A): the head (spec icon, name, points in gold, muted at 0; the per-tree Reset when
     editable and spent), then the client's art at 50% behind the grid, the prerequisite arrows (gold
     when met) on a --bg halo, and tiers × columns of gridcells holding TalentButtons. The grid owns
     the keyboard: one tab stop per tree on its last-focused cell, arrows move, Backspace or Delete
     removes when editable (Enter is the button's own click). Sizes come from the caller's tokens
     (components/ui/talent-tree.css, `data-size` on TalentTrees). -->

<div class="tt">
  <div class="tt-h">
    {#if tree.iconSrc !== ''}<img src={tree.iconSrc} alt="" width="28" height="28" />{/if}
    <h3>{tree.name}</h3>
    <b class:zero={tree.points === 0} data-testid={`tree-points-${tree.id}`}>{tree.points}</b>
    {#if editable && tree.points > 0 && onreset && copy}
      <button
        type="button"
        class="tt-reset"
        aria-label={copy.resetTreeLabel(tree.name)}
        data-testid={`tree-reset-${tree.id}`}
        onclick={onreset}>{copy.resetTree}</button
      >
    {/if}
  </div>
  <div
    class="tt-g"
    role="grid"
    tabindex={-1}
    aria-label={`${tree.name} talents`}
    data-testid={`tree-${tree.id}`}
    style:--cols={tree.columns}
    onkeydown={onGridKey}
  >
    {#if tree.artSrc && !artBroken}
      <img
        class="tt-art"
        src={tree.artSrc}
        alt=""
        aria-hidden="true"
        loading="eager"
        decoding="async"
        data-testid={`tree-art-${tree.id}`}
        onerror={() => (artBroken = true)}
      />
    {/if}
    <svg class="tt-arrows" viewBox={tree.viewBox} preserveAspectRatio="none" aria-hidden="true">
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
                {build}
                {editable}
                focusable={!roving || cell.id === stop}
                {onadd}
                {onremove}
                onfocuscell={(id) => (lastFocused = id)}
              />{/if}</span
          >
        {/each}
      </div>
    {/each}
  </div>
</div>
