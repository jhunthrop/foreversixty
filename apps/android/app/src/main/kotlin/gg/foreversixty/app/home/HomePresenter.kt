package gg.foreversixty.app.home

import android.content.res.Resources
import gg.foreversixty.app.bis.BisLabels
import gg.foreversixty.app.bis.BisRange
import gg.foreversixty.app.items.ItemLook
import gg.foreversixty.app.items.ItemTooltipRequest
import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.BuildTalents
import gg.foreversixty.app.model.CharacterHome
import gg.foreversixty.app.model.HomeUpgrade

/** The character Home is about, from `/v1/me` (lane F1's selector picks it). */
data class HomeCharacter(
    val key: String,
    val region: String,
    val ruleset: String,
    val name: String,
    val race: String?,
    val className: String?,
    /** The primary tree's name as the API gives it (`Beast Mastery`). */
    val specName: String?,
    val level: Int?,
    val guild: String?,
    val faction: String?,
    /** Worn item ids by slot from the last export, for the gear rows the answer does not name. */
    val gear: Map<String, Int> = emptyMap(),
    /** The last export's talent ranks and the build it was made on (lane F6 draws them). */
    val talents: BuildTalents? = null,
    val dataBuild: String? = null,
) {
    /** The path segment of the character routes (`velaren` in `us/normal/velaren`). */
    val slug: String get() = key.substringAfterLast('/')

    companion object {
        fun from(character: AccountCharacter) = HomeCharacter(
            key = character.key, region = character.region, ruleset = character.ruleset, name = character.name,
            race = character.race, className = character.className, specName = character.spec, level = character.level,
            guild = character.guild?.name, faction = character.faction, gear = character.build?.gear.orEmpty(),
            talents = character.build?.talents, dataBuild = character.build?.dataBuild,
        )
    }
}

/**
 * Home's shared formatting: the band it answers from, the rows' presentation order, the API's
 * slugs as words and the tooltip request. It formats and never ranks: every gain is the
 * server's (plan guard rail "one behaviour, one place").
 */
object HomePresenter {
    /**
     * The list Home answers from: the character's level through the site's band rule
     * ([BisRange.forLevel]), the rule BiS opens the character's list with. The server's band
     * stands in only when the level is unknown, and goes through the same rule.
     */
    fun range(home: CharacterHome, character: HomeCharacter): BisRange? =
        (character.level ?: home.band)?.let(BisRange::forLevel)

    /** `60 · raid` for the raid list, the band's words (`50 to 59`) otherwise. */
    fun bandHead(home: CharacterHome, character: HomeCharacter, res: Resources): String? =
        range(home, character)?.let { BisLabels.head(res, it) }

    /**
     * The rows in the coordinator's order (F2 fix round, 2026-10-31), for presentation only:
     * measured gains largest first, then gains not modelled yet, then empty slots. The gains
     * are the server's; this only arranges them. Equal rows keep the server's order.
     */
    fun presentationOrder(upgrades: List<HomeUpgrade>): List<HomeUpgrade> {
        fun group(u: HomeUpgrade) = when {
            u.modelled && u.gain != null -> 0
            u.worn != null -> 1
            else -> 2
        }
        return upgrades.sortedWith(compareBy<HomeUpgrade> { group(it) }.thenByDescending { if (group(it) == 0) it.gain else null })
    }

    /** `night_elf` → `Night Elf`; the API's slugs are lowercase. */
    fun words(slug: String): String =
        slug.split('_', '-', ' ').filter { it.isNotEmpty() }.joinToString(" ") { it.replaceFirstChar(Char::uppercaseChar) }

    /** A tooltip for an item on Home or Your gear, in the answer's build and spec and Home's band. */
    fun tooltipRequest(
        look: ItemLook,
        home: CharacterHome,
        character: HomeCharacter,
        res: Resources,
        sourceLine: String? = null,
        wornId: Int? = null,
    ): ItemTooltipRequest {
        val range = range(home, character)
        return ItemTooltipRequest(
            itemId = look.id, fallback = look, build = home.dataBuild, classSlug = classSlug(home, character),
            spec = home.spec, specName = character.specName, band = range?.band, bandLabel = bandHead(home, character, res),
            preset = range?.preset, faction = character.faction, sourceLine = sourceLine, wornId = wornId,
        )
    }

    /** The items file to read: the class part of the spec slug (`hunter-beast-mastery`), else the character's class. */
    fun classSlug(home: CharacterHome, character: HomeCharacter): String =
        home.spec?.substringBefore('-') ?: character.className.orEmpty()
}
