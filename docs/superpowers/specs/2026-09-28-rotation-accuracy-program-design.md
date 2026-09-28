# Rotation accuracy program: every spec, every level, measured

**Date:** 2026-09-28. **Owner ask:** "make sure that our simulations are 100% accurate with the correct rotations and abilities for every level in WoW Forever."

## What "accurate" can mean here, honestly

Three things are measurable and this program makes each of them a test that fails when it is not true:

1. **Every ability the engine models carries the client's numbers.** Cost, cooldown, cast time, duration, rank ids and learn levels come from the Forever client's own spell tables (`data/builds/<active>/spellconst/<class>.json`, `spellranks.json`). Where the client states a number the engine must match it. Where the client only says "server-side script" (Consecration's tick, Tiger's Fury's bonus) the engine carries the best-known value and the conformance report says so.
2. **Every ability a spec's rotation names is one the engine casts, at every level the character has learned it.** A rotation line that survives the level rewrite but never casts is a defect unless the curated file declares why (execute-only, target-type-only, talent-gated in a bare build).
3. **Every damage ability a class has learned by a level is either in the rotation or deliberately left out.** The ladder lists what is learned but unused, per level, so curation is a review of a list and not a guess.

What is NOT measurable from data: whether Forever's servers apply an effect the way the client tooltip reads. That needs combat logs. The logs product exists; a later phase compares simulated cast sequences and damage per ability against real Forever logs per spec.

## Phase 1: measure (this wave)

### 1a. The rotation ladder (site repo, `sim/request`)

A test builds, for every written spec, a character at levels 10, 20, 30, 38, 40, 50 and 60:

- **Talents:** the spec's guide carries a level-60 build code (`web/src/content/guides/<class>/<spec>.md`, `build:` frontmatter, FS1 `<class>:<race>:<tree1>/<tree2>/<tree3>`). Below 60 the ladder spends `level - 9` points walking the level-60 build's trees from the top row down, the spec's own tree (`data/curated/specs.json` `tree_index`) first, then the others in the order the code lists them, skipping talents whose row is not yet reachable. This is a stated approximation of a leveling build, written into the golden so a reader knows what was simmed.
- **Gear:** a weapon set the engine's item database knows, chosen by level: the highest-item-level main hand (and off hand / ranged / wand where the spec uses one) with `required_level <= level` from `data/builds/<active>/simitems.json`. Bare otherwise. The smoke's bare hunter never shot at all, which hid the whole class's state for weeks.
- **Buffs/consumes:** none (a bare character shows the rotation, not the raid).
- **Run:** 300 iterations of the default encounter per (spec, level); record DPS, every cast by spell name and count per iteration, unresolved ids, and the list of the class's learned damage abilities (from `spellranks.json`, castable ranks with a damage effect) that the cast set does not contain.
- **Golden:** `sim/request/testdata/ladder/<spec>.golden.md`, one table per spec (level, dps, casts...), regenerated with `FOREVER_UPDATE_GOLDEN=1`, committed and reviewed like any diff.
- **Fails when:** a cast action that survived the rewrite has zero casts at a level where its spell is learned and the curated file does not list it under a new `expected_idle` key with a reason; an unresolved id appears; DPS at a level is lower than at the level 10 below it (a rotation that gets worse with level is wrong); a spec has no cast but auto-attack at any level >= 20.

### 1b. Engine conformance report (fork, `sim/core`)

A test walks every registered spell of every class at level 60 and at each ladder level, looks up its ActionID's spell in the client's constants (the fork already has `sim/core/spellconst` reading the pipeline's per-class JSON), and compares: resource cost, cooldown (incl. category cooldown), cast time, GCD, duration, required level. It writes `sim/core/testdata/conformance/<class>.golden.md` and fails only on a *new* mismatch versus the committed golden, so the first run records today's gaps and every later change is judged against them. Each mismatch row says which side is believed (client, or "scripted: engine keeps X").

## Phase 2: implement what Forever has and the engine lacks (starts now, in parallel)

From the class audit (`scratchpad/class-audit/*.md`, 2026-09-28), with the client's numbers as the source:

- Warlock: Wrack (1316697), Incinerate (412758 / 1293812 / 1293813), Decimation (440870) on Soul Fire.
- Mage: Arcane Blast and Missile Barrage (talents present, stubs in the engine); Frost Nova with a Frozen state and Ice Lance / Shatter reading it.
- Rogue: Mutilate (1310703) and Venom (1310707).
- Priest: Shadow Word: Death (1309595 / 1309633 / 1309635 / 1309636) and Early Demise.
- Druid: Rip's duration by combo points (the engine's fixed 12 s makes Ferocious Bite unreachable).

Each lands with a test that a level-60 character of the spec casts it and deals damage, and each rotation that named it inert has the `inert` entry removed and the line proven live by the ladder.

## Phase 3: curate every level

With the ladder's "learned but unused" lists in hand, every spec's curated rotation gets the low-level lines it lacks (builders and fillers a leveling character actually presses), and its guide gets a short "while leveling" paragraph. The ladder golden is the review artifact.

## Phase 4: check against Forever's own logs

Per spec, take a real Forever log the site has parsed, replay the same character through the sim, and compare the cast sequence and damage-per-ability distribution. Differences become engine or rotation work items. This phase needs raids to exist on live servers.

## Rules the program adds

- A rotation change is a curated-file change plus `make apl-sync`; it ships with a fork push and `make engine-pin`.
- A curated `inert` id must be one the pinned engine does not resolve; `expected_idle` (new) is for lines the engine resolves but a bare ladder run cannot cast, with the reason.
- The client wins over Classic knowledge wherever it states a number; "server-side script" values are recorded as such in the conformance golden.
