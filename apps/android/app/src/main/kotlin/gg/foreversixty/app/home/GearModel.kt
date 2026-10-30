package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeItem
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.Item

/** One of the 17 rows of Your gear (spec F2.7): what is worn and the band's pick. */
data class GearRow(
    val slot: String,
    val slotName: String,
    /** Null for an empty slot. */
    val worn: ItemLook?,
    /** The worn item's name, or the server's words for one our data does not know. */
    val wornName: String,
    val wornLevel: String?,
    /** The pick when it differs from what is worn. */
    val pick: ItemLook?,
    val isBestInSlot: Boolean,
)

/**
 * Builds the 17 rows from the home answer: an upgrade row names worn and pick, an already-best
 * row names the one item, and any other slot shows the export's worn item looked up in the
 * build's items file. Nothing is ranked or compared here.
 */
object GearModel {
    fun rows(home: CharacterHome, gear: Map<String, Int>, items: Map<Int, Item>, res: Resources): List<GearRow> {
        val upgrades = home.upgrades.associateBy { it.slot }
        val best = home.alreadyBest.associateBy { it.slot }
        return ItemVocabulary.frameSlots.map { slot ->
            val name = ItemVocabulary.slotName(slot)?.let(res::getString) ?: slot
            upgrades[slot]?.let { return@map upgradeRow(it, name, res) }
            best[slot]?.let { return@map bestRow(it.item, slot, name, res) }
            exportRow(gear[slot], slot, name, items, res)
        }
    }

    private fun upgradeRow(upgrade: HomeUpgrade, slotName: String, res: Resources): GearRow {
        val p = upgrade.pick
        val pick = ItemLook(p.id, p.name, p.quality, p.icon, p.itemLevel)
        val worn = upgrade.worn ?: return GearRow(
            upgrade.slot, slotName, null, res.getString(R.string.gear_emptySlot), null, pick, isBestInSlot = false,
        )
        return GearRow(
            slot = upgrade.slot,
            slotName = slotName,
            worn = ItemLook(worn.id, worn.name, worn.quality, worn.icon, worn.itemLevel),
            wornName = if (worn.known) worn.name ?: worn.label else worn.label,
            wornLevel = worn.itemLevel?.let { res.getString(R.string.items_ilvl, it) },
            pick = pick,
            isBestInSlot = false,
        )
    }

    private fun bestRow(item: HomeItem, slot: String, slotName: String, res: Resources) = GearRow(
        slot = slot,
        slotName = slotName,
        worn = ItemLook(item.id, item.name, item.quality, item.icon, item.itemLevel),
        wornName = item.name ?: res.getString(R.string.items_notInData),
        wornLevel = item.itemLevel?.let { res.getString(R.string.items_ilvl, it) },
        pick = null,
        isBestInSlot = true,
    )

    private fun exportRow(id: Int?, slot: String, slotName: String, items: Map<Int, Item>, res: Resources): GearRow {
        id ?: return GearRow(slot, slotName, null, res.getString(R.string.gear_emptySlot), null, null, isBestInSlot = false)
        val item = items[id] ?: return GearRow(
            slot, slotName, ItemLook(id), res.getString(R.string.items_notInData), null, null, isBestInSlot = false,
        )
        return GearRow(
            slot = slot,
            slotName = slotName,
            worn = ItemLook(id, item.name, item.quality, item.icon, item.itemLevel),
            wornName = item.name,
            wornLevel = res.getString(R.string.items_ilvl, item.itemLevel),
            pick = null,
            isBestInSlot = false,
        )
    }
}
