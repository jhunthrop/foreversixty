<!-- design/mocks/ch_harness/HeaderStatesHarness.svelte
     Acceptance harness for the BUILT shared character header (spec 2026-10-11): every state of
     the real CharacterHeader.svelte on one sheet, laid out like gen_character_header.py's states
     sheet so the two can be composed side by side at the same scale. Not shipped: shoot_character_
     header_built.mjs copies this and ch-states.astro into web/src for the capture and removes them
     again; the page carries `data-session="1"` so the loading skeleton shows, as for a hinted visit. Every character and figure here is an EXAMPLE (the same invented Obnoxious Yell the
     mock uses); the headers are the real component fed through its `source` and `answer` props. -->
<script lang="ts">
  import type { MeCharacter } from '../../lib/account/api';
  import type { CharacterCardState } from '../../lib/bis/character-card-state.svelte';
  import type { AnswerSpec, HeaderAction, HeaderAnswer } from '../../lib/character-header/types';
  import CharacterHeader from './CharacterHeader.svelte';
  import CharacterHeaderSignedOut from './CharacterHeaderSignedOut.svelte';

  const MINUTE = 60_000;
  const DAY = 24 * 60 * MINUTE;
  const iso = (agoMs: number): string => new Date(Date.now() - agoMs).toISOString();

  const EX: MeCharacter = {
    key: 'us/normal/obnoxious-yell',
    region: 'us',
    ruleset: 'normal',
    name: 'Obnoxious Yell',
    class: 'Warrior',
    spec: 'Fury',
    race: 'Gnome',
    level: 60,
    faction: 'alliance',
    realm: 'Living Flame',
    build: { source: 'addon', captured_at: iso(12 * MINUTE), gear: { head: 1 } },
  };
  const NORACE: MeCharacter = { ...EX, race: undefined };
  const LONG: MeCharacter = { ...EX, name: 'Bartholomew Thunderfist' };
  const STALE: MeCharacter = { ...EX, build: { source: 'addon', captured_at: iso(19 * DAY), gear: { head: 1 } } };
  const NOGEAR: MeCharacter = { ...EX, build: undefined };
  const UNDER20: MeCharacter = { ...EX, level: 12 };

  const fixed = (character: MeCharacter | null, status: CharacterCardState['status'] = 'ready'): CharacterCardState => ({
    status,
    error: '',
    me: null,
    character,
    refresh: () => undefined,
  });

  const resolves = (answer: HeaderAnswer) => (): Promise<HeaderAnswer> => Promise.resolve(answer);
  const never = (): Promise<HeaderAnswer> => new Promise(() => undefined);
  const fails = (): Promise<HeaderAnswer> => Promise.reject(new Error('example'));

  const FAILED = 'Your rank did not load.';
  const rankings = (load: AnswerSpec['load']): AnswerSpec => ({ strings: { failed: FAILED }, load });
  const RANK: HeaderAnswer = {
    kind: 'figures',
    label: 'Your rank',
    figures: [{ value: '14th', unit: 'of 212 Fury Warrior kills', color: 'var(--color-parse-purple)' }],
    note: { text: 'On Onyxia, Normal ruleset' },
  };
  const LATEST: HeaderAnswer = {
    kind: 'figures',
    label: 'Latest kill',
    figures: [{ value: '612', unit: 'DPS' }],
    note: { text: 'Onyxia kill · 3 days ago', link: { label: 'Open the fight', href: '#fight' } },
    secondNote: { text: 'Last report 1 day ago', href: '#report' },
  };
  const BIS: HeaderAnswer = {
    kind: 'figures',
    label: 'Best-in-slot slots you wear',
    staleLabel: 'Slots you wore at the last sync',
    figures: [{ value: '6', unit: 'of 17 slots' }],
    bar: 6 / 17,
  };
  const HOME: HeaderAnswer = {
    kind: 'figures',
    label: 'Upgrades in your band',
    figures: [
      { value: '9', unit: 'upgrades' },
      { value: '+11.2', unit: 'DPS' },
    ],
    note: { text: 'Best in slot for 60, together' },
  };
  const bisSpec: AnswerSpec = {
    levelFloor: 20,
    needsGear: true,
    placeholderBar: true,
    strings: {
      belowLevelFloor: 'Best in slot starts at 20.',
      noGear: {
        sentence: 'Add the addon to see how many of these slots you already wear.',
        link: { label: 'Get the addon', href: '#setup' },
      },
      failed: 'Your gear count did not load.',
    },
    load: resolves(BIS),
  };

  const act = (label: string, href = '#'): HeaderAction => ({ label, href });
  const bisAction = (label = 'Fury gear for 60 →') => act(label);
  const UPLOAD: HeaderAction = { kind: 'picker', label: 'Upload a log', href: '#upload' };
</script>

<div class="sheet">
  <h1>Character header: every state, as built</h1>
  <h2>1440, every state</h2>

  <div class="cell">
    <span class="label cap">Ready (Rankings)</span>
    <CharacterHeader source={fixed(EX)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
  </div>
  <div class="cell">
    <span class="label cap">Ready, race not known (Logs)</span>
    <CharacterHeader source={fixed(NORACE)} action={UPLOAD} answer={rankings(resolves(LATEST))} />
  </div>
  <div class="cell">
    <span class="label cap">Ready, home: two figures at the same size</span>
    <CharacterHeader source={fixed(EX)} action={act('See your upgrades →')} answer={rankings(resolves(HOME))} />
  </div>
  <div class="cell">
    <span class="label cap">Loading, nothing known (/v1/me pending)</span>
    <CharacterHeader source={fixed(null, 'loading')} action={act('x')} answer={rankings(never)} />
  </div>
  <div class="cell">
    <span class="label cap">Loading the answer only (also what a selector switch looks like)</span>
    <CharacterHeader source={fixed(EX)} action={act('See your parses →')} answer={rankings(never)} />
  </div>
  <div class="cell">
    <span class="label cap">Stale sync (19 days), BiS</span>
    <CharacterHeader source={fixed(STALE)} action={bisAction()} answer={bisSpec} />
  </div>
  <div class="cell">
    <span class="label cap">Stale sync, a page that does not read gear (Logs)</span>
    <CharacterHeader source={fixed(STALE)} action={UPLOAD} answer={rankings(resolves(LATEST))} />
  </div>
  <div class="cell">
    <span class="label cap">No gear on file, a page that needs gear (BiS)</span>
    <CharacterHeader source={fixed(NOGEAR)} action={bisAction()} answer={bisSpec} />
  </div>
  <div class="cell">
    <span class="label cap">No gear on file, a page that does not (Logs)</span>
    <CharacterHeader source={fixed(NOGEAR)} action={UPLOAD} answer={rankings(resolves(LATEST))} />
  </div>
  <div class="cell">
    <span class="label cap">Under 20, BiS</span>
    <CharacterHeader
      source={fixed(UNDER20)}
      action={bisAction('Fury gear for 20 to 29 →')}
      answer={bisSpec}
    />
  </div>
  <div class="cell">
    <span class="label cap">Under 20, Rankings (identical to ready)</span>
    <CharacterHeader source={fixed(UNDER20)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
  </div>
  <div class="cell">
    <span class="label cap">Empty answer (Rankings)</span>
    <CharacterHeader
      source={fixed(EX)}
      action={act('See your parses →')}
      answer={rankings(
        resolves({
          kind: 'sentence',
          text: 'Not ranked yet on Onyxia. Upload a log of a kill to get on the board.',
        }),
      )}
    />
  </div>
  <div class="cell">
    <span class="label cap">Empty answer (Logs)</span>
    <CharacterHeader
      source={fixed(EX)}
      action={UPLOAD}
      answer={rankings(resolves({ kind: 'sentence', text: 'No logs yet. Upload one, or log live with the companion.' }))}
    />
  </div>
  <div class="cell">
    <span class="label cap">Answer failed</span>
    <CharacterHeader source={fixed(EX)} action={act('See your parses →')} answer={rankings(fails)} />
  </div>
  <div class="cell">
    <span class="label cap">A long two-word name</span>
    <CharacterHeader source={fixed(LONG)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
  </div>
  <div class="cell">
    <span class="label cap">Signed out, no character, or /v1/me failed: the one shared band</span>
    <CharacterHeaderSignedOut />
  </div>

  <h2>390, every state (the card layout)</h2>
  <div class="phones">
    <div class="cell phone">
      <span class="label cap">Ready</span>
      <CharacterHeader layout="card" source={fixed(EX)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
    </div>
    <div class="cell phone">
      <span class="label cap">Loading</span>
      <CharacterHeader layout="card" source={fixed(null, 'loading')} action={act('x')} answer={rankings(never)} />
    </div>
    <div class="cell phone">
      <span class="label cap">Home, two figures</span>
      <CharacterHeader layout="card" source={fixed(EX)} action={act('See your upgrades →')} answer={rankings(resolves(HOME))} />
    </div>
    <div class="cell phone">
      <span class="label cap">Signed out</span>
      <CharacterHeaderSignedOut />
    </div>
    <div class="cell phone">
      <span class="label cap">Stale (BiS)</span>
      <CharacterHeader layout="card" source={fixed(STALE)} action={bisAction()} answer={bisSpec} />
    </div>
    <div class="cell phone">
      <span class="label cap">No gear on file (BiS)</span>
      <CharacterHeader layout="card" source={fixed(NOGEAR)} action={bisAction()} answer={bisSpec} />
    </div>
    <div class="cell phone">
      <span class="label cap">Under 20 (BiS)</span>
      <CharacterHeader layout="card" source={fixed(UNDER20)} action={bisAction('Fury gear for 20 to 29 →')} answer={bisSpec} />
    </div>
    <div class="cell phone">
      <span class="label cap">Empty (Logs)</span>
      <CharacterHeader
        layout="card"
        source={fixed(EX)}
        action={UPLOAD}
        answer={rankings(resolves({ kind: 'sentence', text: 'No logs yet. Upload one, or log live with the companion.' }))}
      />
    </div>
    <div class="cell phone">
      <span class="label cap">Error</span>
      <CharacterHeader layout="card" source={fixed(EX)} action={act('See your parses →')} answer={rankings(fails)} />
    </div>
    <div class="cell phone">
      <span class="label cap">Long name: two lines reserved</span>
      <CharacterHeader layout="card" source={fixed(LONG)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
    </div>
    <div class="cell phone">
      <span class="label cap">Logs, ready</span>
      <CharacterHeader layout="card" source={fixed(EX)} action={UPLOAD} answer={rankings(resolves(LATEST))} />
    </div>
  </div>

  <h2>What each page fills</h2>
  <div class="cell">
    <span class="label cap">Home</span>
    <CharacterHeader source={fixed(EX)} action={act('See your upgrades →')} answer={rankings(resolves(HOME))} />
  </div>
  <div class="cell">
    <span class="label cap">BiS landing</span>
    <CharacterHeader source={fixed(EX)} action={bisAction()} answer={bisSpec} />
  </div>
  <div class="cell">
    <span class="label cap">Rankings</span>
    <CharacterHeader source={fixed(EX)} action={act('See your parses →')} answer={rankings(resolves(RANK))} />
  </div>
  <div class="cell">
    <span class="label cap">Logs</span>
    <CharacterHeader source={fixed(EX)} action={UPLOAD} answer={rankings(resolves(LATEST))} />
  </div>
  <div class="cards">
    <div class="cell card440">
      <span class="label cap">BiS spec page card, 440 px (secondary)</span>
      <CharacterHeader layout="card" emphasis="secondary" source={fixed(EX)} action={act('Send this list to the addon')} answer={bisSpec} />
    </div>
    <div class="cell card440">
      <span class="label cap">Planner card, 440 px (secondary, no answer)</span>
      <CharacterHeader layout="card" emphasis="secondary" source={fixed(EX)} action={act('Send this build to the addon')} />
    </div>
  </div>
</div>

<style>
  .sheet {
    display: flex;
    flex-direction: column;
    gap: 22px;
    box-sizing: border-box;
    max-width: 1344px;
    margin: 0 auto;
    padding: 36px 48px;
  }
  h1 {
    margin: 0;
    font-size: 22px;
  }
  h2 {
    margin: 12px 0 0;
    font-size: 18px;
    letter-spacing: 0.1em;
    text-transform: uppercase;
  }
  .cell {
    display: flex;
    flex-direction: column;
    gap: 8px;
  }
  .cap {
    color: var(--color-gold);
  }
  .phones {
    display: flex;
    flex-wrap: wrap;
    gap: 18px 24px;
    align-items: flex-start;
  }
  .phone {
    width: 390px;
  }
  .cards {
    display: flex;
    flex-wrap: wrap;
    gap: 18px 24px;
    align-items: flex-start;
  }
  .card440 {
    width: 440px;
  }
  @media (max-width: 600px) {
    .sheet {
      padding: 18px;
    }
    .phone,
    .card440 {
      width: 100%;
    }
  }
</style>
