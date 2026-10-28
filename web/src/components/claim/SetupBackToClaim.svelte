<!-- web/src/components/claim/SetupBackToClaim.svelte -->
<!-- Spec 2026-10-27 §6.2: /setup opened from a claim page (`?next=`, checked by safeNextPath)
     offers the way back as its one primary, once the account knows a character in that guild. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../../lib/account/api';
  import { setupReturnTarget } from '../../lib/account/setup-next';
  import { characterSlug } from '../../lib/characters';
  import { guildClaimCopy } from '../../lib/guild/copy';
  import { PRIMARY_BUTTON_FIXED } from '../../lib/planner/styles';

  const target = setupReturnTarget(new URLSearchParams(window.location.search).get('next'));
  let me = $state<Me | null>(null);

  $effect(() => {
    if (target === null) return;
    let live = true;
    void fetchMeOnce()
      .catch(() => null)
      .then((answer) => {
        if (live) me = answer;
      });
    return () => {
      live = false;
    };
  });

  const guildName = $derived(
    target === null
      ? null
      : (me?.guilds.find(
          (g) =>
            g.region === target.guild.region &&
            g.ruleset === target.guild.ruleset &&
            characterSlug(g.name) === target.guild.slug,
        )?.name ?? null),
  );
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
