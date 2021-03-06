package ovirt

import (
	libref "github.com/konveyor/controller/pkg/ref"
	"github.com/konveyor/forklift-controller/pkg/controller/provider/model/base"
)

//
// Kinds
var (
	DataCenterKind  = libref.ToKind(DataCenter{})
	VNICProfileKind = libref.ToKind(NICProfile{})
	ClusterKind     = libref.ToKind(Cluster{})
	HostKind        = libref.ToKind(Host{})
	NetKind         = libref.ToKind(Network{})
	StorageKind     = libref.ToKind(StorageDomain{})
	DiskKind        = "Disk" // TODO: UPDATE WITH MODEL.
	VmKind          = libref.ToKind(VM{})
)

//
// Types.
type Tree = base.Tree
type TreeNode = base.TreeNode
type BranchNavigator = base.BranchNavigator
type ParentNavigator = base.ParentNavigator
