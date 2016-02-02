package destiny_test

import (
	"acceptance-tests/testing/destiny"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Job", func() {
	Describe("SetJobInstanceCount", func() {
		It("sets the correct values for instances and static_ips given a count", func() {
			manifest := destiny.NewEtcd(destiny.Config{
				IAAS: destiny.Warden,
			})
			job := manifest.Jobs[0]
			network := manifest.Networks[0]
			properties := manifest.Properties

			Expect(job.Instances).To(Equal(1))
			Expect(job.Networks[0].StaticIPs).To(HaveLen(1))
			Expect(job.Networks[0].Name).To(Equal(network.Name))
			Expect(properties.Etcd.Machines).To(Equal(job.Networks[0].StaticIPs))

			job, properties = destiny.SetJobInstanceCount(job, network, properties, 3)
			Expect(job.Instances).To(Equal(3))
			Expect(job.Networks[0].StaticIPs).To(HaveLen(3))
			Expect(properties.Etcd.Machines).To(Equal(job.Networks[0].StaticIPs))
		})
	})
})
