<!-- web/src/components/guild/dense/NeedsAWord.svelte -->
<!-- Who the officer talks to before the raid (dense spec 5.1): one Row per raider, the reasons in
     figures, one secondary action: Ping (not answered), Profile (no Discord to ping), Nudge (short
     for the raid, the readiness board's nudge) or Report (rating fell). -->
<script lang="ts">
  import type { Word, WordAction } from '../../../lib/guild/dense/week-board';
  import Button from '../../ui/Button.svelte';
  import Row from '../../ui/Row.svelte';
  import Section from '../../ui/Section.svelte';
  import Crest from '../../ui/Crest.svelte';
  import type { ClassSlug } from '../../../lib/ui/primitives';

  let {
    words,
    hrefOf,
    onPing,
    onNudge,
  }: {
    words: readonly Word[];
    /** Where Profile and Report go for a raider. */
    hrefOf: (word: Word, action: WordAction) => string;
    onPing: () => void;
    /** Copies the raid's nudge for the raider. */
    onNudge: (word: Word) => Promise<void>;
  } = $props();

  let copied = $state<string | null>(null);

  async function nudge(word: Word): Promise<void> {
    await onNudge(word);
    copied = word.key;
  }

  const LABEL: Readonly<Record<WordAction, string>> = {
    profile: 'Profile',
    ping: 'Ping',
    nudge: 'Nudge',
    report: 'Report',
  };
</script>

{#if words.length > 0}
  <Section title="Needs a word" meta={`${words.length} raiders`} panel testid="guild-week-words">
    {#each words as word (word.key)}
      <Row
        title={word.name}
        titleTone={{ classSlug: word.cls as ClassSlug }}
        sub={word.reasons.join(' · ')}
        open
      >
        {#snippet lead()}<Crest cls={word.cls} size={32} />{/snippet}
        {#snippet trail()}
          {#if word.action === 'ping'}
            <Button tone="ghost" size="sm" onclick={onPing}>{LABEL.ping}</Button>
          {:else if word.action === 'nudge'}
            <Button tone="ghost" size="sm" onclick={() => void nudge(word)}
              >{copied === word.key ? 'Copied' : LABEL.nudge}</Button
            >
          {:else}
            <Button tone="ghost" size="sm" href={hrefOf(word, word.action)}>{LABEL[word.action]}</Button>
          {/if}
        {/snippet}
      </Row>
    {/each}
  </Section>
{/if}
