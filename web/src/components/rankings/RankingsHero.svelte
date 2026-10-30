<!-- web/src/components/rankings/RankingsHero.svelte -->
<!-- The page's head (design/specs/2026-10-11-rankings.md section 3, decision 1, in the account
     page's look by site alignment spec 2026-10-29 §5.7): the shared PageHead on the page ground,
     the one Updated stamp as its meta line, and the tier-list pointer as a line plus an
     ActionLink under it. The job line depends on whether anyone is signed in; until the session
     answers, the page's <html data-session> hint chooses it, so it never changes under the
     reader. -->
<script lang="ts">
  import { rankingsCopy } from '../../lib/rankings/copy';
  import ActionLink from '../ui/ActionLink.svelte';
  import PageHead from '../ui/PageHead.svelte';

  let {
    signedIn,
    updated,
  }: {
    /** True or false once the session answered; null while it has not. */
    signedIn: boolean | null;
    /** The board's `updated_at` day, already printed ("Dec 13"); null before the board loads. */
    updated: string | null;
  } = $props();
</script>

{#snippet job()}
  <span data-testid="rankings-job">
    {#if signedIn === null}
      <span class="ch-hint-only">{rankingsCopy.jobSignedIn}</span>
      <span class="ch-hint-out-only">{rankingsCopy.jobSignedOut}</span>
    {:else}
      {signedIn ? rankingsCopy.jobSignedIn : rankingsCopy.jobSignedOut}
    {/if}
  </span>
{/snippet}

{#snippet stamp()}
  <span data-testid="rankings-updated">{rankingsCopy.updated(updated ?? '')}</span>
{/snippet}

{#snippet pointer()}
  <p class="flex flex-wrap items-center gap-x-3 text-[14px]">
    <span class="text-muted">{rankingsCopy.tierQuestion}</span>
    <ActionLink href={rankingsCopy.tierHref} data-testid="rankings-tier-link"
      >{rankingsCopy.tierLink}</ActionLink
    >
  </p>
{/snippet}

<div class="px-[var(--page-inset)] pt-4 md:px-0 md:pt-7" data-testid="rankings-band">
  <PageHead
    eyebrow={rankingsCopy.eyebrow}
    title={rankingsCopy.h1}
    jobContent={job}
    meta={updated === null ? undefined : stamp}
    below={pointer}
    testid="rankings-head"
  />
</div>
