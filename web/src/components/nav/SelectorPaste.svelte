<!-- web/src/components/nav/SelectorPaste.svelte -->
<!-- "Paste an export", inline (spec 3.F): the action row is replaced in place by a labelled
     box and two neutral buttons -- no gold fill. A valid export is handed to `onpaste`, which
     makes it the current character and closes the selector; an invalid one keeps the box and
     says so under it. Focus moves into the box on open. -->
<script lang="ts">
  import { selectorCopy } from '../../lib/character-selector/copy';

  let { onpaste, oncancel }: { onpaste: (input: string) => Promise<boolean>; oncancel: () => void } =
    $props();

  let text = $state('');
  let invalid = $state(false);
  let working = $state(false);
  let box: HTMLTextAreaElement | undefined = $state();

  $effect(() => {
    box?.focus();
  });

  async function submit(): Promise<void> {
    working = true;
    invalid = !(await onpaste(text));
    working = false;
  }

  function onKeydown(event: KeyboardEvent): void {
    if (event.key !== 'Escape') return;
    event.stopPropagation();
    oncancel();
  }
</script>

<div class="csel-paste" data-testid="selector-paste-form" role="presentation" onkeydown={onKeydown}>
  <div class="csel-paste-label">{selectorCopy.pasteLabel}</div>
  <textarea
    bind:this={box}
    bind:value={text}
    placeholder={selectorCopy.pastePlaceholder}
    aria-label={selectorCopy.pasteLabel}
    aria-invalid={invalid ? 'true' : undefined}
    spellcheck="false"
    data-testid="selector-paste-input"></textarea>
  <div class="csel-paste-buttons">
    <button
      type="button"
      class="csel-button"
      disabled={text.trim() === '' || working}
      data-testid="selector-paste-use"
      onclick={() => void submit()}
    >
      {selectorCopy.pasteUse}
    </button>
    <button type="button" class="csel-button" onclick={oncancel}>{selectorCopy.pasteCancel}</button>
  </div>
  {#if invalid}
    <div class="csel-paste-error" role="alert" data-testid="selector-paste-error">
      {selectorCopy.pasteInvalid}
    </div>
  {/if}
</div>
