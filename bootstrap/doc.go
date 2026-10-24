// Package bootstrap implements icon.Provider with Bootstrap Icons
// (https://icons.getbootstrap.com).
//
// Get renders an icon as an <i class="bi bi-NAME"></i> element, which the Bootstrap Icons
// stylesheet draws as a glyph. A few dozen generic names ("add", "delete", "user") map onto
// Bootstrap names, and any other name passes through unchanged as a Bootstrap name.
//
// The name is written into the class attribute without escaping, so pass only names that
// are trusted or validated, such as names matching ^[a-z0-9-]+$.
package bootstrap
