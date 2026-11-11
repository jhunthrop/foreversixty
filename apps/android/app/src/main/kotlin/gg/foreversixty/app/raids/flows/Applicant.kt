package gg.foreversixty.app.raids.flows

import gg.foreversixty.app.designsystem.Codes
import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.auth.session.RelativeAge
import gg.foreversixty.app.flows.FlowFigure
import gg.foreversixty.app.model.ApplicantSheet
import gg.foreversixty.app.model.GuildOfficerApplication
import gg.foreversixty.app.phrase
import gg.foreversixty.app.raids.logic.GameClass
import gg.foreversixty.app.raids.logic.RaidClock
import java.time.Instant
import java.util.Locale

/** A titled line of the applicant's page (Raid ready, Guild history). */
data class ApplicantLine(val title: String, val sub: String)

/** What the applicant's page draws (mock raid-applicant, -accepted). */
data class ApplicantContent(
    val name: String,
    val className: String?,
    val facts: List<String>,
    val figures: List<FlowFigure>,
    val ready: List<ApplicantLine>,
    val history: List<ApplicantLine>,
    val note: String?,
    /** After Accept: the line, the `/ginvite` command and what happens next; null before. */
    val accepted: AcceptedContent?,
)

data class AcceptedContent(val line: String, val command: String, val after: String)

/**
 * The applicant (spec F3): the numbers, raid readiness, the history and their note before Accept;
 * after Accept the `/ginvite` line. Formats the server's sheet; absent parts are left out.
 */
class ApplicantPresenter(private val res: Resources, private val clock: RaidClock, private val now: Instant) {
    fun content(app: GuildOfficerApplication, guild: String, accepted: Boolean): ApplicantContent {
        val character = app.character
        val sheet = app.sheet
        return ApplicantContent(
            name = character.name,
            className = character.className,
            facts = listOfNotNull(kindLine(app), origin(app), applied(app)),
            figures = figures(sheet, character.itemLevel),
            ready = sheet?.let(::ready).orEmpty(),
            history = sheet?.history.orEmpty().map {
                ApplicantLine(it.guildName, res.getString(R.string.raid_applicant_historyLine, res.getQuantityString(R.plurals.raid_applicant_raids, it.raids, it.raids), clock.shortDate(Instant.parse(it.lastRaidAt))))
            },
            note = app.message.takeIf { it.isNotBlank() },
            accepted = if (accepted) {
                AcceptedContent(
                    res.getString(R.string.raid_applicant_accepted, guild, character.name),
                    res.getString(R.string.raid_applicant_ginvite, character.name),
                    res.getString(R.string.raid_applicant_after, character.name),
                )
            } else {
                null
            },
        )
    }

    /** `Holy Priest · 60`. */
    private fun kindLine(app: GuildOfficerApplication): String? {
        val kind = kind(app) ?: return null
        val level = app.character.level ?: return kind
        return res.getString(R.string.raid_applicant_kind, kind, level)
    }

    private fun kind(app: GuildOfficerApplication): String? {
        val cls = app.character.className?.let(GameClass::of)?.let { res.getString(it.label) }
        val spec = app.character.specName?.takeIf { it.isNotBlank() }
        // The application's spec_name already names the class (`Holy Paladin`).
        if (spec != null && cls != null && spec.endsWith(cls)) return spec
        return listOfNotNull(spec, cls).joinToString(" ").ifEmpty { null }
    }

    /** `Alliance Dwarf`: the faction and the race, either alone when the other is unknown. */
    private fun origin(app: GuildOfficerApplication): String? {
        val faction = app.character.faction?.let {
            res.getString(if (it == HORDE) R.string.bis_faction_horde else R.string.bis_faction_alliance)
        }
        return listOfNotNull(faction, app.character.race?.takeIf { it.isNotBlank() }).joinToString(" ").ifEmpty { null }
    }

    private fun applied(app: GuildOfficerApplication): String =
        res.getString(R.string.raid_applicant_applied, RelativeAge.of(Instant.parse(app.createdAt), now).phrase(res))

    /** Home's and Needs you's applicant line: `Mirelle, Holy Priest 60, 2 min ago`. */
    fun line(app: GuildOfficerApplication): String = res.getString(
        R.string.raid_applicant_line,
        app.character.name,
        kind(app).orEmpty(),
        app.character.level ?: 0,
        RelativeAge.of(Instant.parse(app.createdAt), now).phrase(res),
    )

    private fun figures(sheet: ApplicantSheet?, itemLevel: Int?): List<FlowFigure> = listOfNotNull(
        sheet?.score?.let { FlowFigure(res.getString(R.string.raid_applicant_score), it.toString()) },
        sheet?.sim?.let { FlowFigure(res.getString(R.string.raid_applicant_sim), String.format(Locale.US, "%,.0f", it.value), Codes.unit(it.unit)) },
        itemLevel?.let { FlowFigure(res.getString(R.string.raid_applicant_itemLevel), it.toString()) },
    )

    private fun ready(sheet: ApplicantSheet): List<ApplicantLine> = listOfNotNull(
        sheet.attunement?.let { a ->
            when (a.state) {
                OK -> ApplicantLine(res.getString(R.string.raid_applicant_attuned, a.zoneName), a.item?.name.orEmpty())
                MISS -> ApplicantLine(res.getString(R.string.raid_applicant_notAttuned, a.zoneName), a.item?.name.orEmpty())
                else -> ApplicantLine(res.getString(R.string.raid_applicant_attuneUnknown, a.zoneName), "")
            }
        },
        sheet.enchants?.let { e ->
            val line = res.getString(R.string.raid_applicant_enchantsLine, e.enchanted, e.slots)
            val missing = e.missing.map(::slotWord)
            val sub = if (missing.isEmpty()) line else res.getString(R.string.raid_applicant_enchantsMissing, line, and(missing))
            ApplicantLine(res.getString(R.string.raid_applicant_enchants), sub)
        },
        sheet.bisGap?.let { g ->
            ApplicantLine(
                res.getString(R.string.raid_applicant_gap),
                res.getString(R.string.raid_applicant_gapLine, res.getQuantityString(R.plurals.raid_applicant_pieces, g.pieces, g.pieces), g.list),
            )
        },
    )

    private fun slotWord(slot: String): String = SLOTS[slot]?.let(res::getString) ?: slot.replace('_', ' ')

    private fun and(words: List<String>): String =
        if (words.size < 2) words.joinToString() else res.getString(R.string.raid_list_and, words.dropLast(1).joinToString(", "), words.last())

    private companion object {
        const val OK = "ok"
        const val MISS = "miss"
        const val HORDE = "horde"
        val SLOTS = mapOf(
            "back" to R.string.raid_applicant_slot_back,
            "feet" to R.string.raid_applicant_slot_feet,
            "hands" to R.string.raid_applicant_slot_hands,
            "wrist" to R.string.raid_applicant_slot_wrist,
            "chest" to R.string.raid_applicant_slot_chest,
            "legs" to R.string.raid_applicant_slot_legs,
            "shoulder" to R.string.raid_applicant_slot_shoulder,
            "head" to R.string.raid_applicant_slot_head,
            "main_hand" to R.string.raid_applicant_slot_weapon,
            "weapon" to R.string.raid_applicant_slot_weapon,
        )
    }
}
