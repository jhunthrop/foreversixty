<!-- web/src/components/ui/DropZone.svelte -->
<!-- The upload target (logs spec 2026-11-07 section 3A, decision P1): a dashed socket, the one dashed
     surface on the site (as the empty doll socket), 168 px tall, its lines centred. `off` draws it
     at 55% with aria-disabled (signed out: the visitor sees what they get). While a file is dragged
     over it the edge and a 6% ground turn gold. The caller owns the file input and its handlers;
     this draws the target and reports drag and drop. -->
<script lang="ts">
  import type { Snippet } from 'svelte';

  let {
    label = 'Choose a combat log',
    off = false,
    ondrop,
    testid,
    children,
  }: {
    /** The group's accessible name. */
    label?: string;
    off?: boolean;
    /** The dropped files; never called while `off`. */
    ondrop?: (files: FileList | null) => void;
    testid?: string;
    children: Snippet;
  } = $props();

  let dragging = $state(false);

  function over(event: DragEvent): void {
    if (off) return;
    event.preventDefault();
    dragging = true;
  }

  function drop(event: DragEvent): void {
    event.preventDefault();
    dragging = false;
    if (!off) ondrop?.(event.dataTransfer?.files ?? null);
  }
</script>

<div
  class="ui-drop"
  role="group"
  aria-label={label}
  aria-disabled={off ? 'true' : undefined}
  data-off={off ? 'true' : undefined}
  data-dragging={dragging}
  data-testid={testid}
  ondragover={over}
  ondragleave={() => (dragging = false)}
  ondrop={drop}
>
  {@render children()}
</div>

<style>
  .ui-drop {
    display: flex;
    flex-direction: column;
    align-items: center;
    justify-content: center;
    gap: 6px;
    min-height: 168px;
    padding: 20px;
    border: 1px dashed var(--color-line);
    border-radius: 6px;
    text-align: center;
    color: var(--color-muted);
    font-size: 14px;
    line-height: 20px;
  }
  .ui-drop:hover:not([data-off]) {
    border-color: var(--color-nav);
  }
  .ui-drop[data-dragging='true'] {
    border-color: var(--color-gold);
    background: color-mix(in srgb, var(--color-gold) 6%, transparent);
  }
  .ui-drop[data-off] {
    opacity: 0.55;
  }
</style>
