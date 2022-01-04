package k8s

//基于kubebuilder+ apimachinery 实现自定义web hook(设置默认总QPS，且限制单个pod的QPS 为1000）
// https://github.com/zq2599/blog_demos

/*

web hook 创建工作：
kubebuilder create webhook \
--group elasticweb \
--version v1 \
--kind ElasticWeb \
--defaulting \
--programmatic-validation

*/


import (
	"fmt"
	elasticWeb "github.com/zq2599/blog_demos/kubebuilder/elasticweb"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/util/validation/field"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/webhook"
	"strconv"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// log is for logging in this package.
var elasticweblog = logf.Log.WithName("elasticweb-resource")


// 期望状态
type ElasticWebSpec struct {
	// 业务服务对应的镜像，包括名称:tag
	Image string `json:"image"`
	// service占用的宿主机端口，外部请求通过此端口访问pod的服务
	Port *int32 `json:"port"`

	// 单个pod的QPS上限
	SinglePodQPS *int32 `json:"singlePodQPS"`
	// 当前整个业务的总QPS
	TotalQPS *int32 `json:"totalQPS"`
}

// 实际状态，该数据结构中的值都是业务代码计算出来的
type ElasticWebStatus struct {
	// 当前kubernetes中实际支持的总QPS
	RealQPS *int32 `json:"realQPS"`
}

// +kubebuilder:object:root=true

// ElasticWeb is the Schema for the elasticwebs API
type ElasticWeb struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ElasticWebSpec   `json:"spec,omitempty"`
	Status ElasticWebStatus `json:"status,omitempty"`
}

func (in *ElasticWeb) String() string {
	var realQPS string

	if nil == in.Status.RealQPS {
		realQPS = "nil"
	} else {
		realQPS = strconv.Itoa(int(*(in.Status.RealQPS)))
	}

	return fmt.Sprintf("Image [%s], Port [%d], SinglePodQPS [%d], TotalQPS [%d], RealQPS [%s]",
		in.Spec.Image,
		*(in.Spec.Port),
		*(in.Spec.SinglePodQPS),
		*(in.Spec.TotalQPS),
		realQPS)
}

// +kubebuilder:object:root=true

// ElasticWebList contains a list of ElasticWeb
type ElasticWebList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ElasticWeb `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ElasticWeb{}, &ElasticWebList{})
}


func (r *ElasticWeb) Default() {
	elasticweblog.Info("default", "name", r.Name)

	// TODO(user): fill in your defaulting logic.
	// 如果创建的时候没有输入总QPS，就设置个默认值
	if r.Spec.TotalQPS == nil {
		r.Spec.TotalQPS = new(int32)
		*r.Spec.TotalQPS = 1300
		elasticweblog.Info("a. TotalQPS is nil, set default value now", "TotalQPS", *r.Spec.TotalQPS)
	} else {
		elasticweblog.Info("b. TotalQPS exists", "TotalQPS", *r.Spec.TotalQPS)
	}
}

func (r *ElasticWeb) validateElasticWeb() error {
	var allErrs field.ErrorList

	if *r.Spec.SinglePodQPS > 1000 {
		elasticweblog.Info("c. Invalid SinglePodQPS")

		err := field.Invalid(field.NewPath("spec").Child("singlePodQPS"),
			*r.Spec.SinglePodQPS,
			"d. must be less than 1000")

		allErrs = append(allErrs, err)

		return apierrors.NewInvalid(
			schema.GroupKind{Group: "elasticweb.com.bolingcavalry", Kind: "ElasticWeb"},
			r.Name,
			allErrs)
	} else {
		elasticweblog.Info("e. SinglePodQPS is valid")
		return nil
	}
}

// ValidateCreate implements webhook.Validator so a webhook will be registered for the type
func (r *ElasticWeb) ValidateCreate() error {
	elasticweblog.Info("validate create", "name", r.Name)

	// TODO(user): fill in your validation logic upon object creation.

	return r.validateElasticWeb()
}

// ValidateUpdate implements webhook.Validator so a webhook will be registered for the type
func (r *ElasticWeb) ValidateUpdate(old runtime.Object) error {
	elasticweblog.Info("validate update", "name", r.Name)

	// TODO(user): fill in your validation logic upon object update.
	return r.validateElasticWeb()
}



func (r *ElasticWeb) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!

// +kubebuilder:webhook:path=/mutate-elasticweb-com-bolingcavalry-v1-elasticweb,mutating=true,failurePolicy=fail,groups=elasticweb.com.bolingcavalry,resources=elasticwebs,verbs=create;update,versions=v1,name=melasticweb.kb.io

var _ webhook.Defaulter = &ElasticWeb{}

// TODO(user): change verbs to "verbs=create;update;delete" if you want to enable deletion validation.
// +kubebuilder:webhook:verbs=create;update,path=/validate-elasticweb-com-bolingcavalry-v1-elasticweb,mutating=false,failurePolicy=fail,groups=elasticweb.com.bolingcavalry,resources=elasticwebs,versions=v1,name=velasticweb.kb.io

var _ webhook.Validator = &ElasticWeb{}


// ValidateDelete implements webhook.Validator so a webhook will be registered for the type
func (r *ElasticWeb) ValidateDelete() error {
	elasticweblog.Info("validate delete", "name", r.Name)

	// TODO(user): fill in your validation logic upon object deletion.
	return nil
}
