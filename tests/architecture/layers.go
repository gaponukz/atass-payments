package architecture

func domainLayer() []string {
	return []string{
		"payments/internal/domain/entities",
		"payments/internal/domain/errors",
	}
}

func applicationLayer() []string {
	return []string{
		"payments/internal/application/dto",
		"payments/internal/application/cdc_service",
		"payments/internal/application/outbox",
		"payments/internal/application/usecase",
	}
}

func infrastructureLayer() []string {
	return []string{
		"payments/internal/infrastructure/controller",
		"payments/internal/infrastructure/logger",
		"payments/internal/infrastructure/notifier",
		"payments/internal/infrastructure/settings",
		"payments/internal/infrastructure/storage",
	}
}
