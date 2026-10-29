<!-- web/src/components/rankings/RankingsHero.svelte -->
<!-- The page's band (design/specs/2026-10-11-rankings.md section 3, decision 1): eyebrow, h1, one job
     line, one pointer to the tier list, and the one Updated stamp at the right of the h1 (under the
     pointer on a phone). Full-bleed on the shared column, the same recipe as the BiS landing's.
     The job line depends on whether anyone is signed in; until the session answers, the page's
     <html data-session> hint chooses it, so it never changes under the reader. -->
<script lang="ts">
  import { rankingsCopy } from '../../lib/rankings/copy';

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

<div class="rk-band full-bleed" data-testid="rankings-band">
  <header class="rk-hero page-column">
    <p class="label text-gold flex items-center gap-2.5">
      <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>
      {rankingsCopy.eyebrow}
    </p>
    <div class="rk-head">
      <h1 class="rk-h1 heading-page" data-testid="rankings-h1">{rankingsCopy.h1}</h1>
      {#if updated !== null}
        <span class="rk-stamp text-muted text-[12px]" data-testid="rankings-updated"
          >{rankingsCopy.updated(updated)}</span
        >
      {/if}
    </div>
    <p class="rk-job" data-testid="rankings-job">
      {#if signedIn === null}
        <span class="ch-hint-only">{rankingsCopy.jobSignedIn}</span>
        <span class="ch-hint-out-only">{rankingsCopy.jobSignedOut}</span>
      {:else}
        {signedIn ? rankingsCopy.jobSignedIn : rankingsCopy.jobSignedOut}
      {/if}
    </p>
    <p class="flex flex-wrap items-center gap-x-1.5 text-[14px]">
      <span class="text-muted">{rankingsCopy.tierQuestion}</span>
      <a
        class="hover:text-gold-hover inline-flex min-h-11 items-center font-bold underline-offset-4 hover:underline min-[600px]:min-h-0"
        href={rankingsCopy.tierHref}
        data-testid="rankings-tier-link">{rankingsCopy.tierLink}</a
      >
    </p>
  </header>
</div>

<style>
  .rk-band {
    position: relative;
    background: var(--color-raised);
  }
  .rk-band::after {
    content: '';
    position: absolute;
    inset: 0;
    background: linear-gradient(180deg, #26405c66 0%, #07090dcc 70%, #07090d 100%);
    pointer-events: none;
  }
  .rk-hero {
    position: relative;
    z-index: 1;
    display: flex;
    flex-direction: column;
    gap: 12px;
    padding-block: 28px 24px;
  }
  @media (max-width: 767px) {
    .rk-hero {
      padding-inline: 18px;
    }
  }
  .rk-head {
    display: flex;
    align-items: baseline;
    justify-content: space-between;
    gap: 16px;
  }
  .rk-h1 {
    margin: 0;
  }
  .rk-job {
    margin: 0;
    max-width: 62ch;
    font-size: 14px;
    color: #c9c2b2;
  }
  @media (max-width: 599px) {
    .rk-hero {
      gap: 10px;
      padding-block: 22px 20px;
    }
    .rk-h1 {
      font-size: 26px;
    }
    /* The one Updated stamp sits under the tier pointer on a phone. */
    .rk-head {
      display: contents;
    }
    .rk-stamp {
      order: 5;
    }
  }
</style>
