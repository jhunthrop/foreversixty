package gg.foreversixty.app.items

import android.content.res.Resources
import androidx.annotation.StringRes
import androidx.compose.ui.graphics.Color
import gg.foreversixty.app.R
import gg.foreversixty.app.designsystem.FsColor
import java.util.Locale
import kotlin.math.abs

/**
 * The game's words for an item, from the copy catalog: quality, slot, stat and standing names.
 * One table per word family, so the tooltip, Home and the gear list all say the same thing.
 */
object ItemVocabulary {
    /** The 17 character-frame slots, in the in-game frame's order (spec F2.7). */
    val frameSlots = listOf(
        "head", "neck", "shoulder", "back", "chest", "wrist", "hands", "waist", "legs", "feet",
        "finger1", "finger2", "trinket1", "trinket2", "main_hand", "off_hand", "ranged",
    )

    private val slotNames = mapOf(
        "head" to R.string.items_slot_head, "neck" to R.string.items_slot_neck,
        "shoulder" to R.string.items_slot_shoulder, "back" to R.string.items_slot_back,
        "chest" to R.string.items_slot_chest, "wrist" to R.string.items_slot_wrist,
        "hands" to R.string.items_slot_hands, "waist" to R.string.items_slot_waist,
        "legs" to R.string.items_slot_legs, "feet" to R.string.items_slot_feet,
        "finger1" to R.string.items_slot_finger1, "finger2" to R.string.items_slot_finger2,
        "trinket1" to R.string.items_slot_trinket1, "trinket2" to R.string.items_slot_trinket2,
        "main_hand" to R.string.items_slot_mainHand, "off_hand" to R.string.items_slot_offHand,
        "ranged" to R.string.items_slot_ranged, "finger" to R.string.items_slot_finger,
        "trinket" to R.string.items_slot_trinket,
    )

    @StringRes
    fun slotName(slot: String): Int? = slotNames[slot]

    /**
     * A slot's label: two fingers read `Finger 1` and `Finger 2`, two trinkets `Trinket 1` and
     * `Trinket 2`, so the BiS rows can be told apart; an unknown key prints as published.
     */
    fun slotLabel(res: Resources, slot: String): String = slotName(slot)?.let(res::getString) ?: slot

    /** The stats a tooltip lists, in the site's order (`STAT_KEYS`, armor excluded: it has its own line). */
    val statOrder: List<Pair<String, Int>> = listOf(
        "strength" to R.string.items_stat_strength, "agility" to R.string.items_stat_agility,
        "stamina" to R.string.items_stat_stamina, "intellect" to R.string.items_stat_intellect,
        "spirit" to R.string.items_stat_spirit, "crit" to R.string.items_stat_crit, "hit" to R.string.items_stat_hit,
        "spell_power" to R.string.items_stat_spellPower, "healing" to R.string.items_stat_healing,
        "attack_power" to R.string.items_stat_attackPower,
        "ranged_attack_power" to R.string.items_stat_rangedAttackPower,
        "defense" to R.string.items_stat_defense, "dodge" to R.string.items_stat_dodge,
        "parry" to R.string.items_stat_parry, "block" to R.string.items_stat_block,
        "block_value" to R.string.items_stat_blockValue, "mp5" to R.string.items_stat_mp5,
        "spell_penetration" to R.string.items_stat_spellPenetration, "fire_power" to R.string.items_stat_firePower,
        "fire_res" to R.string.items_stat_fireRes, "frost_res" to R.string.items_stat_frostRes,
        "nature_res" to R.string.items_stat_natureRes, "shadow_res" to R.string.items_stat_shadowRes,
        "arcane_res" to R.string.items_stat_arcaneRes,
    )

    private val standings = mapOf(
        "hated" to R.string.items_standing_hated, "hostile" to R.string.items_standing_hostile,
        "unfriendly" to R.string.items_standing_unfriendly, "neutral" to R.string.items_standing_neutral,
        "friendly" to R.string.items_standing_friendly, "honored" to R.string.items_standing_honored,
        "revered" to R.string.items_standing_revered, "exalted" to R.string.items_standing_exalted,
    )

    @StringRes
    fun standingName(standing: String?): Int? = standings[standing]

    @StringRes
    fun bindsLine(bonding: String?): Int? = when (bonding) {
        "pickup" -> R.string.tooltip_bind_pickup
        "equip" -> R.string.tooltip_bind_equip
        "use" -> R.string.tooltip_bind_use
        "quest" -> R.string.items_bind_quest
        else -> null
    }

    private val subclasses = mapOf(
        "axe" to R.string.items_subclass_axe, "bow" to R.string.items_subclass_bow, "gun" to R.string.items_subclass_gun,
        "mace" to R.string.items_subclass_mace, "polearm" to R.string.items_subclass_polearm,
        "sword" to R.string.items_subclass_sword, "staff" to R.string.items_subclass_staff,
        "fist_weapon" to R.string.items_subclass_fistWeapon, "dagger" to R.string.items_subclass_dagger,
        "thrown" to R.string.items_subclass_thrown, "crossbow" to R.string.items_subclass_crossbow,
        "wand" to R.string.items_subclass_wand, "fishing_pole" to R.string.items_subclass_fishingPole,
    )

    @StringRes
    fun subclassName(subclass: String?): Int? = subclasses[subclass]

    private val weaponSlots = mapOf(
        "one_hand" to R.string.items_weaponSlot_oneHand, "two_hand" to R.string.items_weaponSlot_twoHand,
        "main_hand" to R.string.items_weaponSlot_mainHand, "off_hand" to R.string.items_weaponSlot_offHand,
        "ranged" to R.string.items_weaponSlot_ranged, "thrown" to R.string.items_weaponSlot_thrown,
    )

    @StringRes
    fun weaponSlotName(slot: String?): Int? = weaponSlots[slot]

    @StringRes
    fun armorTypeName(armorType: String?): Int? = when (armorType) {
        "cloth" -> R.string.tooltip_armorType_cloth
        "leather" -> R.string.tooltip_armorType_leather
        "mail" -> R.string.tooltip_armorType_mail
        "plate" -> R.string.tooltip_armorType_plate
        "shield" -> R.string.tooltip_armorType_shield
        else -> null
    }
}

/** An item's quality (0 poor to 5 legendary): its name, border colour and text colour. */
data class ItemQuality(val value: Int?) {
    @get:StringRes
    val name: Int?
        get() = when (value) {
            0 -> R.string.items_quality_poor
            1 -> R.string.items_quality_common
            2 -> R.string.items_quality_uncommon
            3 -> R.string.items_quality_rare
            4 -> R.string.items_quality_epic
            5 -> R.string.items_quality_legendary
            else -> null
        }

    /** The icon border (`--color-rarity-*`). */
    val border: Color
        get() = when (value) {
            0 -> FsColor.rarityPoor
            1 -> FsColor.rarityCommon
            2 -> FsColor.rarityUncommon
            3 -> FsColor.rarityRare
            4 -> FsColor.rarityEpic
            5 -> FsColor.rarityLegendary
            else -> FsColor.line
        }

    /** The name colour: the AA text variants for rare and epic, the border colour otherwise. */
    val text: Color
        get() = when (value) {
            3 -> FsColor.rarityRareText
            4 -> FsColor.rarityEpicText
            null -> FsColor.muted
            else -> border
        }
}

/** Number formats the item surfaces share. Formatting only: every figure comes from the API or the static data. */
object ItemNumbers {
    private const val MINUS = "−"

    /** `+11.1` and `−0.8` (a true minus, as the server writes it). */
    fun signedTenths(value: Double): String {
        val magnitude = tenths(abs(value))
        return if (value < 0) "$MINUS$magnitude" else "+$magnitude"
    }

    fun tenths(value: Double): String = String.format(Locale.US, "%.1f", value)

    fun hundredths(value: Double): String = String.format(Locale.US, "%.2f", value)
}
