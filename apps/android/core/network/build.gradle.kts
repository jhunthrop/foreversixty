plugins {
    alias(libs.plugins.forever.android.library)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.forever.contract.fixtures)
}

/** Live API host. Debug builds default to fixture mode so nothing touches production by accident. */
val liveBaseUrl = "https://api.foreversixty.gg"
val debugBaseUrl = providers.gradleProperty("forever.apiBaseUrl").orElse(liveBaseUrl).get()
val debugApiMode = providers.gradleProperty("forever.apiMode").orElse("fixture").get()

android {
    buildFeatures { buildConfig = true }
    buildTypes {
        debug {
            buildConfigField("String", "API_BASE_URL", "\"$debugBaseUrl\"")
            buildConfigField("String", "API_MODE", "\"$debugApiMode\"")
        }
        release {
            buildConfigField("String", "API_BASE_URL", "\"$liveBaseUrl\"")
            buildConfigField("String", "API_MODE", "\"live\"")
        }
    }
}

dependencies {
    api(project(":core:model"))
    api(libs.ktor.client.core)
    implementation(libs.ktor.client.content.negotiation)
    implementation(libs.ktor.serialization.kotlinx.json)
    implementation(libs.ktor.client.okhttp)
    // The fixture engine answers from the contract fixtures without a network.
    implementation(libs.ktor.client.mock)

    testImplementation(libs.junit)
    testImplementation(libs.kotlinx.coroutines.test)
}
