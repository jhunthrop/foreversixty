<!-- web/src/components/HomeGetSetUp.svelte -->
<!-- The "Get set up" banner's signed-in half (review round 1 fix item 3): the three-step
     pitch index.astro renders statically is right for a first-time, signed-out visitor and
     wrong for a returning one who has already done some or all of it -- this island replaces
     it with one slim line naming only what is left, computed from `me` via
     lib/home/get-set-up.ts. Unlike HomeAccountPanel/HomeNextSteps' own occlusion trick (which
     keeps both states the same height on purpose, since a character hub and a sign-in prompt
     are naturally similar heights), the signed-in line here is deliberately much shorter than
     the three-step card it replaces, so this hides the static card with `display: none`
     rather than `visibility: hidden` -- the point is for the block to actually collapse,
     not to reserve the taller card's space forever. -->
<script module lang="ts">
  export const HOME_GET_SET_UP_SIGNED_OUT_ID = 'home-get-set-up-signed-out';
</script>

<script lang="ts">
  import { createHomeHero } from '../lib/account/home-hero.svelte';
  import { getSetUpLine } from '../lib/home/get-set-up';

  // Shares the same cached /v1/me read as the hero and next-steps islands (query()'s own
  // dedupe by URL); this banner needs nothing else off the handle.
  const homeHero = createHomeHero();
  const me = $derived(homeHero.me);
  const ready = $derived(homeHero.ready);

  $effect(() => {
    const el = document.getElementById(HOME_GET_SET_UP_SIGNED_OUT_ID);
    if (el === null) return;
    if (ready && me !== null) {
      el.style.display = 'none';
    } else {
      el.style.display = '';
      // The session hint hid this card before paint (Base.astro); an expired cookie means
      // the three-step pitch is the right thing to show after all.
      if (ready) el.removeAttribute('data-session-hide');
    }
  });

  const line = $derived(me === null ? '' : getSetUpLine(me));
</script>

{#if ready && me !== null}
  <a
    class="text-nav inline-flex min-h-11 items-center text-[13px] font-semibold"
    href="/setup"
    data-testid="home-get-set-up"
  >
    {line}
  </a>
{/if}
