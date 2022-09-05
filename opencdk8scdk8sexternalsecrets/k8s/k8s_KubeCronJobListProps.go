package k8s


// CronJobList is a collection of cron jobs.
// Experimental.
type KubeCronJobListProps struct {
	// items is the list of CronJobs.
	// Experimental.
	Items *[]*KubeCronJobProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

