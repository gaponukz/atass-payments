package architecture

import (
	"testing"

	"github.com/matthewmcnew/archtest"
)

func TestArchitecture(t *testing.T) {
	t.Run("domain layer must have no dependencies", func(t *testing.T) {
		mockT := new(testingT)

		for _, domainPackage := range domainLayer() {
			for _, applicationPackage := range applicationLayer() {
				archtest.Package(t, domainPackage).ShouldNotDependOn(applicationPackage)
			}
		}

		for _, domainPackage := range domainLayer() {
			for _, infrastructurePackage := range infrastructureLayer() {
				archtest.Package(t, domainPackage).ShouldNotDependOn(infrastructurePackage)
			}
		}

		assertNoError(t, mockT)
	})

	t.Run("application layer must not depend on infrastructure", func(t *testing.T) {
		mockT := new(testingT)

		for _, applicationPackage := range applicationLayer() {
			for _, infrastructurePackage := range infrastructureLayer() {
				archtest.Package(t, applicationPackage).ShouldNotDependOn(infrastructurePackage)
			}
		}

		assertNoError(t, mockT)
	})
}
