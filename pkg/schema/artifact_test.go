package schema_test

import (
	"github.com/kairos-io/AuroraBoot/pkg/schema"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Artifact", func() {
	Describe("FileName", func() {
		var artifact *schema.ReleaseArtifact

		BeforeEach(func() {
			artifact = &schema.ReleaseArtifact{
				ArtifactVersion: "v2.4.0",
				Model:           "generic",
				Flavor:          "rockylinux",
				Platform:        "amd64",
				ReleaseVersion:  "v2.4.0",
				Repository:      "kairos-io/kairos",
				Variant:         "core",
			}
		})

		It("should return the correct filename", func() {
			Expect(artifact.FileName()).To(Equal("kairos-core-rockylinux-amd64-generic-v2.4.0"))
		})

		Context("when the container_image is set", func() {
			It("should return an empty string", func() {
				artifact.ContainerImage = "docker://quay.io/kairos/core-rockylinux:latest"
				Expect(artifact.FileName()).To(Equal(""))
			})
		})

		Context("when the model is empty", func() {
			It("should default to generic", func() {
				artifact.Model = ""
				Expect(artifact.FileName()).To(Equal("kairos-core-rockylinux-amd64-generic-v2.4.0"))
			})
		})

		Context("when the platform is empty", func() {
			It("should default to amd64", func() {
				artifact.Platform = ""
				Expect(artifact.FileName()).To(Equal("kairos-core-rockylinux-amd64-generic-v2.4.0"))
			})
		})

		Context("when the variant is empty", func() {
			It("should default to core", func() {
				artifact.Variant = ""
				Expect(artifact.FileName()).To(Equal("kairos-core-rockylinux-amd64-generic-v2.4.0"))
			})

			It("should default to standard when the artifact version contains k3s", func() {
				artifact.Variant = ""
				artifact.ArtifactVersion = "v2.4.0-k3sv1.26.1+k3s1"
				Expect(artifact.FileName()).To(Equal("kairos-standard-rockylinux-amd64-generic-v2.4.0-k3sv1.26.1+k3s1"))
			})
		})

		Context("when the release version is less than v2.4.0", func() {
			It("should use the old filename format", func() {
				artifact.ReleaseVersion = "v2.3.0"
				artifact.ArtifactVersion = "v2.3.0"
				Expect(artifact.FileName()).To(Equal("kairos-rockylinux-v2.3.0"))
			})
		})
	})
})
