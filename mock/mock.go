//go:generate mockgen -package mock -destination ./hn_client.go -mock_names "Client=MockHnClient" goat/pkg/hn Client
//go:generate mockgen -package mock -destination ./generator.go -mock_names "Generator=MockGenerator" goat/pkg/ai Generator

package mock
