import DesignSystem
import SwiftUI

/// You › Notifications › Live surfaces (spec §5.6 Rules): one switch per Live Activity, both on
/// by default. Kept in the App Group so the app and the extension read the same answer.
struct LiveSurfaceSettings: Sendable {
    static let raidKey = "live.raidCountdown"
    static let simKey = "live.simProgress"

    /// The defaults suite: the App Group's in the app, a throwaway one in tests.
    let suiteName: String

    static let shared = LiveSurfaceSettings(suiteName: AppGroup.id)

    private func defaults() -> UserDefaults { UserDefaults(suiteName: suiteName) ?? .standard }

    var raidCountdown: Bool { defaults().object(forKey: Self.raidKey) as? Bool ?? true }
    var simProgress: Bool { defaults().object(forKey: Self.simKey) as? Bool ?? true }

    func set(raidCountdown: Bool) { defaults().set(raidCountdown, forKey: Self.raidKey) }
    func set(simProgress: Bool) { defaults().set(simProgress, forKey: Self.simKey) }
}

/// The `Live surfaces` section: two switches with the spec's words.
struct LiveSurfaceSettingsView: View {
    let settings: LiveSurfaceSettings
    /// The sim switch shows only with `server_sims`: a free player has no server runs (spec §1.1).
    let showsSim: Bool
    var onChange: () -> Void = {}

    @State private var raid: Bool
    @State private var sim: Bool

    init(settings: LiveSurfaceSettings = .shared, showsSim: Bool, onChange: @escaping () -> Void = {}) {
        self.settings = settings
        self.showsSim = showsSim
        self.onChange = onChange
        _raid = State(initialValue: settings.raidCountdown)
        _sim = State(initialValue: settings.simProgress)
    }

    var body: some View {
        VStack(alignment: .leading, spacing: 0) {
            VStack(spacing: 0) {
                Toggle(Copy.liveRaidSetting, isOn: $raid)
                    .frame(minHeight: 52)
                    .onChange(of: raid) { _, on in settings.set(raidCountdown: on); onChange() }
                if showsSim {
                    Divider().overlay(FSColor.lineSoft)
                    Toggle(Copy.liveSimSetting, isOn: $sim)
                        .frame(minHeight: 52)
                        .onChange(of: sim) { _, on in settings.set(simProgress: on); onChange() }
                }
            }
            .font(ScaledType.body)
            .foregroundStyle(FSColor.text)
            .tint(FSColor.kill)
            .padding(.horizontal, 16)
            .background(FSColor.raised, in: .rect(cornerRadius: FSRadius.control))
            .overlay(RoundedRectangle(cornerRadius: FSRadius.control).stroke(FSColor.line))
            .padding(.horizontal, 16)
            .padding(.top, 22)
        }
    }
}

/// You › Live surfaces.
struct LiveSurfacesScreen: View {
    let showsSim: Bool
    var onChange: () -> Void = {}

    var body: some View {
        ScrollView {
            LiveSurfaceSettingsView(showsSim: showsSim, onChange: onChange)
        }
        .background(FSColor.bg)
        .navigationTitle(Copy.liveSettings)
        .navigationBarTitleDisplayMode(.inline)
    }
}
