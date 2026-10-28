<!-- web/src/components/claim/SetupBackToClaim.svelte -->
<!-- Spec 2026-10-27 §6.2: /setup opened from a claim page (`?next=`, checked by safeNextPath)
     offers the way back as its one primary, once the account knows a character in that guild.
     /v1/me is read fresh and again when the tab regains focus, so a save made on this page (or in
     the addon) shows the button without a reload. -->
<script lang="ts">
  import { fetchMeFresh, type Me } from '../../lib/account/api';
  import { setupReturnTarget } from '../../lib/account/setup-next';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import { findMyGuild } from '../../lib/guild/my-guild';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';

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
  <div class="min-h-11" data-testid="setup-back-to-claim">
    {#if guildName !== null}
      <a class={`${PRIMARY_BUTTON_FIXED} w-fit px-5`} href={target.href}>
        {guildClaimCopy.backToClaim(guildName)}
      </a>
    {/if}
  </div>
{/if}
