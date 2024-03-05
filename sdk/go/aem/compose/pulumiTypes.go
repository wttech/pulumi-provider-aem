// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package compose

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/wttech/pulumi-aem-native/sdk/go/aem/internal"
)

var _ = internal.GetEnvOrDefault

type Client struct {
	// Used when trying to connect to the AEM instance machine (often right after creating it). Need to be enough long because various types of connections (like AWS SSM or SSH) may need some time to boot up the agent.
	Action_timeout *string `pulumi:"action_timeout"`
	// Credentials for the connection type
	Credentials map[string]string `pulumi:"credentials"`
	// Settings for the connection type
	Settings map[string]string `pulumi:"settings"`
	// Used when reading the AEM instance state when determining the plan.
	State_timeout *string `pulumi:"state_timeout"`
	// Type of connection to use to connect to the machine on which AEM instance will be running.
	Type string `pulumi:"type"`
}

// ClientInput is an input type that accepts ClientArgs and ClientOutput values.
// You can construct a concrete instance of `ClientInput` via:
//
//	ClientArgs{...}
type ClientInput interface {
	pulumi.Input

	ToClientOutput() ClientOutput
	ToClientOutputWithContext(context.Context) ClientOutput
}

type ClientArgs struct {
	// Used when trying to connect to the AEM instance machine (often right after creating it). Need to be enough long because various types of connections (like AWS SSM or SSH) may need some time to boot up the agent.
	Action_timeout pulumi.StringPtrInput `pulumi:"action_timeout"`
	// Credentials for the connection type
	Credentials pulumi.StringMapInput `pulumi:"credentials"`
	// Settings for the connection type
	Settings pulumi.StringMapInput `pulumi:"settings"`
	// Used when reading the AEM instance state when determining the plan.
	State_timeout pulumi.StringPtrInput `pulumi:"state_timeout"`
	// Type of connection to use to connect to the machine on which AEM instance will be running.
	Type pulumi.StringInput `pulumi:"type"`
}

func (ClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil)).Elem()
}

func (i ClientArgs) ToClientOutput() ClientOutput {
	return i.ToClientOutputWithContext(context.Background())
}

func (i ClientArgs) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientOutput)
}

type ClientOutput struct{ *pulumi.OutputState }

func (ClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Client)(nil)).Elem()
}

func (o ClientOutput) ToClientOutput() ClientOutput {
	return o
}

func (o ClientOutput) ToClientOutputWithContext(ctx context.Context) ClientOutput {
	return o
}

// Used when trying to connect to the AEM instance machine (often right after creating it). Need to be enough long because various types of connections (like AWS SSM or SSH) may need some time to boot up the agent.
func (o ClientOutput) Action_timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Client) *string { return v.Action_timeout }).(pulumi.StringPtrOutput)
}

// Credentials for the connection type
func (o ClientOutput) Credentials() pulumi.StringMapOutput {
	return o.ApplyT(func(v Client) map[string]string { return v.Credentials }).(pulumi.StringMapOutput)
}

// Settings for the connection type
func (o ClientOutput) Settings() pulumi.StringMapOutput {
	return o.ApplyT(func(v Client) map[string]string { return v.Settings }).(pulumi.StringMapOutput)
}

// Used when reading the AEM instance state when determining the plan.
func (o ClientOutput) State_timeout() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Client) *string { return v.State_timeout }).(pulumi.StringPtrOutput)
}

// Type of connection to use to connect to the machine on which AEM instance will be running.
func (o ClientOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Client) string { return v.Type }).(pulumi.StringOutput)
}

type Compose struct {
	// Contents of the AEM Compose YML configuration file.
	Config *string `pulumi:"config"`
	// Script(s) for configuring a launched instance. Must be idempotent as it is executed always when changed. Typically used for installing AEM service packs, setting up replication agents, etc.
	Configure *InstanceScript `pulumi:"configure"`
	// Script(s) for creating an instance or restoring it from a backup. Typically customized to provide AEM library files (quickstart.jar, license.properties, service packs) from alternative sources (e.g., AWS S3, Azure Blob Storage). Instance recreation is forced if changed.
	Create *InstanceScript `pulumi:"create"`
	// Script(s) for deleting a stopped instance.
	Delete *InstanceScript `pulumi:"delete"`
	// Toggle automatic AEM Compose CLI wrapper download. If set to false, assume the wrapper is present in the data directory.
	Download *bool `pulumi:"download"`
	// Version of AEM Compose tool to use on remote machine.
	Version *string `pulumi:"version"`
}

// ComposeInput is an input type that accepts ComposeArgs and ComposeOutput values.
// You can construct a concrete instance of `ComposeInput` via:
//
//	ComposeArgs{...}
type ComposeInput interface {
	pulumi.Input

	ToComposeOutput() ComposeOutput
	ToComposeOutputWithContext(context.Context) ComposeOutput
}

type ComposeArgs struct {
	// Contents of the AEM Compose YML configuration file.
	Config pulumi.StringPtrInput `pulumi:"config"`
	// Script(s) for configuring a launched instance. Must be idempotent as it is executed always when changed. Typically used for installing AEM service packs, setting up replication agents, etc.
	Configure InstanceScriptPtrInput `pulumi:"configure"`
	// Script(s) for creating an instance or restoring it from a backup. Typically customized to provide AEM library files (quickstart.jar, license.properties, service packs) from alternative sources (e.g., AWS S3, Azure Blob Storage). Instance recreation is forced if changed.
	Create InstanceScriptPtrInput `pulumi:"create"`
	// Script(s) for deleting a stopped instance.
	Delete InstanceScriptPtrInput `pulumi:"delete"`
	// Toggle automatic AEM Compose CLI wrapper download. If set to false, assume the wrapper is present in the data directory.
	Download pulumi.BoolPtrInput `pulumi:"download"`
	// Version of AEM Compose tool to use on remote machine.
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ComposeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Compose)(nil)).Elem()
}

func (i ComposeArgs) ToComposeOutput() ComposeOutput {
	return i.ToComposeOutputWithContext(context.Background())
}

func (i ComposeArgs) ToComposeOutputWithContext(ctx context.Context) ComposeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComposeOutput)
}

func (i ComposeArgs) ToComposePtrOutput() ComposePtrOutput {
	return i.ToComposePtrOutputWithContext(context.Background())
}

func (i ComposeArgs) ToComposePtrOutputWithContext(ctx context.Context) ComposePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComposeOutput).ToComposePtrOutputWithContext(ctx)
}

// ComposePtrInput is an input type that accepts ComposeArgs, ComposePtr and ComposePtrOutput values.
// You can construct a concrete instance of `ComposePtrInput` via:
//
//	        ComposeArgs{...}
//
//	or:
//
//	        nil
type ComposePtrInput interface {
	pulumi.Input

	ToComposePtrOutput() ComposePtrOutput
	ToComposePtrOutputWithContext(context.Context) ComposePtrOutput
}

type composePtrType ComposeArgs

func ComposePtr(v *ComposeArgs) ComposePtrInput {
	return (*composePtrType)(v)
}

func (*composePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Compose)(nil)).Elem()
}

func (i *composePtrType) ToComposePtrOutput() ComposePtrOutput {
	return i.ToComposePtrOutputWithContext(context.Background())
}

func (i *composePtrType) ToComposePtrOutputWithContext(ctx context.Context) ComposePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComposePtrOutput)
}

type ComposeOutput struct{ *pulumi.OutputState }

func (ComposeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Compose)(nil)).Elem()
}

func (o ComposeOutput) ToComposeOutput() ComposeOutput {
	return o
}

func (o ComposeOutput) ToComposeOutputWithContext(ctx context.Context) ComposeOutput {
	return o
}

func (o ComposeOutput) ToComposePtrOutput() ComposePtrOutput {
	return o.ToComposePtrOutputWithContext(context.Background())
}

func (o ComposeOutput) ToComposePtrOutputWithContext(ctx context.Context) ComposePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Compose) *Compose {
		return &v
	}).(ComposePtrOutput)
}

// Contents of the AEM Compose YML configuration file.
func (o ComposeOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Compose) *string { return v.Config }).(pulumi.StringPtrOutput)
}

// Script(s) for configuring a launched instance. Must be idempotent as it is executed always when changed. Typically used for installing AEM service packs, setting up replication agents, etc.
func (o ComposeOutput) Configure() InstanceScriptPtrOutput {
	return o.ApplyT(func(v Compose) *InstanceScript { return v.Configure }).(InstanceScriptPtrOutput)
}

// Script(s) for creating an instance or restoring it from a backup. Typically customized to provide AEM library files (quickstart.jar, license.properties, service packs) from alternative sources (e.g., AWS S3, Azure Blob Storage). Instance recreation is forced if changed.
func (o ComposeOutput) Create() InstanceScriptPtrOutput {
	return o.ApplyT(func(v Compose) *InstanceScript { return v.Create }).(InstanceScriptPtrOutput)
}

// Script(s) for deleting a stopped instance.
func (o ComposeOutput) Delete() InstanceScriptPtrOutput {
	return o.ApplyT(func(v Compose) *InstanceScript { return v.Delete }).(InstanceScriptPtrOutput)
}

// Toggle automatic AEM Compose CLI wrapper download. If set to false, assume the wrapper is present in the data directory.
func (o ComposeOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Compose) *bool { return v.Download }).(pulumi.BoolPtrOutput)
}

// Version of AEM Compose tool to use on remote machine.
func (o ComposeOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Compose) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ComposePtrOutput struct{ *pulumi.OutputState }

func (ComposePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Compose)(nil)).Elem()
}

func (o ComposePtrOutput) ToComposePtrOutput() ComposePtrOutput {
	return o
}

func (o ComposePtrOutput) ToComposePtrOutputWithContext(ctx context.Context) ComposePtrOutput {
	return o
}

func (o ComposePtrOutput) Elem() ComposeOutput {
	return o.ApplyT(func(v *Compose) Compose {
		if v != nil {
			return *v
		}
		var ret Compose
		return ret
	}).(ComposeOutput)
}

// Contents of the AEM Compose YML configuration file.
func (o ComposePtrOutput) Config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compose) *string {
		if v == nil {
			return nil
		}
		return v.Config
	}).(pulumi.StringPtrOutput)
}

// Script(s) for configuring a launched instance. Must be idempotent as it is executed always when changed. Typically used for installing AEM service packs, setting up replication agents, etc.
func (o ComposePtrOutput) Configure() InstanceScriptPtrOutput {
	return o.ApplyT(func(v *Compose) *InstanceScript {
		if v == nil {
			return nil
		}
		return v.Configure
	}).(InstanceScriptPtrOutput)
}

// Script(s) for creating an instance or restoring it from a backup. Typically customized to provide AEM library files (quickstart.jar, license.properties, service packs) from alternative sources (e.g., AWS S3, Azure Blob Storage). Instance recreation is forced if changed.
func (o ComposePtrOutput) Create() InstanceScriptPtrOutput {
	return o.ApplyT(func(v *Compose) *InstanceScript {
		if v == nil {
			return nil
		}
		return v.Create
	}).(InstanceScriptPtrOutput)
}

// Script(s) for deleting a stopped instance.
func (o ComposePtrOutput) Delete() InstanceScriptPtrOutput {
	return o.ApplyT(func(v *Compose) *InstanceScript {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(InstanceScriptPtrOutput)
}

// Toggle automatic AEM Compose CLI wrapper download. If set to false, assume the wrapper is present in the data directory.
func (o ComposePtrOutput) Download() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Compose) *bool {
		if v == nil {
			return nil
		}
		return v.Download
	}).(pulumi.BoolPtrOutput)
}

// Version of AEM Compose tool to use on remote machine.
func (o ComposePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Compose) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type InstanceModel struct {
	// Version of the AEM instance. Reflects service pack installations.
	Aem_version string `pulumi:"aem_version"`
	// A brief description of the state details for a specific AEM instance. Possible states include 'created', 'uncreated', 'running', 'unreachable', 'up-to-date', and 'out-of-date'.
	Attributes []string `pulumi:"attributes"`
	// Remote path in which AEM instance is stored.
	Dir string `pulumi:"dir"`
	// Unique identifier of AEM instance defined in the configuration.
	Id string `pulumi:"id"`
	// A list of run modes for a specific AEM instance.
	Run_modes []string `pulumi:"run_modes"`
	// The machine-internal HTTP URL address used for communication with the AEM instance.
	Url string `pulumi:"url"`
}

type InstanceModelOutput struct{ *pulumi.OutputState }

func (InstanceModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceModel)(nil)).Elem()
}

func (o InstanceModelOutput) ToInstanceModelOutput() InstanceModelOutput {
	return o
}

func (o InstanceModelOutput) ToInstanceModelOutputWithContext(ctx context.Context) InstanceModelOutput {
	return o
}

// Version of the AEM instance. Reflects service pack installations.
func (o InstanceModelOutput) Aem_version() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceModel) string { return v.Aem_version }).(pulumi.StringOutput)
}

// A brief description of the state details for a specific AEM instance. Possible states include 'created', 'uncreated', 'running', 'unreachable', 'up-to-date', and 'out-of-date'.
func (o InstanceModelOutput) Attributes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceModel) []string { return v.Attributes }).(pulumi.StringArrayOutput)
}

// Remote path in which AEM instance is stored.
func (o InstanceModelOutput) Dir() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceModel) string { return v.Dir }).(pulumi.StringOutput)
}

// Unique identifier of AEM instance defined in the configuration.
func (o InstanceModelOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceModel) string { return v.Id }).(pulumi.StringOutput)
}

// A list of run modes for a specific AEM instance.
func (o InstanceModelOutput) Run_modes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceModel) []string { return v.Run_modes }).(pulumi.StringArrayOutput)
}

// The machine-internal HTTP URL address used for communication with the AEM instance.
func (o InstanceModelOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v InstanceModel) string { return v.Url }).(pulumi.StringOutput)
}

type InstanceModelArrayOutput struct{ *pulumi.OutputState }

func (InstanceModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstanceModel)(nil)).Elem()
}

func (o InstanceModelArrayOutput) ToInstanceModelArrayOutput() InstanceModelArrayOutput {
	return o
}

func (o InstanceModelArrayOutput) ToInstanceModelArrayOutputWithContext(ctx context.Context) InstanceModelArrayOutput {
	return o
}

func (o InstanceModelArrayOutput) Index(i pulumi.IntInput) InstanceModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstanceModel {
		return vs[0].([]InstanceModel)[vs[1].(int)]
	}).(InstanceModelOutput)
}

type InstanceScript struct {
	// Inline shell commands to be executed
	Inline []string `pulumi:"inline"`
	// Multiline shell script to be executed
	Script *string `pulumi:"script"`
}

// InstanceScriptInput is an input type that accepts InstanceScriptArgs and InstanceScriptOutput values.
// You can construct a concrete instance of `InstanceScriptInput` via:
//
//	InstanceScriptArgs{...}
type InstanceScriptInput interface {
	pulumi.Input

	ToInstanceScriptOutput() InstanceScriptOutput
	ToInstanceScriptOutputWithContext(context.Context) InstanceScriptOutput
}

type InstanceScriptArgs struct {
	// Inline shell commands to be executed
	Inline pulumi.StringArrayInput `pulumi:"inline"`
	// Multiline shell script to be executed
	Script pulumi.StringPtrInput `pulumi:"script"`
}

func (InstanceScriptArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceScript)(nil)).Elem()
}

func (i InstanceScriptArgs) ToInstanceScriptOutput() InstanceScriptOutput {
	return i.ToInstanceScriptOutputWithContext(context.Background())
}

func (i InstanceScriptArgs) ToInstanceScriptOutputWithContext(ctx context.Context) InstanceScriptOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceScriptOutput)
}

func (i InstanceScriptArgs) ToInstanceScriptPtrOutput() InstanceScriptPtrOutput {
	return i.ToInstanceScriptPtrOutputWithContext(context.Background())
}

func (i InstanceScriptArgs) ToInstanceScriptPtrOutputWithContext(ctx context.Context) InstanceScriptPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceScriptOutput).ToInstanceScriptPtrOutputWithContext(ctx)
}

// InstanceScriptPtrInput is an input type that accepts InstanceScriptArgs, InstanceScriptPtr and InstanceScriptPtrOutput values.
// You can construct a concrete instance of `InstanceScriptPtrInput` via:
//
//	        InstanceScriptArgs{...}
//
//	or:
//
//	        nil
type InstanceScriptPtrInput interface {
	pulumi.Input

	ToInstanceScriptPtrOutput() InstanceScriptPtrOutput
	ToInstanceScriptPtrOutputWithContext(context.Context) InstanceScriptPtrOutput
}

type instanceScriptPtrType InstanceScriptArgs

func InstanceScriptPtr(v *InstanceScriptArgs) InstanceScriptPtrInput {
	return (*instanceScriptPtrType)(v)
}

func (*instanceScriptPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceScript)(nil)).Elem()
}

func (i *instanceScriptPtrType) ToInstanceScriptPtrOutput() InstanceScriptPtrOutput {
	return i.ToInstanceScriptPtrOutputWithContext(context.Background())
}

func (i *instanceScriptPtrType) ToInstanceScriptPtrOutputWithContext(ctx context.Context) InstanceScriptPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceScriptPtrOutput)
}

type InstanceScriptOutput struct{ *pulumi.OutputState }

func (InstanceScriptOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstanceScript)(nil)).Elem()
}

func (o InstanceScriptOutput) ToInstanceScriptOutput() InstanceScriptOutput {
	return o
}

func (o InstanceScriptOutput) ToInstanceScriptOutputWithContext(ctx context.Context) InstanceScriptOutput {
	return o
}

func (o InstanceScriptOutput) ToInstanceScriptPtrOutput() InstanceScriptPtrOutput {
	return o.ToInstanceScriptPtrOutputWithContext(context.Background())
}

func (o InstanceScriptOutput) ToInstanceScriptPtrOutputWithContext(ctx context.Context) InstanceScriptPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InstanceScript) *InstanceScript {
		return &v
	}).(InstanceScriptPtrOutput)
}

// Inline shell commands to be executed
func (o InstanceScriptOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v InstanceScript) []string { return v.Inline }).(pulumi.StringArrayOutput)
}

// Multiline shell script to be executed
func (o InstanceScriptOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InstanceScript) *string { return v.Script }).(pulumi.StringPtrOutput)
}

type InstanceScriptPtrOutput struct{ *pulumi.OutputState }

func (InstanceScriptPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InstanceScript)(nil)).Elem()
}

func (o InstanceScriptPtrOutput) ToInstanceScriptPtrOutput() InstanceScriptPtrOutput {
	return o
}

func (o InstanceScriptPtrOutput) ToInstanceScriptPtrOutputWithContext(ctx context.Context) InstanceScriptPtrOutput {
	return o
}

func (o InstanceScriptPtrOutput) Elem() InstanceScriptOutput {
	return o.ApplyT(func(v *InstanceScript) InstanceScript {
		if v != nil {
			return *v
		}
		var ret InstanceScript
		return ret
	}).(InstanceScriptOutput)
}

// Inline shell commands to be executed
func (o InstanceScriptPtrOutput) Inline() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *InstanceScript) []string {
		if v == nil {
			return nil
		}
		return v.Inline
	}).(pulumi.StringArrayOutput)
}

// Multiline shell script to be executed
func (o InstanceScriptPtrOutput) Script() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *InstanceScript) *string {
		if v == nil {
			return nil
		}
		return v.Script
	}).(pulumi.StringPtrOutput)
}

type System struct {
	// Script executed once upon instance connection, often for mounting on VM data volumes from attached disks (e.g., AWS EBS, Azure Disk Storage). This script runs only once, even during instance recreation, as changes are typically persistent and system-wide. If re-execution is needed, it is recommended to set up a new machine.
	Bootstrap *InstanceScript `pulumi:"bootstrap"`
	// Remote root path in which AEM Compose files and unpacked AEM instances will be stored.
	Data_dir *string `pulumi:"data_dir"`
	// Environment variables for AEM instances.
	Env map[string]string `pulumi:"env"`
	// Contents of the AEM system service definition file (systemd).
	Service_config *string `pulumi:"service_config"`
	// System user under which AEM instance will be running. By default, the same as the user used to connect to the machine.
	User *string `pulumi:"user"`
	// Remote root path where provider-related files will be stored.
	Work_dir *string `pulumi:"work_dir"`
}

// SystemInput is an input type that accepts SystemArgs and SystemOutput values.
// You can construct a concrete instance of `SystemInput` via:
//
//	SystemArgs{...}
type SystemInput interface {
	pulumi.Input

	ToSystemOutput() SystemOutput
	ToSystemOutputWithContext(context.Context) SystemOutput
}

type SystemArgs struct {
	// Script executed once upon instance connection, often for mounting on VM data volumes from attached disks (e.g., AWS EBS, Azure Disk Storage). This script runs only once, even during instance recreation, as changes are typically persistent and system-wide. If re-execution is needed, it is recommended to set up a new machine.
	Bootstrap InstanceScriptPtrInput `pulumi:"bootstrap"`
	// Remote root path in which AEM Compose files and unpacked AEM instances will be stored.
	Data_dir pulumi.StringPtrInput `pulumi:"data_dir"`
	// Environment variables for AEM instances.
	Env pulumi.StringMapInput `pulumi:"env"`
	// Contents of the AEM system service definition file (systemd).
	Service_config pulumi.StringPtrInput `pulumi:"service_config"`
	// System user under which AEM instance will be running. By default, the same as the user used to connect to the machine.
	User pulumi.StringPtrInput `pulumi:"user"`
	// Remote root path where provider-related files will be stored.
	Work_dir pulumi.StringPtrInput `pulumi:"work_dir"`
}

func (SystemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*System)(nil)).Elem()
}

func (i SystemArgs) ToSystemOutput() SystemOutput {
	return i.ToSystemOutputWithContext(context.Background())
}

func (i SystemArgs) ToSystemOutputWithContext(ctx context.Context) SystemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemOutput)
}

func (i SystemArgs) ToSystemPtrOutput() SystemPtrOutput {
	return i.ToSystemPtrOutputWithContext(context.Background())
}

func (i SystemArgs) ToSystemPtrOutputWithContext(ctx context.Context) SystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemOutput).ToSystemPtrOutputWithContext(ctx)
}

// SystemPtrInput is an input type that accepts SystemArgs, SystemPtr and SystemPtrOutput values.
// You can construct a concrete instance of `SystemPtrInput` via:
//
//	        SystemArgs{...}
//
//	or:
//
//	        nil
type SystemPtrInput interface {
	pulumi.Input

	ToSystemPtrOutput() SystemPtrOutput
	ToSystemPtrOutputWithContext(context.Context) SystemPtrOutput
}

type systemPtrType SystemArgs

func SystemPtr(v *SystemArgs) SystemPtrInput {
	return (*systemPtrType)(v)
}

func (*systemPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**System)(nil)).Elem()
}

func (i *systemPtrType) ToSystemPtrOutput() SystemPtrOutput {
	return i.ToSystemPtrOutputWithContext(context.Background())
}

func (i *systemPtrType) ToSystemPtrOutputWithContext(ctx context.Context) SystemPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemPtrOutput)
}

type SystemOutput struct{ *pulumi.OutputState }

func (SystemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*System)(nil)).Elem()
}

func (o SystemOutput) ToSystemOutput() SystemOutput {
	return o
}

func (o SystemOutput) ToSystemOutputWithContext(ctx context.Context) SystemOutput {
	return o
}

func (o SystemOutput) ToSystemPtrOutput() SystemPtrOutput {
	return o.ToSystemPtrOutputWithContext(context.Background())
}

func (o SystemOutput) ToSystemPtrOutputWithContext(ctx context.Context) SystemPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v System) *System {
		return &v
	}).(SystemPtrOutput)
}

// Script executed once upon instance connection, often for mounting on VM data volumes from attached disks (e.g., AWS EBS, Azure Disk Storage). This script runs only once, even during instance recreation, as changes are typically persistent and system-wide. If re-execution is needed, it is recommended to set up a new machine.
func (o SystemOutput) Bootstrap() InstanceScriptPtrOutput {
	return o.ApplyT(func(v System) *InstanceScript { return v.Bootstrap }).(InstanceScriptPtrOutput)
}

// Remote root path in which AEM Compose files and unpacked AEM instances will be stored.
func (o SystemOutput) Data_dir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v System) *string { return v.Data_dir }).(pulumi.StringPtrOutput)
}

// Environment variables for AEM instances.
func (o SystemOutput) Env() pulumi.StringMapOutput {
	return o.ApplyT(func(v System) map[string]string { return v.Env }).(pulumi.StringMapOutput)
}

// Contents of the AEM system service definition file (systemd).
func (o SystemOutput) Service_config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v System) *string { return v.Service_config }).(pulumi.StringPtrOutput)
}

// System user under which AEM instance will be running. By default, the same as the user used to connect to the machine.
func (o SystemOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v System) *string { return v.User }).(pulumi.StringPtrOutput)
}

// Remote root path where provider-related files will be stored.
func (o SystemOutput) Work_dir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v System) *string { return v.Work_dir }).(pulumi.StringPtrOutput)
}

type SystemPtrOutput struct{ *pulumi.OutputState }

func (SystemPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**System)(nil)).Elem()
}

func (o SystemPtrOutput) ToSystemPtrOutput() SystemPtrOutput {
	return o
}

func (o SystemPtrOutput) ToSystemPtrOutputWithContext(ctx context.Context) SystemPtrOutput {
	return o
}

func (o SystemPtrOutput) Elem() SystemOutput {
	return o.ApplyT(func(v *System) System {
		if v != nil {
			return *v
		}
		var ret System
		return ret
	}).(SystemOutput)
}

// Script executed once upon instance connection, often for mounting on VM data volumes from attached disks (e.g., AWS EBS, Azure Disk Storage). This script runs only once, even during instance recreation, as changes are typically persistent and system-wide. If re-execution is needed, it is recommended to set up a new machine.
func (o SystemPtrOutput) Bootstrap() InstanceScriptPtrOutput {
	return o.ApplyT(func(v *System) *InstanceScript {
		if v == nil {
			return nil
		}
		return v.Bootstrap
	}).(InstanceScriptPtrOutput)
}

// Remote root path in which AEM Compose files and unpacked AEM instances will be stored.
func (o SystemPtrOutput) Data_dir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *System) *string {
		if v == nil {
			return nil
		}
		return v.Data_dir
	}).(pulumi.StringPtrOutput)
}

// Environment variables for AEM instances.
func (o SystemPtrOutput) Env() pulumi.StringMapOutput {
	return o.ApplyT(func(v *System) map[string]string {
		if v == nil {
			return nil
		}
		return v.Env
	}).(pulumi.StringMapOutput)
}

// Contents of the AEM system service definition file (systemd).
func (o SystemPtrOutput) Service_config() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *System) *string {
		if v == nil {
			return nil
		}
		return v.Service_config
	}).(pulumi.StringPtrOutput)
}

// System user under which AEM instance will be running. By default, the same as the user used to connect to the machine.
func (o SystemPtrOutput) User() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *System) *string {
		if v == nil {
			return nil
		}
		return v.User
	}).(pulumi.StringPtrOutput)
}

// Remote root path where provider-related files will be stored.
func (o SystemPtrOutput) Work_dir() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *System) *string {
		if v == nil {
			return nil
		}
		return v.Work_dir
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInput)(nil)).Elem(), ClientArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComposeInput)(nil)).Elem(), ComposeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComposePtrInput)(nil)).Elem(), ComposeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceScriptInput)(nil)).Elem(), InstanceScriptArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceScriptPtrInput)(nil)).Elem(), InstanceScriptArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemInput)(nil)).Elem(), SystemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemPtrInput)(nil)).Elem(), SystemArgs{})
	pulumi.RegisterOutputType(ClientOutput{})
	pulumi.RegisterOutputType(ComposeOutput{})
	pulumi.RegisterOutputType(ComposePtrOutput{})
	pulumi.RegisterOutputType(InstanceModelOutput{})
	pulumi.RegisterOutputType(InstanceModelArrayOutput{})
	pulumi.RegisterOutputType(InstanceScriptOutput{})
	pulumi.RegisterOutputType(InstanceScriptPtrOutput{})
	pulumi.RegisterOutputType(SystemOutput{})
	pulumi.RegisterOutputType(SystemPtrOutput{})
}
