package gg.foreversixty.app.items

import android.content.res.Resources
import gg.foreversixty.app.R

/**
 * The words every item surface shares: slot labels and quality names (spec §4.3). One place, so
 * the BiS list, the tooltip sheet and Home name an item the same way.
 */
object ItemWords {
    /**
     * The paperdoll label of a BiS or item slot key. Two rings read `Ring 1` and `Ring 2`,
     * two trinkets `Trinket 1` and `Trinket 2`, so the rows can be told apart.
     */
    fun slot(res: Resources, key: String): String = when (key) {
        "head" -> res.getString(R.string.items_slot_head)
        "neck" -> res.getString(R.string.items_slot_neck)
        "shoulder" -> res.getString(R.string.items_slot_shoulder)
        "back" -> res.getString(R.string.items_slot_back)
        "chest" -> res.getString(R.string.items_slot_chest)
        "wrist" -> res.getString(R.string.items_slot_wrist)
        "hands" -> res.getString(R.string.items_slot_hands)
        "waist" -> res.getString(R.string.items_slot_waist)
        "legs" -> res.getString(R.string.items_slot_legs)
        "feet" -> res.getString(R.string.items_slot_feet)
        "finger" -> res.getString(R.string.items_slot_finger)
        "finger1" -> res.getString(R.string.items_slot_finger1)
        "finger2" -> res.getString(R.string.items_slot_finger2)
        "trinket" -> res.getString(R.string.items_slot_trinket)
        "trinket1" -> res.getString(R.string.items_slot_trinket1)
        "trinket2" -> res.getString(R.string.items_slot_trinket2)
        "main_hand" -> res.getString(R.string.items_slot_mainHand)
        "off_hand" -> res.getString(R.string.items_slot_offHand)
        "ranged" -> res.getString(R.string.items_slot_ranged)
        else -> key
    }

    /** [quality] is 0 (poor) to 5 (legendary). */
    fun quality(res: Resources, quality: Int): String = res.getString(
        when (quality) {
            0 -> R.string.items_quality_poor
            1 -> R.string.items_quality_common
            2 -> R.string.items_quality_uncommon
            3 -> R.string.items_quality_rare
            4 -> R.string.items_quality_epic
            else -> R.string.items_quality_legendary
        },
    )
}
