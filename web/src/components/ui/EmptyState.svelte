<!-- web/src/components/ui/EmptyState.svelte -->
<!-- The one empty state (component audit 2026-11-10 cluster 10), for a region whose data loaded with
     nothing in it. Three shapes, one look:
     - the line (default): the muted sentence and at most one action that leads to the thing that
       would fill it, plus a quieter second link;
     - with `fill`: the designed empty (the Rankings board's, owner 2026-11-09 "every empty state
       designed and honest"): what is missing in strong, what fills it muted under it, then the
       action as the primary button and the second link as a gold link;
     - `inset`: a panel's one muted line with the panel's 16 px inset (a report table or rail with
       nothing in the window); `section` takes back a Section head's own 12 px. -->
<script lang="ts">
  import type { Snippet } from 'svelte';
  import { LINK } from '../../lib/ui/primitives';
  import Button from './Button.svelte';

  let {
    message,
    fill,
    action,
    secondary,
    inset = false,
    section = false,
    role,
    testid = 'empty-state',
    children,
  }: {
    message: string;
    /** What would fill it, under the message: the designed empty. */
    fill?: string;
    action?: { label: string; href: string };
    /** A quieter second link beside the action. */
    secondary?: { label: string; href: string };
    /** A panel's muted line with its 16 px inset. */
    inset?: boolean;
    /** `inset` under a Section head: takes back the head's 12 px. */
    section?: boolean;
    role?: 'status' | 'alert';
    testid?: string;
    /** More under the actions (the bosses ranked elsewhere), behind a hairline. */
    children?: Snippet;
  } = $props();
</script>

{#if inset}
  <p class={['text-muted p-4 text-[14px] leading-6', section && '-mt-3']} {role} data-testid={testid}>{message}</p>
{:else if fill !== undefined}
  <div class="flex flex-col gap-1 py-1" {role} data-testid={testid}>
    <p class="text-strong font-semibold">{message}</p>
    <p class="text-muted text-[14px] leading-5">{fill}</p>
    {#if action !== undefined || secondary !== undefined}
      <div class="mt-3 flex flex-wrap items-center gap-x-5 gap-y-2">
        {#if action !== undefined}<Button tone="primary" size="sm" href={action.href} testid="{testid}-action"
            >{action.label}</Button
          >{/if}
        {#if secondary !== undefined}<a class={LINK} href={secondary.href} data-testid="{testid}-secondary"
            >{secondary.label} ›</a
          >{/if}
      </div>
    {/if}
    {#if children}<div class="border-line-soft mt-4 border-t pt-4">{@render children()}</div>{/if}
  </div>
{:else}
  <div class="flex min-h-11 flex-wrap items-center gap-3 text-[14px]" {role} data-testid={testid}>
    <span class="text-muted">{message}</span>
    {#if action !== undefined}
      <a
        class="text-text inline-flex min-h-11 items-center text-[13px] font-semibold underline md:min-h-0"
        href={action.href}
        data-testid="{testid}-action"
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
    {@render children?.()}
  </div>
{/if}
