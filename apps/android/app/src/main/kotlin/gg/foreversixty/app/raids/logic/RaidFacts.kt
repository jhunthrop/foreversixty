package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidView
import gg.foreversixty.app.model.RaidViewerCharacter
import java.time.Instant

/** A sign-up answer, as the API spells it. */
enum class SignupState(val wire: String, @param:StringRes val label: Int) {
    In("in", R.string.raid_state_in),
    Late("late", R.string.raid_state_late),
    Bench("bench", R.string.raid_state_bench),
    Out("out", R.string.raid_state_out),
    ;

    /** In, Late and Bench carry a spec; Out never asks (spec F4.3). */
    val carriesSpec: Boolean get() = this != Out

    /** In and Late draw kill green when checked; Bench and Out draw warm. */
    val isComing: Boolean get() = this == In || this == Late

    companion object {
        fun of(wire: String?): SignupState? = entries.firstOrNull { it.wire == wire }
    }
}

/** The raid's chip (spec F4.2): its status, or the derived Ended once the end has passed. */
enum class RaidChip(@param:StringRes val label: Int) {
    Planned(R.string.raid_status_planned),
    Locked(R.string.raid_status_locked),
    Cancelled(R.string.raid_status_cancelled),
    Ended(R.string.raid_status_ended),
    Done(R.string.raid_status_done),
}

@StringRes
fun roleLabel(role: String): Int = when (role) {
    "tank" -> R.string.raid_role_tank
    "healer" -> R.string.raid_role_healer
    else -> R.string.raid_role_dps
}

/** The game's class order (the roster, the classes strip) and each class's names. */
enum class GameClass(@param:StringRes val label: Int, @param:StringRes val none: Int) {
    Warrior(R.string.raids_class_warrior, R.string.raids_classNone_warrior),
    Paladin(R.string.raids_class_paladin, R.string.raids_classNone_paladin),
    Hunter(R.string.raids_class_hunter, R.string.raids_classNone_hunter),
    Rogue(R.string.raids_class_rogue, R.string.raids_classNone_rogue),
    Priest(R.string.raids_class_priest, R.string.raids_classNone_priest),
    Shaman(R.string.raids_class_shaman, R.string.raids_classNone_shaman),
    Mage(R.string.raids_class_mage, R.string.raids_classNone_mage),
    Warlock(R.string.raids_class_warlock, R.string.raids_classNone_warlock),
    Druid(R.string.raids_class_druid, R.string.raids_classNone_druid),
    ;

    val slug: String get() = name.lowercase()

    companion object {
        fun of(slug: String): GameClass? = entries.firstOrNull { it.slug == slug.lowercase() }
        fun rank(slug: String): Int = of(slug)?.ordinal ?: entries.size
    }
}

/** The planned row's one chip. */
sealed interface RowChip {
    data class Status(val chip: RaidChip) : RowChip
    data class Signed(val state: SignupState, val role: String) : RowChip
    data object NotSigned : RowChip
}

/**
 * Read-only facts of one raid view at an instant. Nothing here re-derives a rule the API owns: it
 * orders, counts what the view lists, and picks the copy.
 */
class RaidFacts(val raid: RaidView, val now: Instant) {
    val startsAt: Instant = Instant.parse(raid.startsAt)
    val inviteAt: Instant? = raid.inviteAt?.let(Instant::parse)
    val endsAt: Instant = startsAt.plusSeconds(raid.durationMinutes * SECONDS_PER_MINUTE)

    val chip: RaidChip = when (raid.status) {
        "cancelled" -> RaidChip.Cancelled
        "done" -> RaidChip.Done
        else -> when {
            !now.isBefore(endsAt) -> RaidChip.Ended
            raid.status == "locked" -> RaidChip.Locked
            else -> RaidChip.Planned
        }
    }

    /** Planned or locked, and not ended: a raid the Raids tab offers (spec F4.1). */
    val isSignable: Boolean get() = chip == RaidChip.Planned || chip == RaidChip.Locked

    private val can: Set<String> = raid.viewer.can.orEmpty().toSet()
    val canLock: Boolean get() = "lock" in can
    val canMove: Boolean get() = "move" in can
    val showsOfficerTools: Boolean get() = canLock || canMove || raid.viewer.toolsEnded == true

    /** The countdown runs to the invite time when there is one, else to the pull. */
    val countdownTarget: Instant get() = inviteAt ?: startsAt

    /** Signed up counts In plus Late; Bench and Out are not counted. */
    val signedUp: Int get() = raid.counts.signedIn + raid.counts.late

    fun countLine(res: Resources): String = raid.size?.let { res.getString(R.string.raid_count, signedUp, it) }
        ?: res.getString(R.string.raid_countNoSize, signedUp)

    /** The bar's fill, 0 to 1; null (no bar) while the size is unknown. */
    val fill: Float? get() = raid.size?.takeIf { it > 0 }?.let { (signedUp.toFloat() / it).coerceAtMost(1f) }

    /** `24/40`, beside the roster heading. */
    val ratio: String get() = raid.size?.let { "$signedUp/$it" } ?: "$signedUp"

    /** The breakdown's parts; the late clause only when someone is late. */
    fun breakdownParts(res: Resources): List<String> {
        val counts = raid.counts
        val parts = listOf(
            res.getQuantityString(R.plurals.raids_part_tanks, counts.tanks, counts.tanks),
            res.getQuantityString(R.plurals.raids_part_healers, counts.healers, counts.healers),
            res.getString(R.string.raids_part_dps, counts.dps),
        )
        return if (counts.late > 0) parts + res.getString(R.string.raids_part_late, counts.late) else parts
    }

    /** `3 tanks · 6 healers · 14 DPS · 1 late`. */
    fun breakdown(res: Resources): String = breakdownParts(res).joinToString(SEPARATOR)

    /** The planned row's chip: Cancelled, Locked, Ended, else the viewer's own answer. */
    fun rowChip(currentKey: String?): RowChip = when (chip) {
        RaidChip.Cancelled -> RowChip.Status(RaidChip.Cancelled)
        RaidChip.Locked -> RowChip.Status(RaidChip.Locked)
        RaidChip.Ended, RaidChip.Done -> RowChip.Status(RaidChip.Ended)
        RaidChip.Planned -> viewerCharacters(currentKey).firstNotNullOfOrNull { c ->
            SignupState.of(c.state)?.let { RowChip.Signed(it, c.role) }
        } ?: RowChip.NotSigned
    }

    /** The viewer's characters, the current one first, then in the view's order. */
    fun viewerCharacters(currentKey: String?): List<RaidViewerCharacter> {
        val characters = raid.viewer.characters
        return characters.filter { it.characterKey == currentKey } + characters.filter { it.characterKey != currentKey }
    }

    companion object {
        const val SEPARATOR = " · "
        private const val SECONDS_PER_MINUTE = 60L
    }
}

/** The Raids tab's raids: the next signable one, and the rest of the planned list. */
class PlannedRaids(planned: List<RaidView>, now: Instant) {
    private val sorted = planned.sortedBy { Instant.parse(it.startsAt) }
    val next: RaidView? = sorted.firstOrNull { RaidFacts(it, now).isSignable }
    val also: List<RaidView> = sorted.filter { it.id != next?.id }
}
