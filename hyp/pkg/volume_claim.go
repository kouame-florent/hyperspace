package hyp

import (
	"context"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"k8s.io/client-go/kubernetes"
)

type VolumeClaimObject struct {
	ObjectMeta
}

type VolumeClaimTemplate struct {
	TemplateMeta
	ClaimSize string
}

func NewVolumeClaimTemplate(name, namespace string, claimSize string) VolumeClaimTemplate {
	return VolumeClaimTemplate{
		TemplateMeta: TemplateMeta{
			Name:      name,
			Namespace: namespace,
		},
		ClaimSize: claimSize,
	}
}

func (v *VolumeClaimTemplate) CreateVolumeClaim(ctx context.Context, cls kubernetes.Interface) (VolumeClaimObject, error) {

	pvc := &corev1.PersistentVolumeClaim{
		ObjectMeta: metav1.ObjectMeta{
			Name: v.Name,
		},

		Spec: corev1.PersistentVolumeClaimSpec{
			AccessModes: []corev1.PersistentVolumeAccessMode{
				corev1.ReadWriteMany,
			},
			Resources: corev1.ResourceRequirements{
				Requests: corev1.ResourceList{
					corev1.ResourceStorage: resource.MustParse(v.ClaimSize),
				},
			},
		},
	}

	ctrVol, err := cls.CoreV1().PersistentVolumeClaims(v.Name).Create(ctx, pvc, metav1.CreateOptions{})
	if err != nil {
		return VolumeClaimObject{}, err
	}

	return VolumeClaimObject{
		ObjectMeta: ObjectMeta{
			UID:  string(ctrVol.UID),
			Name: ctrVol.Name,
		},
	}, nil
}
