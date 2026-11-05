<!-- web/src/components/guild/GuildNotFound.svelte -->
<!-- A guild address that parses but names no guild the site knows (guild round 3): a written line
     in place of the API's bare "no such guild" and a retry that could only fail again. It names the
     region and ruleset looked in, since the wrong ruleset is the usual miss, and links back to the
     lookup. -->
<script lang="ts">
  import { rulesetLabel, type CharacterPath } from '../../lib/characters';
  import { guildHomeCopy } from '../../lib/guild/copy';
  import { INLINE_LINK } from '../../lib/guild/field-styles';

  let { path }: { path: CharacterPath } = $props();
</script>

<p class="text-[14px]" data-testid="guild-not-found">
  {guildHomeCopy.notFound(`${path.region.toUpperCase()} ${rulesetLabel(path.ruleset)}`)}
  <a class={INLINE_LINK} href="/guild">{guildHomeCopy.lookUpAgain}</a>
</p>
