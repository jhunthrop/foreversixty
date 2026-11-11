package gg.foreversixty.app.gear

import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.LootFile

/**
 * Where an item comes from (`/data/{build}/loot.json`): the source's kind and name, the drop chance when
 * known, the boss, a crafted source's profession and a reputation source's standing.
 */
data class ItemSource(
    val kind: String,
    val name: String,
    val chance: Double?,
    val boss: String? = null,
    val profession: String? = null,
    val standing: String? = null,
    /** A dungeon's or raid's level range, the client's Group Finder row ("15-24"). */
    val levels: IntRange? = null,
    /** A dungeon's or raid's icon. */
    val icon: String? = null,
    /** A crafted item's profession skill (`required_skill`), when the loot file carries it. */
    val skill: Int? = null,
)

/** One upgrade the player can go and get: the item, its gain over what is worn, and where it comes from. */
data class RunPick(val slot: String, val itemId: Int, val itemName: String, val gain: Double, val source: ItemSource)

/**
 * Where to go tonight (spec 2026-11-08 F5), computed on the device from the cached list and loot files,
 * the way the web's `lib/bis` reads them: every candidate of a slot (the pick and its runners-up) with a
 * positive gain over the worn item, placed at its sources. Home's Best run line is the best dungeon pick.
 */
object WhereToGo {
    const val DUNGEON = "dungeon"
    const val RAID = "raid"

    /** Every item's sources, from the loot file's instances (their bosses and trash) and its other sources. */
    fun sources(loot: LootFile): Map<Int, List<ItemSource>> {
        val out = linkedMapOf<Int, MutableList<ItemSource>>()
        fun add(id: Int, source: ItemSource) {
            out.getOrPut(id) { mutableListOf() }.add(source)
        }
        loot.sources.forEach { source ->
            val levels = source.levelMin?.let { min -> source.levelMax?.let { max -> min..max } }
            val place = ItemSource(source.kind, source.name, null, levels = levels, icon = source.icon)
            source.bosses.orEmpty().forEach { boss ->
                boss.items.forEach { id ->
                    add(id, place.copy(chance = chance(boss.itemChances, id) ?: chance(source.itemChances, id), boss = boss.name.ifBlank { null }))
                }
            }
            source.trash.orEmpty().forEach { id -> add(id, place.copy(chance = chance(source.itemChances, id))) }
            source.items.orEmpty().forEach { id ->
                add(
                    id,
                    ItemSource(
                        source.kind, source.name, chance(source.itemChances, id),
                        profession = source.profession, standing = source.standing, skill = source.requiredSkill?.get(id.toString()),
                    ),
                )
            }
        }
        return out
    }

    /**
     * The candidates worth a trip: for each upgrade the server measured, the slot's pick and runners-up
     * whose gain over the worn item stays positive, at every source of theirs. Largest gain first.
     */
    fun picks(upgrades: List<HomeUpgrade>, slots: List<BisSlot>, sources: Map<Int, List<ItemSource>>): List<RunPick> {
        val bySlot = slots.associateBy { it.slot }
        return upgrades.flatMap { upgrade ->
            val gain = upgrade.gain?.takeIf { upgrade.modelled } ?: return@flatMap emptyList()
            val slot = bySlot[upgrade.slot] ?: return@flatMap emptyList()
            candidates(slot, gain).flatMap { (id, name, itemGain) ->
                sources[id].orEmpty().map { RunPick(upgrade.slot, id, name, itemGain, it) }
            }
        }.sortedByDescending { it.gain }
    }

    /** Home's Best run: the dungeon pick with the largest gain, or null when no upgrade drops in a dungeon. */
    fun bestRun(picks: List<RunPick>): RunPick? = picks.filter { it.source.kind == DUNGEON }.maxByOrNull { it.gain }

    private fun candidates(slot: BisSlot, pickGain: Double): List<Triple<Int, String, Double>> {
        val pick = slot.itemId?.takeIf { it > 0 }?.let { Triple(it, slot.itemName.orEmpty(), pickGain) }
        val alternatives = slot.alternatives.orEmpty().map { Triple(it.itemId, it.itemName, pickGain + it.dpsDelta) }
        return listOfNotNull(pick).plus(alternatives).filter { it.third > 0 }
    }

    private fun chance(chances: Map<String, Double>?, id: Int): Double? = chances?.get(id.toString())
}
