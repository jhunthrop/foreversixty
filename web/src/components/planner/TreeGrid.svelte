<!-- web/src/components/planner/TreeGrid.svelte -->
<!-- One tree as a tier-by-column grid. Owns the roving tabindex: exactly one cell in the
     tree is tabbable and arrow keys move it. Enter adds a point through the cell's own
     native button activation; Backspace and Delete remove one, handled here. -->
<script lang="ts">
  import type { BandDiffView } from '../../lib/planner/band-compare';
  import { connectorsFor, gridViewBox } from '../../lib/planner/connectors';
  import { bandCompareCopy } from '../../lib/planner/copy';
  import { gridCells, gridSize, moveFocus, type GridCell } from '../../lib/planner/grid';
  import { dataUrl } from '../../lib/planner/load';
  import { CONNECTOR_STROKE } from '../../lib/planner/styles';
  import type { PlannerStore } from '../../lib/planner/store.svelte';
  import type { TalentTree } from '../../lib/planner/types';
  import TalentCell from './TalentCell.svelte';

  let {
    store,
    tree,
    readOnly = false,
    bandDiff = null,
  }: {
    store: PlannerStore;
    tree: TalentTree;
    readOnly?: boolean;
    /** Rebuild spec §4.D/§4.E.4: `BandCompare`'s own diff, once a band has loaded, so this
     *  grid can mark every cell whose rank differs from it. `null` before any band has
     *  loaded -- no cell ever claims a difference without real data (tenet 8). */
    bandDiff?: BandDiffView | null;
  } = $props();

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

<!-- One box, sized and centred on the grid itself, not the panel's full width (build review
     round 1, finding 1: art that fills a wider panel than the grid needs crops down to a
     tiny cluster of icons in a sea of background). `w-fit` on this shared box is what makes
     the panel's own height the grid's height too: the art is `absolute inset-0` inside it,
     so it never has a size of its own to grow the box past what the grid (a normal, in-flow
     child) already gives it, and it renders cropped to the grid's own bounds rather than the
     panel's. preserveAspectRatio="none" lets one viewBox in base-breakpoint units cover the
     slightly larger md and lg cells, and non-scaling-stroke keeps the line weight identical
     at every size. -->
<div class="relative mx-auto w-fit">
  {#if tree.background && !artBroken}
    <!-- The client's own panel, cropped and processed in the pipeline. `object-cover`
         fills the whole box whatever its aspect and crops the excess evenly around the
         centre, so the art reads as the grid's own backdrop rather than a strip beside it.

         `loading="eager"` (fix round 3, ux-designer review): this art, like the talent
         icons below it (`TalentCell.svelte`'s own identical comment), is this page's own
         content, in the first viewport at every desktop width this lane supports (all
         three trees render side by side from `md`) -- `lazy` left it undecoded at capture
         time in three of the acceptance screenshots, a real, visible bug, not just a
         screenshot artifact. Three small webps is a trivial cost next to that. -->
    <img
      src={backgroundSrc}
      alt=""
      aria-hidden="true"
      width="300"
      height="331"
      loading="eager"
      decoding="async"
      data-testid={`tree-art-${tree.id}`}
      class="rounded-control pointer-events-none absolute inset-0 h-full w-full object-cover object-center"
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
      <!-- Rebuild spec §4.E.2: a dark halo under the coloured line, sharing the same `d` --
         two stacked paths, not a new shape, so the halo moves and sizes with the line it
         sits under for free. -->
      <path
        d={connector.d}
        aria-hidden="true"
        fill="none"
        stroke-width="5"
        stroke-linecap="round"
        style="vector-effect: non-scaling-stroke"
        class={CONNECTOR_STROKE.halo}
      />
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
    onkeydown={readOnly ? undefined : onKeyDown}
  >
    {#each Array.from({ length: size.tiers }, (_, tier) => tier) as tier (tier)}
      <!-- `contents` generates no box, so the row satisfies grid's required children without
         taking part in the layout: the cells stay direct items of the CSS grid above. -->
      <div role="row" class="contents">
        {#each Array.from({ length: size.columns }, (_, column) => column) as column (column)}
          {@const cell = cells.find((c) => c.tier === tier && c.column === column)}
          <!-- The grid is the panel's dominant element (build review round 1, finding 1 and
             4): h-11/w-11 through md's h-12/w-12 grows again to h-14/w-14 from lg, where the
             extra desktop width this lane also gave the panel (finding 2) would otherwise
             leave the same small cluster of icons, just inside a wider box. -->
          <div role="gridcell" class="flex h-11 w-11 md:h-12 md:w-12 lg:h-14 lg:w-14">
            {#if cell}
              <TalentCell
                {store}
                talent={cell.talent}
                focused={focused?.talent.id === cell.talent.id}
                onfocuscell={() => (focusedId = cell.talent.id)}
                {readOnly}
                differs={bandDiff !== null && bandDiff.diffTalentIds.has(cell.talent.id)}
                bandRankLine={bandDiff !== null
                  ? bandCompareCopy.bandRankTooltipLine(
                      bandDiff.bandLabel,
                      bandDiff.bandRankById.get(cell.talent.id) ?? 0,
                      cell.talent.max_rank,
                    )
                  : null}
              />
            {/if}
          </div>
        {/each}
      </div>
    {/each}
  </div>
</div>
