<!-- web/src/components/CopyForGame.svelte -->
<!-- Task L8.1: the one "Copy for the game" button. Every surface that holds a build or a sim
     result -- the planner's share panel, Top Gear's results, a saved sim -- renders this with
     the FSB1 code `buildCodeFor` (lib/addon/build-code.ts) made for it, so the label, the
     3 s confirmation and the failure fallback are the same everywhere.

     `code` empty disables the button: the page has no build it can encode yet (talent data
     still loading, or nothing in the build has a record to encode). When the clipboard
     refuses, the code is shown in a read-only field to select by hand; the field follows
     `code`, so it never offers a stale build. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import { addonCopy } from '../lib/addon/copy';
  import { copyForGameLabel, createCopyForGame, type CopyForGameState } from '../lib/addon/copy-for-game';
  import { SECONDARY_BUTTON } from '../lib/planner/styles';

  let { code, testid = 'copy-for-game' }: { code: string; testid?: string } = $props();

  let copyState = $state<CopyForGameState>('idle');
  const copier = createCopyForGame((next) => (copyState = next));
  onDestroy(copier.dispose);
</script>

<div class="flex flex-col gap-2">
  <button
    type="button"
    class="{SECONDARY_BUTTON} border-line-warm text-nav w-fit px-4"
    data-testid={testid}
    disabled={code === ''}
    onclick={() => void copier.copy(code)}
  >
    {copyForGameLabel(copyState)}
  </button>
  {#if copyState === 'failed'}
    <p class="text-muted text-[13px]" role="status" data-testid={`${testid}-failed`}>
      {addonCopy.copyFailed}
    </p>
    <input
      type="text"
      readonly
      value={code}
      aria-label={addonCopy.buildCodeLabel}
      class="border-line-warm rounded-control bg-raised text-text h-11 w-full max-w-[420px] border px-3 font-mono text-[12px]"
      data-testid={`${testid}-field`}
      onfocus={(event) => event.currentTarget.select()}
    />
  {/if}
</div>
