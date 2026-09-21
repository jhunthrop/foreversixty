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

<!-- `client:visible` (astro/dist/runtime/client/visible.js) observes this island's own
     `el.children` -- real Elements only, never Svelte's comment-node placeholders -- so
     when `guild` starts null (every load, until the fetch resolves) the block below alone
     renders nothing and the IntersectionObserver has no target to watch, which means the
     island never hydrates at all: not on this render, not once a guild loads. This anchor
     is the fix -- an always-present, zero-size, aria-hidden span with no visual footprint
     (confirmed empirically: a 0x0 empty span still reports `isIntersecting` correctly),
     so the observer always has a real element the moment this island mounts. -->
<span aria-hidden="true"></span>
{#if guild !== null}
  <span class="text-[13px]">
    <span class="text-muted">·</span>
    <!-- Testid on the `<a>` itself, not a wrapping span -- the same convention
         SessionNav.svelte's own `session-my-guild` link already uses, so either one is a
         real `href` a caller can assert against directly rather than needing a nested
         locator. -->
    <a href={guildHref(guild.region, guild.ruleset, guild.name)} data-testid="home-my-guild">My guild</a>
  </span>
{/if}
