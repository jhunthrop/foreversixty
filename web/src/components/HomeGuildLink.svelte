<!-- web/src/components/HomeGuildLink.svelte -->
<!-- The home page's "Your guild" panel content (spec 2026-09-23 §2 item 5, replacing the
     small "· My guild" popular-link this file used to render): signed out or signed in with
     no guild, the static claim sentence and the Battle.net sign-in link (plan Ruling 3: no
     guild directory page exists to send a visitor to instead); signed in with a guild, the
     guild's name, a progression figure and a link to its home -- the same `killed / total`
     read `Guild.svelte`'s own progression heading already computes from `GuildPage.progression`.
     `client:visible` (astro/dist/runtime/client/visible.js) observes this island's own
     `el.children` -- real Elements only, never Svelte's comment-node placeholders -- so the
     root below always renders a real (if empty) element for every state, the same anchor
     trick this file used before the guild fetch existed. -->
<script lang="ts">
  import { fetchMeOnce, type Me, type MeGuild } from '../lib/account/api';
  import { fetchGuild, type GuildPage } from '../lib/rankings/api';
  import { characterSlug, guildHref, type CharacterPath } from '../lib/characters';
  import { homeGuildPanelCopy } from '../lib/home-landing-copy';
  import EmptyState from './ui/EmptyState.svelte';
  import LoadError from './ui/LoadError.svelte';
  import Skeleton from './ui/Skeleton.svelte';

  type Status = 'loading' | 'claim' | 'guild-loading' | 'ready' | 'me-failed' | 'guild-failed';

  let status = $state<Status>('loading');
  let guild = $state<MeGuild | null>(null);
  let progression = $state<{ killed: number; total: number } | null>(null);
  let attempt = $state(0);

  /** `fetchGuild`'s path wants the same URL-safe slug `guildHref` builds into the link below
   *  -- `characterSlug`, the one place that derivation lives -- never the raw display name,
   *  which a multi-word guild name like "The Last Watch" would send straight into the path
   *  unencoded and unslugged. */
  function guildPath(g: MeGuild): CharacterPath {
    return {
      region: g.region as CharacterPath['region'],
      ruleset: g.ruleset as CharacterPath['ruleset'],
      slug: characterSlug(g.name),
    };
  }

  async function loadGuild(g: MeGuild): Promise<void> {
    status = 'guild-loading';
    try {
      const page: GuildPage = await fetchGuild(guildPath(g));
      progression = {
        killed: page.progression.filter((row) => row.kills > 0).length,
        total: page.progression.length,
      };
      status = 'ready';
    } catch {
      status = 'guild-failed';
    }
  }

  async function loadMe(): Promise<void> {
    status = 'loading';
    let me: Me | null;
    try {
      me = await fetchMeOnce();
    } catch {
      status = 'me-failed';
      return;
    }
    const first = me !== null && me.guilds.length > 0 ? me.guilds[0] : null;
    guild = first;
    if (first === null) {
      status = 'claim';
      return;
    }
    await loadGuild(first);
  }

  $effect(() => {
    void attempt;
    void loadMe();
  });

  function retry(): void {
    attempt += 1;
  }
</script>

<!-- Always-present anchor so `client:visible`'s observer has a real element from mount. -->
<span aria-hidden="true"></span>
{#if status === 'loading' || status === 'guild-loading'}
  <Skeleton lines={2} rowHeight="h-4" testid="home-guild-skeleton" />
{:else if status === 'me-failed'}
  <LoadError message={homeGuildPanelCopy.failed} onRetry={retry} testid="home-guild-error" />
{:else if status === 'claim'}
  <div class="reveal flex flex-col gap-3" data-testid="home-guild-claim">
    <p class="text-[14px] leading-relaxed">{homeGuildPanelCopy.claimSentence}</p>
    <a href={homeGuildPanelCopy.claimHref} class="text-nav w-fit text-[13px] font-semibold">
      {homeGuildPanelCopy.claimLinkLabel}
    </a>
  </div>
{:else if status === 'guild-failed' && guild !== null}
  <div class="reveal flex flex-col gap-3" data-testid="home-guild-partial">
    <a
      href={guildHref(guild.region, guild.ruleset, guild.name)}
      class="text-strong w-fit text-[15px] font-semibold"
      data-testid="home-my-guild"
    >
      {guild.name}
    </a>
    <EmptyState message={homeGuildPanelCopy.failed} testid="home-guild-progression-error" />
  </div>
{:else if status === 'ready' && guild !== null && progression !== null}
  <div class="reveal flex flex-col gap-2" data-testid="home-guild-ready">
    <a
      href={guildHref(guild.region, guild.ruleset, guild.name)}
      class="text-strong w-fit text-[15px] font-semibold"
      data-testid="home-my-guild"
    >
      {guild.name}
    </a>
    <span class="text-muted tabular font-mono text-[13px]">
      {homeGuildPanelCopy.progressionOf(progression.killed, progression.total)}
    </span>
  </div>
{/if}
