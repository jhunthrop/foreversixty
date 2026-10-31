package gg.foreversixty.app.premium

import gg.foreversixty.app.network.AppConfigRepository
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.asStateFlow
import java.util.Locale

/**
 * Decision P: no in-app purchase. `See plans` links to web checkout only where the store allows
 * it, which the API names (`GET /v1/app/config`'s `web_checkout_regions`). Every unknown fails
 * closed: no config, no storefront, or the feature off means no link.
 */
object WebCheckoutPolicy {
    private val ISO_REGIONS: Set<String> = Locale.getISOCountries().toSet()

    fun mayLink(regions: List<String>?, storefront: String?, featureOn: Boolean): Boolean {
        val country = storefront?.let(::alpha2) ?: return false
        return featureOn && regions != null && country in regions
    }

    /** Play Billing names the country in ISO 3166-1 alpha-2 (`US`); anything else is not a region. */
    fun alpha2(code: String): String? = code.uppercase(Locale.ROOT).takeIf { it in ISO_REGIONS }
}

/** Where the storefront country comes from: Play Billing in the app, a fixed answer in tests. */
fun interface StorefrontReader {
    suspend fun countryCode(): String?
}

/**
 * Whether this install may show `See plans` and `Manage billing`, resolved once the config and
 * the storefront have answered. Starts at no link and stays there on any failure.
 */
class PlansLink(
    private val appConfig: AppConfigRepository,
    private val storefront: StorefrontReader,
    private val featureOn: () -> Boolean,
) {
    private val state = MutableStateFlow(false)
    val mayLink: StateFlow<Boolean> = state.asStateFlow()

    suspend fun resolve() {
        if (appConfig.current == null) appConfig.refresh()
        val country = storefront.countryCode()
        state.value = WebCheckoutPolicy.mayLink(appConfig.current?.webCheckoutRegions, country, featureOn())
    }
}
