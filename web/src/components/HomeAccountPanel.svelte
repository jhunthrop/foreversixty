<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in hub summary (spec 2026-09-23 §2 item 2): while this is loading
     or the visitor is signed out, it renders nothing and the sky band's own server-rendered
     signed-out hero (index.astro) stays exactly where it is -- the same "server shell first,
     island swaps in place" trick SessionNav.svelte's header link uses, so the hero never
     shows two competing versions. Both this island's root and index.astro's signed-out block
     share `[grid-area:1/1]` in a shared grid wrapper, so once this mounts signed-in it
     visually occludes the signed-out block (a solid background over the same cell) instead
     of the two stacking and reflowing the page underneath. -->
<script lang="ts">
  import type { MeCharacter } from '../lib/account/api';
  import { classColorVar } from '../lib/report/format';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import CharacterPortrait from './character/CharacterPortrait.svelte';
  import CharacterGuildLine from './character/CharacterGuildLine.svelte';
  import { ratingCopy } from '../lib/rating/copy';
  import { HOME_CHIP_LIMIT, HOME_SIGNED_OUT_ID, homePanelCopy } from '../lib/home-panel-copy';
  import { createHomeHero } from '../lib/account/home-hero.svelte';
  import { armorySimHref } from '../lib/sim/url';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';

  // The session read, hero derivation and rating fetch all live in one shared composable
  // (lib/account/home-hero.svelte.ts) so this island and HomeNextSteps.svelte agree on who
  // "the main character" is without each fetching /v1/me or the rating separately.
  const homeHero = createHomeHero();
  const me = $derived(homeHero.me);
  // 'ready' before this rewrite meant "the fetch attempt finished, whichever way" -- true on
  // both the old `.then` and `.catch` branches -- so it maps to createQueryState's two
  // terminal statuses, not just the successful one.
  const ready = $derived(homeHero.ready);

  /**
   * The grid-overlay CLS trick (this component's root and index.astro's signed-out block
   * share the same `[grid-area:1/1]` cell) only ever covers the signed-out row visually --
   * it stays mounted underneath, so without this a signed-in keyboard/screen-reader user
   * could still tab to, or hear, a duplicate "Sign in with Battle.net" link sitting behind
   * the visible strip. Reaches outside this component's own root via `document`, the same
   * cross-island DOM-reach pattern `syncTabHrefs` in `lib/sim/tabs.ts` uses to coordinate
   * with a sibling shell element it doesn't own. `ready && me !== null && hero !== null`
   * never reverts within one mount today (`fetchMeOnce` resolves once), but the else branch
   * clears both attributes anyway so this stays correct if that ever changes. Gated on
   * `hero !== null` too (review fix): a signed-in visitor with zero characters yet -- a real
   * state, `main-character.ts`'s own `mainCharacter([])` returns null for it -- has no hub
   * summary to show, so the signed-out block (sentence + Battle.net button) must stay live
   * and focusable rather than being hidden behind a hero that renders nothing.
   */
  $effect(() => {
    const signedOut = document.getElementById(HOME_SIGNED_OUT_ID);
    if (signedOut === null) return;
    if (ready && me !== null && hero !== null) {
      signedOut.setAttribute('inert', '');
      signedOut.setAttribute('aria-hidden', 'true');
    } else {
      signedOut.removeAttribute('inert');
      signedOut.removeAttribute('aria-hidden');
      // The session hint hid this block before paint (Base.astro); an expired cookie or an
      // account with no characters means the block is the right thing to show after all.
      if (ready) signedOut.removeAttribute('data-session-hide');
    }
  });

  /**
   * The hero character: `heroCharacter()`'s own rule (a current-character pointer naming a
   * listed character, never a guess) first -- the same rule /account's hero band uses --
   * falling back to `mainCharacter()` (the hub's own well-defined "one character to point
   * at on arrival" algorithm, not a guess either) so a signed-in visitor with characters but
   * no current-character pointer yet still sees a hero, the way a first-ever sign-in does.
   */
  const hero = $derived(homeHero.hero);
  function switchTo(character: MeCharacter): void {
    homeHero.switchTo(character);
  }
  /** Every other character, in the account's order, capped for the row; the rest is "+N more". */
  const others = $derived(
    hero === null || me === null ? [] : me.characters.filter((c) => c.key !== hero.key),
  );
  const shownOthers = $derived(others.slice(0, HOME_CHIP_LIMIT));
  const hiddenCount = $derived(others.length - shownOthers.length);

  // The latest rating figure, when one exists (spec 2026-09-23 §2 item 2): chained off the
  // hero rather than blocking it, since this island is already deferred (`client:visible`)
  // and never sits on the LCP path. Never shown until it resolves with a real sample -- an
  // absent figure, not an invented one.
  const rating = $derived(homeHero.rating);
  const ratingFigure = $derived(
    rating !== null && rating.sample_size > 0 && rating.latest !== null
      ? `${ratingCopy.panelHeading} ${rating.latest.overall.toFixed(2)}`
      : '',
  );
</script>

<!-- The root always renders, even empty: the island hydrates with client:visible (an
     eager island cost the home page one animation step of LCP), and an observer needs a
     box to see. Empty and pointer-events-none, it occludes nothing until signed in. -->
{#if ready && me !== null && hero !== null}
  <div class="flex flex-col gap-4 bg-[var(--color-bg)] [grid-area:1/1]" data-testid="home-account-panel">
    <div class="flex flex-wrap items-end gap-5">
      {#if hero.render_url !== undefined}
        <img
          class="hidden max-h-[320px] w-auto shrink-0 object-contain lg:block"
          src={hero.render_url}
          alt=""
          loading="lazy"
          data-testid="home-hero-render"
        />
      {:else}
        <div class="hidden lg:block">
          <CharacterPortrait character={hero} size="lg" testid="home-hero-portrait" />
        </div>
      {/if}
      <div class="flex flex-col gap-2">
        <CharacterIdentity character={hero} size="lg" descriptor="full" heading testid="home-hero" />
        {#if hero.guild !== undefined}
          <CharacterGuildLine guild={hero.guild} testid="home-hero-guild" />
        {/if}
        <div class="flex flex-wrap items-center gap-3 pt-1">
          {#if hero.build !== undefined}
            <a
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4`}
              href={armorySimHref(hero.key)}
            >
              {homePanelCopy.simCharacter(hero.name)}
            </a>
          {:else}
            <a
              class={`${SECONDARY_BUTTON_FIXED} border-line-warm-strong text-strong px-4`}
              href="/account#characters"
            >
              {homePanelCopy.getTheBuild}
            </a>
          {/if}
          <a class={`${SECONDARY_BUTTON_FIXED} border-line-warm text-text px-4`} href="/planner">
            {homePanelCopy.planTalents}
          </a>
          <a class="text-nav text-[13px] font-semibold" href="/logs">{homePanelCopy.logs}</a>
          <a class="text-nav text-[13px] font-semibold" href="/account">{homePanelCopy.yourCharacters}</a>
          {#if ratingFigure !== ''}
            <span class="text-muted tabular font-mono text-[12px]" data-testid="home-hero-rating"
              >{ratingFigure}</span
            >
          {/if}
        </div>
      </div>
    </div>
    {#if shownOthers.length > 0}
      <ul class="flex w-full flex-wrap gap-2" data-testid="home-character-chips">
        {#each shownOthers as other (other.key)}
          <li>
            <button
              type="button"
              class="border-line-warm bg-raised rounded-control inline-flex min-h-11 items-center gap-2 border px-2 text-[13px] md:min-h-8"
              onclick={() => switchTo(other)}
              aria-label={homePanelCopy.switchTo(other.name)}
              data-testid="home-character-chip"
            >
              <CharacterPortrait character={other} size="sm" testid="home-chip" />
              <span class="font-semibold" style={`color: ${classColorVar(other.class)}`}>{other.name}</span>
              {#if other.level !== undefined}<span class="text-muted">{other.level}</span>{/if}
            </button>
          </li>
        {/each}
        {#if hiddenCount > 0}
          <li>
            <a
              class="text-nav inline-flex min-h-11 items-center text-[13px] font-semibold md:min-h-8"
              href="/account">{homePanelCopy.moreCharacters(hiddenCount)}</a
            >
          </li>
        {/if}
      </ul>
    {/if}
  </div>
{:else}
  <div class="pointer-events-none min-h-[52px] [grid-area:1/1]" aria-hidden="true"></div>
{/if}
