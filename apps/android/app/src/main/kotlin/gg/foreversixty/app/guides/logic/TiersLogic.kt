package gg.foreversixty.app.guides.logic

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.TierList
import gg.foreversixty.app.model.TierRow
import gg.foreversixty.app.model.TiersFile

/** Which list the tier page shows: a role (`dps`, `tank`, `healer`) and a faction. */
data class TierChoice(val role: String, val faction: String) {
    companion object {
        const val DPS = "dps"
        const val TANK = "tank"
        const val HEALER = "healer"
        const val HORDE = "horde"
        const val ALLIANCE = "alliance"
        val roles = listOf(DPS, TANK, HEALER)

        /** The web's order: Alliance, then Horde. */
        val factions = listOf(ALLIANCE, HORDE)
    }
}

/** A tier letter's colour (the web's tier colours): S legendary, A epic, B rare, C uncommon, D muted. */
enum class TierTone { Legendary, Epic, Rare, Uncommon, Muted }

/**
 * One drawn row: the file's row, its spec and class names apart (`Survival`, `Hunter`), the specs it
 * ties with (by spec name) and whether it is your spec.
 */
data class TierRowModel(
    val row: TierRow,
    val specName: String,
    val className: String,
    val tone: TierTone?,
    val ties: List<String>,
    val isYours: Boolean,
)

/** A letter's group (`S  Within 5% of the top  1 spec`): [start] and [end] are its band in percent behind the top. */
data class TierGroup(val letter: String, val tone: TierTone, val start: Int, val end: Int?, val rows: List<TierRowModel>)

/** The head's answer for your spec: its row, its place of [count], the list's top row and [letter]. */
data class TierAnswer(val row: TierRowModel, val count: Int, val top: TierRow)

/** The tier page's rules. Order, gaps, letters and ties are the file's (the web's `rankRole`); nothing is computed here. */
object Tiers {
    private const val TIE_ABOVE = "above"
    private const val TIE_BELOW = "below"

    /** The letters with the percent behind the top each one starts at (the web's tier-rules.ts). */
    private val BANDS = listOf("S" to 0, "A" to 5, "B" to 10, "C" to 20, "D" to 30)

    /** The margin a tie is within, in percent (the web's TIE_PCT). */
    const val TIE_PERCENT = 1

    /** Where the letters S to C end, in percent behind the top: 5, 10, 20, 30. */
    val bandEnds: List<Int> get() = BANDS.drop(1).map { it.second }

    /** The current character's role (from its guide) and faction; no character: DPS and Alliance, as the site opens. */
    fun defaultChoice(index: GuideIndex, character: AccountCharacter?): TierChoice {
        val role = Guides.yourGuide(index, character)?.role?.takeIf { it in TierChoice.roles } ?: TierChoice.DPS
        val faction = character?.faction?.takeIf { it in TierChoice.factions } ?: TierChoice.ALLIANCE
        return TierChoice(role, faction)
    }

    fun list(file: TiersFile, choice: TierChoice): TierList {
        val roles = if (choice.faction == TierChoice.ALLIANCE) file.factions.alliance else file.factions.horde
        return when (choice.role) {
            TierChoice.TANK -> roles.tank
            TierChoice.HEALER -> roles.healer
            else -> roles.dps
        }
    }

    /** How many specs each role's list has for [faction] (the role segments' counts). */
    fun counts(file: TiersFile, faction: String): Map<String, Int> =
        TierChoice.roles.associateWith { list(file, TierChoice(it, faction)).rows.size }

    fun rows(list: TierList, index: GuideIndex?, yourSpec: String?): List<TierRowModel> = list.rows.mapIndexed { i, row ->
        val spec = index?.let { Guides.spec(it, row.spec) }
        val className = index?.let { Guides.classOf(it, row.spec) }?.className ?: row.name.substringAfterLast(' ')
        val specName = spec?.specName ?: row.name.removeSuffix(" $className")
        val ties = listOfNotNull(
            list.rows.getOrNull(i - 1)?.takeIf { row.tie == TIE_ABOVE },
            list.rows.getOrNull(i + 1)?.takeIf { row.tie == TIE_BELOW },
        ).map { other -> index?.let { Guides.spec(it, other.spec)?.specName } ?: other.name }
        TierRowModel(row, specName, className, tone(row.tier), ties, isYours = row.spec == yourSpec)
    }

    /** The lettered lists in letter groups, empty letters left out; a list without letters (tank, healer) has none. */
    fun groups(rows: List<TierRowModel>): List<TierGroup> {
        if (rows.any { it.row.tier == null }) return emptyList()
        return BANDS.mapIndexedNotNull { i, (letter, start) ->
            val members = rows.filter { it.row.tier == letter }
            val tone = tone(letter)
            if (members.isEmpty() || tone == null) null else TierGroup(letter, tone, start, BANDS.getOrNull(i + 1)?.second, members)
        }
    }

    fun answer(rows: List<TierRowModel>): TierAnswer? {
        val yours = rows.firstOrNull { it.isYours } ?: return null
        return TierAnswer(yours, rows.size, rows.first().row)
    }

    fun tone(letter: String?): TierTone? = when (letter) {
        "S" -> TierTone.Legendary
        "A" -> TierTone.Epic
        "B" -> TierTone.Rare
        "C" -> TierTone.Uncommon
        "D" -> TierTone.Muted
        else -> null
    }
}
