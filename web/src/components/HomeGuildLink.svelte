<!-- web/src/components/HomeGuildLink.svelte -->
<!-- "My guild" on the homepage (spec section 4.3), mounted client:visible in the hero
     (plan ruling 5: no existing homepage section reads GET /v1/me to extend, and this must
     never be client:load on a content page). Renders nothing before the session resolves
     and nothing for the common signed-out/no-guild visitor, so there is nothing to shift
     for the case Lighthouse actually measures (lighthouserc.json runs unauthenticated) --
     only a real signed-in member with a guild sees this line appear. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { guildHref } from '../lib/characters';

  let me = $state<Me | null>(null);

  $effect(() => {
    void fetchMeOnce()
      .then((result) => {
        me = result;
      })
      .catch(() => {
        me = null;
      });
  });

  const guild = $derived(me !== null && me.guilds.length > 0 ? me.guilds[0] : null);
</script>

{#if guild !== null}
  <span class="text-[13px]" data-testid="home-my-guild">
    <span class="text-muted">·</span>
    <a href={guildHref(guild.region, guild.ruleset, guild.name)}>My guild</a>
  </span>
{/if}
