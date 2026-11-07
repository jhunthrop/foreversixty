<!-- web/src/components/ui/TalentTrees.svelte -->
<!-- A build's three talent trees side by side at the compact cell (guides spec 2026-11-08 §3A):
     24 px cells (28 from 1760, 22 below 768), 4 px gaps, the three under 400 px wide. The
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

<style>
  .ui-trees {
    display: grid;
    gap: 16px;
    grid-template-columns: repeat(3, max-content);
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
    grid-template-columns: repeat(4, 24px);
    gap: 4px;
  }
  .ui-tree-grid > [role='row'] {
    display: contents;
  }
  .ui-tree-cell {
    position: relative;
    display: block;
    width: 24px;
    height: 24px;
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
    .ui-trees {
      grid-template-columns: repeat(3, minmax(0, max-content));
      gap: 10px;
    }
    .ui-tree-grid {
      grid-template-columns: repeat(4, 22px);
      gap: 3px;
    }
    .ui-tree-cell {
      width: 22px;
      height: 22px;
    }
  }
  @media (min-width: 1760px) {
    .ui-tree-grid {
      grid-template-columns: repeat(4, 28px);
    }
    .ui-tree-cell {
      width: 28px;
      height: 28px;
    }
  }
</style>
