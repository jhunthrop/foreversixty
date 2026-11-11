package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import androidx.annotation.StringRes
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidView

/**
 * The first role under the raid's comp target in tank, healer, DPS order (round 4 s1, the server's
 * `notify.ShortRoles` rule: the role an officer can fill, not the largest raw gap, so Needs you and the
 * raid.short push's Mention name the same role): its wire role, its group label (`Healers`), its
 * lower-case word (`healers`) and how many it is short.
 */
data class ShortRole(val wire: String, @param:StringRes val group: Int, @param:StringRes val lower: Int, val missing: Int) {
    companion object {
        const val TANK = "tank"
        const val HEALER = "healer"
        const val DPS = "dps"

        /** Null when the raid has no comp target or every role is met. */
        fun of(raid: RaidView): ShortRole? {
            val target = raid.compTarget ?: return null
            val c = ComingCounts.of(raid)
            return listOf(
                ShortRole(TANK, R.string.raid_group_tanks, R.string.home_officer_role_tanks, target.tanks - c.tanks),
                ShortRole(HEALER, R.string.raid_group_healers, R.string.home_officer_role_healers, target.healers - c.healers),
                ShortRole(DPS, R.string.raid_group_damage, R.string.home_officer_role_damage, target.melee + target.ranged - c.melee - c.ranged),
            ).firstOrNull { it.missing > 0 }
        }

        /** A sign-up's role group: tank, healer, else damage (melee and ranged). */
        fun group(role: String): String = when (role) {
            TANK, HEALER -> role
            else -> DPS
        }

        /** `2 healers`, `1 tank`, `3 DPS`. */
        fun people(res: Resources, role: String?, n: Int): String = when (role) {
            TANK -> res.getQuantityString(R.plurals.raids_part_tanks, n, n)
            HEALER -> res.getQuantityString(R.plurals.raids_part_healers, n, n)
            else -> res.getString(R.string.raids_part_dps, n)
        }
    }
}
