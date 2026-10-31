plugins {
    alias(libs.plugins.forever.android.library)
    alias(libs.plugins.kotlin.serialization)
    alias(libs.plugins.forever.contract.fixtures)
}

dependencies {
    api(project(":core:network"))
    implementation(libs.androidx.browser)
    implementation(libs.androidx.datastore.core)
    implementation(libs.tink.android)
    implementation(libs.kotlinx.coroutines.core)

    testImplementation(libs.junit)
    testImplementation(libs.kotlinx.coroutines.test)
    testImplementation(libs.ktor.client.mock)
}
