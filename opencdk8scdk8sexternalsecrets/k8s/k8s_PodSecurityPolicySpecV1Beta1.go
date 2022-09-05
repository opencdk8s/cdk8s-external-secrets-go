package k8s


// PodSecurityPolicySpec defines the policy enforced.
// Experimental.
type PodSecurityPolicySpecV1Beta1 struct {
	// fsGroup is the strategy that will dictate what fs group is used by the SecurityContext.
	// Experimental.
	FsGroup *FsGroupStrategyOptionsV1Beta1 `field:"required" json:"fsGroup" yaml:"fsGroup"`
	// runAsUser is the strategy that will dictate the allowable RunAsUser values that may be set.
	// Experimental.
	RunAsUser *RunAsUserStrategyOptionsV1Beta1 `field:"required" json:"runAsUser" yaml:"runAsUser"`
	// seLinux is the strategy that will dictate the allowable labels that may be set.
	// Experimental.
	SeLinux *SeLinuxStrategyOptionsV1Beta1 `field:"required" json:"seLinux" yaml:"seLinux"`
	// supplementalGroups is the strategy that will dictate what supplemental groups are used by the SecurityContext.
	// Experimental.
	SupplementalGroups *SupplementalGroupsStrategyOptionsV1Beta1 `field:"required" json:"supplementalGroups" yaml:"supplementalGroups"`
	// allowedCapabilities is a list of capabilities that can be requested to add to the container.
	//
	// Capabilities in this field may be added at the pod author's discretion. You must not list a capability in both allowedCapabilities and requiredDropCapabilities.
	// Experimental.
	AllowedCapabilities *[]*string `field:"optional" json:"allowedCapabilities" yaml:"allowedCapabilities"`
	// AllowedCSIDrivers is an allowlist of inline CSI drivers that must be explicitly set to be embedded within a pod spec.
	//
	// An empty value indicates that any CSI driver can be used for inline ephemeral volumes. This is a beta field, and is only honored if the API server enables the CSIInlineVolume feature gate.
	// Experimental.
	AllowedCsiDrivers *[]*AllowedCsiDriverV1Beta1 `field:"optional" json:"allowedCsiDrivers" yaml:"allowedCsiDrivers"`
	// allowedFlexVolumes is an allowlist of Flexvolumes.
	//
	// Empty or nil indicates that all Flexvolumes may be used.  This parameter is effective only when the usage of the Flexvolumes is allowed in the "volumes" field.
	// Experimental.
	AllowedFlexVolumes *[]*AllowedFlexVolumeV1Beta1 `field:"optional" json:"allowedFlexVolumes" yaml:"allowedFlexVolumes"`
	// allowedHostPaths is an allowlist of host paths.
	//
	// Empty indicates that all host paths may be used.
	// Experimental.
	AllowedHostPaths *[]*AllowedHostPathV1Beta1 `field:"optional" json:"allowedHostPaths" yaml:"allowedHostPaths"`
	// AllowedProcMountTypes is an allowlist of allowed ProcMountTypes.
	//
	// Empty or nil indicates that only the DefaultProcMountType may be used. This requires the ProcMountType feature flag to be enabled.
	// Experimental.
	AllowedProcMountTypes *[]*string `field:"optional" json:"allowedProcMountTypes" yaml:"allowedProcMountTypes"`
	// allowedUnsafeSysctls is a list of explicitly allowed unsafe sysctls, defaults to none.
	//
	// Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of allowed sysctls. Single * means all unsafe sysctls are allowed. Kubelet has to allowlist all allowed unsafe sysctls explicitly to avoid rejection.
	//
	// Examples: e.g. "foo/*" allows "foo/bar", "foo/baz", etc. e.g. "foo.*" allows "foo.bar", "foo.baz", etc.
	// Experimental.
	AllowedUnsafeSysctls *[]*string `field:"optional" json:"allowedUnsafeSysctls" yaml:"allowedUnsafeSysctls"`
	// allowPrivilegeEscalation determines if a pod can request to allow privilege escalation.
	//
	// If unspecified, defaults to true.
	// Experimental.
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// defaultAddCapabilities is the default set of capabilities that will be added to the container unless the pod spec specifically drops the capability.
	//
	// You may not list a capability in both defaultAddCapabilities and requiredDropCapabilities. Capabilities added here are implicitly allowed, and need not be included in the allowedCapabilities list.
	// Experimental.
	DefaultAddCapabilities *[]*string `field:"optional" json:"defaultAddCapabilities" yaml:"defaultAddCapabilities"`
	// defaultAllowPrivilegeEscalation controls the default setting for whether a process can gain more privileges than its parent process.
	// Experimental.
	DefaultAllowPrivilegeEscalation *bool `field:"optional" json:"defaultAllowPrivilegeEscalation" yaml:"defaultAllowPrivilegeEscalation"`
	// forbiddenSysctls is a list of explicitly forbidden sysctls, defaults to none.
	//
	// Each entry is either a plain sysctl name or ends in "*" in which case it is considered as a prefix of forbidden sysctls. Single * means all sysctls are forbidden.
	//
	// Examples: e.g. "foo/*" forbids "foo/bar", "foo/baz", etc. e.g. "foo.*" forbids "foo.bar", "foo.baz", etc.
	// Experimental.
	ForbiddenSysctls *[]*string `field:"optional" json:"forbiddenSysctls" yaml:"forbiddenSysctls"`
	// hostIPC determines if the policy allows the use of HostIPC in the pod spec.
	// Experimental.
	HostIpc *bool `field:"optional" json:"hostIpc" yaml:"hostIpc"`
	// hostNetwork determines if the policy allows the use of HostNetwork in the pod spec.
	// Experimental.
	HostNetwork *bool `field:"optional" json:"hostNetwork" yaml:"hostNetwork"`
	// hostPID determines if the policy allows the use of HostPID in the pod spec.
	// Experimental.
	HostPid *bool `field:"optional" json:"hostPid" yaml:"hostPid"`
	// hostPorts determines which host port ranges are allowed to be exposed.
	// Experimental.
	HostPorts *[]*HostPortRangeV1Beta1 `field:"optional" json:"hostPorts" yaml:"hostPorts"`
	// privileged determines if a pod can request to be run as privileged.
	// Experimental.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// readOnlyRootFilesystem when set to true will force containers to run with a read only root file system.
	//
	// If the container specifically requests to run with a non-read only root file system the PSP should deny the pod. If set to false the container may run with a read only root file system if it wishes but it will not be forced to.
	// Experimental.
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// requiredDropCapabilities are the capabilities that will be dropped from the container.
	//
	// These are required to be dropped and cannot be added.
	// Experimental.
	RequiredDropCapabilities *[]*string `field:"optional" json:"requiredDropCapabilities" yaml:"requiredDropCapabilities"`
	// RunAsGroup is the strategy that will dictate the allowable RunAsGroup values that may be set.
	//
	// If this field is omitted, the pod's RunAsGroup can take any value. This field requires the RunAsGroup feature gate to be enabled.
	// Experimental.
	RunAsGroup *RunAsGroupStrategyOptionsV1Beta1 `field:"optional" json:"runAsGroup" yaml:"runAsGroup"`
	// runtimeClass is the strategy that will dictate the allowable RuntimeClasses for a pod.
	//
	// If this field is omitted, the pod's runtimeClassName field is unrestricted. Enforcement of this field depends on the RuntimeClass feature gate being enabled.
	// Experimental.
	RuntimeClass *RuntimeClassStrategyOptionsV1Beta1 `field:"optional" json:"runtimeClass" yaml:"runtimeClass"`
	// volumes is an allowlist of volume plugins.
	//
	// Empty indicates that no volumes may be used. To allow all volumes you may use '*'.
	// Experimental.
	Volumes *[]*string `field:"optional" json:"volumes" yaml:"volumes"`
}

