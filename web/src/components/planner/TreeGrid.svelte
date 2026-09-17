<!-- web/src/components/planner/TreeGrid.svelte -->
<!-- One tree as a tier-by-column grid. Owns the roving tabindex: exactly one cell in the
     tree is tabbable and arrow keys move it. Enter adds a point through the cell's own
     native button activation; Backspace and Delete remove one, handled here. -->
<script lang="ts">
  import { connectorsFor, gridViewBox } from '../../lib/planner/connectors';
  import { gridCells, gridSize, moveFocus, type GridCell } from '../../lib/planner/grid';
  import { dataUrl } from '../../lib/planner/load';
  import { CONNECTOR_STROKE } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { TalentTree } from '../../lib/planner/types';
  import TalentCell from './TalentCell.svelte';

  let { store, tree }: { store: PlannerStore; tree: TalentTree } = $props();

  const cells = $derived(gridCells(tree));
  const size = $derived(gridSize(tree));
  const connectors = $derived(connectorsFor(tree, store.ranks));
  const viewBox = $derived(gridViewBox(size.tiers, size.columns));
  const backgroundSrc = $derived(dataUrl(store.treeVersion, `trees/${tree.background}.webp`));
  // A build with no art still has to render a usable tree, so a missing image
  // leaves the panel's own background rather than a broken-image box.
  let artBroken = $state(false);
  let focusedId = $state<number | null>(null);
  const focused = $derived<GridCell | null>(
    cells.find((cell) => cell.talent.id === focusedId) ?? cells[0] ?? null,
  );

  let root: HTMLDivElement;

  function focusCell(cell: GridCell): void {
    focusedId = cell.talent.id;
    root.querySelector<HTMLButtonElement>(`[data-testid="talent-${cell.talent.id}"]`)?.focus();
  }

  const ARROWS: Record<string, [number, number]> = {
    ArrowRight: [1, 0],
    ArrowLeft: [-1, 0],
    ArrowDown: [0, 1],
    ArrowUp: [0, -1],
  };

  function onKeyDown(event: KeyboardEvent): void {
    if (!focused) return;
    const delta = ARROWS[event.key];
    if (delta) {
      event.preventDefault();
      focusCell(moveFocus(cells, focused, delta[0], delta[1]));
      return;
    }
    // Enter and Space are deliberately not handled here: the focused cell is a real
    // <button>, so the browser turns them into a click, which already adds a point.
    // Handling them here too would spend two points per press.
    if (event.key === 'Backspace' || event.key === 'Delete') {
      event.preventDefault();
      store.removePoint(focused.talent.id);
    }
  }
</script>

<!-- `w-fit` so the overlay is exactly the grid's box: the grid's columns are
     max-content, so a full-width wrapper would stretch every line sideways.
     preserveAspectRatio="none" lets one viewBox in base-breakpoint units cover
     the slightly larger md cells, and non-scaling-stroke keeps the line weight
     identical at both sizes. -->
<div class="relative w-fit">
  {#if tree.background && !artBroken}
    <!-- The client's own 320x384 panel, processed in the pipeline. The source texture
         carries a black margin of game padding along its bottom and right edges, so this
         is `object-cover` anchored top-left rather than `object-fill`: stretching would
         squash that padding into the box instead of cropping it away, and a black band
         would show. -->
    <img
      src={backgroundSrc}
      alt=""
      aria-hidden="true"
      width="320"
      height="384"
      loading="lazy"
      decoding="async"
      data-testid={`tree-art-${tree.id}`}
      class="rounded-control pointer-events-none absolute inset-0 h-full w-full object-cover object-left-top"
      onerror={() => (artBroken = true)}
    />
  {/if}
  <svg
    class="pointer-events-none absolute inset-0 h-full w-full"
    {viewBox}
    preserveAspectRatio="none"
    aria-hidden="true"
  >
    {#each connectors as connector (connector.id)}
      <path
        d={connector.d}
        data-testid={`connector-${connector.id}`}
        data-met={connector.met}
        fill="none"
        stroke-width="2"
        stroke-linecap="round"
        style="vector-effect: non-scaling-stroke"
        class={connector.met ? CONNECTOR_STROKE.met : CONNECTOR_STROKE.unmet}
      />
    {/each}
  </svg>
  <!-- The roving tabindex lives on the cells; the container's own -1 keeps it out of the tab
       order while still giving the composite widget a focus target of its own. -->
  <div
    bind:this={root}
    role="grid"
    tabindex={-1}
    aria-label={`${tree.name} talents`}
    data-testid={`tree-${tree.id}`}
    class="relative grid gap-2"
    style={`grid-template-columns: repeat(${size.columns}, minmax(0, max-content));`}
    onkeydown={onKeyDown}
  >
    {#each Array.from({ length: size.tiers }, (_, tier) => tier) as tier (tier)}
      <!-- `contents` generates no box, so the row satisfies grid's required children without
           taking part in the layout: the cells stay direct items of the CSS grid above. -->
      <div role="row" class="contents">
        {#each Array.from({ length: size.columns }, (_, column) => column) as column (column)}
          {@const cell = cells.find((c) => c.tier === tier && c.column === column)}
          <div role="gridcell" class="flex h-11 w-11 md:h-12 md:w-12">
            {#if cell}
              <TalentCell
                {store}
                talent={cell.talent}
                focused={focused?.talent.id === cell.talent.id}
                onfocuscell={() => (focusedId = cell.talent.id)}
              />
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
