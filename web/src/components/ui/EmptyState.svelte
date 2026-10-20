<!-- web/src/components/ui/EmptyState.svelte -->
<!-- The empty state every island shows when its data loaded and there is nothing in it: one
     sentence, and at most one action that leads to the thing that would fill it. -->
<script lang="ts">
  let {
    message,
    action = undefined,
    secondary = undefined,
    testid = 'empty-state',
  }: {
    message: string;
    action?: { label: string; href: string };
    /** A quieter second link beside the action. */
    secondary?: { label: string; href: string };
    testid?: string;
  } = $props();
</script>

<div class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" data-testid={testid}>
  <span class="text-muted">{message}</span>
  {#if action !== undefined}
    <a
      class="text-text inline-flex min-h-11 items-center text-[13px] font-semibold underline md:min-h-0"
      href={action.href}
    >
      {action.label}
    </a>
  {/if}
  {#if secondary !== undefined}
    <a
      class="text-muted inline-flex min-h-11 items-center text-[13px] underline md:min-h-0"
      href={secondary.href}
      data-testid="{testid}-secondary"
    >
      {secondary.label}
    </a>
  {/if}
</div>
