package k8s


// CronJob represents the configuration of a single cron job.
// Experimental.
type KubeCronJobProps struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of a cron job, including the schedule.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// Experimental.
	Spec *CronJobSpec `field:"optional" json:"spec" yaml:"spec"`
}

