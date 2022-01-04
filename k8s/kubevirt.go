package k8s

//基于kubevirt 对k8s的探索


import (
	"github.com/kubevirt/client-go"
	v1 "kubevirt.io/api/core/v1"
)

type DomainManager interface {
	//SyncVMI 为创建虚拟机
	SyncVMI(*v1.VirtualMachineInstance, bool, *cmdv1.VirtualMachineOptions) (*api.DomainSpec, error)
	//暂停VMI
	PauseVMI(*v1.VirtualMachineInstance) error
	//恢复暂停的VMI
	UnpauseVMI(*v1.VirtualMachineInstance) error
	KillVMI(*v1.VirtualMachineInstance) error
	//删除VMI
	DeleteVMI(*v1.VirtualMachineInstance) error
	SignalShutdownVMI(*v1.VirtualMachineInstance) error
	MarkGracefulShutdownVMI(*v1.VirtualMachineInstance) error
	ListAllDomains() ([]*api.Domain, error)
	//迁移VMI
	MigrateVMI(*v1.VirtualMachineInstance, *cmdclient.MigrationOptions) error
	PrepareMigrationTarget(*v1.VirtualMachineInstance, bool) error
	GetDomainStats() ([]*stats.DomainStats, error)
	//取消迁移
	CancelVMIMigration(*v1.VirtualMachineInstance) error
	//如下需要启用Qemu guest agent，没启用会包VMI does not have guest agent
	connectedGetGuestInfo() (v1.VirtualMachineInstanceGuestAgentInfo, error)
	GetUsers() ([]v1.VirtualMachineInstanceGuestOSUser, error)
	GetFilesystems() ([]v1.VirtualMachineInstanceFileSystem, error)
	SetGuestTime(*v1.VirtualMachineInstance) error
}

