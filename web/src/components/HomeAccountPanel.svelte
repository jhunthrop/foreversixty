<!-- web/src/components/HomeAccountPanel.svelte -->
<!-- The home page's signed-in hub summary, rebuilt to the home rebuild spec (2026-09-30)
     §3.B.1/§3.B.2/§3.B.4: the character hero (name, descriptor, sync line), the three
     next-action cards (Best in slot / Talents / Simulator), and the Switch character panel
     that fills the hero's right column. While this is loading or the visitor is signed
     out, it renders nothing and the sky band's own server-rendered signed-out hero
     (index.astro) stays exactly where it is -- the same "server shell first, island swaps
     in place" trick SessionNav.svelte's header link uses, so the hero never shows two
     competing versions. Both this island's root and index.astro's signed-out block share
     `[grid-area:1/1]` in a shared grid wrapper, so once this mounts signed-in it visually
     occludes the signed-out block (a solid background over the same cell) instead of the
     two stacking and reflowing the page underneath.

     The old class-tree art backdrop and the big Battle.net character render image are
     gone: the new right column is the opaque Switch character panel, which would have sat
     behind or fought with either for the same space, and neither appears anywhere in the
     home rebuild spec's own regions. -->
<script lang="ts">
  import { mount, unmount } from 'svelte';
  import type { MeCharacter } from '../lib/account/api';
  import CharacterIdentity from './character/CharacterIdentity.svelte';
  import { factionMarkSrc } from '../lib/faction-mark';
  import { guildHref } from '../lib/characters';
  import { ratingCopy } from '../lib/rating/copy';
  import { homePanelCopy, homeHeroCardsCopy, HOME_SIGNED_OUT_ID } from '../lib/home-panel-copy';
  import { createHomeHero } from '../lib/account/home-hero.svelte';
  import { armorySimHref } from '../lib/sim/url';
  import { listMySims } from '../lib/sim/api';
  import { simCardLine } from '../lib/home/next-steps';
  import type { SimListRow } from '../lib/sim/types';
  import { homeHeroLevelRaceClassLine } from '../lib/account/character-descriptor';
  import { relativeTime } from '../lib/sim/sources';
  import { accountPageCopy } from '../lib/account/account-page-copy';
  import Skeleton from './ui/Skeleton.svelte';

  // The session read, hero derivation and rating fetch all live in one shared composable
  // (lib/account/home-hero.svelte.ts) so this island and HomeNextSteps.svelte agree on who
  // "the main character" is without each fetching /v1/me or the rating separately.
  const homeHero = createHomeHero();
  const me = $derived(homeHero.me);
  // Same rule as the account page's hero band: when no character has a build, the reason is
  // Blizzard's, and saying so once here keeps "No build yet" from reading as the site's failure.
  const noBattlenetData = $derived(me !== null && me.characters.every((c) => c.build === undefined));
  const ready = $derived(homeHero.ready);
  const hero = $derived(homeHero.hero);

  /**
   * The grid-overlay CLS trick (this component's root and index.astro's signed-out block
   * share the same `[grid-area:1/1]` cell) only ever covers the signed-out row visually --
   * it stays mounted underneath, so without this a signed-in keyboard/screen-reader user
   * could still tab to, or hear, a duplicate "Sign in with Battle.net" link sitting behind
   * the visible strip.
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
      if (ready) signedOut.removeAttribute('data-session-hide');
    }
  });

  /**
   * The hero's right column (index.astro's `home-switch-character-slot`) is filled from
   * here, not by its own island: this island already holds the session, and a second
   * island in the hero put its module requests on the largest paint's path for every
   * visitor (Lighthouse, 05f3e478). The panel's module is imported only once /v1/me has
   * answered signed-in, so a signed-out page never fetches it.
   */
  $effect(() => {
    if (!ready || me === null) return;
    const slot = document.querySelector<HTMLElement>('[data-testid="home-switch-character-slot"]');
    if (slot === null) return;
    const account = me;
    const currentKey = hero?.key ?? null;
    let panel: Record<string, unknown> | null = null;
    let cancelled = false;
    void import('./character/HomeSwitchCharacterPanel.svelte').then(
      ({ default: HomeSwitchCharacterPanel }) => {
        if (cancelled) return;
        panel = mount(HomeSwitchCharacterPanel, {
          target: slot,
          props: { me: account, currentKey, onswitch: switchTo },
        });
      },
    );
    return () => {
      cancelled = true;
      if (panel !== null) void unmount(panel);
    };
  });

  function switchTo(character: MeCharacter): void {
    homeHero.switchTo(character);
  }

  /** Home rebuild spec §3.B.1's descriptor line: "Level 24 Troll Hunter" (spec name
   *  omitted -- see homeHeroLevelRaceClassLine's own doc) plus, inline, a FactionMark and
   *  the faction's own coloured word, plus the realm-region clause the mock shows as
   *  "Skyborne-US". Composed here rather than through `characterDescriptor` because that
   *  function's plain-text output is shared by three unrelated surfaces this lane does not
   *  touch, and cannot carry this line's icon/colour markup anyway. */
  const levelRaceClass = $derived(hero === null ? '' : homeHeroLevelRaceClassLine(hero));
  const realmRegion = $derived(hero?.realm === undefined ? '' : `${hero.realm}-${hero.region.toUpperCase()}`);
  const factionLabel = $derived(
    hero?.faction === 'alliance' ? 'Alliance' : hero?.faction === 'horde' ? 'Horde' : '',
  );

  /** §3.B.1's sync line: "{relative time} from the addon · gear and talents in sync" --
   *  only when a build exists at all (never invented), and never the mock's "bags in sync"
   *  clause, which no field on `MeCharacter` backs yet (tenet 8). Past 24h the line gains a
   *  second, un-alarmed sentence per §4's stale-state row. */
  const syncedAgo = $derived(hero?.build === undefined ? '' : relativeTime(hero.build.captured_at));
  const syncedHoursAgo = $derived(
    hero?.build === undefined ? 0 : (Date.now() - new Date(hero.build.captured_at).getTime()) / 3_600_000,
  );

  // The Simulator card's own data: the visitor's latest saved sim, the one figure of the
  // three next-action cards this site can compute honestly today (§3.B.2's own ruling).
  // Independent per-card loading, the same rule HomeNextSteps.svelte's own cards follow.
  let latestSim = $state<SimListRow | null>(null);
  let simStatus = $state<'loading' | 'ready' | 'failed'>('loading');

  $effect(() => {
    const current = hero;
    if (current === null) return;
    simStatus = 'loading';
    void listMySims(1)
      .then((page) => {
        if (current !== hero) return;
        latestSim = page.rows[0] ?? null;
        simStatus = 'ready';
      })
      .catch(() => {
        if (current !== hero) return;
        simStatus = 'failed';
      });
  });

  // The latest rating figure, when one exists: chained off the hero rather than blocking
  // it, since this island is already deferred (`client:idle-after-load`) and never sits on
  // the LCP path. Never shown until it resolves with a real sample -- an absent figure, not
  // an invented one.
  const rating = $derived(homeHero.rating);
  const ratingFigure = $derived(
    rating !== null && rating.sample_size > 0 && rating.latest !== null
      ? `${ratingCopy.panelHeading} ${rating.latest.overall.toFixed(2)}`
      : '',
  );

  const CARD_CLASS =
    'flex flex-col gap-[6px] p-[14px_16px] rounded-panel border border-line bg-gradient-to-b from-card-top to-raised shadow-[inset_0_-1px_0_rgba(229,185,85,.35)] text-strong hover:border-gold-deep transition-colors duration-150';
</script>

<!-- The root always renders, even empty: the island hydrates with client:idle-after-load,
     and an observer needs a box to see. Empty and pointer-events-none, it occludes nothing
     until signed in. -->
{#if ready && me !== null && hero !== null}
  <div class="relative flex flex-col gap-4 [grid-area:1/1]" data-testid="home-account-panel">
    <p class="label text-gold flex items-center gap-[10px]" data-testid="home-hero-eyebrow">
      <i class="bg-gold inline-block h-px w-7" aria-hidden="true"></i>
      {homePanelCopy.yourCharacterEyebrow}
    </p>
    <div class="relative flex flex-wrap items-center gap-[18px]">
      <CharacterIdentity character={hero} size="xl" descriptor="none" heading testid="home-hero">
        {#snippet below()}
          <!-- Review round 1 item 1: one line, the spec's own shape -- "Level 24 Troll
               Marksmanship Hunter · <faction emblem> Horde · <Sample Guild> · Skyborne-US".
               The guild is inline (angle brackets, same convention the mock's own
               "&lt;Sample Guild&gt;" uses), not CharacterGuildLine's separate block-level
               line: that component's own "Verified" text is dropped here rather than given
               a second pill, since the mock's hero line carries no verified indicator at
               all. -->
          <span
            class="flex flex-wrap items-center gap-x-1 text-[15px] text-[#c9c2b2]"
            data-testid="home-hero-descriptor"
          >
            <span>{levelRaceClass}</span>
            {#if factionLabel !== ''}
              <span>·</span>
              <img
                src={factionMarkSrc(hero.faction === 'alliance' ? 'alliance' : 'horde')}
                alt=""
                width="16"
                height="16"
                loading="lazy"
                class="inline-block align-[-2px]"
              />
              <span
                class="font-semibold"
                style={`color: ${hero.faction === 'alliance' ? 'var(--color-alliance)' : 'var(--color-horde)'}`}
                >{factionLabel}</span
              >
            {/if}
            {#if hero.guild !== undefined}
              <span>·</span>
              <span data-testid="home-hero-guild"
                >&lt;<a href={guildHref(hero.region, hero.ruleset, hero.guild.name)} class="text-nav"
                  >{hero.guild.name}</a
                >&gt;</span
              >
            {/if}
            {#if realmRegion !== ''}
              <span>·</span>
              <span>{realmRegion}</span>
            {/if}
          </span>
          {#if syncedAgo !== ''}
            <p class="text-muted text-[12px]" data-testid="home-hero-sync">
              <span class="mono">{syncedAgo}</span>
              {homePanelCopy.syncedFromAddon}
              {#if syncedHoursAgo > 24}
                <br />{homePanelCopy.reopenAddonToRefresh}
              {/if}
            </p>
          {/if}
          {#if noBattlenetData}
            <p class="text-muted text-[13px]" data-testid="home-hero-no-bnet-data">
              {accountPageCopy.noBattlenetDataForRealm}
            </p>
          {/if}
          {#if ratingFigure !== ''}
            <span class="text-muted tabular font-mono text-[12px]" data-testid="home-hero-rating"
              >{ratingFigure}</span
            >
          {/if}
        {/snippet}
      </CharacterIdentity>
    </div>

    <!-- The three next-action cards (§3.B.2): Best in slot and Talents are not yet
         computable against live data (no worn-gear or talent-compare source exists), so
         they show one settled, honest line instead of a figure that would have to be
         invented -- see homeHeroCardsCopy's own doc. Simulator alone carries real data. -->
    <div class="grid grid-cols-1 gap-3 sm:grid-cols-3" data-testid="home-hero-cards">
      <a href="#upgrades" class={CARD_CLASS} data-testid="home-hero-card-bis">
        <span class="label text-gold">{homeHeroCardsCopy.bestInSlotLabel}</span>
        <span class="text-muted text-[13px]">{homeHeroCardsCopy.bestInSlotNotAvailable}</span>
      </a>
      <a href="/planner" class={CARD_CLASS} data-testid="home-hero-card-talents">
        <span class="label text-gold">{homeHeroCardsCopy.talentsLabel}</span>
        <span class="text-muted text-[13px]">{homeHeroCardsCopy.talentsNotAvailable}</span>
      </a>
      <a href={armorySimHref(hero.key)} class={CARD_CLASS} data-testid="home-hero-card-sim">
        <span class="label text-gold">{homeHeroCardsCopy.simulatorLabel}</span>
        {#if simStatus === 'loading'}
          <Skeleton lines={2} rowHeight="h-3" testid="home-hero-card-sim-skeleton" />
        {:else if simStatus === 'failed'}
          <span class="text-muted text-[12px]">{homePanelCopy.noSimYet}</span>
        {:else if latestSim !== null}
          <span class="text-strong text-[13px] font-semibold" data-testid="home-hero-card-sim-value"
            >{simCardLine(latestSim)}</span
          >
        {:else}
          <span class="text-muted text-[13px]">{homePanelCopy.noSimYet}</span>
          <span class="text-nav text-[12px] font-semibold">{homePanelCopy.runAction}</span>
        {/if}
      </a>
    </div>
  </div>
{:else}
  <div class="pointer-events-none min-h-[220px] [grid-area:1/1]" aria-hidden="true"></div>
{/if}
