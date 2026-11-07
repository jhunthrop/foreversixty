<!-- web/src/components/claim/SetupBackToClaim.svelte -->
<!-- Spec 2026-10-27 §6.2, setup spec 2026-11-07 §4.6: /setup opened from a claim page (`?next=`,
     checked by safeNextPath) offers the way back as a Notice under the head, once the account
     knows a character in that guild; the plate's primary stays the first step not done. Its 40 px
     line is reserved while the guild is unknown, so nothing below moves when it lands. /v1/me is
     read fresh and again when the tab regains focus, so a save made on this page (or in the addon)
     shows the way back without a reload. -->
<script lang="ts">
  import { fetchMeFresh, type Me } from '../../lib/account/api';
  import { setupReturnTarget } from '../../lib/account/setup-next';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import { findMyGuild } from '../../lib/guild/my-guild';
  import { REGION } from '../../lib/ui/regions';
  import Notice from '../ui/Notice.svelte';

  const target = setupReturnTarget(
    typeof window === 'undefined' ? null : new URLSearchParams(window.location.search).get('next'),
  );
  let me = $state<Me | null>(null);

  $effect(() => {
    if (target === null) return;
    let live = true;
    const read = (): void => {
      void fetchMeFresh((fresh) => {
        if (live) me = fresh;
      })
        .catch(() => null)
        .then((answer) => {
          if (live && answer !== null) me = answer;
        });
    };
    read();
    window.addEventListener('focus', read);
    return () => {
      live = false;
      window.removeEventListener('focus', read);
    };
  });

  const guildName = $derived(target === null ? null : (findMyGuild(me, target.guild)?.name ?? null));
</script>

{#if target !== null}
  <div
    class="setup-claim min-h-10 px-[var(--page-inset)]"
    data-region={REGION.summary}
    data-testid="setup-back-to-claim"
  >
    {#if guildName !== null}
      <Notice
        text={guildClaimCopy.setupClaimReady(guildName)}
        action={{ label: guildClaimCopy.backToClaim(guildName), href: target.href }}
        testid="setup-back-to-claim-notice"
      />
    {/if}
  </div>
{/if}

<style>
  /* The way back is a 44 px target: 12 px of padding above and below, taken back by the same
     negative margin, so the Notice keeps its one line. */
  .setup-claim :global(.ui-notice a) {
    display: inline-block;
    margin-block: -12px;
    padding-block: 12px;
  }
</style>
