package gg.foreversixty.app.raids.logic

import gg.foreversixty.app.model.RaidView

/**
 * Who is coming by role (round 3, row 1): the one count model Home's raid card, the Raid tab's Who is
 * coming and the short role read. Counted from the sign-ups answering In or Late, the same people the
 * names list, so `Damage` always equals melee plus ranged and every surface reads the same figure.
 */
data class ComingCounts(val tanks: Int, val healers: Int, val melee: Int, val ranged: Int) {
    val damage: Int get() = melee + ranged

    companion object {
        const val MELEE = "melee"
        const val RANGED = "ranged"

        fun of(raid: RaidView): ComingCounts {
            val groups = raid.signups.filter { SignupState.of(it.state)?.isComing == true }.groupingBy { groupOf(it.role, it.group) }.eachCount()
            return ComingCounts(
                tanks = groups[ShortRole.TANK] ?: 0,
                healers = groups[ShortRole.HEALER] ?: 0,
                melee = groups[MELEE] ?: 0,
                ranged = groups[RANGED] ?: 0,
            )
        }

        /** Tank and healer by role; damage by its group, ranged when the server sends no melee group. */
        fun groupOf(role: String, group: String): String = when (role) {
            ShortRole.TANK, ShortRole.HEALER -> role
            else -> if (group == MELEE) MELEE else RANGED
        }
    }
}
