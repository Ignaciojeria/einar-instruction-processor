package archetype

import (
	_ "instruction-processor/app/adapter/in/nats_subscription"
	"instruction-processor/app/shared/archetype/container"
	_ "instruction-processor/app/shared/archetype/echo_server"
	_ "instruction-processor/app/shared/archetype/nats"
	"instruction-processor/app/shared/config"
)

// ARCHETYPE CONFIGURATION
func Setup() error {
	if !config.Installations.EnableCobraCli {
		if err := config.Setup(); err != nil {
			return err
		}
	}
	if err := InjectInstallations(); err != nil {
		return err
	}

	if err := injectOutboundAdapters(); err != nil {
		return err
	}

	if err := injectInboundAdapters(); err != nil {
		return err
	}

	if !config.Installations.EnableHTTPServer {
		return nil
	}
	if err := container.HTTPServerContainer.LoadDependency(); err != nil {
		return err
	}
	return nil
}

func InjectInstallations() error {
	for _, v := range container.InstallationsContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

// CUSTOM INITIALIZATION OF YOUR DOMAIN COMPONENTS
func injectOutboundAdapters() error {
	for _, v := range container.OutboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}

func injectInboundAdapters() error {
	for _, v := range container.InboundAdapterContainer {
		if err := v.LoadDependency(); err != nil {
			return err
		}
	}
	return nil
}
