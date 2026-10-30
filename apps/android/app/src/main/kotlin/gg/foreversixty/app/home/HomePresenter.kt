package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemNumbers
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.FinishedSim
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.HomeWorn
import kotlin.math.roundToInt

/** The character Home is about, from `/v1/me` (lane F1's selector picks it). */
data class HomeCharacter(
    val key: String,
    val region: String,
    val ruleset: String,
    val name: String,
    val race: String?,
    val className: String?,
    /** The primary tree's name as the API gives it (`Beast Mastery`). */
    val specName: String?,
    val level: Int?,
    val guild: String?,
    val faction: String?,
    /** Worn item ids by slot from the last export, for the gear rows the answer does not name. */
    val gear: Map<String, Int> = emptyMap(),
) {
    /** The path segment of the character routes (`velaren` in `us/normal/velaren`). */
    val slug: String get() = key.substringAfterLast('/')

    companion object {
        fun from(character: AccountCharacter) = HomeCharacter(
            key = character.key, region = character.region, ruleset = character.ruleset, name = character.name,
            race = character.race, className = character.className, specName = character.spec, level = character.level,
            guild = character.guild?.name, faction = character.faction, gear = character.build?.gear.orEmpty(),
        )
    }
}

/** One upgrade row (spec F2.1 item 2). */
data class HomeUpgradeRow(
    val slot: String,
    val pick: ItemLook,
    val wornId: Int?,
    val meta: String,
    val replaces: String,
    val gain: Gain,
    val sourceLine: String,
) {
    sealed interface Gain {
        /** `+5.3` over `DPS`. */
        data class Measured(val figure: String, val unit: String) : Gain

        /** The server's words for a gain it did not model (`Gain not modelled yet`). */
        data class Words(val text: String) : Gain
    }
}

/** A button on an answer sentence. */
data class HomeAction(val kind: Kind, val title: String, val primary: Boolean) {
    enum class Kind { GetAddon, OpenList, OpenPlanner }
}

/** What the answer region says (spec F2.1 to F2.4 and the empty states). */
sealed interface HomeAnswerBody {
    data class Upgrades(val rows: List<HomeUpgradeRow>, val footer: String) : HomeAnswerBody
    data class AllMatch(val text: String) : HomeAnswerBody
    data class Message(val text: String, val buttons: List<HomeAction>, val note: String?) : HomeAnswerBody
}

/** One "Next for {name}" row. */
data class HomeNextRow(val kind: Kind, val title: String, val line: String, val tag: Tag?) {
    enum class Kind { Talents, Sims }
    enum class Tag { Matches, Differs }
}

/** Everything Home draws for one answer. Built by [HomePresenter], rendered by `HomeAnswer`. */
data class HomeAnswerModel(
    val heroMeta: String,
    val count: String?,
    val gainFigure: String?,
    val gainUnit: String?,
    val note: String?,
    val synced: String?,
    val bandHead: String?,
    val body: HomeAnswerBody,
    val already: String?,
    val nextTitle: String,
    val next: List<HomeNextRow>,
)

/**
 * Turns the API's home answer into words. It formats and never ranks: the rows keep the server's
 * order and every gain is the server's (plan guard rail "one behaviour, one place").
 */
class HomePresenter(
    val character: HomeCharacter,
    private val finishedSims: List<FinishedSim>,
    private val time: HomeTime,
    private val res: Resources,
) {
    fun model(home: CharacterHome): HomeAnswerModel {
        val isAnswer = home.state == STATE_OK && home.upgradeCount > 0
        return HomeAnswerModel(
            heroMeta = heroMeta,
            count = if (isAnswer) res.getQuantityString(R.plurals.home_answer_upgrades, home.upgradeCount, home.upgradeCount) else null,
            gainFigure = if (isAnswer) home.totalGain?.let(ItemNumbers::signedTenths) else null,
            gainUnit = if (isAnswer && home.totalGain != null) home.unitWord.orEmpty() else null,
            note = if (isAnswer) note(home) else null,
            synced = synced(home),
            bandHead = bandHead(home, res),
            body = body(home),
            already = already(home),
            nextTitle = res.getString(R.string.home_next_title, character.name),
            next = listOf(talentsRow(home), simsRow(home.spec)),
        )
    }

    /** The Next rows alone, for the error state with nothing cached. */
    fun nextRowsWithoutAnswer(): List<HomeNextRow> = listOf(simsRow(null))

    /** `Orc Beast Mastery Hunter · Level 60 · Olympus XXVII`. */
    val heroMeta: String
        get() {
            val who = listOfNotNull(character.race, character.specName, character.className).joinToString(" ") { words(it) }
            return listOfNotNull(
                who.ifEmpty { null },
                character.level?.let { res.getString(R.string.home_hero_level, it) },
                character.guild,
            ).joinToString(" · ")
        }

    private fun note(home: CharacterHome): String? {
        val label = home.bandLabel ?: return null
        val base = res.getString(R.string.home_answer_note, label)
        val unmodelled = home.notModelledCount
        return if (unmodelled > 0) base + res.getQuantityString(R.plurals.home_answer_notModelled, unmodelled, unmodelled) else base
    }

    private fun synced(home: CharacterHome): String? {
        val at = parseInstant(home.syncedAt) ?: return null
        val age = time.age(at, res)
        return res.getString(if (home.syncedFrom == "blizzard") R.string.home_answer_synced_bnet else R.string.home_answer_synced_addon, age)
    }

    private fun body(home: CharacterHome): HomeAnswerBody {
        val spec = character.specName.orEmpty()
        val band = home.bandLabel.orEmpty()
        return when (home.state) {
            "no_gear" -> HomeAnswerBody.Message(
                res.getString(R.string.home_answer_noGear),
                listOf(
                    HomeAction(HomeAction.Kind.GetAddon, res.getString(R.string.signin_getAddon), primary = true),
                    HomeAction(HomeAction.Kind.OpenList, res.getString(R.string.home_answer_listButton, spec, band), primary = false),
                ),
                res.getString(R.string.home_answer_noGearNote),
            )
            "under_20" -> HomeAnswerBody.Message(
                res.getString(R.string.home_answer_under20),
                listOf(HomeAction(HomeAction.Kind.OpenList, res.getString(R.string.home_answer_under20Button, spec), primary = false)),
                null,
            )
            "no_spec" -> HomeAnswerBody.Message(
                res.getString(R.string.home_answer_pickSpec),
                listOf(HomeAction(HomeAction.Kind.OpenPlanner, res.getString(R.string.home_answer_openPlanner), primary = false)),
                null,
            )
            "no_list" -> HomeAnswerBody.Message(res.getString(R.string.home_answer_noList), emptyList(), null)
            else -> upgradesBody(home, band)
        }
    }

    private fun upgradesBody(home: CharacterHome, band: String): HomeAnswerBody {
        if (home.upgradeCount <= 0) return HomeAnswerBody.AllMatch(res.getString(R.string.home_answer_allMatch, band))
        val rows = home.upgrades.take(VISIBLE_ROWS).map { row(it, home.unitWord.orEmpty()) }
        val more = home.upgradeCount - rows.size
        val footer = if (more > 0) {
            res.getQuantityString(R.plurals.home_upgrades_more, more, more)
        } else {
            res.getString(R.string.home_upgrades_allSlots)
        }
        return HomeAnswerBody.Upgrades(rows, footer)
    }

    private fun row(upgrade: HomeUpgrade, unit: String): HomeUpgradeRow {
        val pick = upgrade.pick
        val gain = upgrade.gain
        return HomeUpgradeRow(
            slot = upgrade.slot,
            pick = ItemLook(pick.id, pick.name, pick.quality, pick.icon, pick.itemLevel),
            wornId = upgrade.worn?.id,
            meta = "${upgrade.slotLabel} · ${pick.sourceLine}",
            replaces = replaces(upgrade.worn, res),
            gain = if (upgrade.modelled && gain != null) {
                HomeUpgradeRow.Gain.Measured(ItemNumbers.signedTenths(gain), unit)
            } else {
                HomeUpgradeRow.Gain.Words(upgrade.gainWords)
            },
            sourceLine = pick.sourceLine,
        )
    }

    private fun already(home: CharacterHome): String? =
        home.alreadyBest.takeIf { it.isNotEmpty() }
            ?.let { best -> res.getString(R.string.home_upgrades_already, best.joinToString(", ") { it.slotLabel }) }

    private fun talentsRow(home: CharacterHome): HomeNextRow {
        val band = home.bandLabel.orEmpty()
        val talents = home.talents
        val title = res.getString(R.string.home_next_talents)
        return when (talents.state) {
            "optimized" -> HomeNextRow(
                HomeNextRow.Kind.Talents, title,
                res.getString(R.string.home_next_talentsMatch, character.specName.orEmpty(), band), HomeNextRow.Tag.Matches,
            )
            "differs" -> HomeNextRow(
                HomeNextRow.Kind.Talents, title,
                res.getString(R.string.home_next_talentsDiffer, talents.pointsDiffer ?: 0, talents.bandPoints ?: 0, band),
                HomeNextRow.Tag.Differs,
            )
            "no_build" -> HomeNextRow(HomeNextRow.Kind.Talents, title, res.getString(R.string.home_next_talentsNoBuild), null)
            else -> HomeNextRow(HomeNextRow.Kind.Talents, title, res.getString(R.string.home_next_talentsNone), null)
        }
    }

    /** The newest unseen finished server run for this spec (`/v1/me` `finished_sims`, newest first). */
    private fun simsRow(spec: String?): HomeNextRow {
        val sim = finishedSims.firstOrNull { it.spec == spec }
        val line = sim?.let {
            res.getString(R.string.home_next_simFinished, character.specName.orEmpty(), it.dps.roundToInt().toString())
        } ?: res.getString(R.string.home_next_noSim)
        return HomeNextRow(HomeNextRow.Kind.Sims, res.getString(R.string.home_next_sims), line, null)
    }

    companion object {
        const val VISIBLE_ROWS = 5
        private const val STATE_OK = "ok"

        /** `60 · raid` for the raid preset, the band label alone otherwise. */
        fun bandHead(home: CharacterHome, res: Resources): String? {
            val label = home.bandLabel ?: return null
            return if (home.preset == "raid") res.getString(R.string.home_upgrades_bandRaid, label) else label
        }

        /**
         * `Replaces Belt of the Archmage, ilvl 62`; the server's own label for an item our data
         * does not know (`Not in our item data yet`); `Empty slot` when nothing is worn.
         */
        fun replaces(worn: HomeWorn?, res: Resources): String {
            worn ?: return res.getString(R.string.home_upgrades_empty)
            val name = worn.name
            val level = worn.itemLevel
            if (!worn.known || name == null || level == null) return worn.label
            return res.getString(R.string.home_upgrades_replaces, name, level)
        }

        /** `night_elf` → `Night Elf`; the API's slugs are lowercase. */
        fun words(slug: String): String =
            slug.split('_', '-', ' ').filter { it.isNotEmpty() }.joinToString(" ") { it.replaceFirstChar(Char::uppercaseChar) }

        /** A tooltip for an item on Home or Your gear, in the answer's build, spec and band. */
        fun tooltipRequest(
            look: ItemLook,
            home: CharacterHome,
            character: HomeCharacter,
            sourceLine: String? = null,
            wornId: Int? = null,
        ) = ItemTooltipRequest(
            itemId = look.id, fallback = look, build = home.dataBuild, classSlug = classSlug(home, character),
            spec = home.spec, specName = character.specName, band = home.band, bandLabel = home.bandLabel,
            preset = home.preset, faction = character.faction, sourceLine = sourceLine, wornId = wornId,
        )

        /** The items file to read: the class part of the spec slug (`hunter-beast-mastery`), else the character's class. */
        fun classSlug(home: CharacterHome, character: HomeCharacter): String =
            home.spec?.substringBefore('-') ?: character.className.orEmpty()
    }
}
