<!-- web/src/components/CopyForGame.svelte -->
<!-- Task L8.1: the one "Copy for the game" button. Every surface that holds a build or a sim
     result -- the planner's In the game section, Top Gear's results, a saved sim -- renders this
     with the FSB1 code `buildCodeFor` (lib/addon/build-code.ts) made for it, so the label, the
     two-second confirmation and the failure fallback are the same everywhere.

     The system Button (`primary` where it is the page's one primary: the planner signed out,
     spec 2026-11-08 ruling Q6). `code` empty disables it: the page has no build it can encode
     yet. When the clipboard refuses, the code is shown in a read-only CodeField to select by
     hand; the field follows `code`, so it never offers a stale build. -->
<script lang="ts">
  import { onDestroy } from 'svelte';
  import { addonCopy } from '../lib/addon/copy';
  import { copyForGameLabel, createCopyForGame, type CopyForGameState } from '../lib/addon/copy-for-game';
  import Button from './ui/Button.svelte';
  import CodeField from './ui/CodeField.svelte';

  let {
    code,
    primary = false,
    testid = 'copy-for-game',
  }: { code: string; primary?: boolean; testid?: string } = $props();

  let copyState = $state<CopyForGameState>('idle');
  const copier = createCopyForGame((next) => (copyState = next));
  onDestroy(copier.dispose);
</script>

<div class="flex flex-col items-start gap-[10px]">
  <Button
    tone={primary ? 'primary' : 'secondary'}
    disabled={code === ''}
    onclick={() => void copier.copy(code)}
    {testid}>{copyForGameLabel(copyState)}</Button
  >
  {#if copyState === 'failed'}
    <p class="text-small text-muted" role="status" data-testid={`${testid}-failed`}>{addonCopy.copyFailed}</p>
    <CodeField value={code} label={addonCopy.buildCodeLabel} short readonly testid={`${testid}-field`} />
  {/if}
</div>
