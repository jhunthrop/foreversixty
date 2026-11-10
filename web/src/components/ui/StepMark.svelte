<!-- web/src/components/ui/StepMark.svelte -->
<!-- A step's 26 px round mark (setup spec 2026-11-07 section 3A), shared by Steps and by any compact
     step list (the claim page's setup list): `ok` done (kill, a check), `next` the first step not
     done (gold, its number), `todo` not done (a ring, the number muted), `warn` needs the player
     (gold "!"), `miss` failed (wipe "×"), `busy` running (a ring with a gold arc turning once a
     second), `skip` passed over on purpose (a muted dash in a ring; guild setup 2026-11-09). -->
<script lang="ts" module>
  export type StepsMark = 'ok' | 'next' | 'todo' | 'warn' | 'miss' | 'busy' | 'skip';

  /** The mark's accessible word: "Step 2: Next". */
  export const STEP_MARK_WORD: Readonly<Record<StepsMark, string>> = {
    ok: 'Done',
    next: 'Next',
    todo: 'Not done',
    warn: 'To do',
    miss: 'Failed',
    busy: 'In progress',
    skip: 'Skipped',
  };
</script>

<script lang="ts">
  let { n, mark }: { n: number; mark: StepsMark } = $props();

  const GLYPH: Partial<Record<StepsMark, string>> = { ok: '✓', warn: '!', miss: '×', busy: '', skip: '–' };
</script>

<span class="ui-step-mark" data-mark={mark} role="img" aria-label={`Step ${n}: ${STEP_MARK_WORD[mark]}`}
  >{GLYPH[mark] ?? n}</span
>

<style>
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
  .ui-step-mark[data-mark='todo'],
  .ui-step-mark[data-mark='skip'] {
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
</style>
