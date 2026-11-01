package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeItem
import gg.foreversixty.app.model.HomeUpgrade
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.paperdoll.Paperdoll

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
    /** An empty off hand because the main hand is two-handed: the Paperdoll says `Two-hand in main hand`. */
    val offHandBlocked: Boolean = false,
)

/**
 * Builds the 17 rows from the home answer: an upgrade row names worn and pick, an already-best
 * row names the one item, and any other slot shows the export's worn item looked up in the
 * build's items file. Nothing is ranked or compared here.
 */
object GearModel {
    /** [offHandBlockedBy] names the two-hander that leaves the off hand empty (see [offHandBlocker]). */
    fun rows(
        home: CharacterHome,
        gear: Map<String, Int>,
        items: Map<Int, Item>,
        res: Resources,
        offHandBlockedBy: String? = null,
    ): List<GearRow> {
        val upgrades = home.upgrades.associateBy { it.slot }
        val best = home.alreadyBest.associateBy { it.slot }
        return ItemVocabulary.frameSlots.map { slot ->
            val name = ItemVocabulary.slotName(slot)?.let(res::getString) ?: slot
            upgrades[slot]?.let { return@map upgradeRow(it, name, res) }
            best[slot]?.let { return@map bestRow(it.item, slot, name, res) }
            if (slot == OFF_HAND && gear[slot] == null && offHandBlockedBy != null) {
                return@map GearRow(
                    slot, name, null, res.getString(R.string.items_offHandTwoHand, offHandBlockedBy), null, null, isBestInSlot = false,
                    offHandBlocked = true,
                )
            }
            exportRow(gear[slot], slot, name, items, res)
        }
    }

    private const val OFF_HAND = "off_hand"
    private const val MAIN_HAND = "main_hand"
    private const val TWO_HAND_EQUIPPED = "two_hand_equipped"

    /**
     * The two-hander that leaves the off hand empty, as the Discord panel words it: the worn main
     * hand when the items file says it is two-handed, else the band's main-hand pick when the
     * band's off hand is empty for that reason (`two_hand_equipped`).
     */
    fun offHandBlocker(home: CharacterHome, gear: Map<String, Int>, items: Map<Int, Item>, band: BisBand?): String? {
        if (gear[OFF_HAND] != null) return null
        val wornMain = home.upgrades.firstOrNull { it.slot == MAIN_HAND }?.worn?.id
            ?: home.alreadyBest.firstOrNull { it.slot == MAIN_HAND }?.item?.id
            ?: gear[MAIN_HAND]
        wornMain?.let { items[it] }?.takeIf { it.twoHand }?.let { return it.name }
        band ?: return null
        if (band.slots.firstOrNull { it.slot == OFF_HAND }?.emptyReason != TWO_HAND_EQUIPPED) return null
        val main = band.slots.firstOrNull { it.slot == MAIN_HAND } ?: return null
        val id = main.itemId ?: return null
        return items[id]?.name ?: main.itemName
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
            // An unknown item carries no name, so it reads `{Slot} · Item {id}` everywhere it is drawn.
            worn = ItemLook(worn.id, if (worn.known) worn.name else null, worn.quality, worn.icon, worn.itemLevel),
            wornName = if (worn.known) worn.name ?: worn.label else Paperdoll.unknownName(slotName, worn.id, res),
            wornLevel = worn.itemLevel?.let { res.getString(R.string.items_ilvl, it) },
            pick = pick,
            isBestInSlot = false,
        )
    }

    private fun bestRow(item: HomeItem, slot: String, slotName: String, res: Resources) = GearRow(
        slot = slot,
        slotName = slotName,
        worn = ItemLook(item.id, item.name, item.quality, item.icon, item.itemLevel),
        wornName = item.name ?: Paperdoll.unknownName(slotName, item.id, res),
        wornLevel = item.itemLevel?.let { res.getString(R.string.items_ilvl, it) },
        pick = null,
        isBestInSlot = true,
    )

    private fun exportRow(id: Int?, slot: String, slotName: String, items: Map<Int, Item>, res: Resources): GearRow {
        id ?: return GearRow(slot, slotName, null, res.getString(R.string.gear_emptySlot), null, null, isBestInSlot = false)
        val item = items[id] ?: return GearRow(
            slot, slotName, ItemLook(id), Paperdoll.unknownName(slotName, id, res), null, null, isBestInSlot = false,
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
