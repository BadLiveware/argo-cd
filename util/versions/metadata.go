package versions

// RevisionResolutionType represents the type of revision resolution that occurred
type RevisionResolutionType string

const (
	// RevisionResolutionDirect indicates the revision was resolved directly (exact match)
	RevisionResolutionDirect RevisionResolutionType = "direct"
	// RevisionResolutionRange indicates the revision was resolved from a semver constraint/range
	RevisionResolutionRange RevisionResolutionType = "range"
	// RevisionResolutionSymbolicReference indicates the revision was resolved from a symbolic reference (e.g., HEAD)
	RevisionResolutionSymbolicReference RevisionResolutionType = "symbolic_reference"
	// RevisionResolutionTruncatedCommitSHA indicates the revision was assumed to be a truncated commit SHA
	RevisionResolutionTruncatedCommitSHA RevisionResolutionType = "truncated_commit_sha"
	// RevisionResolutionVersion indicates the revision was resolved as a specific version
	RevisionResolutionVersion RevisionResolutionType = "version"
)

// RevisionMetadata contains metadata about how a revision was resolved
type RevisionMetadata struct {
	// OriginalRevision is the original revision string provided by the user
	OriginalRevision string
	// ResolutionType indicates how the revision was resolved
	ResolutionType RevisionResolutionType
	// ResolvedTag is the actual tag/version that was resolved (if applicable)
	ResolvedTag string
	// ResolvedTo is the target of a symbolic reference resolution (Git-specific)
	ResolvedTo string
}

// IsEmpty returns true if the metadata has no meaningful data
func (m *RevisionMetadata) IsEmpty() bool {
	return m == nil || (m.OriginalRevision == "" && m.ResolutionType == "" && m.ResolvedTag == "" && m.ResolvedTo == "")
}

// NewRevisionMetadata creates a new RevisionMetadata with the given parameters
func NewRevisionMetadata(originalRevision string, resolutionType RevisionResolutionType) *RevisionMetadata {
	return &RevisionMetadata{
		OriginalRevision: originalRevision,
		ResolutionType:   resolutionType,
	}
}

// WithResolvedTag sets the resolved tag and returns the metadata for chaining
func (m *RevisionMetadata) WithResolvedTag(tag string) *RevisionMetadata {
	m.ResolvedTag = tag
	return m
}

// WithResolvedTo sets the resolved target and returns the metadata for chaining
func (m *RevisionMetadata) WithResolvedTo(target string) *RevisionMetadata {
	m.ResolvedTo = target
	return m
}

// ToEnvVars converts the revision metadata to environment variable key-value pairs
// with the ARGOCD_ prefix for use in build environments
func (m *RevisionMetadata) ToEnvVars() map[string]string {
	envVars := make(map[string]string)

	if m.OriginalRevision != "" {
		envVars["ARGOCD_ORIGINAL_REVISION"] = m.OriginalRevision
	}
	if m.ResolutionType != "" {
		envVars["ARGOCD_RESOLUTION_TYPE"] = string(m.ResolutionType)
	}
	if m.ResolvedTag != "" {
		envVars["ARGOCD_RESOLVED_TAG"] = m.ResolvedTag
	}
	if m.ResolvedTo != "" {
		envVars["ARGOCD_RESOLVED_TO"] = m.ResolvedTo
	}

	return envVars
}
