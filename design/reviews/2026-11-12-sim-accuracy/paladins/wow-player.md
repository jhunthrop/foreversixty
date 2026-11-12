# wow-player verdict: paladin guide and card text, sim-accuracy 2026-11-12

Scope: the sentences the paladins lane changed in `web/src/content/guides/paladin/retribution.md` (Judgements in a raid, The Seal twist, the Consecration rank list, The execute, the Levels 20 to 60 Consecration paragraph), `web/src/content/guides/paladin/protection.md` (Taunts and swaps, the leveling Consecration sentence) and the card notes in `data/curated/apl/paladin-retribution.json` (Judgement, Seal of Command with Unsealed Wrath, the twist, Hammer of Wrath) and `data/curated/apl/paladin-protection.json` (Consecration).

## Round 1: NO-SHIP

- MUST: "The execute" argued against a rule the player never saw ("with Judgement a ranged attack, rank 1 no longer edges rank 2 at 53 or 54"). Replaced verbatim: "Hammer of Wrath at your top rank at every level."
- SHOULD 1 (taken verbatim): the Protection Consecration sentence and card note said rank 1 "goes down far more often", which reads as expiring; now "costs less and you can drop it far more often on a dungeon pull", and the guide carries "at 60 without raid buffs go back to rank 1".
- SHOULD 2 (taken): "Judgement is a ranged attack now" lost its patch-note "now".
- SHOULD 3 (taken verbatim): "Only your swings and Holy Strike feed those" became "can fire those".
- SHOULD 4 and the optionals (quantifying "your newest rank" as rank 4; "Judge it as soon as the pull starts"; the Holy Conduit per-point wording): not taken, the wording predates the lane and is used across the guide.
- Data queries: Judgement does not consume the Seal in Forever (Blizzard's Deep Dive recap; the line note is right). The top-level `notes` paragraph in `paladin-retribution.json` carries stale engine history the lane did not touch.

## Round 2: SHIP

No remaining MUST. One non-blocking check, "from 39 use your newest rank" against the rank-40 training level: deliberate, the measured crossover is 39 (rank 1 ties rank 2 there and loses to rank 3 from 40).
