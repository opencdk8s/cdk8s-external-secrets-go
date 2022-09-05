package k8s


// CronJobList is a collection of cron jobs.
// Experimental.
type KubeCronJobListV1Beta1Props struct {
	// items is the list of CronJobs.
	// Experimental.
	Items *[]*KubeCronJobV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// Experimental.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

