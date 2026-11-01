package gg.foreversixty.app.items

import android.content.res.Resources
import gg.foreversixty.app.R
import gg.foreversixty.app.bis.BisListBuilder
import gg.foreversixty.app.bis.BisUnit
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import gg.foreversixty.app.model.BisSlot
import gg.foreversixty.app.model.Item
import gg.foreversixty.app.model.ItemSet
import kotlin.math.abs
import kotlin.math.roundToInt

/** An item's name, quality, icon and level as the API sends them; every field but id may be null. */
data class ItemLook(
    val id: Int,
    val name: String? = null,
    val quality: Int? = null,
    val icon: String? = null,
    val itemLevel: Int? = null,
)

/**
 * Who asked for a tooltip and what they already know: the item, the build its static data lives
 * in, and the band the pick belongs to. Built by Home and the gear list.
 */
data class ItemTooltipRequest(
    val itemId: Int,
    /** The name and look to draw while the static file loads, or when it lacks the item. */
    val fallback: ItemLook,
    val build: String,
    val classSlug: String,
    val spec: String?,
    val specName: String?,
    val band: Int?,
    val bandLabel: String?,
    val preset: String?,
    val faction: String?,
    /** The server's source words for this item, when the caller has them (a Home pick). */
    val sourceLine: String? = null,
    /** The item worn in the same slot, so the next best can say "You wear this". */
    val wornId: Int? = null,
)

/** The static files a tooltip reads, already loaded. */
data class ItemTooltipSources(
    val items: Map<Int, Item>,
    val sets: List<ItemSet>,
    val bis: BisFile?,
)

/** One runner-up under "The next best": its delta against the pick and where it comes from. */
data class ItemTooltipAlternative(val look: ItemLook, val meta: String, val source: String?)

/**
 * Everything the tooltip sheet draws, in the game's order (spec F2.6, the site's
 * `lib/items/tooltip.ts`). Pure: the sheet renders it and the tests read it.
 */
data class ItemTooltipModel(
    val look: ItemLook,
    val headerMeta: String,
    /** False when the build's file does not list the item; the panel then says so. */
    val known: Boolean,
    val binds: String? = null,
    val unique: Boolean = false,
    val slotName: String? = null,
    val typeName: String? = null,
    val armor: String? = null,
    val weaponDamage: String? = null,
    val weaponSpeed: String? = null,
    val weaponDps: String? = null,
    /** The white `+N Stat` lines. */
    val stats: List<String> = emptyList(),
    /** The client's green `Equip:` sentences for spell power, healing, mp5 and the rest. */
    val equipLines: List<String> = emptyList(),
    val effect: String? = null,
    val setName: String? = null,
    val setBonuses: List<String> = emptyList(),
    val requiresLevel: String? = null,
    val itemLevel: String? = null,
    val source: String? = null,
    val alternatives: List<ItemTooltipAlternative> = emptyList(),
    val pickFooter: String? = null,
) {
    companion object {
        private val weaponSlots = setOf("main_hand", "off_hand", "ranged")

        /** Spec F2.6: up to two runners-up. */
        private const val MAX_ALTERNATIVES = 2

        fun build(request: ItemTooltipRequest, sources: ItemTooltipSources, res: Resources): ItemTooltipModel {
            val item = sources.items[request.itemId] ?: return unknown(request, res)
            val quality = ItemQuality(item.quality)
            val slotName = ItemVocabulary.slotName(item.slot)?.let(res::getString)
            val band = band(request, sources.bis)
            val pickSlot = band?.slots?.firstOrNull { it.itemId == item.id }
            val weapon = item.slot in weaponSlots && item.damageMax > 0
            val set = item.setId?.let { id -> sources.sets.firstOrNull { it.id == id } }
            return ItemTooltipModel(
                look = ItemLook(item.id, item.name, item.quality, item.icon, item.itemLevel),
                headerMeta = res.getString(R.string.items_header, quality.name?.let(res::getString).orEmpty(), slotName.orEmpty()),
                known = true,
                binds = ItemVocabulary.bindsLine(item.bonding)?.let(res::getString),
                unique = item.unique,
                slotName = if (weapon) weaponSlotName(item, res) else slotName,
                typeName = (if (weapon) ItemVocabulary.subclassName(item.subclass) else ItemVocabulary.armorTypeName(item.armorType))
                    ?.let(res::getString),
                armor = item.armor.takeIf { it > 0 }?.let { res.getString(R.string.tooltip_armor, it) },
                weaponDamage = if (weapon) {
                    res.getString(R.string.items_weapon_damage, item.damageMin.roundToInt(), item.damageMax.roundToInt())
                } else {
                    null
                },
                weaponSpeed = if (weapon) res.getString(R.string.items_weapon_speed, ItemNumbers.hundredths(item.speed)) else null,
                weaponDps = if (weapon) res.getString(R.string.items_weapon_dps, ItemNumbers.tenths(item.dps)) else null,
                stats = statLines(item, res),
                equipLines = equipLines(item, res),
                effect = item.effectText.split(Regex("\\s+")).filter { it.isNotEmpty() }.joinToString(" ").ifEmpty { null },
                setName = set?.name,
                setBonuses = set?.bonuses.orEmpty().map { res.getString(R.string.items_setBonus, it.pieces, it.description) },
                requiresLevel = item.requiredLevel.takeIf { it > 0 }?.let { res.getString(R.string.tooltip_requires, it) },
                itemLevel = res.getString(R.string.tooltip_itemLevel, item.itemLevel),
                source = source(item, request, band, res),
                alternatives = alternatives(pickSlot, request, sources.items, res, BisUnit.forRole(band?.role)),
                pickFooter = pickSlot?.let { footer(ItemVocabulary.slotLabel(res, it.slot), request, res) },
            )
        }

        private fun unknown(request: ItemTooltipRequest, res: Resources): ItemTooltipModel =
            ItemTooltipModel(
                look = request.fallback,
                headerMeta = ItemQuality(request.fallback.quality).name?.let(res::getString).orEmpty(),
                known = false,
                itemLevel = request.fallback.itemLevel?.let { res.getString(R.string.tooltip_itemLevel, it) },
                source = request.sourceLine,
            )

        /** The band the pick belongs to: the request's band and preset, the character's faction when known. */
        fun band(request: ItemTooltipRequest, bis: BisFile?): BisBand? {
            val level = request.band ?: return null
            val candidates = bis?.bands.orEmpty().filter { it.band == level && it.preset == (request.preset ?: it.preset) }
            return candidates.firstOrNull { it.faction == request.faction } ?: candidates.firstOrNull()
        }

        /**
         * The left word of a weapon's slot line as the client prints it (`Two-Hand`, `Main Hand`):
         * the client's InventoryType when the data carries it, else what `two_hand` and the slot say.
         */
        private fun weaponSlotName(item: Item, res: Resources): String = res.getString(
            ItemVocabulary.weaponSlotName(item.weaponSlot) ?: when {
                item.twoHand -> R.string.items_weaponSlot_twoHand
                item.slot == "ranged" -> R.string.items_weaponSlot_ranged
                else -> R.string.items_weaponSlot_oneHand
            },
        )

        /**
         * `+22 Intellect`, `-15 Parry`: each stat's own share (the stat less what an on-equip spell
         * states, `equip_stats`), the client's sign, the site's stat order. Without `equip_stats`
         * (older data) every stat prints here.
         */
        private fun statLines(item: Item, res: Resources): List<String> =
            ItemVocabulary.statOrder.mapNotNull { (key, name) ->
                val stat = item.stats[key]?.takeIf { it != 0 } ?: return@mapNotNull null
                if (percentLine(item, key, res) != null) return@mapNotNull null
                val value = stat - (item.equipStats?.get(key) ?: 0)
                if (value == 0) return@mapNotNull null
                "${if (value < 0) "-" else "+"}${abs(value)} ${res.getString(name)}"
            }

        /** The green `Equip:` sentences: the percent ones, then one per positive on-equip share. */
        private fun equipLines(item: Item, res: Resources): List<String> {
            val keys = ItemVocabulary.statOrder.map { it.first }
            val percent = keys.filter { (item.stats[it] ?: 0) != 0 }.mapNotNull { percentLine(item, it, res) }
            val shares = equipShares(item.equipStats.orEmpty())
            val flat = keys.mapNotNull { key ->
                val value = shares[key]?.takeIf { it > 0 } ?: return@mapNotNull null
                ItemVocabulary.equipStats[key]?.let { res.getString(it, value) }
            }
            return percent + flat
        }

        /**
         * What each on-equip spell states. One spell giving equal spell power and healing (`Increase
         * Spell Dam`) is one sentence: healing keeps only what it adds beyond spell power.
         */
        fun equipShares(equip: Map<String, Int>): Map<String, Int> {
            val spellPower = equip[SPELL_POWER] ?: 0
            return equip + (HEALING to ((equip[HEALING] ?: 0) - spellPower).coerceAtLeast(0))
        }

        private const val SPELL_POWER = "spell_power"
        private const val HEALING = "healing"

        private fun percentLine(item: Item, key: String, res: Resources): String? {
            val line = ItemVocabulary.percentEquipStats[key] ?: return null
            val percent = item.percentStats[key]?.takeIf { it != 0 && it == item.stats[key] } ?: return null
            return res.getString(line, "$percent%")
        }

        /**
         * Where it comes from: the site's formatted line (the BiS row's, else the caller's, which
         * the API formats the same way), plus the standing a vendor requires (` · requires
         * Friendly`). A reputation reward's line already names its standing.
         */
        private fun source(item: Item, request: ItemTooltipRequest, band: BisBand?, res: Resources): String? {
            val (line, kind) = bisSource(item.id, band) ?: (request.sourceLine to null)
            line ?: return null
            val standing = ItemVocabulary.standingName(item.minStanding)?.let(res::getString)
            return if (kind == "vendor" && standing != null) line + res.getString(R.string.tooltip_requiresStanding, standing) else line
        }

        private fun bisSource(id: Int, band: BisBand?): Pair<String?, String?>? {
            band ?: return null
            band.slots.firstOrNull { it.itemId == id && (it.sourceLine ?: it.source) != null }
                ?.let { return (it.sourceLine ?: it.source) to it.sourceKind }
            band.slots.forEach { slot ->
                slot.alternatives?.firstOrNull { it.itemId == id }?.let { return (it.sourceLine ?: it.source) to it.sourceKind }
            }
            return null
        }

        /**
         * Up to two runners-up from the BiS row, each in the band's unit (DPS, HPS or score); the
         * worn one leads with "You wear this · ".
         */
        private fun alternatives(
            slot: BisSlot?,
            request: ItemTooltipRequest,
            items: Map<Int, Item>,
            res: Resources,
            unit: BisUnit,
        ): List<ItemTooltipAlternative> {
            val unitWord = res.getString(unit.word)
            return slot?.alternatives.orEmpty().take(MAX_ALTERNATIVES).map { alt ->
                val row = items[alt.itemId]
                val look = ItemLook(alt.itemId, row?.name ?: alt.itemName, row?.quality, row?.icon, row?.itemLevel)
                val delta = BisListBuilder.signedDelta(alt.dpsDelta)?.let { res.getString(R.string.tooltip_delta, it, unitWord) }
                    ?: res.getString(R.string.items_sameDps, unitWord)
                val worn = if (alt.itemId == request.wornId) res.getString(R.string.tooltip_youWear) else ""
                ItemTooltipAlternative(look, worn + delta, alt.sourceLine ?: alt.source)
            }
        }

        /**
         * `Best in slot for Finger 1 · Fury Warrior 60 · Phase 1 raid`, named by the BiS row's slot; the
         * band label carries the phase, never a preset id.
         */
        private fun footer(slotName: String, request: ItemTooltipRequest, res: Resources): String? {
            if (request.specName == null || request.bandLabel == null) return null
            return res.getString(R.string.tooltip_pickOf, slotName, request.specName, request.bandLabel)
        }
    }
}
