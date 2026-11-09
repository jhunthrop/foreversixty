<!-- web/src/components/ui/Steps.svelte -->
<!-- A numbered checklist with an action per step (setup spec 2026-11-07 section 3A): Checklist's
     open Rows with a 26 px numbered mark and the step's one action in the trail. Marks: `ok` done
     (kill, a check), `next` the first step not done (gold, its number; its row carries the page's
     one primary Button), `todo` not done (a ring, the number muted), `warn` needs the player (gold
     "!"), `miss` failed (wipe "×"), `busy` running (a ring with a gold arc turning once a second).
     The rows are not focusable; Tab reaches each action in order. The Plate around it names the
     list ("Your setup"). -->
<script lang="ts" module>
  export type StepsMark = 'ok' | 'next' | 'todo' | 'warn' | 'miss' | 'busy';

  /** The mark's accessible word: "Step 2: Next". */
  export const STEP_MARK_WORD: Readonly<Record<StepsMark, string>> = {
    ok: 'Done',
    next: 'Next',
    todo: 'Not done',
    warn: 'To do',
    miss: 'Failed',
    busy: 'In progress',
  };
</script>

<script lang="ts">
  import type { Snippet } from 'svelte';
  import Row from './Row.svelte';

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

  const GLYPH: Partial<Record<StepsMark, string>> = { ok: '✓', warn: '!', miss: '×', busy: '' };
</script>

<div class="ui-steps" data-testid={testid}>
  {#each items as item (item.key)}
    <div class="ui-step" data-mark={item.mark}>
      <Row title={item.title} testid={testid === undefined ? undefined : `${testid}-${item.key}`}>
        {#snippet lead()}
          <span
            class="ui-step-mark"
            data-mark={item.mark}
            role="img"
            aria-label={`Step ${item.n}: ${STEP_MARK_WORD[item.mark]}`}>{GLYPH[item.mark] ?? item.n}</span
          >
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
  .ui-step-mark {
    position: relative;
    display: inline-flex;
    flex: none;
    align-items: center;
    justify-content: center;
    width: 26px;
    height: 26px;
    border-radius: 50%;
    font:
      700 13px/26px var(--font-body),
      sans-serif;
    font-variant-numeric: tabular-nums;
  }
  .ui-step-mark[data-mark='ok'] {
    background: var(--color-kill);
    color: var(--color-bg);
  }
  .ui-step-mark[data-mark='next'],
  .ui-step-mark[data-mark='warn'] {
    background: var(--color-gold);
    color: var(--color-bg);
  }
  .ui-step-mark[data-mark='miss'] {
    background: var(--color-wipe);
    color: var(--color-bg);
  }
  .ui-step-mark[data-mark='todo'] {
    box-shadow: inset 0 0 0 1.5px var(--color-line);
    color: var(--color-muted);
  }
  .ui-step-mark[data-mark='busy'] {
    box-shadow: inset 0 0 0 2px var(--color-line);
  }
  .ui-step-mark[data-mark='busy']::after {
    content: '';
    position: absolute;
    inset: 0;
    border: 2px solid transparent;
    border-top-color: var(--color-gold);
    border-radius: 50%;
    animation: ui-step-spin 1s linear infinite;
  }
  @keyframes ui-step-spin {
    to {
      transform: rotate(360deg);
    }
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
