import Foundation
import SwiftUI

/// The API's sanitised `notes_html` as attributed text (spec F4.2 item 7): paragraphs, line
/// breaks, list items, bold, italic and links. The API has already sanitised it; anything else
/// is dropped to its text.
enum NotesText {
    static func attributed(_ html: String, base: Font, bold: Font, link: Color) -> AttributedString {
        var out = AttributedString()
        var isBold = false, isItalic = false
        var href: URL?
        var pendingBreak = false
        let scanner = Scanner(string: html)
        scanner.charactersToBeSkipped = nil
        while !scanner.isAtEnd {
            if let text = scanner.scanUpToString("<"), !text.isEmpty {
                let decoded = decode(text)
                if !decoded.trimmingCharacters(in: .whitespacesAndNewlines).isEmpty {
                    if pendingBreak, !out.characters.isEmpty { out.append(AttributedString("\n")) }
                    pendingBreak = false
                    var run = AttributedString(decoded.replacingOccurrences(of: "\n", with: " "))
                    run.font = isBold ? bold : base
                    if isItalic { run.font = (isBold ? bold : base).italic() }
                    if let href {
                        run.link = href
                        run.foregroundColor = link
                        run.underlineStyle = .single
                    }
                    out.append(run)
                }
            }
            guard scanner.scanString("<") != nil, let tag = scanner.scanUpToString(">") else { break }
            _ = scanner.scanString(">")
            let name = tag.split(separator: " ").first.map { $0.lowercased() } ?? ""
            switch name {
            case "strong", "b": isBold = true
            case "/strong", "/b": isBold = false
            case "em", "i": isItalic = true
            case "/em", "/i": isItalic = false
            case "a": href = attribute("href", in: tag).flatMap(URL.init(string:)).flatMap { $0.scheme == "https" ? $0 : nil }
            case "/a": href = nil
            case "/p", "br", "br/", "/li", "/ul", "/ol", "/h1", "/h2", "/h3": pendingBreak = true
            case "li":
                if !out.characters.isEmpty { out.append(AttributedString("\n")) }
                pendingBreak = false
                var bullet = AttributedString("• ")
                bullet.font = base
                out.append(bullet)
            default: break
            }
        }
        return out
    }

    private static func attribute(_ name: String, in tag: String) -> String? {
        guard let range = tag.range(of: "\(name)=\"") else { return nil }
        let rest = tag[range.upperBound...]
        return rest.firstIndex(of: "\"").map { decode(String(rest[..<$0])) }
    }

    private static func decode(_ text: String) -> String {
        [("&amp;", "&"), ("&lt;", "<"), ("&gt;", ">"), ("&quot;", "\""), ("&#39;", "'"), ("&#x27;", "'"), ("&nbsp;", " ")]
            .reduce(text) { $0.replacingOccurrences(of: $1.0, with: $1.1) }
    }
}
