<!-- web/src/components/guild/GuildNotFound.svelte -->
<!-- A guild address that parses but names no guild the site knows (guild round 3): a written line
     in place of the API's bare "no such guild" and a retry that could only fail again. It names the
     region and ruleset looked in, since the wrong ruleset is the usual miss, and links back to the
     lookup. When the account's guilds or the Top guilds board hold a name that starts with or
     contains what was typed, it offers that one (guild loop round 3): no search endpoint, only
     reads the site already makes. -->
<script lang="ts">
  import { onMount } from 'svelte';
  import { fetchMeOnce } from '../../lib/account/api';
  import { rulesetLabel, type CharacterPath } from '../../lib/characters';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { INLINE_LINK } from '../../lib/guild/field-styles';
  import { guildSuggestion, type GuildCandidate } from '../../lib/guild/guild-suggest';
  import { fetchGuildRankings } from '../../lib/rankings/api';

  let {
    path,
    candidates: givenCandidates = null,
  }: {
    path: CharacterPath;
    /** The guilds to match against, when a test hands them in; the page reads its own on mount. */
    candidates?: readonly GuildCandidate[] | null;
  } = $props();

  let read = $state<readonly GuildCandidate[]>([]);
  const candidates = $derived(givenCandidates ?? read);
  const suggestion = $derived(guildSuggestion(path, candidates));

  /** The account's guilds, then the Top guilds board; a read that fails adds nothing. */
  async function readCandidates(): Promise<readonly GuildCandidate[]> {
    const [me, board] = await Promise.all([
      fetchMeOnce().catch(() => null),
      fetchGuildRankings({ encounter: '', kind: 'progress' }).catch(() => ({ rows: [] })),
    ]);
    return [...(me?.guilds ?? []), ...board.rows.map((row) => row.guild)];
  }

  onMount(() => {
    if (givenCandidates !== null) return;
    void readCandidates().then((found) => (read = found));
  });
</script>

<div class="flex flex-col gap-2">
  <p class="text-[14px]" data-testid="guild-not-found">
    {guildHomeCopy.notFound(`${path.region.toUpperCase()} ${rulesetLabel(path.ruleset)}`)}
    <a class={INLINE_LINK} href="/guild">{guildHomeCopy.lookUpAgain}</a>
  </p>
  {#if suggestion !== null}
    <p class="text-[14px]" data-testid="guild-not-found-suggestion">
      {guildHomeCopy.didYouMean}<a class={INLINE_LINK} href={suggestion.href}
        >{guildHomeCopy.suggestionLink(suggestion.name, suggestion.where)}</a
      >?
    </p>
  {/if}
</div>
