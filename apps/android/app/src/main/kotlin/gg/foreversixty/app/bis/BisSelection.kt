package gg.foreversixty.app.bis

import gg.foreversixty.app.model.AccountCharacter
import gg.foreversixty.app.model.ApiJson
import gg.foreversixty.app.model.BisBand
import gg.foreversixty.app.model.BisFile
import kotlinx.serialization.SerialName
import kotlinx.serialization.Serializable
import kotlinx.serialization.SerializationException
import kotlinx.serialization.builtins.ListSerializer

/**
 * One spec of the bundled catalogue (data/curated/specs.json, copied into the app: the CDN does
 * not publish it). Only the fields the picker reads; the file is the source of truth.
 */
@Serializable
data class SpecInfo(
    /** The spec key, `shaman-elemental`: the BiS file's name. */
    @SerialName("spec") val spec: String,
    @SerialName("class_slug") val classSlug: String,
    /** The display name, `Elemental`. */
    @SerialName("name") val name: String,
    /** `dps`, `healer` or `tank`. */
    @SerialName("role") val role: String,
    /** The icon stem on the CDN (`/data/<build>/icons/<icon>.webp`). */
    @SerialName("icon") val icon: String,
)

/** The specs of every class, in the catalogue's order. */
data class SpecCatalog(val specs: List<SpecInfo>) {
    fun specsOf(classSlug: String): List<SpecInfo> = specs.filter { it.classSlug == classSlug }

    fun spec(key: String): SpecInfo? = specs.firstOrNull { it.spec == key }

    /** The spec a character names by its class slug and tree name (`warrior`, `Fury`). */
    fun spec(classSlug: String, name: String): SpecInfo? =
        specs.firstOrNull { it.classSlug == classSlug && it.name.equals(name, ignoreCase = true) }

    /** Class slugs in first-seen order, for a picker with no classes file. */
    val classSlugs: List<String> get() = specs.map { it.classSlug }.distinct()

    companion object {
        const val ASSET = "bundled/specs.json"

        /** Empty for a body that does not decode; the bundle test catches a broken build. */
        fun decode(json: String): SpecCatalog = try {
            SpecCatalog(ApiJson.decodeFromString(ListSerializer(SpecInfo.serializer()), json))
        } catch (e: SerializationException) {
            SpecCatalog(emptyList())
        } catch (e: IllegalArgumentException) {
            SpecCatalog(emptyList())
        }
    }
}

/** A level range of the picker: a band of the BiS file and its preset (spec F3.2). */
@Serializable
data class BisRange(val band: Int, val preset: String) {
    companion object {
        const val BARE = "bare"
        const val RAID = "raid"
        const val FIRST_BAND = 20
        const val LAST_BAND = 60
        const val BAND_STEP = 10

        /** `20 to 29` through `60`, then `60 · Phase 1 raid`. */
        val all: List<BisRange> =
            (FIRST_BAND..LAST_BAND step BAND_STEP).map { BisRange(it, BARE) } + BisRange(LAST_BAND, RAID)

        /**
         * The site's band rule (`bandForLevel`): rounded down to the step and clamped to 20 to 60.
         * A level-60 character reads the raid list, the site's headline preset.
         */
        fun forLevel(level: Int): BisRange {
            val band = level.coerceIn(FIRST_BAND, LAST_BAND) / BAND_STEP * BAND_STEP
            return BisRange(band, if (band == LAST_BAND) RAID else BARE)
        }
    }
}

/** Which list the BiS tab shows. */
@Serializable
data class BisSelection(val classSlug: String, val spec: String, val range: BisRange, val faction: String) {
    fun withClass(slug: String, catalog: SpecCatalog): BisSelection =
        copy(classSlug = slug, spec = catalog.specsOf(slug).firstOrNull()?.spec ?: spec)

    companion object {
        val factions = listOf("alliance", "horde")
        private const val WARRIOR = "warrior"

        /** The first visit, signed out: Warrior preselected (spec F3.2), the headline list. */
        fun signedOutDefault(catalog: SpecCatalog): BisSelection = BisSelection(
            classSlug = WARRIOR,
            spec = catalog.specsOf(WARRIOR).firstOrNull()?.spec ?: "$WARRIOR-arms",
            range = BisRange.all.last(),
            faction = factions.first(),
        )

        /** The character's own list, when the catalogue knows its class and spec. */
        fun forCharacter(character: AccountCharacter, catalog: SpecCatalog): BisSelection? {
            val classSlug = character.className ?: return null
            val spec = character.spec?.let { catalog.spec(classSlug, it) } ?: return null
            val level = character.level ?: return null
            val faction = character.faction?.takeIf { it in factions } ?: return null
            return BisSelection(classSlug, spec.spec, BisRange.forLevel(level), faction)
        }
    }
}

/**
 * The band the site shows for a range and faction (`presets.ts` `selectBand`): the asked preset,
 * else the bare list. Nothing is re-ranked; the file's own entry is returned.
 */
fun BisFile.band(range: BisRange, faction: String): BisBand? {
    val candidates = bands.filter { it.band == range.band && it.faction == faction }
    return candidates.firstOrNull { it.preset == range.preset } ?: candidates.firstOrNull { it.preset == BisRange.BARE }
}
