<!-- web/src/components/character/CharacterMissing.svelte -->
<!-- No character by that name on that realm (character spec 2026-11-06 C10): the name as the
     Title, one line saying so, the same name elsewhere first when the search finds one, the lookup
     filled with the name, then how to add the character (sign in, the addon and the companion). -->
<script lang="ts">
  import { ctaCopy } from '../../lib/addon/cta-copy';
  import { searchCharacters, type CharacterSuggestion } from '../../lib/character-search-api';
  import { characterHrefs, characterPageCopy as copy } from '../../lib/character/page-view';
  import { rulesetLabel, type CharacterPath } from '../../lib/characters';
  import { specName } from '../../lib/guild/spec-name';
  import { classSlugFromName } from '../../lib/report/tree-sizes';
  import type { ClassSlug } from '../../lib/ui/primitives';
  import { REGION } from '../../lib/ui/regions';
  import Button from '../ui/Button.svelte';
  import Field from '../ui/Field.svelte';
  import FormRow from '../ui/FormRow.svelte';
  import Notice from '../ui/Notice.svelte';
  import Page from '../ui/Page.svelte';
  import PageHeader from '../ui/PageHeader.svelte';
  import Row from '../ui/Row.svelte';
  import Section from '../ui/Section.svelte';
  import Crest from '../ui/Crest.svelte';

  let { path, name, signedIn }: { path: CharacterPath; name: string; signedIn: boolean } = $props();

  let elsewhere = $state<CharacterSuggestion[]>([]);
  /** The lookup box, filled with the name the address asked for. */
  let query = $derived(name);
  $effect(() => {
    const requested = name;
    const controller = new AbortController();
    elsewhere = [];
    void searchCharacters(requested, { signal: controller.signal, limit: 5 })
      .then((rows) => {
        elsewhere = rows.filter(
          (row) =>
            row.name.toLowerCase() === requested.toLowerCase() &&
            !(row.region === path.region && row.ruleset === path.ruleset),
        );
      })
      .catch(() => {
        elsewhere = [];
      });
    return () => controller.abort();
  });

  function sub(row: CharacterSuggestion): string {
    const who = [
      row.level === null ? '' : `Level ${row.level}`,
      row.spec ? specName(row.spec) : '',
      row.class ?? '',
    ]
      .filter(Boolean)
      .join(' ');
    const where = `${rulesetLabel(row.ruleset)} ${row.region.toUpperCase()}`;
    return who === '' ? where : `${who} · ${where}`;
  }

  function lookup(event: SubmitEvent): void {
    event.preventDefault();
    if (query.trim() !== '') window.location.assign(characterHrefs.lookup(query.trim()));
  }
</script>

<Page layout="ledger" testid="character-not-found">
  {#snippet header()}<PageHeader title={name} testid="character-head" />{/snippet}
  {#snippet bar()}
    <div class="character-summary" data-region={REGION.summary}>
      <Notice
        text={copy.notFound(name, path.region, path.ruleset)}
        tone="neutral"
        testid="character-not-found-line"
      />
    </div>
  {/snippet}
  {#if elsewhere.length > 0}
    <Section title={copy.elsewhere} testid="character-elsewhere">
      {#each elsewhere as row (row.key)}
        <Row
          title={row.name}
          titleTone={row.class ? { classSlug: classSlugFromName(row.class) as ClassSlug } : undefined}
          sub={sub(row)}
          href={characterHrefs.character(row.region, row.ruleset, row.name)}
          open
        >
          {#snippet lead()}<Crest cls={row.class ?? ''} size={40} />{/snippet}
        </Row>
      {/each}
    </Section>
  {/if}
  <Section title={copy.lookup} testid="character-lookup">
    <FormRow label={copy.lookup} onsubmit={lookup}>
      <Field id="character-lookup-name" label={copy.lookupName} bind:value={query} grow />
      {#snippet submit()}<Button tone="primary" type="submit">{copy.lookupSubmit}</Button>{/snippet}
    </FormRow>
  </Section>
  <Section title={copy.add} testid="character-add">
    {#if !signedIn}<Row title={copy.signIn} sub={copy.signInSub} href={characterHrefs.login} open />{/if}
    <Row title={ctaCopy.getAddonAndCompanion} sub={copy.addonSub} href={characterHrefs.setup} open />
  </Section>
</Page>

<style>
  .character-summary {
    padding-inline: var(--page-inset);
  }
</style>
