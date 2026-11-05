<!-- web/src/components/guild/GuildNotFound.svelte -->
<!-- A guild address that parses but names no guild the site knows (guild round 3): a written line
     in place of the API's bare "no such guild" and a retry that could only fail again. It names the
     region and ruleset looked in, since the wrong ruleset is the usual miss, and links back to the
     lookup. Under it, up to three guilds the address may have meant (guild loop round 3): the
     account's guilds and the Top guilds board whose names start with or contain what was typed,
     then the guild search's rows, which reach a stranger with no guilds of their own. No search
     under three letters; a read that fails adds nothing. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchMeOnce } from '../../lib/account/api';
  import { rulesetLabel, type CharacterPath } from '../../lib/characters';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { searchGuilds } from '../../lib/guild/guild-search';
  import { guildSuggestions, searchTerm, type GuildCandidate } from '../../lib/guild/guild-suggest';
  import { fetchGuildRankings } from '../../lib/rankings/api';

  let {
    path,
    candidates: givenKnown = null,
    searched: givenSearched = null,
  }: {
    path: CharacterPath;
    /** The account's and the board's guilds, when a test hands them in; the page reads its own on mount. */
    candidates?: readonly GuildCandidate[] | null;
    /** The guild search's rows, when a test hands them in. */
    searched?: readonly GuildCandidate[] | null;
  } = $props();

  let readKnown = $state<readonly GuildCandidate[]>([]);
  let readSearched = $state<readonly GuildCandidate[]>([]);
  const suggestions = $derived(
    guildSuggestions(path, givenKnown ?? readKnown, givenSearched ?? readSearched),
  );

  /** The account's guilds, then the Top guilds board; a read that fails adds nothing. */
  async function readCandidates(): Promise<readonly GuildCandidate[]> {
    const [me, board] = await Promise.all([
      fetchMeOnce().catch(() => null),
      fetchGuildRankings({ encounter: '', kind: 'progress' }).catch(() => ({ rows: [] })),
    ]);
    return [...(me?.guilds ?? []), ...board.rows.map((row) => row.guild)];
  }

  /** The guild search for the typed words in the typed region; silent on failure. */
  async function readSearch(): Promise<readonly GuildCandidate[]> {
    const term = searchTerm(path);
    if (term === null) return [];
    return searchGuilds(term, path.region).catch(() => []);
  }

  onMount(() => {
    if (givenKnown === null) void readCandidates().then((found) => (readKnown = found));
    if (givenSearched === null) void readSearch().then((found) => (readSearched = found));
  });
</script>

<div class="flex flex-col gap-2">
  <p class="text-[14px]" data-testid="guild-not-found">
    {guildHomeCopy.notFound(`${path.region.toUpperCase()} ${rulesetLabel(path.ruleset)}`)}
    <a class={INLINE_LINK} href="/guild">{guildHomeCopy.lookUpAgain}</a>
  </p>
  {#if suggestions.length > 0}
    <p class="text-[14px]" data-testid="guild-not-found-suggestion">
      {guildHomeCopy.didYouMean}{#each suggestions as suggestion, index (suggestion.href)}{guildHomeCopy.suggestionJoin(
          index,
          suggestions.length,
        )}<a class={INLINE_LINK} href={suggestion.href}
          >{guildHomeCopy.suggestionLink(suggestion.name, suggestion.where)}</a
        >{/each}?
    </p>
  {/if}
</div>
