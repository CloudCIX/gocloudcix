# Compute

## Backups

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupUpdateParam">ComputeBackupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackup">ComputeBackup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ListMetadata">ListMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupListResponse">ComputeBackupListResponse</a>

Methods:

- <code title="post /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupNewParams">ComputeBackupNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/{id}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/backups/{id}/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupUpdateParams">ComputeBackupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupResponse">ComputeBackupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/backups/">client.Compute.Backups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupListParams">ComputeBackupListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeBackupListResponse">ComputeBackupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## GPUs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUUpdateParam">ComputeGPUUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPU">ComputeGPU</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUListResponse">ComputeGPUListResponse</a>

Methods:

- <code title="get /compute/gpus/{id}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/gpus/{id}/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUUpdateParams">ComputeGPUUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUListParams">ComputeGPUListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUListResponse">ComputeGPUListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /compute/gpus/">client.Compute.GPUs.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUService.Attach">Attach</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUAttachParams">ComputeGPUAttachParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeGPUResponse">ComputeGPUResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Images

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImage">ComputeImage</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageGetResponse">ComputeImageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageListResponse">ComputeImageListResponse</a>

Methods:

- <code title="get /compute/images/{id}/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageGetResponse">ComputeImageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/images/">client.Compute.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageListParams">ComputeImageListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeImageListResponse">ComputeImageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Instances

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceUpdateParam">ComputeInstanceUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#Bom">Bom</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstance">ComputeInstance</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceListResponse">ComputeInstanceListResponse</a>

Methods:

- <code title="post /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceNewParams">ComputeInstanceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/{id}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/instances/{id}/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceUpdateParams">ComputeInstanceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceResponse">ComputeInstanceResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/instances/">client.Compute.Instances.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceListParams">ComputeInstanceListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeInstanceListResponse">ComputeInstanceListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Snapshots

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotUpdateParam">ComputeSnapshotUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshot">ComputeSnapshot</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>

Methods:

- <code title="post /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotNewParams">ComputeSnapshotNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/{id}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /compute/snapshots/{id}/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotUpdateParams">ComputeSnapshotUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotResponse">ComputeSnapshotResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /compute/snapshots/">client.Compute.Snapshots.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotListParams">ComputeSnapshotListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ComputeSnapshotListResponse">ComputeSnapshotListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Network

## Firewalls

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallUpdateParam">NetworkFirewallUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewall">NetworkFirewall</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallListResponse">NetworkFirewallListResponse</a>

Methods:

- <code title="post /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallNewParams">NetworkFirewallNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/{id}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/firewalls/{id}/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallUpdateParams">NetworkFirewallUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallResponse">NetworkFirewallResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/firewalls/">client.Network.Firewalls.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallListParams">NetworkFirewallListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkFirewallListResponse">NetworkFirewallListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## IPGroups

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupUpdateParam">NetworkIPGroupUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroup">NetworkIPGroup</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>

Methods:

- <code title="post /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupNewParams">NetworkIPGroupNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupUpdateParams">NetworkIPGroupUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupResponse">NetworkIPGroupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/ip_groups/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupListParams">NetworkIPGroupListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupListResponse">NetworkIPGroupListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /network/ip_groups/{id}/">client.Network.IPGroups.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkIPGroupService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Routers

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterUpdateParam">NetworkRouterUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#BaseIPAddress">BaseIPAddress</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouter">NetworkRouter</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterListResponse">NetworkRouterListResponse</a>

Methods:

- <code title="post /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterNewParams">NetworkRouterNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/{id}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/routers/{id}/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterUpdateParams">NetworkRouterUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterResponse">NetworkRouterResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/routers/">client.Network.Routers.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterListParams">NetworkRouterListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkRouterListResponse">NetworkRouterListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Vpns

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnUpdateParam">NetworkVpnUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpn">NetworkVpn</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnListResponse">NetworkVpnListResponse</a>

Methods:

- <code title="post /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnNewParams">NetworkVpnNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/{id}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /network/vpns/{id}/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnUpdateParams">NetworkVpnUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnResponse">NetworkVpnResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /network/vpns/">client.Network.Vpns.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnListParams">NetworkVpnListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#NetworkVpnListResponse">NetworkVpnListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Project

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectUpdateParam">ProjectUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#Project">Project</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectResponse">ProjectResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectListResponse">ProjectListResponse</a>

Methods:

- <code title="post /project/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectNewParams">ProjectNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/{id}/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /project/{id}/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectUpdateParams">ProjectUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectResponse">ProjectResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /project/">client.Project.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectListParams">ProjectListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#ProjectListResponse">ProjectListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Storage

## Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumesUpdateParam">StorageVolumesUpdateParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumes">StorageVolumes</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeListResponse">StorageVolumeListResponse</a>

Methods:

- <code title="post /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeNewParams">StorageVolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/{id}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /storage/volumes/{id}/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, id <a href="https://pkg.go.dev/builtin#int64">int64</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeUpdateParams">StorageVolumeUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumesResponse">StorageVolumesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /storage/volumes/">client.Storage.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeListParams">StorageVolumeListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go">gocloudcix</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/gocloudcix-go#StorageVolumeListResponse">StorageVolumeListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
