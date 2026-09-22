<!-- web/src/components/GuildStatus.svelte -->
<!-- The one loading/failed rendering every /guild/* view (Guild, GuildJoin, GuildClaim,
     GuildSettings, reached through GuildShell.svelte) used to duplicate four times: a
     Skeleton sized to that view's own ready height while its fetch is in flight, and a
     LoadError with a retry that re-fires the same load() on failure. Never mounted on its
     own -- each of the four hosts renders its own ready view and its own extra statuses
     (Character.svelte-style 'missing', GuildSettings's 'forbidden') outside this
     component, since those are not shared across all four. -->
<script lang="ts">
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  let {
    status,
    error,
    onRetry,
    lines,
    rowHeight = 'h-4',
    minHeight,
    testid,
  }: {
    status: 'loading' | 'failed';
    error: string;
    onRetry: () => void;
    lines: number;
    rowHeight?: string;
    minHeight: string;
    testid: string;
  } = $props();
</script>

{#if status === 'loading'}
  <Skeleton {lines} {rowHeight} {minHeight} testid={`${testid}-skeleton`} />
{:else}
  <LoadError message={error} {onRetry} testid={`${testid}-error`} />
{/if}
