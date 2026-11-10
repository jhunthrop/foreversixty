<!-- web/src/components/ui/Steps.svelte -->
<!-- A numbered checklist with an action per step (setup spec 2026-11-07 section 3A): Checklist's
     open Rows with a 26 px numbered mark and the step's one action in the trail. Marks: `ok` done
     (kill, a check), `next` the first step not done (gold, its number; its row carries the page's
     one primary Button), `todo` not done (a ring, the number muted), `warn` needs the player (gold
     "!"), `miss` failed (wipe "×"), `busy` running (a ring with a gold arc turning once a second),
     `skip` passed over on purpose (a muted dash in a ring; guild setup 2026-11-09).
     The rows are not focusable; Tab reaches each action in order. The Plate around it names the
     list ("Your setup"). -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import Row from './Row.svelte';
  import StepMark, { type StepsMark } from './StepMark.svelte';

  interface StepItem {
    key: string;
    n: number;
    mark: StepsMark;
    title: string;
  }

  let {
    items,
    sub,
    action,
    testid,
  }: {
    items: readonly StepItem[];
    /** The step's lines under its title. */
    sub: Snippet<[StepItem]>;
    /** The step's one action (nothing for a done step). */
    action?: Snippet<[StepItem]>;
    testid?: string;
  } = $props();
</script>

<div class="ui-steps" data-testid={testid}>
  {#each items as item (item.key)}
    <div class="ui-step" data-mark={item.mark}>
      <Row title={item.title} testid={testid === undefined ? undefined : `${testid}-${item.key}`}>
        {#snippet lead()}
          <StepMark n={item.n} mark={item.mark} />
        {/snippet}
        {#snippet subContent()}{@render sub(item)}{/snippet}
        {#snippet trail()}{#if action}{@render action(item)}{/if}{/snippet}
      </Row>
    </div>
  {/each}
</div>

<style>
  .ui-step :global(.ui-row) {
    min-height: 64px;
    padding-block: 12px;
  }
  /* Each step is its own Row's only child, so the hairline between steps is set here. */
  .ui-step:not(:last-child) :global(.ui-row) {
    border-bottom-width: 1px;
  }
  .ui-step :global(.ui-row > span:nth-child(2) > span + span) {
    margin-top: 2px;
    font-size: 14px;
    line-height: 20px;
  }
  @media (max-width: 599px) {
    .ui-step :global(.ui-row) {
      align-items: start;
      grid-template-columns: auto minmax(0, 1fr);
    }
    /* A done step has no action: its empty trail takes no line of its own. */
    .ui-step :global(.ui-row > span:nth-child(3):not(:empty)) {
      grid-column: 1 / -1;
      justify-self: start;
      margin-top: 8px;
    }
  }
</style>
