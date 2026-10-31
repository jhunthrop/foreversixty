package gg.foreversixty.app.raids.logic

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.model.RaidComposition

/**
 * The raid leader's read (coordinator ruling, player #5), drawn from the API's `composition` block.
 * The app only words it; every count and every present or missing comes from the API.
 */
data class CompositionRead(
    val classes: List<ClassLine>,
    /** `Tanks 3/4` against the raid's targets, else `Tanks 3`. */
    val roles: List<String>,
    val buffs: List<BuffLine>,
    /** `Missing: Blessings`, or `Every raid buff is covered.` */
    val buffSummary: String,
    val allCovered: Boolean,
    /** `Classes coming: Warrior 5, Paladin 0, …` */
    val classesAccessibilityLabel: String,
) {
    data class ClassLine(val className: String, val name: String, val count: Int)
    data class BuffLine(val key: String, val name: String, val label: String, val icon: String, val present: Boolean, val accessibilityLabel: String)

    companion object {
        fun of(composition: RaidComposition?, res: Resources): CompositionRead? {
            composition ?: return null
            val classes = composition.classes.map { entry ->
                ClassLine(entry.className, GameClass.of(entry.className)?.let { res.getString(it.label) } ?: entry.className, entry.count)
            }
            val groups = listOf(R.string.raid_group_tanks, R.string.raid_group_healers, R.string.raid_group_dps).map(res::getString)
            val counts = composition.roles.let { listOf(it.tanks, it.healers, it.dps) }
            val targets = composition.targets?.let { listOf(it.tanks, it.healers, it.dps) }
            val roles = groups.indices.map { i ->
                if (targets != null) res.getString(R.string.raids_comp_role, groups[i], counts[i], targets[i])
                else res.getString(R.string.raids_comp_roleCount, groups[i], counts[i])
            }
            val buffs = composition.buffs.map { buff ->
                BuffLine(
                    buff.key, buff.name, buff.label, buff.icon, buff.present,
                    if (buff.present) res.getString(R.string.raids_a11y_buffPresent, buff.name, buff.count)
                    else res.getString(R.string.raids_a11y_buffMissing, buff.name),
                )
            }
            val missing = composition.buffs.filterNot { it.present }.map { it.name }
            return CompositionRead(
                classes = classes,
                roles = roles,
                buffs = buffs,
                buffSummary = if (missing.isEmpty()) res.getString(R.string.raids_comp_covered)
                else res.getString(R.string.raids_comp_missing, missing.joinToString(", ")),
                allCovered = missing.isEmpty(),
                classesAccessibilityLabel = res.getString(
                    R.string.raids_a11y_classes,
                    classes.joinToString(", ") { res.getString(R.string.raids_a11y_classCount, it.name, it.count) },
                ),
            )
        }
    }
}
