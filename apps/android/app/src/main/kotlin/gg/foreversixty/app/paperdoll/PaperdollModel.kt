package gg.foreversixty.app.paperdoll

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.home.GearRow
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemQuality
import gg.foreversixty.app.items.ItemVocabulary
import gg.foreversixty.app.model.CharacterGearSlot
import gg.foreversixty.app.model.Item

/** One socket of the Paperdoll (spec W3 §4, design/specs/2026-10-31-paperdoll.md). */
data class PaperdollSlot(
    val slot: String,
    val slotName: String,
    /** The worn item; null for an empty slot. */
    val item: ItemLook?,
    /** The worn item's name, `Empty`, or the off hand's words under a two-hander. */
    val name: String,
    /** `Head · ilvl 61`, or the slot word alone with no item level. */
    val detail: String,
    /** The gold dot: the BiS list has a better item for this slot. */
    val better: Boolean,
) {
    val isEmpty: Boolean get() = item == null
}

/** The phone frame: two mirrored columns and the weapons in a row under them. */
data class PaperdollColumns(val left: List<PaperdollSlot>, val right: List<PaperdollSlot>, val weapons: List<PaperdollSlot>)

/**
 * The one Paperdoll model (tenet 10): Home's Your gear and the character page build the same 17
 * sockets, in the in-game frame's order, and the frame lays them out.
 */
object Paperdoll {
    /** The character frame's left column (shirt and tabard are not in the data). */
    val LEFT = listOf("head", "neck", "shoulder", "back", "chest", "wrist")
    val RIGHT = listOf("hands", "waist", "legs", "feet", "finger1", "finger2", "trinket1", "trinket2")
    val WEAPONS = listOf("main_hand", "off_hand", "ranged")

    private const val MAIN_HAND = "main_hand"
    private const val OFF_HAND = "off_hand"

    fun columns(slots: List<PaperdollSlot>): PaperdollColumns {
        val bySlot = slots.associateBy { it.slot }
        return PaperdollColumns(LEFT.mapNotNull(bySlot::get), RIGHT.mapNotNull(bySlot::get), WEAPONS.mapNotNull(bySlot::get))
    }

    /** Home's Your gear (F2.7): GearModel's rows as sockets, the gold dot where Home has an upgrade. */
    fun fromGear(rows: List<GearRow>, res: Resources): List<PaperdollSlot> = rows.map { row ->
        PaperdollSlot(
            slot = row.slot,
            slotName = row.slotName,
            item = row.worn,
            name = row.wornName,
            detail = detail(row.slotName, row.worn?.itemLevel, res),
            better = row.pick != null && row.pick.id != row.worn?.id,
        )
    }

    /**
     * A character card's gear (F11.5): each worn item looked up in its class's items file for its
     * icon and quality; the gold dot where the band's pick would gain points.
     */
    fun fromProfile(gear: List<CharacterGearSlot>, items: Map<Int, Item>, res: Resources): List<PaperdollSlot> {
        val bySlot = gear.associateBy { it.slot }
        val main = bySlot[MAIN_HAND]?.worn?.let { items[it.id] }
        return ItemVocabulary.frameSlots.map { slot ->
            val slotName = ItemVocabulary.slotLabel(res, slot)
            val entry = bySlot[slot]
            val worn = entry?.worn
            val blocked = slot == OFF_HAND && worn == null && main?.twoHand == true
            val item = worn?.let { w ->
                items[w.id]?.let { ItemLook(it.id, it.name, it.quality, it.icon, it.itemLevel) } ?: ItemLook(w.id, w.name, null, null, w.itemLevel)
            }
            PaperdollSlot(
                slot = slot,
                slotName = slotName,
                item = item,
                name = when {
                    blocked -> res.getString(R.string.items_offHandTwoHand, main.name)
                    item == null -> res.getString(R.string.gear_emptySlot)
                    else -> item.name ?: res.getString(R.string.items_notInData)
                },
                detail = detail(slotName, item?.itemLevel, res),
                better = entry.isBetter(),
            )
        }
    }

    /** A pick that is not what is worn and is not scored as no gain. */
    private fun CharacterGearSlot?.isBetter(): Boolean {
        val pick = this?.pick ?: return false
        val points = gap
        return pick.id != worn?.id && (points == null || points > 0.0)
    }

    private fun detail(slotName: String, level: Int?, res: Resources): String =
        level?.let { res.getString(R.string.paperdoll_detail, slotName, it) } ?: slotName

    /** `Head, Blue Dragonscale Helm, Epic, item level 61. A better item is on your best in slot list.` */
    fun label(slot: PaperdollSlot, res: Resources): String {
        val item = slot.item ?: return res.getString(R.string.paperdoll_a11y_empty, slot.slotName)
        val quality = ItemQuality(item.quality).name?.let(res::getString).orEmpty()
        val base = res.getString(R.string.paperdoll_a11y_socket, slot.slotName, slot.name, quality, item.itemLevel?.toString().orEmpty())
        return if (slot.better) "$base ${res.getString(R.string.paperdoll_a11y_better)}" else base
    }
}
