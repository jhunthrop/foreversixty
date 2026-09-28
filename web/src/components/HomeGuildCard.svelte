<!-- web/src/components/HomeGuildCard.svelte -->
<!-- The hero's right column (spec 2026-09-28: replaces the horizontally scrolling date
     strip, "it's pretty worthless"), signed-in only. The left column already carries the
     signed-out pitch (index.astro's `home-signed-out` block), so this island simply renders
     nothing once it learns the visitor is signed out -- the same "always-present anchor,
     conditionally empty" trick `HomeAccountPanel.svelte` uses so `client:visible`
     (astro/dist/runtime/client/visible.js) always has a real Element to observe from mount.

     Data: `fetchMeOnce()` for the account's most recent guild membership, then
     `fetchGuildHome(id)` (member-gated) for the claim state, this week's reports and the
     roster; `fetchGuild(path)` alongside it for the same `killed / total` progression
     figure the guild's own public page computes, so a plain member sees the identical
     numbers there. Every rule about which state wins and what it says lives in
     `lib/guild/home-card.ts`'s `homeGuildCardView` -- this component only fetches, and
     renders whatever that pure function returns. -->
<script lang="ts">
  import { fetchMeOnce, type Me } from '../lib/account/api';
  import { sessionHinted } from '../lib/data/query';
  import { fetchGuild } from '../lib/rankings/api';
  import { fetchGuildHome, type GuildHome } from '../lib/guild/api';
  import { characterSlug, guildHref, type CharacterPath } from '../lib/characters';
  import { homeGuildCardCopy } from '../lib/guild/copy';
  import { homeGuildCardView } from '../lib/guild/home-card';
  import { SECONDARY_BUTTON_FIXED } from '../lib/planner/styles';
  import Skeleton from './ui/Skeleton.svelte';

  /** `sessionCookie` exists only so the SSR test can pin the pre-paint decision below
   *  (the same reason HomeTimeline.astro takes `now`); index.astro never passes it, so in
   *  production the card always reads the real document.cookie. */
  let { sessionCookie }: { sessionCookie?: string } = $props();

  type Status = 'loading' | 'signed-out' | 'guild-loading' | 'ready';

  const hinted = (): boolean =>
    sessionCookie === undefined ? sessionHinted() : sessionHinted(sessionCookie);

  /**
   * Decided from the session cookie's readable half before the first paint, never from
   * the /v1/me round trip: Lighthouse (signed out) saw a skeleton panel appear and then
   * vanish once /v1/me answered, and on a phone that collapse moved everything under the
   * hero -- a layout shift over the page's budget (web.yml's verify job, 842aa320). No
   * cookie means nothing to fetch and nothing to draw, the same pre-paint hint
   * Base.astro's inline script and lib/data/query.ts's private cache already trust.
   */
  let status = $state<Status>(hinted() ? 'loading' : 'signed-out');
  let me = $state<Me | null>(null);
  let home = $state<GuildHome | null>(null);
  let progression = $state<{ killed: number; total: number } | null>(null);

  /** `fetchGuild`'s path wants the same URL-safe slug `guildHref` builds into the link
   *  below -- `characterSlug`, the one place that derivation lives -- never the raw
   *  display name. */
  function guildPath(g: { region: string; ruleset: string; name: string }): CharacterPath {
    return {
      region: g.region as CharacterPath['region'],
      ruleset: g.ruleset as CharacterPath['ruleset'],
      slug: characterSlug(g.name),
    };
  }

  /** Home and progression are two independent fetches (a member-gated read and a public
   *  one); either can fail on its own without sinking the other -- a progression failure
   *  only drops the bosses figure from the steady line (`homeGuildCardView`'s own rule),
   *  and a home failure still leaves the guild's name and a link to it. */
  async function loadGuildData(guildId: number, path: CharacterPath): Promise<void> {
    status = 'guild-loading';
    const [homeResult, progressionResult] = await Promise.allSettled([
      fetchGuildHome(guildId),
      fetchGuild(path),
    ]);
    home = homeResult.status === 'fulfilled' ? homeResult.value : null;
    progression =
      progressionResult.status === 'fulfilled'
        ? {
            killed: progressionResult.value.progression.filter((row) => row.kills > 0).length,
            total: progressionResult.value.progression.length,
          }
        : null;
    status = 'ready';
  }

  async function load(): Promise<void> {
    if (!hinted()) {
      status = 'signed-out';
      return;
    }
    status = 'loading';
    home = null;
    progression = null;
    let result: Me | null;
    try {
      result = await fetchMeOnce();
    } catch {
      result = null;
    }
    me = result;
    if (result === null) {
      status = 'signed-out';
      return;
    }
    const guild = result.guilds.length > 0 ? result.guilds[0] : null;
    if (guild === null) {
      status = 'ready';
      return;
    }
    await loadGuildData(guild.id, guildPath(guild));
  }

  $effect(() => {
    void load();
  });

  const view = $derived(status === 'ready' ? homeGuildCardView(me, home, progression) : null);
  const guild = $derived(me !== null && me.guilds.length > 0 ? me.guilds[0] : null);
</script>

<!-- Always-present anchor so `client:visible`'s observer has a real element from mount. -->
<span aria-hidden="true"></span>
{#if status === 'loading' || status === 'guild-loading'}
  <!-- The same min-height as the ready card below, so the skeleton never gives way to a
       taller or shorter panel and the hero's neighbours never move under a signed-in
       visitor either. -->
  <div
    class="bg-raised/85 border-line rounded-panel min-h-[168px] border px-[18px] py-[14px]"
    data-testid="home-guild-card-skeleton"
  >
    <Skeleton lines={3} rowHeight="h-4" testid="home-guild-card-skeleton-rows" />
  </div>
{:else if status === 'ready' && view !== null}
  <div
    class="reveal bg-raised/85 border-line rounded-panel flex min-h-[168px] min-w-0 flex-col gap-3 border px-[18px] py-[14px]"
    data-testid="home-guild-card"
  >
    <div class="flex items-center gap-2">
      <span
        class="bg-gold h-2 w-2 shrink-0 rounded-full shadow-[0_0_10px_rgba(229,185,85,.8)]"
        aria-hidden="true"
      ></span>
      <span class="label text-gold">{homeGuildCardCopy.eyebrow}</span>
    </div>
    {#if view.state === 'no-guild'}
      <p class="text-strong truncate text-[15px] font-semibold" data-testid="home-guild-card-name">
        {homeGuildCardCopy.noGuildLead}
      </p>
    {:else if guild !== null}
      <a
        href={guildHref(guild.region, guild.ruleset, guild.name)}
        class="text-strong w-fit truncate text-[15px] font-semibold"
        data-testid="home-guild-card-name"
      >
        {guild.name}
      </a>
    {/if}
    {#if view.line !== ''}
      <p class="text-muted text-[13px] leading-relaxed" data-testid="home-guild-card-line">{view.line}</p>
    {/if}
    {#if view.action !== null}
      {#if view.state === 'no-guild'}
        <a
          href={view.action.href}
          class="text-nav w-fit text-[13px] font-semibold"
          data-testid="home-guild-card-action"
        >
          {view.action.label}
        </a>
      {:else}
        <a
          href={view.action.href}
          class="{SECONDARY_BUTTON_FIXED} border-line-warm text-text w-fit px-3"
          data-testid="home-guild-card-action"
        >
          {view.action.label}
        </a>
      {/if}
    {/if}
  </div>
{/if}
