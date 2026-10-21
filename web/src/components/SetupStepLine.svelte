<!-- web/src/components/SetupStepLine.svelte -->
<!-- One step's line on /setup, under its action: signed in, "Signed in as hunthrop#1894"; the addon,
     "Bow Jackzon arrived from the addon 1 hour ago" or what to do next; the companion, the device and
     when it was last seen, or "No companion signed in yet". Signed out it draws nothing. Its row is
     reserved from the first paint for a hinted (signed-in) visitor, so the answer never moves the
     page. -->
<script lang="ts">
  import {
    addonLine,
    companionLine,
    signInLine,
    type SetupState,
    type StepLine,
  } from '../lib/setup/progress';
  import { loadSetupState } from '../lib/setup/state';

  let { step }: { step: 'signin' | 'addon' | 'companion' } = $props();

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

  const line = $derived<StepLine | null>(
    state === null
      ? null
      : step === 'signin'
        ? signInLine(state.me)
        : step === 'addon'
          ? addonLine(state.me)
          : companionLine(state.me, state.devices),
  );
</script>

<p
  class="setup-when-in flex min-h-[22px] items-center gap-2 text-[14px]"
  class:text-muted={line !== null && !line.done}
  data-testid={`setup-line-${step}`}
  data-done={line === null ? undefined : String(line.done)}
>
  {#if line !== null}
    {#if line.done}
      <svg
        class="text-gold shrink-0"
        width="14"
        height="14"
        viewBox="0 0 24 24"
        fill="none"
        stroke="currentColor"
        stroke-width="3"
        stroke-linecap="round"
        stroke-linejoin="round"
        aria-hidden="true"><path d="M5 12.5l4.5 4.5L19 7.5" /></svg
      >
    {/if}
    <span>{line.text}</span>
  {/if}
</p>
