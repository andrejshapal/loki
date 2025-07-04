package distributor

import (
	"time"

	"github.com/grafana/loki/v3/pkg/compactor/retention"
	"github.com/grafana/loki/v3/pkg/distributor/shardstreams"
	"github.com/grafana/loki/v3/pkg/loghttp/push"
)

// Limits is an interface for distributor limits/related configs
type Limits interface {
	retention.Limits
	MaxLineSize(userID string) int
	MaxLineSizeTruncate(userID string) bool
	MaxLineSizeTruncateIdentifier(userID string) string

	MaxLabelNamesPerSeries(userID string) int
	MaxLabelNameLength(userID string) int
	MaxLabelValueLength(userID string) int

	CreationGracePeriod(userID string) time.Duration
	RejectOldSamples(userID string) bool
	RejectOldSamplesMaxAge(userID string) time.Duration

	IncrementDuplicateTimestamps(userID string) bool
	DiscoverServiceName(userID string) []string
	DiscoverGenericFields(userID string) map[string][]string
	DiscoverLogLevels(userID string) bool
	LogLevelFields(userID string) []string
	LogLevelFromJSONMaxDepth(userID string) int

	ShardStreams(userID string) shardstreams.Config
	IngestionRateStrategy() string
	IngestionRateBytes(userID string) float64
	IngestionBurstSizeBytes(userID string) int
	AllowStructuredMetadata(userID string) bool
	MaxStructuredMetadataSize(userID string) int
	MaxStructuredMetadataCount(userID string) int
	OTLPConfig(userID string) push.OTLPConfig

	BlockIngestionUntil(userID string) time.Time
	BlockIngestionStatusCode(userID string) int
	BlockIngestionPolicyUntil(userID string, policy string) time.Time
	EnforcedLabels(userID string) []string
	PolicyEnforcedLabels(userID string, policy string) []string

	IngestionPartitionsTenantShardSize(userID string) int

	SimulatedPushLatency(userID string) time.Duration
}
