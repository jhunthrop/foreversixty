package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidNotSigned
import gg.foreversixty.app.model.RaidSignup
import gg.foreversixty.app.model.RaidView
import java.time.Instant
import java.time.ZoneId

/** One roster row: a sign-up, or a member not signed up (state null). */
data class RosterRow(
    val characterKey: String,
    val name: String,
    val className: String,
    val spec: String,
    val specName: String,
    val specIcon: String?,
    val role: String,
    val state: SignupState?,
    val note: String,
    val own: Boolean,
) {
    companion object {
        fun of(signup: RaidSignup) = RosterRow(
            signup.characterKey, signup.name, signup.className, signup.spec, signup.specName.orEmpty(),
            signup.specIcon, signup.role, SignupState.of(signup.state), signup.note, signup.own,
        )

        fun of(member: RaidNotSigned, own: Boolean) = RosterRow(
            member.characterKey, member.name, member.className, member.spec, member.specName.orEmpty(),
            member.specIcon, member.role ?: "dps", null, "", own,
        )
    }
}

/** A roster group's kind and its copy. */
enum class RosterKind(@param:StringRes val label: Int, @param:StringRes val emptyLine: Int?) {
    Tanks(R.string.raid_group_tanks, R.string.raid_emptyRole_tank),
    Healers(R.string.raid_group_healers, R.string.raid_emptyRole_healer),
    Dps(R.string.raid_group_dps, R.string.raid_emptyRole_dps),
    Late(R.string.raid_group_late, null),
    Bench(R.string.raid_group_bench, null),
    Out(R.string.raid_group_out, null),
}

data class RosterGroup(val kind: RosterKind, val rows: List<RosterRow>)

/**
 * The roster's groups (spec F4.2 item 8; the site's raid-roster.ts): In by role, then Late, Bench,
 * Out and Not signed up. Within a group: class in the game's order, then name.
 */
class RosterGroups(raid: RaidView, overrides: Map<String, SignupState?> = emptyMap()) {
    val groups: List<RosterGroup>
    val notSigned: List<RosterRow>

    init {
        val ownKeys = raid.viewer.characters.map { it.characterKey }.toSet()
        val base = raid.signups.map(RosterRow::of) + raid.notSigned.map { RosterRow.of(it, it.characterKey in ownKeys) }
        val rows = base.map { row -> if (row.characterKey in overrides) row.copy(state = overrides[row.characterKey]) else row }
            .sortedWith(ORDER)
        fun inState(state: SignupState, role: String? = null) =
            rows.filter { it.state == state && (role == null || it.role == role) }
        groups = listOf(
            RosterGroup(RosterKind.Tanks, inState(SignupState.In, "tank")),
            RosterGroup(RosterKind.Healers, inState(SignupState.In, "healer")),
            RosterGroup(RosterKind.Dps, inState(SignupState.In, "dps")),
            RosterGroup(RosterKind.Late, inState(SignupState.Late)),
            RosterGroup(RosterKind.Bench, inState(SignupState.Bench)),
            RosterGroup(RosterKind.Out, inState(SignupState.Out)),
        ).filter { it.rows.isNotEmpty() || it.kind.emptyLine != null }
        notSigned = rows.filter { it.state == null }
    }

    companion object {
        val ORDER: Comparator<RosterRow> = compareBy<RosterRow> { GameClass.rank(it.className) }.thenBy { it.name }
    }
}

/** The classes strip (spec F4.2 item 6): In plus Late per class, in the game's order. */
class ClassStrip(raid: RaidView) {
    val entries: List<Pair<GameClass, Int>> = run {
        val coming = raid.signups.filter { SignupState.of(it.state)?.isComing == true }
        GameClass.entries.map { cls -> cls to coming.count { it.className.lowercase() == cls.slug } }
    }

    /** `Classes coming: Warrior 4, Hunter 2, …, no Paladins`. */
    fun accessibilityLabel(res: Resources): String {
        val present = entries.filter { it.second > 0 }
            .map { res.getString(R.string.raids_a11y_classCount, res.getString(it.first.label), it.second) }
        val missing = entries.filter { it.second == 0 }.map { res.getString(it.first.none) }
        return res.getString(R.string.raids_a11y_classes, (present + missing).joinToString(", "))
    }
}

/** The officer's Copy the list (the site's copyList). */
object RosterCopy {
    fun text(raid: RaidView, clock: RaidClock, res: Resources, now: Instant): String {
        val facts = RaidFacts(raid, now)
        val guildZone = runCatching { ZoneId.of(raid.guildTimezone) }.getOrDefault(clock.zone)
        val header = listOf(
            raid.title,
            res.getString(R.string.raids_when_date, clock.shortDate(facts.startsAt), clock.guildClock(facts.startsAt, guildZone)),
            "<t:${facts.startsAt.epochSecond}:F>",
        ).joinToString(RaidFacts.SEPARATOR)
        val size = raid.size?.let { res.getString(R.string.raids_copy_size, facts.signedUp, it) }
            ?: res.getString(R.string.raid_countNoSize, facts.signedUp)
        val groups = RosterGroups(raid).groups.filter { it.rows.isNotEmpty() }.map { group ->
            res.getString(R.string.raids_copy_group, res.getString(group.kind.label), group.rows.joinToString(", ") { it.name })
        }
        val count = res.getString(R.string.raids_copy_count, size, facts.breakdownParts(res).joinToString(", "))
        return (listOf(header, count) + groups).joinToString("\n")
    }
}
