<!-- web/src/components/guild/GuildSettingsTab.svelte -->
<!-- Guild control-centre spec §4.G: Settings absorbs GuildSettings.svelte's full surface,
     restyled into the tab strip rather than a separate route. Every field and copy string
     is GuildSettings.svelte's own, unchanged (tenet 10 -- this tab is a restyle, not a
     rewrite of settings logic): this is a thin wrapper, not a second implementation. -->
<script lang="ts">
  import type { CharacterPath } from '../../lib/characters';
  import GuildSettings from '../GuildSettings.svelte';

  let { path }: { path: CharacterPath } = $props();
</script>

<section data-testid="guild-settings-tab">
  <GuildSettings {path} />
</section>
