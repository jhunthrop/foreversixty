<!-- web/src/components/ui/EmptyState.svelte -->
<!-- The empty state every island shows when its data loaded and there is nothing in it: one
     sentence, and at most one action that leads to the thing that would fill it. -->
<script lang="ts">
  let {
    message,
    action = undefined,
    testid = 'empty-state',
  }: { message: string; action?: { label: string; href: string }; testid?: string } = $props();
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
</div>
