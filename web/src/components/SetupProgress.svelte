<!-- web/src/components/SetupProgress.svelte -->
<!-- The three progress chips at the top of /setup: "1 Sign in", "2 Addon", "3 Companion", each an
     anchor to its step, gold with a check once the account shows the step done and a ring with its
     number until then (UX review r5, finding 2; the done wording is the player review's N11).
     Server-rendered, it reads the same <html> hints the steps use (global.css), so a signed-in
     visitor sees step 1 done before any script runs; once the session and device reads answer the
     chips follow them. -->
<script lang="ts">
  import { addonCopy } from '../lib/addon/copy';
  import { chipsDone, type SetupState } from '../lib/setup/progress';
  import { loadSetupState } from '../lib/setup/state';

  let state = $state<SetupState | null>(null);

  $effect(() => {
    let live = true;
    void loadSetupState().then((answer) => {
      if (live) state = answer;
    });
    return () => {
      live = false;
    };
  });

  const done = $derived(state === null ? null : chipsDone(state));
  /** Which <html> hint a chip follows until the reads answer: the session cookie, the stored character, or none. */
  const HINTS = ['session', 'pointer', 'none'] as const;
</script>

<ol
  class="setup-chips flex flex-wrap gap-2"
  data-testid="setup-status"
  aria-label={addonCopy.setupProgressLabel}
>
  {#each addonCopy.setupChips as chip, index (chip.step)}
    <li
      class="setup-chip"
      data-step={chip.step}
      data-hint={HINTS[index]}
      data-done={done === null ? undefined : String(done[index])}
    >
      <a
        class="setup-chip-link border-line-warm text-muted inline-flex min-h-11 items-center gap-2 rounded-full border px-3 text-[13px] font-semibold"
        href={chip.href}
      >
        <span class="setup-chip-mark" aria-hidden="true">
          <span
            class="chip-todo border-line-warm text-muted inline-flex size-5 items-center justify-center rounded-full border text-[11px]"
            >{chip.step}</span
          >
          <span class="chip-done bg-gold text-bg inline-flex size-5 items-center justify-center rounded-full">
            <svg
              width="12"
              height="12"
              viewBox="0 0 24 24"
              fill="none"
              stroke="currentColor"
              stroke-width="3.5"
              stroke-linecap="round"
              stroke-linejoin="round"><path d="M5 12.5l4.5 4.5L19 7.5" /></svg
            >
          </span>
        </span>
        <span class="chip-todo">{chip.todo}</span>
        <span class="chip-done text-strong">{chip.done}</span>
      </a>
    </li>
  {/each}
</ol>
