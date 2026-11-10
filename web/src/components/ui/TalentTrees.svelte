<!-- web/src/components/ui/TalentTrees.svelte -->
<!-- A build's three talent trees side by side at the compact cell (guides spec 2026-11-08 §3A):
     the cell sizes itself from the width it is given (a container query: the three trees fill it,
     from 18 to 32 px a cell), so the trees never run out of a lane or a plate column at any page
     width; under the 18 px floor the trees stack. The
     planner's read-only semantics: a grid per tree, a row per tier, a gridcell per column, a group
     per talent; a learned talent carries its rank badge and a gold (full) or green (partial) ring,
     an unlearned one is greyed. Each talent opens the game tooltip (`data-game-tip`) when the
     page's tooltip host carries it. Static: the server renders every cell. -->
<script lang="ts" module>
  /** One talent cell: where it sits, its icon, the rank shown and its tooltip key. */
  export interface TalentCell {
    tier: number;
    column: number;
    name: string;
    icon: string;
    rank: number;
    maxRank: number;
    tipKey?: string;
  }
  export interface TalentTree {
    name: string;
    points: number;
    talents: readonly TalentCell[];
  }
</script>

<script lang="ts">
  import { dataUrl } from '../../lib/planner/load';

  let {
    trees,
    build,
    tiers = 7,
    columns = 4,
    testid,
  }: {
    trees: readonly TalentTree[];
    build: string;
    tiers?: number;
    columns?: number;
    testid?: string;
  } = $props();

  const cellAt = (tree: TalentTree, tier: number, column: number): TalentCell | undefined =>
    tree.talents.find((talent) => talent.tier === tier && talent.column === column);
  const state = (cell: TalentCell): 'full' | 'part' | 'off' =>
    cell.rank === 0 ? 'off' : cell.rank >= cell.maxRank ? 'full' : 'part';
</script>

<div class="ui-trees-box" style:--cols={columns}>
  <div class="ui-trees" data-testid={testid}>
    {#each trees as tree (tree.name)}
      <div class="ui-tree">
        <h3>{tree.name} <b class="tabular-nums">{tree.points}</b></h3>
        <div class="ui-tree-grid" role="grid" aria-label={`${tree.name} talents`}>
          {#each { length: tiers }, tier (tier)}
            <div role="row">
              {#each { length: columns }, column (column)}
                {@const cell = cellAt(tree, tier, column)}
                <span role="gridcell"
                  >{#if cell}<span
                      class="ui-tree-cell"
                      data-state={state(cell)}
                      role="group"
                      aria-label={`${cell.name} ${cell.rank}/${cell.maxRank}`}
                      tabindex={cell.tipKey ? 0 : undefined}
                      data-game-tip={cell.tipKey}
                      ><img
                        src={dataUrl(build, `icons/${cell.icon}.webp`)}
                        alt=""
                        width="24"
                        height="24"
                      />{#if cell.rank > 0}<b>{cell.rank}/{cell.maxRank}</b>{/if}</span
                    >{/if}</span
                >
              {/each}
            </div>
          {/each}
        </div>
      </div>
    {/each}
  </div>
</div>

<style>
  /* The trees fill their box: three trees of --cols cells, --cell-gap between cells and
     --tree-gap between trees, the cell what is left over, kept between the floor and the cap. */
  .ui-trees-box {
    --cell-gap: 4px;
    --tree-gap: 16px;
    --cell-min: 18px;
    --cell-max: 32px;
    container-type: inline-size;
  }
  .ui-trees {
    --cell: clamp(
      var(--cell-min),
      calc(
        (100cqw - 2 * var(--tree-gap) - 3 * (var(--cols) - 1) * var(--cell-gap)) / (3 * var(--cols)) - 0.5px
      ),
      var(--cell-max)
    );
    /* Each tree is as wide as its cells, never its name: a long name ("Marksmanship 20") may run
       into the gap, and the trees still fit the box. */
    display: grid;
    gap: var(--tree-gap);
    grid-template-columns: repeat(3, calc(var(--cols) * var(--cell) + (var(--cols) - 1) * var(--cell-gap)));
  }
  .ui-tree h3 {
    display: flex;
    gap: 8px;
    align-items: baseline;
    margin: 0;
    padding-bottom: 6px;
    font-size: 13px;
    line-height: 18px;
    font-weight: 700;
    color: var(--color-strong);
    white-space: nowrap;
  }
  .ui-tree h3 b {
    color: var(--color-gold);
  }
  .ui-tree-grid {
    display: grid;
    grid-template-columns: repeat(var(--cols), var(--cell));
    gap: var(--cell-gap);
  }
  .ui-tree-grid > [role='row'] {
    display: contents;
  }
  .ui-tree-cell {
    position: relative;
    display: block;
    width: var(--cell);
    height: var(--cell);
    cursor: pointer;
  }
  .ui-tree-cell img {
    display: block;
    width: 100%;
    height: 100%;
    border-radius: 3px;
  }
  .ui-tree-cell[data-state='off'] img {
    filter: grayscale(1) brightness(0.45);
  }
  .ui-tree-cell[data-state='full'] img {
    box-shadow: 0 0 0 2px var(--color-gold);
  }
  .ui-tree-cell[data-state='part'] img {
    box-shadow: 0 0 0 2px var(--color-kill);
  }
  .ui-tree-cell b {
    position: absolute;
    right: -3px;
    bottom: -4px;
    font-size: 10px;
    line-height: 12px;
    font-weight: 700;
    color: var(--color-strong);
    text-shadow:
      0 0 3px #000,
      0 0 3px #000;
  }
  .ui-tree-cell:focus-visible {
    outline: 2px solid var(--color-gold);
    outline-offset: 2px;
  }
  @media (max-width: 767px) {
    /* A tree's name may take two lines in a 97 px column ("Marksmanship 11"): the three heads
       share one height so the grids start level. */
    .ui-tree h3 {
      white-space: normal;
      flex-wrap: wrap;
      column-gap: 6px;
      row-gap: 0;
      min-height: 42px;
      align-content: flex-end;
    }
    .ui-trees-box {
      --cell-gap: 3px;
      --tree-gap: 10px;
    }
  }
  /* Under the floor (three 18 px trees no longer fit) the trees stack, one under the other. */
  @container (max-width: 270px) {
    .ui-trees {
      grid-template-columns: calc(var(--cols) * var(--cell) + (var(--cols) - 1) * var(--cell-gap));
    }
  }
</style>
