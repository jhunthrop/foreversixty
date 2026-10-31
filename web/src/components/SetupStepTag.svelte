<!-- web/src/components/SetupStepTag.svelte -->
<!-- The state tag in a /setup step card's head: "Done" (success) or "To do" (muted). It reads
     the same session and device answers as the progress chips and, until they answer, the same
     <html> hints (global.css, .setup-step-tag), so a signed-in visitor never sees step 1 flip. -->
<script lang="ts">
  import { addonCopy } from '../lib/addon/copy';
  import { chipsDone, STEP_HINTS, type SetupState } from '../lib/setup/progress';
  import { loadSetupState } from '../lib/setup/state';
  import Tag from './ui/Tag.svelte';

  let { step }: { step: 1 | 2 | 3 } = $props();

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

  const done = $derived(state === null ? null : chipsDone(state)[step - 1]);
</script>

<span
  class="setup-step-tag inline-flex"
  data-testid={`setup-step-tag-${step}`}
  data-hint={STEP_HINTS[step - 1]}
  data-done={done === null ? undefined : String(done)}
>
  <span class="chip-done"><Tag tone="success">{addonCopy.setupTagDone}</Tag></span>
  <span class="chip-todo"><Tag tone="muted">{addonCopy.setupTagTodo}</Tag></span>
</span>
