package gg.foreversixty.app.guides.logic

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.GuideIndex
import gg.foreversixty.app.model.TierList
import gg.foreversixty.app.model.TierRow
import gg.foreversixty.app.model.TiersFile

/** Which list the tier segment shows: a role (`dps`, `tank`, `healer`) and a faction. */
data class TierChoice(val role: String, val faction: String) {
    companion object {
        const val DPS = "dps"
        const val TANK = "tank"
        const val HEALER = "healer"
        const val HORDE = "horde"
        const val ALLIANCE = "alliance"
        val roles = listOf(DPS, TANK, HEALER)
        val factions = listOf(HORDE, ALLIANCE)
    }
}

/** A tier letter's colour (spec §4 tier chip): S legendary, A epic, B rare, C uncommon, D muted. */
enum class TierTone { Legendary, Epic, Rare, Uncommon, Muted }

/** One drawn row: the file's row, its gap words (`≈ −7.5%` on a tie) and whether it is your spec. */
data class TierRowModel(val row: TierRow, val gap: String, val tone: TierTone?, val isYours: Boolean)

/** The callout (spec F9.4): your spec's row and the list's race line under it. */
data class TierCallout(val row: TierRow, val note: String?)

/** F9.4's rules. Order, gaps and letters are the file's (the web's `rankRole`); nothing is computed here. */
object Tiers {
    private const val TIE_MARK = "≈"

    /** The current character's role (from its guide) and faction; no character: DPS and Horde. */
    fun defaultChoice(index: GuideIndex, character: AccountCharacter?): TierChoice {
        val role = Guides.yourGuide(index, character)?.role?.takeIf { it in TierChoice.roles } ?: TierChoice.DPS
        val faction = character?.faction?.takeIf { it in TierChoice.factions } ?: TierChoice.HORDE
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

    fun rows(list: TierList, yourSpec: String?): List<TierRowModel> = list.rows.map { row ->
        val gap = if (row.tie != null) "$TIE_MARK ${row.gapText}" else row.gapText
        TierRowModel(row, gap, tone(row.tier), isYours = row.spec == yourSpec)
    }

    fun callout(list: TierList, yourSpec: String?): TierCallout? {
        val row = list.rows.firstOrNull { it.spec == yourSpec } ?: return null
        return TierCallout(row, list.notes.lastOrNull())
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
