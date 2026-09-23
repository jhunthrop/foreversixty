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
  import { fetchMeOnce, type Me, type MeCharacter } from '../lib/account/api';
  import { readCurrent } from '../lib/current-character';
  import { classColorVar } from '../lib/report/format';
  import { classSquare, classIconUrl, characterDescriptor } from '../lib/account/character-descriptor';
  import { heroCharacter } from '../lib/account/hero-character';
  import { mainCharacter } from '../lib/account/main-character';
  import { parseCharacterPath } from '../lib/characters';
  import { fetchCharacterRating } from '../lib/rankings/api';
  import type { CharacterRating } from '../lib/rating/types';
  import { ratingCopy } from '../lib/rating/copy';
  import { HOME_CHIP_LIMIT, HOME_SIGNED_OUT_ID, homePanelCopy } from '../lib/home-panel-copy';
  import { pointerForCharacter } from '../lib/account/main-character';
  import { writeCurrent, CURRENT_CHARACTER_CHANGED } from '../lib/current-character';

  let me = $state<Me | null>(null);
  let ready = $state(false);

  $effect(() => {
    void fetchMeOnce()
      .then((result) => {
        me = result;
        ready = true;
      })
      .catch(() => {
        ready = true;
      });
  });

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
    }
  });

  /**
   * The hero character: `heroCharacter()`'s own rule (a current-character pointer naming a
   * listed character, never a guess) first -- the same rule /account's hero band uses --
   * falling back to `mainCharacter()` (the hub's own well-defined "one character to point
   * at on arrival" algorithm, not a guess either) so a signed-in visitor with characters but
   * no current-character pointer yet still sees a hero, the way a first-ever sign-in does.
   */
  // Bumped whenever this panel writes the pointer, so `hero` re-reads it: a chip click makes
  // that character current for the planner, the simulator and the account page alike.
  let pointerVersion = $state(0);
  const hero = $derived<MeCharacter | null>(
    me === null ? null : (heroCharacter(readCurrentIfReady(), me.characters) ?? mainCharacter(me.characters)),
  );
  function readCurrentIfReady() {
    void pointerVersion;
    if (!ready || me === null) return null;
    return readCurrent();
  }
  function switchTo(character: MeCharacter): void {
    writeCurrent(pointerForCharacter(character));
    pointerVersion += 1;
    window.dispatchEvent(new Event(CURRENT_CHARACTER_CHANGED));
  }
  /** Every other character, in the account's order, capped for the row; the rest is "+N more". */
  const others = $derived(
    hero === null || me === null ? [] : me.characters.filter((c) => c.key !== hero.key),
  );
  const shownOthers = $derived(others.slice(0, HOME_CHIP_LIMIT));
  const hiddenCount = $derived(others.length - shownOthers.length);
  const heroPath = $derived(hero === null ? null : parseCharacterPath(`/character/${hero.key}`));
  const descriptor = $derived(hero === null ? '' : characterDescriptor(hero));
  const square = $derived(hero === null ? null : classSquare(hero));
  const classIcon = $derived(hero === null ? undefined : classIconUrl(hero));

  // The latest rating figure, when one exists (spec 2026-09-23 §2 item 2): a second fetch,
  // chained off the hero rather than blocking it, since this island is already deferred
  // (`client:visible`) and never sits on the LCP path. Never shown until it resolves with a
  // real sample -- an absent figure, not an invented one.
  let rating = $state<CharacterRating | null>(null);
  $effect(() => {
    const path = heroPath;
    rating = null;
    if (path === null) return;
    void fetchCharacterRating(path)
      .then((result) => {
        if (path !== heroPath) return;
        rating = result;
      })
      .catch(() => {
        rating = null;
      });
  });
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
  <div
    class="flex flex-wrap items-center gap-3 bg-[var(--color-bg)] [grid-area:1/1]"
    data-testid="home-account-panel"
  >
    {#if hero.avatar_url !== undefined}
      <img
        class="h-10 w-10 shrink-0 rounded-[3px] object-cover"
        src={hero.avatar_url}
        alt=""
        loading="lazy"
        data-testid="home-hero-avatar"
      />
    {:else if square !== null}
      <span
        class="relative flex h-10 w-10 shrink-0 items-center justify-center rounded-[3px] text-[16px] font-bold"
        style={`background-color: color-mix(in srgb, ${square.color} 22%, transparent); color: ${square.color}`}
        data-testid="home-hero-avatar-fallback"
      >
        {square.letter}
        {#if classIcon !== undefined}
          <img
            class="absolute inset-0 h-10 w-10 rounded-[3px] object-cover"
            src={classIcon}
            alt=""
            loading="lazy"
          />
        {/if}
      </span>
    {/if}
    <div class="flex min-w-0 flex-col gap-0.5">
      <span class="text-[15px] font-semibold" style={`color: ${classColorVar(hero.class)}`}>{hero.name}</span>
      {#if descriptor !== ''}
        <span class="text-muted text-[12px]">{descriptor}</span>
      {/if}
    </div>
    <a class="text-nav text-[13px] font-semibold" href="/planner">{homePanelCopy.openInPlanner}</a>
    <a class="text-nav text-[13px] font-semibold" href="/sim">{homePanelCopy.openInSimulator}</a>
    <a class="text-nav text-[13px] font-semibold" href="/logs">{homePanelCopy.logs}</a>
    <a class="text-nav text-[13px] font-semibold" href="/account">{homePanelCopy.yourCharacters}</a>
    {#if ratingFigure !== ''}
      <span class="text-muted tabular font-mono text-[12px]" data-testid="home-hero-rating"
        >{ratingFigure}</span
      >
    {/if}
    {#if shownOthers.length > 0}
      <ul class="flex w-full flex-wrap gap-2" data-testid="home-character-chips">
        {#each shownOthers as other (other.key)}
          {@const chip = classSquare(other)}
          {@const icon = classIconUrl(other)}
          <li>
            <button
              type="button"
              class="border-line-warm bg-raised rounded-control inline-flex min-h-11 items-center gap-2 border px-2 text-[13px] md:min-h-8"
              onclick={() => switchTo(other)}
              aria-label={homePanelCopy.switchTo(other.name)}
              data-testid="home-character-chip"
            >
              {#if other.avatar_url !== undefined}
                <img
                  class="h-6 w-6 rounded-[3px] object-cover"
                  src={other.avatar_url}
                  alt=""
                  loading="lazy"
                />
              {:else}
                <span
                  class="relative flex h-6 w-6 items-center justify-center rounded-[3px] text-[11px] font-bold"
                  style={`background-color: color-mix(in srgb, ${chip.color} 22%, transparent); color: ${chip.color}`}
                >
                  {chip.letter}
                  {#if icon !== undefined}
                    <img
                      class="absolute inset-0 h-6 w-6 rounded-[3px] object-cover"
                      src={icon}
                      alt=""
                      loading="lazy"
                    />
                  {/if}
                </span>
              {/if}
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
